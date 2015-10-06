package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

const (
	gameQuery   = "http://www.boardgamegeek.com/xmlapi2/thing?id={gameid}&stats=1&ratingcomments=1"
	searchQuery = "http://www.boardgamegeek.com/xmlapi2/search"
)

var (
	addr = flag.String("addr", "127.0.0.1:8181", "`host:port` to listen on")
)

func main() {
	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc("/game/{gameid}", handleGame)
	r.HandleFunc("/search", handleSearch)
	r.HandleFunc("/", handleMain)

	http.Handle("/", r)
	err := http.ListenAndServe(":8181", r)
	if err != nil {
		log.Fatal(err)
	}
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	games, err := GetTopRanked()
	if err != nil {
		httperr(w, "failed to retrieve top rated games", err)
		return
	}

	err = template.Must(template.New("toprankpage").Parse(mainPage)).Execute(w, games)
	if err != nil {
		httperr(w, "template error", err)
		return
	}
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	// build search request to send to BGG
	query := r.FormValue("query")

	form := url.Values{}
	form.Set("query", query)
	u, _ := url.ParseRequestURI(searchQuery)
	urlstr := fmt.Sprint(u)

	fmt.Println(urlstr)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlstr, bytes.NewBufferString(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// send request and parse data
	resp, err := client.Do(req)
	if err != nil {
		httperr(w, "failed to run BGG search", err)
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		httperr(w, "could not download search data", err)
		return
	}

	root := &RootSearch{}
	err = xml.Unmarshal(data, &root)
	if err != nil {
		httperr(w, "failed to parse search data", err)
		return
	}

	games, err := NewSearchList(root.Search)
	if err != nil {
		httperr(w, "failed to retrieve game details", err)
		return
	}

	err = template.Must(template.New("searchpage").Parse(searchPage)).Execute(w, games)
	if err != nil {
		httperr(w, "template error", err)
		return
	}
}

func httperr(w http.ResponseWriter, msg string, err error) {
	http.Error(w, msg, http.StatusInternalServerError)
	log.Print(msg+":", err)
	return
}

func handleGame(w http.ResponseWriter, r *http.Request) {
	gameid, err := strconv.Atoi(mux.Vars(r)["gameid"])
	if err != nil {
		httperr(w, "invalid game id", err)
		return
	}

	log.Println("getting game ", gameid)
	g, err := RetrieveGame(gameid)
	if err != nil {
		httperr(w, "failed to get game info", err)
		return
	}

	err = template.Must(template.New("gamepage").Parse(gamePage)).Execute(w, g)
	if err != nil {
		httperr(w, "template error", err)
		return
	}
}

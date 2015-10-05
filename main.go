package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

const gameQuery = "http://www.boardgamegeek.com/xmlapi2/thing?id={gameid}&stats=1&ratingcomments=1"

var (
	addr = flag.String("addr", "127.0.0.1:8181", "`host:port` to listen on")
)

func main() {
	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc("/game/{gameid}", handleGame)

	http.Handle("/", r)
	err := http.ListenAndServe(":8181", r)
	if err != nil {
		log.Fatal(err)
	}
}

func handleGame(w http.ResponseWriter, r *http.Request) {
	gameid := mux.Vars(r)["gameid"]
	gameurl := strings.Replace(gameQuery, "{gameid}", gameid, -1)
	log.Print("requesting game info from url " + gameurl)

	resp, err := http.Get(gameurl)
	if err != nil {
		http.Error(w, "could not retrieve raw game info", http.StatusNotFound)
		log.Print(err)
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "could not download game data", http.StatusInternalServerError)
		log.Print(err)
		return
	}

	game := &Root{}
	err = xml.Unmarshal(data, &game)
	if err != nil {
		http.Error(w, "failed to parse game data", http.StatusInternalServerError)
		log.Print(err)
		return
	}

	fmt.Fprintf(w, "%+v", game)
	//fmt.Fprintf(w, "%s", data)
}

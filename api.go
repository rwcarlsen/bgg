package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var topgame = regexp.MustCompile(`href="/boardgame/([0-9]+)`)
var starttbl = regexp.MustCompile(`Board Game Rank`)

func GetTopRanked() ([]*Game, error) {
	resp, err := http.Get("http://boardgamegeek.com/browse/boardgame")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	start := starttbl.FindIndex(data)[0]
	data = data[start:]
	idstrs := topgame.FindAllSubmatch(data, -1)
	ids := []int{}

	for _, match := range idstrs {
		id, err := strconv.Atoi(string(match[1]))
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	uniqids := []int{}
	for i, id := range ids {
		if i%2 == 0 {
			uniqids = append(uniqids, id)
		}
	}
	ids = uniqids

	var wg sync.WaitGroup
	wg.Add(len(ids))
	games := make([]*Game, len(ids))
	for i, id := range ids {
		go func(i, id int) {
			defer wg.Done()
			var err error
			games[i], err = RetrieveGame(id)
			if err != nil {
				log.Print(err)
				games[i] = &Game{Id: id, Name: strconv.Itoa(id)}
			}
		}(i, id)
	}
	wg.Wait()
	return games, nil
}

type RootSearch struct {
	Search []SearchRaw `xml:"item"`
}

type SearchRaw struct {
	Id            int     `xml:"id,attr"`
	Name          AttrVal `xml:"name"`
	YearPublished AttrVal `xml:"yearpublished"`
}

type SearchList []*Game

func NewSearchList(s []SearchRaw) (SearchList, error) {
	var wg sync.WaitGroup
	wg.Add(len(s))

	list := make(SearchList, len(s))
	for i, item := range s {
		go func(index int, item SearchRaw) {
			defer wg.Done()
			var err error
			list[index], err = RetrieveGame(item.Id)
			if err != nil {
				year, _ := strconv.Atoi(item.YearPublished.Val)
				log.Print(err)
				list[index] = &Game{Id: item.Id, Name: item.Name.Val, YearPublished: year}
			}
		}(i, item)
	}
	wg.Wait()

	return list, nil
}

func RetrieveGame(id int) (*Game, error) {
	gameurl := strings.Replace(gameQuery, "{gameid}", strconv.Itoa(id), -1)

	resp, err := http.Get(gameurl)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve game info: %v", err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not download game data: %v", err)
	}

	root := &Root{}
	err = xml.Unmarshal(data, &root)
	if err != nil {
		return nil, fmt.Errorf("failed to parse game data: %v", err)
	}

	game, err := NewGame(&root.Game)
	if err != nil {
		return nil, fmt.Errorf("failed to generate game struct: %v", err)
	}
	return game, nil
}

type Game struct {
	Id            int
	Name          string
	ThumbPath     string
	ImagePath     string
	MinPlayers    int
	MaxPlayers    int
	YearPublished int
	Description   []string
	PlayTime      time.Duration
	MinAge        int
	Links         []Link
	NUsersRated   int
	AverageRating float64
	RatingStddev  float64
	Rank          int
}

func imgpath(s string) string {
	s = strings.Trim(s, "/")
	if !strings.HasPrefix(s, "http://") {
		s = "http://" + s
	}
	return s
}

func NewGame(raw *RawGame) (*Game, error) {
	var err error
	g := &Game{
		ThumbPath:   imgpath(raw.ThumbPath),
		ImagePath:   imgpath(raw.ImagePath),
		Description: strings.Split(raw.Description, "&#10;"),
		Links:       raw.Links,
	}

	g.Id, err = strconv.Atoi(raw.Id)
	if err != nil {
		return nil, err
	}

	for _, name := range raw.Names {
		if strings.ToLower(name.Type) == "primary" {
			g.Name = name.Name
			break
		}
	}

	g.MinPlayers, err = strconv.Atoi(raw.MinPlayers.Val)
	if err != nil {
		log.Print(err)
	}

	g.MaxPlayers, err = strconv.Atoi(raw.MaxPlayers.Val)
	if err != nil {
		log.Print(err)
	}

	g.MinAge, err = strconv.Atoi(raw.MinAge.Val)
	if err != nil {
		log.Print(err)
	}

	g.YearPublished, err = strconv.Atoi(raw.YearPublished.Val)
	if err != nil {
		log.Print(err)
	}

	v, err := strconv.Atoi(raw.PlayingTime.Val)
	if err != nil {
		log.Print(err)
	}
	g.PlayTime = time.Duration(v) * time.Minute

	g.NUsersRated, err = strconv.Atoi(raw.Ratings.UsersRated.Val)
	if err != nil {
		log.Print(err)
	}

	g.AverageRating, err = strconv.ParseFloat(raw.Ratings.Average.Val, 64)
	if err != nil {
		log.Print(err)
	}

	g.RatingStddev, err = strconv.ParseFloat(raw.Ratings.Stddev.Val, 64)
	if err != nil {
		log.Print(err)
	}

	for _, rank := range raw.Ratings.Ranks {
		if rank.Name == "boardgame" {
			g.Rank, err = strconv.Atoi(rank.Value)
			if err != nil {
				log.Print(err)
			}
			break
		}
	}

	return g, nil
}

type Root struct {
	Game RawGame `xml:"item"`
}

type RawGame struct {
	Id            string  `xml:"id,attr"`
	ThumbPath     string  `xml:"thumbnail"`
	ImagePath     string  `xml:"image"`
	Names         []Name  `xml:"name"`
	MinPlayers    AttrVal `xml:"minplayers"`
	MaxPlayers    AttrVal `xml:"maxplayers"`
	YearPublished AttrVal `xml:"yearpublished"`
	Description   string  `xml:"description"`
	PlayingTime   AttrVal `xml:"playingtime"`
	MinAge        AttrVal `xml:"minage"`
	Links         []Link  `xml:"link"`
	Ratings       Ratings `xml:"statistics>ratings"`
}

type Link struct {
	Type  string `xml:"type,attr"`
	Id    string `xml:"id,attr"`
	Value string `xml:"value,attr"`
}

type Name struct {
	Type string `xml:"type,attr"`
	Name string `xml:"value,attr"`
}

type AttrVal struct {
	Val string `xml:"value,attr"`
}

type Ratings struct {
	UsersRated AttrVal `xml:"usersrated"`
	Average    AttrVal `xml:"average"`
	Stddev     AttrVal `xml:"stddev"`
	Ranks      []Rank  `xml:"ranks>rank"`
}

type Rank struct {
	Type         string `xml:"type,attr"`
	Id           string `xml:"id,attr"`
	Name         string `xml:"name,attr"`
	FriendlyName string `xml:"friendlyname,attr"`
	Value        string `xml:"value,attr"`
}

package main

import (
	"strconv"
	"strings"
	"time"
)

type Game struct {
	Id            int
	Name          string
	ThumbPath     string
	ImagePath     string
	MinPlayers    int
	MaxPlayers    int
	YearPublished int
	Description   string
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
		Description: strings.Replace(raw.Description, "&#10;", "</p><p>", -1),
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
		return nil, err
	}

	g.MaxPlayers, err = strconv.Atoi(raw.MaxPlayers.Val)
	if err != nil {
		return nil, err
	}

	g.MinAge, err = strconv.Atoi(raw.MinAge.Val)
	if err != nil {
		return nil, err
	}

	g.YearPublished, err = strconv.Atoi(raw.YearPublished.Val)
	if err != nil {
		return nil, err
	}

	v, err := strconv.Atoi(raw.PlayingTime.Val)
	if err != nil {
		return nil, err
	}
	g.PlayTime = time.Duration(v) * time.Minute

	g.NUsersRated, err = strconv.Atoi(raw.Ratings.UsersRated.Val)
	if err != nil {
		return nil, err
	}

	g.AverageRating, err = strconv.ParseFloat(raw.Ratings.Average.Val, 64)
	if err != nil {
		return nil, err
	}

	g.RatingStddev, err = strconv.ParseFloat(raw.Ratings.Stddev.Val, 64)
	if err != nil {
		return nil, err
	}

	for _, rank := range raw.Ratings.Ranks {
		if rank.Name == "boardgame" {
			g.Rank, err = strconv.Atoi(rank.Value)
			if err != nil {
				return nil, err
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

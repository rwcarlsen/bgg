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

func NewGame(raw *RawGame) (*Game, error) {
	var err error
	g := &Game{
		ThumbPath:   raw.ThumbPath,
		ImagePath:   raw.ImagePath,
		Description: raw.Description.Val,
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
	Description   AttrVal `xml:"description"`
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

/*
<?xml version="1.0" encoding="utf-8"?>
<items termsofuse="http://boardgamegeek.com/xmlapi/termsofuse">
  <item type="boardgame" id="124742">
    <thumbnail>//cf.geekdo-images.com/images/pic1324609_t.jpg</thumbnail>
    <image>//cf.geekdo-images.com/images/pic1324609.jpg</image>
    <name type="primary" sortindex="1" value="Android: Netrunner"/>
    <name type="alternate" sortindex="1" value="ネットランナー"/>
    <name type="alternate" sortindex="1" value="安卓纪元：矩阵潜袭"/>
    <description>short description</description>
    <yearpublished value="2012"/>
    <minplayers value="2"/>
    <maxplayers value="2"/>
    <playingtime value="45"/>
    <minplaytime value="45"/>
    <maxplaytime value="45"/>
    <minage value="14"/>
    <link type="boardgamecategory" id="1023" value="Bluffing"/>
    <link type="boardgamecategory" id="1002" value="Card Game"/>
    <link type="boardgamecategory" id="1016" value="Science Fiction"/>
    <statistics page="1">
      <ratings>
        <usersrated value="15414"/>
        <average value="8.16265"/>
        <bayesaverage value="7.97223"/>
        <ranks>
          <rank type="subtype" id="1" name="boardgame" friendlyname="Board Game Rank" value="7" bayesaverage="7.97223"/>
          <rank type="family" id="4667" name="cgs" friendlyname="Customizable Rank" value="2" bayesaverage="7.98248"/>
        </ranks>
        <stddev value="1.60795"/>
        <median value="0"/>
        <owned value="24556"/>
        <trading value="695"/>
        <wanting value="633"/>
        <wishing value="3281"/>
        <numcomments value="3226"/>
        <numweights value="1270"/>
        <averageweight value="3.3079"/>
      </ratings>
    </statistics>
  </item>
</items>
*/

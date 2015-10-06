package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"testing"
)

func TestGameParse(t *testing.T) {
	v := &Root{}
	err := xml.Unmarshal([]byte(netrunner), &v)
	if err != nil {
		log.Fatal(err)
	}

	g, err := NewGame(&v.Game)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", g)
}

func TestSearchParse(t *testing.T) {
	v := &RootSearch{}
	err := xml.Unmarshal([]byte(testsearch), &v)
	if err != nil {
		log.Fatal(err)
	}

	t.Logf("%+v\n", v.Search)
}

func TestGetTopRanked(t *testing.T) {
	games, err := GetTopRanked()
	if err != nil {
		t.Fatal(err)
	}
	for _, g := range games {
		fmt.Printf("%+v\n", g)
	}
}

const testsearch = `
<items total="67" termsofuse="http://boardgamegeek.com/xmlapi/termsofuse">
  <item type="boardgame" id="124742">
    <name type="primary" value="Android: Netrunner"/>
    <yearpublished value="2012"/>
  </item>
  <item type="boardgame" id="135103">
    <name type="primary" value="Android: Netrunner – A Study in Static"/>
    <yearpublished value="2013"/>
  </item>
  <item type="boardgame" id="160683">
    <name type="primary" value="Android: Netrunner – All That Remains"/>
    <yearpublished value="2014"/>
  </item>
</items>
`

const netrunner = `
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
`

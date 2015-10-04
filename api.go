package main

import "encoding/xml"

type Raw struct {
	Game Game `xml:"item"`
}

type Game struct {
	XMLName       xml.Name `xml:"item"`
	Id            string   `xml:"id,attr"`
	ThumbPath     string   `xml:"thumbnail"`
	ImagePath     string   `xml:"image"`
	Name          []Name   `xml:"name"`
	MinPlayers    AttrVal  `xml:"minplayers"`
	MaxPlayers    AttrVal  `xml:"maxplayers"`
	YearPublished AttrVal  `xml:"yearpublished"`
	Description   AttrVal  `xml:"description"`
	PlayingTime   AttrVal  `xml:"playingtime"`
	MinAge        AttrVal  `xml:"minage"`
	Links         []Link   `xml:"link"`
	Ratings       Ratings  `xml:"statistics>ratings"`
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
	Val string `xml:"value,attr`
}

type Ratings struct {
	XMLName    xml.Name `xml:"ratings"`
	UsersRated AttrVal  `xml:"usersrated"`
	Average    AttrVal  `xml:"average"`
	Stddev     AttrVal  `xml:"stddev"`
	//Ranks      []Rank   `xml:ranks`
}

//type Ratings struct {
//	Data RatingsInner `xml:"ratings"`
//	Page string       `xml:"page,attr"`
//}

/*
<?xml version="1.0" encoding="utf-8"?>
<items termsofuse="http://boardgamegeek.com/xmlapi/termsofuse">
  <item type="boardgame" id="124742">
    <thumbnail>//cf.geekdo-images.com/images/pic1324609_t.jpg</thumbnail>
    <image>//cf.geekdo-images.com/images/pic1324609.jpg</image>
    <name type="primary" sortindex="1" value="Android: Netrunner"/>
    <name type="alternate" sortindex="1" value="ネットランナー"/>
    <name type="alternate" sortindex="1" value="安卓纪元：矩阵潜袭"/>
    <description>Game description from the publisher&amp;#10;&amp;#10;Welcome to New Angeles, home of the Beanstalk. From our branch offices in this monument of human achievement, NBN proudly broadcasts all your favorite media programming. We offer fully comprehensive streaming in music and threedee, news and sitcoms, classic movies and sensies. We cover it all. Ours is a brave new age, and as humanity hurtles into space and the future with an astonishing series of new advances every day, NBN and our affiliates are keeping pace, bringing you all the vid that's fit to view.&amp;#10;&amp;#10;Android: Netrunner is an asymmetrical Living Card Game for two players. Set in the cyberpunk future of Android and Infiltration, the game pits a megacorporation and its massive resources against the subversive talents of lone runners.&amp;#10;&amp;#10;Corporations seek to score agendas by advancing them. Doing so takes time and credits. To buy the time and earn the credits they need, they must secure their servers and data forts with &amp;quot;ice&amp;quot;. These security programs come in different varieties, from simple barriers, to code gates and aggressive sentries. They serve as the corporation's virtual eyes, ears, and machine guns on the sprawling information superhighways of the network.&amp;#10;&amp;#10;In turn, runners need to spend their time and credits acquiring a sufficient wealth of resources, purchasing the necessary hardware, and developing suitably powerful ice-breaker programs to hack past corporate security measures. Their jobs are always a little desperate, driven by tight timelines, and shrouded in mystery. When a runner jacks-in and starts a run at a corporate server, he risks having his best programs trashed or being caught by a trace program and left vulnerable to corporate countermeasures. It's not uncommon for an unprepared runner to fail to bypass a nasty sentry and suffer massive brain damage as a result. Even if a runner gets through a data fort's defenses, there's no telling what it holds. Sometimes, the runner finds something of value. Sometimes, the best he can do is work to trash whatever the corporation was developing.&amp;#10;&amp;#10;The first player to seven points wins the game, but not likely before he suffers some brain damage or bad publicity.&amp;#10;&amp;#10;</description>
    <yearpublished value="2012"/>
    <minplayers value="2"/>
    <maxplayers value="2"/>
    <poll name="suggested_numplayers" title="User Suggested Number of Players" totalvotes="290">
      <results numplayers="1">
        <result value="Best" numvotes="2"/>
        <result value="Recommended" numvotes="4"/>
        <result value="Not Recommended" numvotes="196"/>
      </results>
      <results numplayers="2">
        <result value="Best" numvotes="266"/>
        <result value="Recommended" numvotes="19"/>
        <result value="Not Recommended" numvotes="2"/>
      </results>
      <results numplayers="2+">
        <result value="Best" numvotes="2"/>
        <result value="Recommended" numvotes="4"/>
        <result value="Not Recommended" numvotes="195"/>
      </results>
    </poll>
    <playingtime value="45"/>
    <minplaytime value="45"/>
    <maxplaytime value="45"/>
    <minage value="14"/>
    <poll name="suggested_playerage" title="User Suggested Player Age" totalvotes="143">
      <results>
        <result value="2" numvotes="0"/>
        <result value="3" numvotes="0"/>
        <result value="4" numvotes="0"/>
        <result value="5" numvotes="0"/>
        <result value="6" numvotes="0"/>
        <result value="8" numvotes="1"/>
        <result value="10" numvotes="16"/>
        <result value="12" numvotes="70"/>
        <result value="14" numvotes="52"/>
        <result value="16" numvotes="3"/>
        <result value="18" numvotes="0"/>
        <result value="21 and up" numvotes="1"/>
      </results>
    </poll>
    <poll name="language_dependence" title="Language Dependence" totalvotes="164">
      <results>
        <result level="1" value="No necessary in-game text" numvotes="0"/>
        <result level="2" value="Some necessary text - easily memorized or small crib sheet" numvotes="0"/>
        <result level="3" value="Moderate in-game text - needs crib sheet or paste ups" numvotes="7"/>
        <result level="4" value="Extensive use of text - massive conversion needed to be playable" numvotes="145"/>
        <result level="5" value="Unplayable in another language" numvotes="12"/>
      </results>
    </poll>
    <link type="boardgamecategory" id="1023" value="Bluffing"/>
    <link type="boardgamecategory" id="1002" value="Card Game"/>
    <link type="boardgamecategory" id="1016" value="Science Fiction"/>
    <link type="boardgamemechanic" id="2001" value="Action Point Allowance System"/>
    <link type="boardgamemechanic" id="2040" value="Hand Management"/>
    <link type="boardgamemechanic" id="2016" value="Secret Unit Deployment"/>
    <link type="boardgamemechanic" id="2015" value="Variable Player Powers"/>
    <link type="boardgamefamily" id="17106" value="Android"/>
    <link type="boardgamefamily" id="18254" value="Android: Netrunner LCG"/>
    <link type="boardgamefamily" id="5611" value="Cyberpunk"/>
    <link type="boardgamefamily" id="20584" value="Hackers"/>
    <link type="boardgamefamily" id="5337" value="Living Card Game®"/>
    <link type="boardgameexpansion" id="135103" value="Android: Netrunner – A Study in Static"/>
    <link type="boardgameexpansion" id="160683" value="Android: Netrunner – All That Remains"/>
    <link type="boardgameexpansion" id="172404" value="Android: Netrunner – Breaker Bay"/>
    <link type="boardgameexpansion" id="182375" value="Android: Netrunner – Business First"/>
    <link type="boardgameexpansion" id="172620" value="Android: Netrunner – Chrome City"/>
    <link type="boardgameexpansion" id="139596" value="Android: Netrunner – Creation and Control"/>
    <link type="boardgameexpansion" id="133500" value="Android: Netrunner – Cyber Exodus"/>
    <link type="boardgameexpansion" id="178866" value="Android: Netrunner – Data and Destiny"/>
    <link type="boardgameexpansion" id="149085" value="Android: Netrunner – Double Time"/>
    <link type="boardgameexpansion" id="147739" value="Android: Netrunner – Fear and Loathing"/>
    <link type="boardgameexpansion" id="157482" value="Android: Netrunner – First Contact"/>
    <link type="boardgameexpansion" id="137863" value="Android: Netrunner – Future Proof"/>
    <link type="boardgameexpansion" id="152314" value="Android: Netrunner – Honor and Profit"/>
    <link type="boardgameexpansion" id="136147" value="Android: Netrunner – Humanity's Shadow"/>
    <link type="boardgameexpansion" id="180303" value="Android: Netrunner – Kala Ghoda"/>
    <link type="boardgameexpansion" id="144640" value="Android: Netrunner – Mala Tempora"/>
    <link type="boardgameexpansion" id="175436" value="Android: Netrunner – Old Hollywood"/>
    <link type="boardgameexpansion" id="142727" value="Android: Netrunner – Opening Moves"/>
    <link type="boardgameexpansion" id="164314" value="Android: Netrunner – Order and Chaos"/>
    <link type="boardgameexpansion" id="143694" value="Android: Netrunner – Second Thoughts"/>
    <link type="boardgameexpansion" id="162359" value="Android: Netrunner – The Source"/>
    <link type="boardgameexpansion" id="156112" value="Android: Netrunner – The Spaces Between"/>
    <link type="boardgameexpansion" id="173909" value="Android: Netrunner – The Underway"/>
    <link type="boardgameexpansion" id="176426" value="Android: Netrunner – The Universe of Tomorrow"/>
    <link type="boardgameexpansion" id="168823" value="Android: Netrunner – The Valley"/>
    <link type="boardgameexpansion" id="132005" value="Android: Netrunner – Trace Amount"/>
    <link type="boardgameexpansion" id="146243" value="Android: Netrunner – True Colors"/>
    <link type="boardgameexpansion" id="159073" value="Android: Netrunner – Up and Over"/>
    <link type="boardgameexpansion" id="154544" value="Android: Netrunner – Upstalk"/>
    <link type="boardgameexpansion" id="130806" value="Android: Netrunner – What Lies Ahead"/>
    <link type="boardgameimplementation" id="1301" value="Netrunner" inbound="true"/>
    <link type="boardgamedesigner" id="14" value="Richard Garfield"/>
    <link type="boardgamedesigner" id="62803" value="Lukas Litzsinger"/>
    <link type="boardgameartist" id="56728" value="Bruno Balixa"/>
    <link type="boardgameartist" id="20316" value="Ralph Beisner"/>
    <link type="boardgameartist" id="66328" value="Del Borovic"/>
    <link type="boardgameartist" id="29100" value="Gong Studios"/>
    <link type="boardgameartist" id="24858" value="Henning Ludvigsen"/>
    <link type="boardgameartist" id="70695" value="Ed Mattinian"/>
    <link type="boardgameartist" id="48601" value="Adam Schumpert"/>
    <link type="boardgameartist" id="19229" value="Mark Anthony Taduran"/>
    <link type="boardgamepublisher" id="17" value="Fantasy Flight Games"/>
    <link type="boardgamepublisher" id="3475" value="Arclight"/>
    <link type="boardgamepublisher" id="2973" value="Edge Entertainment"/>
    <link type="boardgamepublisher" id="4617" value="Galakta"/>
    <link type="boardgamepublisher" id="12540" value="Game Harbor"/>
    <link type="boardgamepublisher" id="5530" value="Giochi Uniti"/>
    <link type="boardgamepublisher" id="264" value="Heidelberger Spieleverlag"/>
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

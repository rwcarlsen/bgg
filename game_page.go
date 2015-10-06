package main

const cssstyle = `
@media screen and (min-width:600px) {
	#menu {
		margin-left: 0;
		max-width: 200px;
		position: fixed;
		top: 0;
		left: 0;
		bottom: 0;
	}

	#layout {
		padding-left: 200px;
	}

	.content {
		padding-right: 2em;
		padding-left: 2em;
	}
}

@media screen and (max-width:599px) {
	.content {
		padding-right: 0.4em;
		padding-left: 0.4em;
	}
}

#menu {
	z-index: 1000;
	overflow-y: auto;
	background: #191818;
}

#menu a {
	color: #999;
	display: block;
	border: none;
	white-space: normal;
	padding: 0.625em 1em;
	margin: auto;
	text-decoration: none;
}

#menu > a {
	font-size: 1.25em;
}

#menu ul {
	border: none;
	background: transparent;
	border-top: 1px solid #333;
	margin-top: 0;
	padding: 0;
}

#menu li {
	display: block;
	margin: auto;
}

#menu li a:hover,
#menu li a:focus {
	display: block;
	background: #333;
}

.searchlist,
.game {
	margin-left: auto;
	margin-right: auto;
	max-width: 768px;
}

#searchbox {
	border: none;
	display: block;
	width: 80%;
	margin: auto;
	padding: 0;
}

#searchform {
	width: 100%;
	margin: auto;
	border-top: 1px solid #333;
	padding: 10px 0;
}

#submitbutton {
	display: none;
}
`

const menu = `
<div id="menu">
    <a href="/">pBGG</a>

	<div id="searchform">
		<form method="get" action="/search">
			<input id="searchbox" type="text" name="query">
			<input id="submitbutton" type="submit" value="Go">
		</form>
	</div>

    <ul>
        <li><a href="/">List by Rank</a></li>
        <li><a href="/game/124742">Android Netrunner</a></li>
    </ul>
</div>
`

const gamePage = `
<!DOCTYPE html>
<html class="no-js" lang="en-US">
<head>
	<title>pBGG</title>
    <style type="text/css">
	` + cssstyle + `
    </style>
</head>
<body lang="en">

` + menu + `

<div class="game">
    <br>
    <img src="{{.ImagePath}}" style="width:250px">

    <br>
    <h2>{{.Name}} ({{.YearPublished}})</h2>
    <i><b>{{.AverageRating}} average rating ({{.NUsersRated}} reviews), rank {{.Rank}}</b></i>

    <br>
    {{range .Description}}
	<p>{{.}}</p>
	{{end}}

    <ul>
    <li>{{.MinPlayers}} to {{.MaxPlayers}} players</li>
    <li>{{.PlayTime}} play-time</li>
    <li>{{.MinAge}}+ years</li>
    </ul>
</div>

</body>
</html>
`

const searchPage = `
<!DOCTYPE html>
<html class="no-js" lang="en-US">
<head>
	<title>pBGG</title>
    <style type="text/css">
	` + cssstyle + tablecss + `
    </style>
</head>
<body lang="en">

` + menu + `

<div class="searchlist">
	<table>
		<tr><th></th><th>Name</th><th>Rank</th><th>Rating</th></tr>

		{{ range . }}
		
			<tr class="searchrow">
				<td style="text-align: center;"> <a href="/game/{{.Id}}"><img src="{{.ThumbPath}}" style="maxheight:70px; maxwidth:70px;"></a></td>
				<td style="text-align: left"> <a href="/game/{{.Id}}">{{.Name}} ({{.YearPublished}})</a> </td>
				<td style="text-align: center"> {{if gt .Rank 0}} {{.Rank}} {{end}}</td>
				<td style="text-align: left"> {{.AverageRating}} <br> ({{.NUsersRated}} users) </td>
			</tr>
		</a>
		{{ end }}
	</table>
</div>

</body>
</html>
`

const mainPage = `
<!DOCTYPE html>
<html class="no-js" lang="en-US">
<head>
	<title>pBGG</title>
    <style type="text/css">
	` + cssstyle + tablecss + `
    </style>
</head>
<body lang="en">

` + menu + `

<div class="searchlist">
	<table>
		<tr><th></th><th>Rank</th><th>Name</th><th>Rating</th></tr>

		{{ range . }}
		
			<tr class="searchrow">
				<td style="text-align: center;"> <a href="/game/{{.Id}}"><img src="{{.ThumbPath}}" style="maxheight:70px; maxwidth:70px;"></a></td>
				<td style="text-align: center"> {{if gt .Rank 0}} {{.Rank}} {{end}}</td>
				<td style="text-align: left"> <a href="/game/{{.Id}}">{{.Name}} ({{.YearPublished}})</a> </td>
				<td style="text-align: left"> {{.AverageRating}} <br> ({{.NUsersRated}} users) </td>
			</tr>
		</a>
		{{ end }}
	</table>
</div>

</body>
</html>
`

const tablecss = `
table {
	width:100%;
	border-color:#a9a9a9;
	color:#333333;
	border-collapse:collapse;
	margin:auto;
	border-width:1px;
	text-align:left;
}
th {
	padding:4px;
	border-style:solid;
	border-color:#a9a9a9;
	border-width:1px;
	background-color:#b8b8b8;
	text-align:left;
}
tr {
	maxwidth=50px;
	height=70px;
	background-color:#ffffff;
	text-align:left;
}
td {
	padding:4px;
	border-color:#a9a9a9;
	border-style:solid;
	border-width:1px;
	text-align:center;
}
`

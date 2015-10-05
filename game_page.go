package main

const gamePage = `
<!DOCTYPE html>
<html class="no-js" lang="en-US">
<head>
	<title>pBGG</title>
    <link href="/style.css" rel="stylesheet" type="text/css">
</head>
<body lang="en">

<br>

<img src="{{.ImagePath}}" style="width:250px">

<br>

<h2>{{.Name}}</h2>
<i><b>{{.AverageRating}} average rating ({{.NUsersRated}} reviews), rank {{.Rank}}</b></i>

<br>

<p>{{.Description}}</p>

<ul>
<li>{{.MinPlayers}} to {{.MaxPlayers}} players</li>
<li>{{.PlayTime}} play-time</li>
<li>Min. age {{.MinAge}} years</li>
</ul>

</body>
</html>
`

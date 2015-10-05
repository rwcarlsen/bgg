package main

const gamePage = `
<!DOCTYPE html>
<html class="no-js" lang="en-US">
<head>
	<title>pBGG</title>
    <link href="/style.css" rel="stylesheet" type="text/css">
    <style type="text/css">
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

        .game {
            margin-left: auto;
            margin-right: auto;
            max-width: 768px;
        }

    </style>
</head>
<body lang="en">

<div id="menu">
    <a href="/">pBGG</a>
    <ul>
        <li><a href="/game/124742">Android Netrunner</a></li>
    </ul>
</div>

<div class="game">
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
</div>

</body>
</html>
`

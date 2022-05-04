<p align="center">
<img src="https://img.shields.io/github/go-mod/go-version/radenrishwan/anime-api" alt="go">
<img src="https://img.shields.io/badge/Fiber-2.32.0-blue" alt="fiber">
<img src="https://img.shields.io/badge/Goquery-1.8.0-blue" alt="goquery">
<img src="https://img.shields.io/badge/Chromedp-0.8.1-red" alt="goquery">
</p>

# 👋 GOGOANIME API

an unofficial anime api scrap from GOGOANIME

# 👻 ShowCase

COMMING SOON

# 🪛 Installation

before you running api, please make sure your computer has already installed Go

## 🐳 Using Docker

this project running on port 3000. if you want to change the port, change on docker-compose file.

```bash
$ docker-compose -f docker-compose.yml up -d
```

now you can access this link to make sure app is running :

```bash
$ curl http://127.0.0.1:3000/
```

and the output:

```bash
{"author":"Raden Mohamad Rishwan","github":"radenrishwan","status":"Heyy its works"}
```

## <img src="https://go.dev/blog/gopher/gopher.png" width=22 height=22> Build using GO

if you didnt install docker, you can build this project using go.

1. first, you need to download all dependencies
    ```bash
    $ go mod download
    ```
2. after that, you can build an executable file or running using GO
    ```bash
    $ go build -o main .
    ```
   or
    ```bash
    $ go main.exe
    ```
3. if you build to executable file, you need running app with:
    ```bash
    $ ./main.exe
    ```

## ❤️‍🔥API DOC

ENDPOINT :

```bash
/api/v1/anime
```

### GET RECENT ANIME

- Url : `/recent-anime`
- Method : GET
- Response :

```json
{
  "code": 200,
  "message": "Success",
  "data": [
    {
      "title": "Devidol!",
      "image": "https://gogocdn.net/cover/devidol.png",
      "last_episode": "Episode 12",
      "url": "https://gogoanime.sk/devidol-episode-12"
    },
    {
      "title": "Araiguma Rascal",
      "image": "https://gogocdn.net/cover/araiguma-rascal.png",
      "last_episode": "Episode 11",
      "url": "https://gogoanime.sk/araiguma-rascal-episode-11"
    },
    {
      "title": "Madou King Granzort",
      "image": "https://gogocdn.net/cover/madou-king-granzort.png",
      "last_episode": "Episode 32",
      "url": "https://gogoanime.sk/madou-king-granzort-episode-32"
    },
    {
      "title": "Kaginado Season 2",
      "image": "https://gogocdn.net/cover/kaginado-season-2.png",
      "last_episode": "Episode 4",
      "url": "https://gogoanime.sk/kaginado-season-2-episode-4"
    },
    {
      "title": "Tomodachi Game",
      "image": "https://gogocdn.net/cover/tomodachi-game.png",
      "last_episode": "Episode 5",
      "url": "https://gogoanime.sk/tomodachi-game-episode-5"
    }
  ]
}
```

### GET LIST ANIME

- Url : `/list-anime/:page` *page must be number
- Method : GET
- Response :

```json
{
  "code": 200,
  "message": "Success",
  "data": [
    {
      "title": ".Hack//G.U. Returner",
      "slug": "hackgu-returner",
      "url": "https://gogoanime.sk/category/hackgu-returner"
    },
    {
      "title": ".hack//G.U. Trilogy",
      "slug": "hackgu-trilogy",
      "url": "https://gogoanime.sk/category/hackgu-trilogy"
    },
    {
      "title": ".hack//Gift",
      "slug": "hackgift",
      "url": "https://gogoanime.sk/category/hackgift"
    },
    {
      "title": ".hack//Legend of the Twilight",
      "slug": "hacklegend-of-the-twilight",
      "url": "https://gogoanime.sk/category/hacklegend-of-the-twilight"
    },
    {
      "title": ".hack//Liminality",
      "slug": "hackliminality",
      "url": "https://gogoanime.sk/category/hackliminality"
    }
  ]
}
```

### GET ANIME INFO

- Url : `/:name`
- Method : GET
- Response :

```json
{
  "code": 200,
  "message": "Success",
  "data": {
    "title": "86",
    "alternative_name": "Eighty Six; 86―エイティシックス―",
    "img": "https://gogocdn.net/cover/86.png",
    "type": "Spring 2021 Anime",
    "plot": "The Republic of San Magnolia.\n\nFor a long time, this country has been besieged by its neighbor, the Giadian Empire, which created a series of unmanned drones called the Legion. After years of painstaking research, the Republic finally developed autonomous drones of their own, turning the one-sided struggle into a war without casualties—or at least, that's what the government claims.\n\nIn truth, there is no such thing as a bloodless war. Beyond the fortified walls protecting the eighty-five Republic territories lies the \"nonexistent\" Eighty-Sixth Sector. The young men and women of this forsaken land are branded the Eighty-Six and, stripped of their humanity, pilot the \"unmanned\" weapons into battle...\n\nShinn directs the actions of a detachment of young Eighty-Sixers while on the battlefield. Lena is a \"handler\" who commands the detachment from the remote rear with the help of special communication.\n\nThe farewell story of the severe and sad struggle of these two begins!",
    "genre": "Drama, Sci-Fi",
    "released": "2021",
    "status": "Completed"
  }
}
```

### GET DOWNLOAD LINK

- Url : `/:name/downloads`
- Method : GET
- Response :

```json
{
  "code": 200,
  "message": "Success",
  "data": {
    "name": "86",
    "total_episode": 11,
    "episodes": [
      {
        "episode": "5",
        "url": "https://gogoanime.sk/86-episode-5"
      },
      {
        "episode": "4",
        "url": "https://gogoanime.sk/86-episode-4"
      },
      {
        "episode": "3",
        "url": "https://gogoanime.sk/86-episode-3"
      },
      {
        "episode": "2",
        "url": "https://gogoanime.sk/86-episode-2"
      },
      {
        "episode": "1",
        "url": "https://gogoanime.sk/86-episode-1"
      }
    ]
  }
}
```

### ERROR RESPONSE

response :

```json
{
  "code": "errorcode",
  "message": "errormessage",
  "data": "some message error"
}
```

# 📃 TODO
- Add stream link
- Showcase using this API
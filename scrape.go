package main

import (
	"context"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/radenrishwan/anime-api/model"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type animeScrape struct {
}

func NewAnimeScrape() AnimeScrape {
	return &animeScrape{}
}

type AnimeScrape interface {
	GetRecentAnime() ([]model.AnimeBanner, error)
	GetListAnime(page int) ([]model.ListAnime, error)
	GetAnimeInfoByName(name string) (model.Anime, error)
	GetDownloadPage(name string) (model.Episodes, error)
	GetListGenre() ([]model.Genre, error)
	GetGenre(name string, page int) ([]model.SearchBanner, error)
	FindAnime(name string) ([]model.SearchBanner, error)
}

func (scrape *animeScrape) GetRecentAnime() ([]model.AnimeBanner, error) {
	response := GetHTML("https://gogoanime.sk/")

	reader, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var animes []model.AnimeBanner
	reader.Find(".items li").Each(func(_ int, selection *goquery.Selection) {
		var anime model.AnimeBanner

		img, exists := selection.Find(".img a img").Attr("src")
		url, exists := selection.Find(".img a").Attr("href")
		name := selection.Find(".name a").Text()
		lastEpisode := selection.Find(".episode").Text()
		slug, exists := selection.Find(".img a").Attr("href")

		if exists {
			anime.Image = img
			anime.Url = "https://" + ENDPOINT + url
			anime.Title = name
			anime.LastEpisode = lastEpisode
			anime.Slug = slug[1 : len(slug)-(len(lastEpisode)+1)]

			animes = append(animes, anime)
		}
	})

	if animes == nil {
		return animes, errors.New("anime not found")

	}

	return animes, nil
}

func (scrape *animeScrape) GetListAnime(page int) ([]model.ListAnime, error) {
	response := GetHTML("https://gogoanime.sk/anime-list.html?page=" + strconv.Itoa(page))

	reader, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var animes []model.ListAnime
	reader.Find(".listing li").Each(func(_ int, selection *goquery.Selection) {
		var anime model.ListAnime

		url, exists := selection.Find("a").Attr("href")

		if exists {
			anime.Url = "https://" + ENDPOINT + url
			anime.Title = selection.Find("a").Text()
			anime.Slug = strings.ReplaceAll(url, "/category/", "")

			animes = append(animes, anime)
		}
	})

	if animes == nil {
		return animes, errors.New("anime not found")
	}

	return animes, nil
}

func (scrape *animeScrape) GetAnimeInfoByName(name string) (model.Anime, error) {
	response := GetHTML("https://gogoanime.sk/category/" + name)

	reader, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var anime model.Anime
	reader.Find(".anime_info_body_bg").Each(func(_ int, selection *goquery.Selection) {
		img, exists := selection.ChildrenFiltered("img").Attr("src")
		title := selection.Find("h1").Text()

		selection.Find(".type").Each(func(index int, selection *goquery.Selection) {
			info := selection.Find("span").Text()
			element := strings.TrimSpace(selection.Text()[len(info):len(selection.Text())]) // trim space and delete categories chars
			switch index {
			case 0:
				anime.Type = element
			case 1:
				anime.Plot = element
			case 2:
				anime.Genre = element
			case 3:
				anime.Released = element
			case 4:
				anime.Status = element
			default:
				anime.AlternativeName = element
			}
		})

		if exists {
			anime.Img = img
			anime.Title = title
		}
	})

	if anime.Img == "" {
		return anime, errors.New("anime not found")
	}

	return anime, nil
}

func (scrape *animeScrape) GetDownloadPage(name string) (model.Episodes, error) {
	var result model.Episodes
	var episodes []model.Episode

	anime, err := scrape.GetAnimeInfoByName(name)
	if err != nil {
		return result, err
	}

	result.Title = anime.Title

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var response string
	err = chromedp.Run(ctx,
		chromedp.Navigate("https://gogoanime.sk/category/"+name),
		chromedp.InnerHTML("#episode_related", &response),
	)

	if err != nil {
		log.Fatalln(err)
	}

	reader, err := goquery.NewDocumentFromReader(strings.NewReader(response))
	if err != nil {
		log.Fatalln(err)
	}

	reader.Find("li").Each(func(index int, selection *goquery.Selection) {
		urlDump, exists := selection.Find("a").Attr("href")
		url := "https://" + ENDPOINT + strings.TrimSpace(urlDump)

		epsDump := selection.Find("a .name").Text()
		eps := strings.TrimSpace(epsDump[len(selection.Find("a .name span").Text()):]) // delete EP and space

		var episode model.Episode
		if exists {
			episode.Url = url
			episode.Episode = eps

			episodes = append(episodes, episode)
		}
	})

	result.Episodes = episodes
	result.TotalEpisode = len(episodes)

	if episodes == nil {
		return result, errors.New("episodes not found")
	}

	return result, nil
}

func (scrape *animeScrape) GetListGenre() ([]model.Genre, error) {
	var genres []model.Genre

	response, err := http.Get("https://gogoanime.sk")
	if err != nil {
		log.Fatalln(err)
	}

	reader, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	reader.Find(".genre ul li a").Each(func(_ int, selection *goquery.Selection) {
		url, exists := selection.Attr("href")
		info := selection.Text()

		if exists {
			genres = append(genres, model.Genre{
				Url:   "https://" + ENDPOINT + url,
				Genre: strings.ToLower(info),
			})
		}
	})

	if genres == nil {
		return genres, errors.New("genres not found")
	}

	return genres, nil
}

func (scrape *animeScrape) GetGenre(name string, page int) ([]model.SearchBanner, error) {
	var genres []model.SearchBanner

	response, err := http.Get("https://gogoanime.sk/genre/" + name + "?page=" + strconv.Itoa(page))
	if err != nil {
		log.Fatalln(err)
	}

	reader, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	reader.Find(".last_episodes .items li").Each(func(_ int, selection *goquery.Selection) {
		img, exists := selection.Find(".img a img").Attr("src")
		url, exists := selection.Find(".img a").Attr("href")
		title := selection.Find(".name").Text()
		release := selection.Find(".released").Text()[9:]

		if exists {
			genres = append(genres, model.SearchBanner{
				Image:    img,
				Url:      "https://" + ENDPOINT + url,
				Title:    title,
				Released: strings.TrimSpace(release),
				Slug:     url[10:],
			})
		}
	})

	if genres == nil {
		return genres, errors.New("genre not found")
	}

	return genres, nil
}

func (scrape *animeScrape) FindAnime(name string) ([]model.SearchBanner, error) {
	var genres []model.SearchBanner

	response, err := http.Get("https://gogoanime.sk//search.html?keyword=" + name)
	if err != nil {
		log.Fatalln(err)
	}

	reader, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	reader.Find(".last_episodes .items li").Each(func(_ int, selection *goquery.Selection) {
		img, exists := selection.Find(".img a img").Attr("src")
		url, exists := selection.Find(".img a").Attr("href")
		title := selection.Find(".name").Text()
		release := selection.Find(".released").Text()[9:]

		if exists {
			genres = append(genres, model.SearchBanner{
				Image:    img,
				Url:      "https://" + ENDPOINT + url,
				Title:    title,
				Released: strings.TrimSpace(release),
				Slug:     url[10:],
			})
		}
	})

	if genres == nil {
		return genres, errors.New("genre not found")
	}

	return genres, nil
}

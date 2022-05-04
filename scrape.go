package main

import (
	"context"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/radenrishwan/anime-api/model"
	"log"
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
	//GetDownloadLink(name string)
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

		if exists {
			anime.Image = img
			anime.Url = "https://" + ENDPOINT + url
			anime.Title = name
			anime.LastEpisode = lastEpisode

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

//func (scrape animeScrape) GetDownloadLink(name string) {
//	ctx, cancel := chromedp.NewContext(context.Background())
//	defer cancel()
//
//	episodes, err := scrape.GetDownloadPage(name)
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	a := 1
//	for _, episode := range episodes {
//		responseDownloadUrl, err := http.Get(episode.Url)
//		if err != nil {
//			log.Fatalln(err)
//		}
//
//		readerDownloadUrl, err := goquery.NewDocumentFromReader(responseDownloadUrl.Body)
//		if err != nil {
//			log.Fatalln(err)
//		}
//
//		downloadUrl, exists := readerDownloadUrl.Find(".dowloads a").Attr("href")
//		fmt.Println(downloadUrl)
//		if exists && a == 1 {
//			a += 1
//
//			var result string
//			err := chromedp.Run(ctx,
//				chromedp.Navigate(downloadUrl),
//				chromedp.WaitVisible(".mirror_link"),
//				chromedp.InnerHTML(".content_c_bg", &result),
//			)
//
//			//response := strings.NewReader(result)
//			//reader, err := goquery.NewDocumentFromReader(response)
//			//
//			//reader.Find("#content-download .mirror_link").Each(func(_ int, selection *goquery.Selection) {
//			//	//url, existUrl := selection.Find("a").Attr("href")
//			//
//			//	fmt.Println(selection.Find("h6").Text())
//			//})
//
//			if err != nil {
//				log.Fatalln(err)
//			}
//
//		}
//	}
//}

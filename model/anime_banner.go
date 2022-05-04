package model

type AnimeBanner struct {
	Title       string `json:"title"`
	Image       string `json:"image"`
	LastEpisode string `json:"last_episode"`
	Url         string `json:"url"`
	Slug        string `json:"slug"`
}

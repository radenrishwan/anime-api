package model

type Episodes struct {
	Title        string    `json:"name"`
	TotalEpisode int       `json:"total_episode"`
	Episodes     []Episode `json:"episodes"`
}

type Episode struct {
	Episode string `json:"episode"`
	Url     string `json:"url"`
}

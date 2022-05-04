package model

type SearchBanner struct {
	Title    string `json:"title"`
	Image    string `json:"image"`
	Released string `json:"released"`
	Url      string `json:"url"`
	Slug     string `json:"slug"`
}

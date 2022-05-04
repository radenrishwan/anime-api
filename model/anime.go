package model

type Anime struct {
	Title           string `json:"title"`
	AlternativeName string `json:"alternative_name"`
	Img             string `json:"img"`
	Type            string `json:"type"`
	Plot            string `json:"plot"`
	Genre           string `json:"genre"`
	Released        string `json:"released"`
	Status          string `json:"status"`
}

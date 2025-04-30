package model

type Anime struct {
	Title string `json:"title"`
	URL   string `json:"url"`
	Id  int `json:"mal_id"`
	Images struct {
		JPG struct {
			ImageURL string `json:"image_url"`
		} `json:"jpg"`
	} `json:"images"`
}
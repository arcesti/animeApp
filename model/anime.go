package model

type Anime struct {
	Title  string `json:"title"`
	URL    string `json:"url"`
	Id     int    `json:"mal_id"`
	Images struct {
		JPG struct {
			ImageURL string `json:"image_url"`
		} `json:"jpg"`
	} `json:"images"`
}

type Personagens struct {
	Character struct {
		Id     int    `json:"mal_id"`
		URL    string `json:"url"`
		Name   string `json:"name"`
		Images struct {
			JPG struct {
				ImageURL string `json:"image_url"`
			} `json:"jpg"`
		} `json:"images"`
	} `json:"character"`
	Role string `json:"role"`
}

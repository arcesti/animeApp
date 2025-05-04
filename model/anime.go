package model

type Genre struct {
	Mal_id int    `json:"mal_id"`
	Name   string `json:"name"`
}

type Anime struct {
	Title   string  `json:"title"`
	URL     string  `json:"url"`
	Id      int     `json:"mal_id"`
	Ano     int     `json:"year"`
	Score   float64 `json:"score"`
	Synopse string  `json:"synopsis"`
	Genres  []Genre `json:"genres"`
	Trailer struct {
		Images struct {
			ImageTrailer string `json:"maximum_image_url"`
		}
	}
	Images struct {
		WEBP struct {
			ImageURL string `json:"large_image_url"`
		} `json:"webp"`
	} `json:"images"`
}

type Personagens struct {
	Character struct {
		Id     int    `json:"mal_id"`
		URL    string `json:"url"`
		Name   string `json:"name"`
		Images struct {
			WEBP struct {
				ImageURL string `json:"image_url"`
			} `json:"webp"`
		} `json:"images"`
	} `json:"character"`
	Role    string `json:"role"`
	IdAnime int    `json:"anime_mal_id"`
}

package model

type Movie struct {
	ImdbId string `json:"imdb_id"`
	Title  string `json:"title"`
	Url    string `json:"url"`
}

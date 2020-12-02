package model

type Result struct {
	Items             []SearchItem      `json:"items"`
	SearchInformation searchInformation `json:"searchInformation"`
}

type searchInformation struct {
	TotalResults string `json:"totalResults"`
}

type SearchItem struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

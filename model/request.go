package model

import (
	"net/url"
	"os"
	"strings"
)

type Request struct {
	Query          string
	ApiKey         string
	SearchEngineId string
}

func (r *Request) GenerateLink() string {
	return os.Getenv("API_URL") + "?key=" + r.ApiKey + "&cx=" + r.SearchEngineId + "&q=" + r.GetQuery()
}

func (r *Request) GetQuery() string {
	return url.QueryEscape(strings.ToLower(r.Query))
}

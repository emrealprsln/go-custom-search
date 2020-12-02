package model

import (
	"fmt"
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
	return fmt.Sprintf("%s?key=%s&cx=%s&q=%s", os.Getenv("API_URL"), r.ApiKey, r.SearchEngineId, r.GetQuery())
}

func (r *Request) GetQuery() string {
	return url.QueryEscape(strings.ToLower(r.Query))
}

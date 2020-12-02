package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/emrealprsln/go-custom-search/driver"
	"github.com/emrealprsln/go-custom-search/model"
	"github.com/emrealprsln/go-custom-search/util"
)

const (
	searchIdPrefix = "tt"
	regexPattern   = `^https?:\/\/www.imdb.com\/title\/(tt\d+)\/$`
	searchCacheKey = "search:"
)

type SearchService interface {
	SearchByName(name string) ([]model.Movie, util.RestError)
}

type searchService struct{}

func NewSearchService() SearchService {
	return &searchService{}
}

func (s searchService) SearchByName(name string) ([]model.Movie, util.RestError) {
	r := model.Request{
		Query:          name,
		ApiKey:         os.Getenv("API_KEY"),
		SearchEngineId: os.Getenv("SEARCH_ENGINE_ID"),
	}

	if cachedSearch := s.getCachedSearch(r.GetQuery()); cachedSearch != nil {
		return cachedSearch, nil
	}

	res, err := http.Get(r.GenerateLink())
	if err != nil {
		return nil, util.NewRestError(util.RequestErr, util.RequestErrMsg)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, util.NewRestError(util.InvalidParamsErr, util.InvalidParamsErrMsg)
	}

	result, resultErr := s.parseResponse(res)
	if resultErr != nil {
		return nil, util.NewRestError(util.UnknownErr, resultErr.Error())
	}
	if result.SearchInformation.TotalResults == "0" {
		return nil, util.NewRestError(util.NotFound, util.NotFoundMsg)
	}

	movies := s.processMovies(result.Items)
	if len(movies) > 0 {
		driver.SetValue(searchCacheKey+r.GetQuery(), movies)
	}
	return movies, nil
}

func (s searchService) parseResponse(r *http.Response) (*model.Result, error) {
	var result model.Result

	bytes, byteErr := ioutil.ReadAll(r.Body)
	if byteErr != nil {
		return nil, byteErr
	}
	if jsonErr := json.Unmarshal(bytes, &result); jsonErr != nil {
		return nil, jsonErr
	}
	return &result, nil
}

func (s searchService) processMovies(items []model.SearchItem) []model.Movie {
	var movies []model.Movie

	re := regexp.MustCompile(regexPattern)
	for _, value := range items {
		match := re.FindStringSubmatch(value.Link)

		if len(match) == 2 && strings.HasPrefix(match[1], searchIdPrefix) {
			movies = append(movies, model.Movie{ImdbId: match[1], Title: value.Title, Url: value.Link})
		}
	}
	return movies
}

func (s searchService) getCachedSearch(query string) []model.Movie {
	var movies []model.Movie

	result := driver.GetValue(searchCacheKey + query)
	if result == "" {
		return nil
	}
	err := json.Unmarshal([]byte(result), &movies)
	if err != nil {
		return nil
	}
	return movies
}

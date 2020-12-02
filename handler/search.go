package handler

import (
	"net/http"
	"strings"

	"github.com/emrealprsln/go-custom-search/service"
	"github.com/emrealprsln/go-custom-search/util"
	"github.com/gin-gonic/gin"
)

const (
	nameRequired = "the name field is required"
)

type SearchHandler interface {
	Search(c *gin.Context)
}

type searchHandler struct {
	searchService service.SearchService
}

func NewSearchHandler(s service.SearchService) SearchHandler {
	return &searchHandler{
		searchService: s,
	}
}

func (s searchHandler) Search(c *gin.Context) {
	name := strings.TrimSpace(c.Query("name"))
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": util.NewRestError(util.InvalidParamsErr, nameRequired)})
		return
	}

	results, err := s.searchService.SearchByName(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, results)
}

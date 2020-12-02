package main

import (
	"os"

	"github.com/emrealprsln/go-custom-search/driver"
	"github.com/emrealprsln/go-custom-search/handler"
	"github.com/emrealprsln/go-custom-search/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	rdb, rdbErr := driver.InitCache()
	if rdbErr != nil {
		panic(rdbErr)
	}
	defer rdb.Close()

	r := gin.Default()
	s := handler.NewSearchHandler(service.NewSearchService())

	r.GET("/search", s.Search)
	r.Run(":" + os.Getenv("APP_PORT"))
}

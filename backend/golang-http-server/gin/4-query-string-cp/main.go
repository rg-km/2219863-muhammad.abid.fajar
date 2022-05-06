package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	// TODO: answer here
	router.GET("/movie", func(c *gin.Context) {
		genre := c.DefaultQuery("genre", "general")
		country := c.Query("country")
		if country != "" {
			c.String(http.StatusOK, "Here result of query of movie with genre %s in country %s", genre, country)
		} else {
			c.String(http.StatusOK, "Here result of query of movie with genre %s", genre)
		}
		
	})

	return router
}

func main() {
	router := GetRouter()
	router.Run(":8080")
}

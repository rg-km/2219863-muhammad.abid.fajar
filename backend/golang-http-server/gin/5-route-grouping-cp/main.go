package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	Title string
}

var movies = map[int]Movie{
	1: {
		"Spiderman",
	},
	2: {
		"Joker",
	},
	3: {
		"Escape Plan",
	},
	4: {
		"Lord of the Rings",
	},
}

var MovieListHandler = func(c *gin.Context) {
	c.JSON(200, gin.H{}) // TODO: replace this
}

var MovieGetHandler = func(c *gin.Context) {
	no := c.Param("no")
	g, err := strconv.Atoi(no)
	if err != nil {
		return
	}
	fmt.Println(no)
	for k := 1; k <= len(movies); k++ {
		if g >= len(movies) {
			c.String(http.StatusNotFound, "data not found")
			return
		} else {
			c.String(http.StatusOK, "{\"data\":{\"Title\":\"%s\"}}", movies[g].Title)
			return
		}
	}
}

func GetRouter() *gin.Engine {
	router := gin.Default()
	// TODO: answer here
	v1 := router.Group("/movie")
	{
		v1.GET("/list", MovieListHandler)
		v1.GET("/get/:no", MovieGetHandler)
	}
	return router
}

func main() {
	router := GetRouter()

	router.Run(":8080")
}

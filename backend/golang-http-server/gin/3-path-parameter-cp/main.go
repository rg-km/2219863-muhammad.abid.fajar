package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name    string
	Country string
	Age     int
}

var data = map[int]User{
	1: {
		Name:    "Roger",
		Country: "United States",
		Age:     70,
	},
	2: {
		Name:    "Tony",
		Country: "United States",
		Age:     40,
	},
	3: {
		Name:    "Asri",
		Country: "Indonesia",
		Age:     30,
	},
}

func ProfileHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		no := c.Param("no")
		g, err := strconv.Atoi(no)
		if err != nil {
			return
		}
		fmt.Println(no)
		for k := 1; k <= len(data); k++ {
			if g >= len(data) {
				c.String(http.StatusNotFound, "data not found")
				return
			} else {
				c.String(http.StatusOK, "Name: %s, Country: %s, Age: %d", data[g].Name, data[g].Country, data[g].Age)
				return
			}
		}
	}
}

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/profile/:no", ProfileHandler())

	return r
}

func main() {
	router := GetRouter()
	router.Run(":8080")
}

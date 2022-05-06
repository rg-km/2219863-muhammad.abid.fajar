package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var RequestMethodHandler = func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": c.Request.Method,
	})
}

func GetGinRoute() *gin.Engine {
	r := gin.Default()

	r.GET("/get", RequestMethodHandler)
	r.POST("/post", RequestMethodHandler)
	r.PUT("/put", RequestMethodHandler)
	r.DELETE("/delete", RequestMethodHandler)
	r.PATCH("/patch", RequestMethodHandler)
	r.HEAD("/head", RequestMethodHandler)
	r.OPTIONS("/options", RequestMethodHandler)

	return r // TODO: replace this
}

func main() {
	r := GetGinRoute()
	r.Run("localhost:3000")
}

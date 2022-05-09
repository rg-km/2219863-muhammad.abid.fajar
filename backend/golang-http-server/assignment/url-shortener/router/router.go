package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/handlers"
)

func SetupRouter(urlHandler handlers.URLHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/nonexist", func(c *gin.Context) {
		c.String(404, "Error 404")
	})
	router.POST("/", urlHandler.Get)
	router.POST("/custom", urlHandler.CreateCustom)
	router.GET("/pawgrammers", urlHandler.Create)

	return router // TODO: replace this
}

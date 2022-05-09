package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/handlers"
)

func SetupRouter(urlHandler handlers.URLHandler) *gin.Engine {
	r := gin.Default()
	r.GET("/:path", urlHandler.Get)
	r.POST("/", urlHandler.Create)
	r.POST("/:path", urlHandler.CreateCustom)
	return r
	//return &gin.Engine{} // TODO: replace this
}

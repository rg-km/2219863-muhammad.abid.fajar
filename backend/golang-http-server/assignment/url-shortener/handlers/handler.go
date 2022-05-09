package handlers

import (
	"fmt"
	"net/http"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/entity"
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/repository"

	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	repo *repository.URLRepository
}

func NewURLHandler(repo *repository.URLRepository) URLHandler {
	return URLHandler{
		repo: repo,
	}
}

func (h *URLHandler) Get(c *gin.Context) {
	// TODO: answer here
	path := c.Param("path")
	url, err := h.repo.Get(path)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if url.LongURL == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": entity.ErrURLNotFound})
		return
	}

	c.Redirect(http.StatusFound, url.LongURL)

}

func (h *URLHandler) Create(c *gin.Context) {
	// TODO: answer here
	var url entity.URL
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	urls, err := h.repo.Create(url.LongURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"longURL":  urls.LongURL,
		"shortURL": urls.ShortURL,
	})
}

func (h *URLHandler) CreateCustom(c *gin.Context) {
	// TODO: answer here
	var url entity.URL
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	if url.LongURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": entity.ErrURLNotFound,
		})
	}
	urls, err := h.repo.CreateCustom(url.LongURL, url.ShortURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"longURL":  url.LongURL,
		"shortURL": url.ShortURL,
	})

	fmt.Println("longnya", urls.LongURL)
	fmt.Println("shortnya", urls.ShortURL)
	fmt.Println("test ::", url)
}

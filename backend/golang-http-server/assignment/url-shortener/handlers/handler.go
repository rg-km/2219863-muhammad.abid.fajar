package handlers

import (
	// "errors"
	// "net/http"

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
	var g entity.URL
	m , err := h.repo.Get(g.LongURL)
	if err != nil {
		return
	}
	h.repo.Data[m.LongURL] = m.LongURL
}

func (h *URLHandler) Create(c *gin.Context) {
	var g entity.URL
	m, err := h.repo.Create(g.LongURL)
	if err != nil {
		return
	}
	h.repo.Data[m.LongURL] = m.LongURL
	
}

func (h *URLHandler) CreateCustom(c *gin.Context) {
	// TODO: answer here
}

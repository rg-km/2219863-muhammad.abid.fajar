package repository

import (
	"sync"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/entity"
)

type URLRepository struct {
	mu   sync.Mutex
	Data map[string]string
}

func NewMapRepository() URLRepository {
	data := make(map[string]string, 0)
	return URLRepository{
		Data: data,
	}
}

func (r *URLRepository) Get(path string) (*entity.URL, error) {

	return &entity.URL{} , nil // TODO: replace this
}

func (r *URLRepository) Create(longURL string) (*entity.URL, error) {
	if longURL == "" {
		return nil, entity.ErrURLNotFound
	}
	var url *entity.URL
	url.LongURL = longURL
	url.ShortURL = entity.GetRandomShortURL(longURL)
	

	return url , nil // TODO: replace this
}

func (r *URLRepository) CreateCustom(longURL, customPath string) (*entity.URL, error) {
	if longURL == "" {
		return nil, entity.ErrCustomURLIsExists
	}


	return &entity.URL{} , nil // TODO: replace this
}

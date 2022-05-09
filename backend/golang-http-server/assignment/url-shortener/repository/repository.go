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
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.Data[path]; !ok {
		return nil, entity.ErrURLNotFound
	}
	return &entity.URL{
		ShortURL: path,
		LongURL:  r.Data[path],
	}, nil
	//&entity.URL{} , nil // TODO: replace this
}

func (r *URLRepository) Create(longURL string) (*entity.URL, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	url := &entity.URL{
		LongURL:  longURL,
		ShortURL: entity.GetRandomShortURL(longURL),
	}
	r.Data[url.ShortURL] = url.LongURL

	return url, nil
	//&entity.URL{} , nil // TODO: replace this
}

func (r *URLRepository) CreateCustom(longURL, customPath string) (*entity.URL, error) {
	//&entity.URL{} , nil // TODO: replace this
	r.mu.Lock()
	defer r.mu.Unlock()
	// if longURL == "" {
	// 	return nil, entity.ErrURLNotFound
	// }
	// url := &entity.URL{
	// 	LongURL: longURL,
	// }
	// if customPath != "" {
	// 	url.ShortURL = customPath
	// }
	// r.Data[url.ShortURL] = url.LongURL
	// return url, nil

	url := &entity.URL{
		LongURL:  longURL,
		ShortURL: customPath,
	}
	r.Data[url.ShortURL] = url.LongURL
	return url, nil
}

package service

import "url-shortener/internal/store"

type URLService struct {
	store store.URLStore
}

func NewURLService(s store.URLStore) *URLService {
	return &URLService{store: s}
}

func (s *URLService) Shorten(url string) (code string, err error) {
	return s.store.Save(url)
}

func (s *URLService) Resolve(code string) (url string, err error) {
	return s.store.Get(code)
}

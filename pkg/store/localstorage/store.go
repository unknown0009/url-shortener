package localstorage

import "github.com/fidesy/url-shortener/pkg/shortener"

type LocalStorage struct {
	originalUrls map[string]string
	shortUrls    map[string]string
}

func New() *LocalStorage {
	return &LocalStorage{
		originalUrls: make(map[string]string),
		shortUrls:    make(map[string]string),
	}
}

func (s *LocalStorage) GetOriginalUrl(short string) string {
	return s.shortUrls[short]
}

func (s *LocalStorage) CreateShortUrl(url string) string {
	shortUrl, found := s.originalUrls[url]
	if found {
		return shortUrl
	}

	short := shortener.GenerateURL()
	if _, found := s.shortUrls[short]; !found {
		s.shortUrls[short] = url
		s.originalUrls[url] = short
		return short
	}

	return s.CreateShortUrl(url)
}

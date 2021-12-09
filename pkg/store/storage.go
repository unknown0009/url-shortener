package store

import (
	"github.com/fidesy/url-shortener/pkg/store/localstorage"
	"github.com/fidesy/url-shortener/pkg/store/postgresql"
)

type Store interface {
	GetOriginalUrl(string) string
	CreateShortUrl(string) string
}

func New(database, dbUrl string) Store {
	if database == "postgresql" {
		return postgresql.New(dbUrl)
	} else {
		return localstorage.New()
	}
}

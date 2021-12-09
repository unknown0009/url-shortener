package postgresql

import (
	"log"

	"github.com/fidesy/url-shortener/pkg/shortener"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sqlx.DB
}

func New(databaseUrl string) *Postgres {
	db, err := sqlx.Open("postgres", databaseUrl)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return &Postgres{db: db}
}

func (p *Postgres) Close() {
	p.db.Close()
}

type UrlDB struct {
	Url      string
	ShortUrl string
}

func (p *Postgres) GetOriginalUrl(short string) string {
	var urls []UrlDB
	err := p.db.Select(&urls, "SELECT url FROM urls WHERE shorturl=$1", short)
	if err != nil {
		log.Panic(err)
	}

	if len(urls) == 0 {
		return ""
	}
	return urls[0].Url
}

const (
	createShortUrl = `INSERT INTO urls (url, shorturl) VALUES ($1, $2)`
)

func (p *Postgres) CreateShortUrl(url string) string {
	shortUrl := shortener.GenerateURL()

	var urls []UrlDB
	err := p.db.Select(&urls, "SELECT * FROM urls WHERE url=$1 or shorturl=$2",
		url, shortUrl)
	if err != nil {
		log.Panic(err)
	}

	if len(urls) == 0 {
		p.db.MustExec(createShortUrl, url, shortUrl)
		return shortUrl
	} else if urls[0].Url == url {
		return urls[0].ShortUrl
	} else {
		return p.CreateShortUrl(url)
	}
}

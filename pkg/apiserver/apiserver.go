package apiserver

import (
	"net/http"

	"github.com/fidesy/url-shortener/pkg/store"
	"github.com/fidesy/url-shortener/pkg/store/postgresql"
	"github.com/gorilla/mux"
)

type APIServer struct {
	config *Config
	router *mux.Router
	store  store.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	s.configureRouter()
	s.configureStore()

	switch s.store.(type) {
	case *postgresql.Postgres:
		defer s.store.(*postgresql.Postgres).Close()
	}

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/{url}", s.GetOriginalUrl()).Methods("GET")
	s.router.HandleFunc("/", s.CreateShortUrl()).Methods("POST")
}

func (s *APIServer) configureStore() {
	st := store.New(s.config.Storage, s.config.DatabaseUrl)
	s.store = st
}

func (s *APIServer) GetOriginalUrl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		url := mux.Vars(r)["url"]
		original_url := s.store.GetOriginalUrl(url)
		if original_url == "" {
			w.Write([]byte("no such url found\n"))
		} else {
			w.Write([]byte(original_url + "\n"))
		}
	}
}

func (s *APIServer) CreateShortUrl() http.HandlerFunc {
	const host = "http://127.0.0.1/"
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		url := r.FormValue("url")
		short_url := s.store.CreateShortUrl(url)
		w.Write([]byte(host + short_url + "\n"))
	}
}

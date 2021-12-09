package apiserver

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIServer_GetOriginalUrlNoFound(t *testing.T) {
	s := New(NewConfig())
	s.configureStore()

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/some", nil)

	s.GetOriginalUrl().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "no such url found")
}

func TestAPIServer_CreateShortUrl(t *testing.T) {
	s := New(NewConfig())
	s.configureStore()

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/?url=http://google.com", nil)

	s.CreateShortUrl().ServeHTTP(rec, req)

	short_url := rec.Body.String()
	assert.Equal(t, strings.HasPrefix(short_url, "http://127.0.0.1/"), true)
}

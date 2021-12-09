package localstorage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocalStorage_GetOriginalUrl(t *testing.T) {
	storage := New()
	original_url := "http://google.com"
	short_url := storage.CreateShortUrl(original_url)
	assert.Equal(t, storage.GetOriginalUrl(short_url), original_url)
}

func TestLocalStorage_CreateShortUrl(t *testing.T) {
	storage := New()
	original_url := "http://google.com"
	short_url_1 := storage.CreateShortUrl(original_url)
	short_url_2 := storage.CreateShortUrl(original_url)
	assert.Equal(t, short_url_1, short_url_2)
}

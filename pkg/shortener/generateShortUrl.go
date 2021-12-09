package shortener

import (
	"math/rand"
	"time"
)

const (
	signs  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	length = len(signs)
)

func GenerateURL() string {
	rand.Seed(time.Now().Unix())
	url := ""
	for i := 0; i < 10; i++ {
		url += string(signs[rand.Intn(length)])
	}

	return url
}

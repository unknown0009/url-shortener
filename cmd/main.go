package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"github.com/fidesy/url-shortener/pkg/apiserver"
)

func loadConfig() *apiserver.Config {
	req, err := ioutil.ReadFile("configs/config.json")
	if err != nil {
		panic(err)
	}
	var config *apiserver.Config
	err = json.Unmarshal(req, &config)
	if err != nil {
		panic(err)
	}

	return config
}

func main() {
	config := loadConfig()

	storage := flag.String("db", "", "define storage")
	flag.Parse()
	if *storage != "" {
		config.Storage = *storage
	}
	
	server := apiserver.New(config)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}

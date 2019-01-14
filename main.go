package main

import (
	"log"

	"github.com/go-resty/resty"
	"github.com/joho/godotenv"
)

func main() {
	log.Print("starting vaultier ...")

	// load configuration variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	resty.R().Get("http://httpbin.org/get")
}

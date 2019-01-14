package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-resty/resty"
	"github.com/joho/godotenv"
)

func main() {
	log.Print("starting vaultier ...")

	// load configuration variables
	e := godotenv.Load()
	if e != nil {
		log.Fatal("Error loading .env file")
	}

	// open specs file
	specsFile, e := ioutil.ReadFile(os.Getenv("PLUGIN_SECRET_SPECS_PATH"))
	if e != nil {
		log.Fatal(fmt.Sprintf("Error loading specs file:\n%s", e))
	}

	// parse JSON
	var specs Specs
	e = json.Unmarshal(specsFile, &specs)
	if e != nil {
		log.Fatal(fmt.Sprintf("Error parsing specs:\n%s", e))
	}

	fmt.Printf("%+v\n", specs)

	resty.R().Get("http://httpbin.org/get")
}

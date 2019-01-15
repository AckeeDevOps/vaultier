package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"

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

	// parse YAML
	var specs Specs
	e = yaml.Unmarshal(specsFile, &specs)
	if e != nil {
		log.Fatal(fmt.Sprintf("Error parsing specs:\n%s", e))
	}

	// validate specification
	e = specs.validate()
	if e != nil {
		log.Fatal(fmt.Sprintf("Error validating specs:\n%s", e))
	}

	fmt.Printf("%+v\n", specs)
}

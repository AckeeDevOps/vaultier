package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/vranystepan/vaultier/client"

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

	vaultAddr := os.Getenv("PLUGIN_VAULT_ADDR")
	vaultToken := os.Getenv("PLUGIN_VAULT_TOKEN")

	client := client.New(vaultAddr, vaultToken, false)

	results := []map[string]interface{}{}
	for _, branch := range specs.Specs {
		for _, secret := range branch.Secrets {
			res, err := client.Get(secret.Path, secret.KeyMap)
			if err != nil {
				log.Fatal(fmt.Sprintf("error getting secrets:\n%s", err))
			}
			results = append(results, res)
		}
	}

	final := mergeResults(results)
	log.Print(final)

}

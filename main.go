package main

import (
	"encoding/json"
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

	// get env variables
	vaultAddr := os.Getenv("PLUGIN_VAULT_ADDR")
	vaultToken := os.Getenv("PLUGIN_VAULT_TOKEN")
	currentBranch := os.Getenv("PLUGIN_BRANCH")
	cause := os.Getenv("PLUGIN_RUN_CAUSE")
	specsPath := os.Getenv("PLUGIN_SECRET_SPECS_PATH")

	// open specs file
	specsFile, e := ioutil.ReadFile(specsPath)
	if e != nil {
		log.Fatal(fmt.Sprintf("Error loading specs file:\n%s", e))
	}

	// parse YAML
	var specs Specs
	e = yaml.Unmarshal(specsFile, &specs)
	if e != nil {
		log.Fatal(fmt.Sprintf("Error parsing specs:\n%s", e))
	}

	var specsSelection []SecretPathEntry
	if cause == "delivery" {
		for _, b := range specs.Branches {
			if b.Name == currentBranch {
				specsSelection = b.Secrets
				break
			}
		}
	} else if cause == "test" {
		specsSelection = specs.TestConfig.Secrets
	} else {
		log.Fatal("unknown PLUGIN_RUN_CAUSE value")
	}

	if cap(specsSelection) == 0 {
		log.Fatal(fmt.Sprintf("%s configuration is empty", cause))
	}

	// create a new Vault client
	client := client.New(vaultAddr, vaultToken, false)

	// go through specification and push results to single map
	results := []map[string]interface{}{}
	for _, secret := range specsSelection {
		res, err := client.Get(secret.Path, secret.KeyMap)
		if err != nil {
			log.Fatal(fmt.Sprintf("error getting secrets:\n%s", err))
		}
		results = append(results, res)
	}

	final := mergeResults(results)
	finalJSON, err := json.Marshal(final)
	if err != nil {
		log.Fatal("failed to marshal final results")
	}

	log.Print(string(finalJSON))

	// push secrets back to JSON

}

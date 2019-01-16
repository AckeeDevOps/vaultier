package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"

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

	spew.Dump(specs)

	// validate specification
	//e = specs.validate()
	//if e != nil {
	//	log.Fatal(fmt.Sprintf("Error validating specs:\n%s", e))
	//}

	vaultAddr := os.Getenv("PLUGIN_VAULT_ADDR")
	vaultToken := os.Getenv("PLUGIN_VAULT_TOKEN")
	currentBranch := os.Getenv("PLUGIN_BRANCH")

	client := client.New(vaultAddr, vaultToken, false)

	// go through specification and push results to single map
	results := []map[string]interface{}{}
	for _, branch := range specs.Branches {
		if branch.Name == currentBranch {
			for _, secret := range branch.Secrets {
				res, err := client.Get(secret.Path, secret.KeyMap)
				if err != nil {
					log.Fatal(fmt.Sprintf("error getting secrets:\n%s", err))
				}
				results = append(results, res)
			}
		} else {
			log.Print(fmt.Sprintf("skipping branch '%s'", branch.Name))
		}
	}

	final := mergeResults(results)
	finalJSON, err := json.Marshal(final)
	if err != nil {
		log.Fatal("failed to marshal final results")
	}

	log.Print(string(finalJSON))

	// push secrets back to JSON

}

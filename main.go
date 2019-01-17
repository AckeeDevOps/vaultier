package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/vranystepan/vaultier/client"
	"github.com/vranystepan/vaultier/config"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	log.Print("starting vaultier ...")

	// get and validate config
	cfg := config.Create()
	err := cfg.Validate()
	if err != nil {
		log.Fatal(err)
	}

	// open specs file
	specsFile, e := ioutil.ReadFile(cfg.SpecsPath)
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
	if cfg.Cause == "delivery" {
		for _, b := range specs.Branches {
			if b.Name == cfg.Branch {
				specsSelection = b.Secrets
				break
			}
		}
	} else if cfg.Cause == "test" {
		specsSelection = specs.TestConfig.Secrets
	} else {
		log.Fatal("unknown PLUGIN_RUN_CAUSE value")
	}

	if cap(specsSelection) == 0 {
		log.Fatal(fmt.Sprintf("%s configuration is empty", cfg.Cause))
	}

	// create a new Vault client
	final := collectSecrets(specsSelection, cfg.VaultAddr, cfg.VaultToken, false)

	finalJSON, err := json.Marshal(final)
	if err != nil {
		log.Fatal("failed to marshal final results")
	}

	log.Print(string(finalJSON))

	// push secrets back to JSON

}

func collectSecrets(secrets []SecretPathEntry, vaultAddr string, vaultToken string, insecure bool) map[string]interface{} {
	client := client.New(vaultAddr, vaultToken, insecure)
	results := []map[string]interface{}{}

	for _, secret := range secrets {
		res, err := client.Get(secret.Path, secret.KeyMap)
		if err != nil {
			log.Fatal(fmt.Sprintf("error getting secrets:\n%s", err))
		}
		results = append(results, res)
	}

	return mergeResults(results)
}

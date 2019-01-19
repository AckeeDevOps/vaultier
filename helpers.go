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

type helmManifestFotmat struct {
	Secrets map[string]string `json:"secrets"`
}

// merge multiple results into single map
func mergeResults(maps []map[string]string) map[string]string {
	result := map[string]string{}
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

// get configuration from the specs file based on PLUGIN_RUN_CAUSE
func getSelection(s Specs, c *config.PluginConfig) []SecretPathEntry {
	var specsSelection []SecretPathEntry
	if c.Cause == "delivery" {
		for _, b := range s.Branches {
			if b.Name == c.Branch {
				specsSelection = b.Secrets
				break
			}
		}
	} else if c.Cause == "test" {
		specsSelection = s.TestConfig.Secrets
	} else {
		log.Fatal("unknown PLUGIN_RUN_CAUSE value")
	}

	if cap(specsSelection) == 0 {
		log.Fatal(fmt.Sprintf("%s configuration is empty", c.Cause))
	}

	return specsSelection
}

// parse provided configuraton
func getConfig() *config.PluginConfig {
	cfg := config.Create()
	err := cfg.Validate()
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}

// read specs file
func getSpecs(c *config.PluginConfig) Specs {
	log.Printf("getting secrets configuration from %s", c.SpecsPath)
	// open specs file
	specsFile, e := ioutil.ReadFile(c.SpecsPath)
	if e != nil {
		log.Fatal(fmt.Sprintf("Error loading specs file:\n%s", e))
	}

	// parse YAML
	var specs Specs
	e = yaml.Unmarshal(specsFile, &specs)
	if e != nil {
		log.Fatal(fmt.Sprintf("Error parsing specs:\n%s", e))
	}

	return specs
}

// generate secrets manifest in the requested format
func generateManifest(c *config.PluginConfig, s map[string]string) []byte {
	var finalObj interface{}
	if c.Cause == "delivery" {
		finalObj = helmManifestFotmat{
			Secrets: s,
		}
	} else {
		finalObj = s
	}

	finalJSON, err := json.Marshal(finalObj)
	if err != nil {
		log.Fatal("failed to marshal final results")
	}

	return finalJSON
}

// go through specs and call vault client
func collectSecrets(secrets []SecretPathEntry, vaultAddr string, vaultToken string, insecure bool) map[string]string {
	client := client.New(vaultAddr, vaultToken, insecure)
	results := []map[string]string{}

	for _, secret := range secrets {
		res, err := client.Get(secret.Path, secret.KeyMap)
		log.Printf("Getting secrets from %s", secret.Path)
		if err != nil {
			log.Fatal(fmt.Sprintf("error getting secrets:\n%s", err))
		}
		results = append(results, res)
	}

	return mergeResults(results)
}

// write results to the file
func writeFile(c *config.PluginConfig, s []byte) {
	err := ioutil.WriteFile(c.OutputPath, s, 0644)
	if err != nil {
		log.Fatalf("could not create output file %s", c.OutputPath)
	}

	log.Printf("data successfully written to %s", c.OutputPath)
}

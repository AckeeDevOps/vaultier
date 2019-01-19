package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/vranystepan/vaultier/config"
	yaml "gopkg.in/yaml.v2"
)

// merge multiple results into single map
func mergeResults(maps []map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

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

func getConfig() *config.PluginConfig {
	cfg := config.Create()
	err := cfg.Validate()
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}

func getSpecs(c *config.PluginConfig) Specs {
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

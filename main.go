package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/vranystepan/vaultier/client"
)

func main() {
	log.Print("starting vaultier ...")

	// get and validate config
	cfg := getConfig()

	// get secrets specification
	var specs = getSpecs(cfg)

	// select current config
	var specsSelection = getSelection(specs, cfg)

	// collect secrets from Vault
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

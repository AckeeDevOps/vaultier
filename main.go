package main

import (
	"log"
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

	// prepare helm or .env structure
	finalJSON := generateManifest(cfg, final)

	log.Print(string(finalJSON))

	// push secrets back to JSON

}

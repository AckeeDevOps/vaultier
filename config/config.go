package config

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

type PluginConfig struct {
	VaultAddr  string
	VaultToken string
	Cause      string
	Branch     string
	SpecsPath  string
	OutputPath string
}

func Create() *PluginConfig {
	p := PluginConfig{}

	vaultAddr := os.Getenv("PLUGIN_VAULT_ADDR")          // required
	vaultToken := os.Getenv("PLUGIN_VAULT_TOKEN")        // required
	currentBranch := os.Getenv("PLUGIN_BRANCH")          // required
	cause := os.Getenv("PLUGIN_RUN_CAUSE")               // optional, default=delivery
	specsPath := os.Getenv("PLUGIN_SECRET_SPECS_PATH")   // optional, default=./secrets.yaml
	outputPath := os.Getenv("PLUGIN_SECRET_OUTPUT_PATH") // required

	p.VaultAddr = strings.ToLower(vaultAddr)
	p.Branch = strings.ToLower(currentBranch)
	p.VaultToken = vaultToken
	p.Cause = strings.ToLower(cause)
	p.SpecsPath = specsPath
	p.OutputPath = outputPath

	return &p
}

func (c *PluginConfig) Validate() error {
	errors := []string{}

	// validate Vault URL
	_, err := url.ParseRequestURI(c.VaultAddr)
	if err != nil {
		errors = append(errors, "invalid Vault address")
	}

	// validate token
	if c.VaultToken == "" {
		errors = append(errors, "empty Vault token")
	}

	// validate branch
	if c.Branch == "" {
		errors = append(errors, "empty branch name")
	}

	// validate branch
	if c.OutputPath == "" {
		errors = append(errors, "empty output path")
	}

	// validate run cause
	if c.Cause != "delivery" && c.Cause != "test" {
		log.Print("using default run cause: delivery")
		c.Cause = "delivery"
	}

	// validate path
	if c.SpecsPath == "" {
		log.Print("using default specs path: secrets.yaml")
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal("could not get current working directory")
		}
		c.SpecsPath = fmt.Sprintf("%s/secrets.yaml", cwd)
	}

	if cap(errors) != 0 {
		return fmt.Errorf("Validation failed: %s", strings.Join(errors[:], ", "))
	}

	return nil
}

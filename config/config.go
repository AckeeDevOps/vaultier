package config

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

// PluginConfig contains configuration options received from the environment variables
type PluginConfig struct {
	VaultAddr    string
	VaultToken   string
	Cause        string
	Branch       string
	SpecsPath    string
	OutputPath   string
	OutputFormat string
}

// Create creates a new object with the actual configuration options
func Create() *PluginConfig {
	p := PluginConfig{}

	vaultAddr := os.Getenv("VAULTIER_VAULT_ADDR")          // required
	vaultToken := os.Getenv("VAULTIER_VAULT_TOKEN")        // required
	currentBranch := os.Getenv("VAULTIER_BRANCH")          // required
	cause := os.Getenv("VAULTIER_RUN_CAUSE")               // optional, default=delivery
	outputFormat := os.Getenv("VAULTIER_OUTPUT_FORMAT")    // optional, default=delivery
	specsPath := os.Getenv("VAULTIER_SECRET_SPECS_PATH")   // optional, default=./secrets.yaml
	outputPath := os.Getenv("VAULTIER_SECRET_OUTPUT_PATH") // required

	p.VaultAddr = strings.ToLower(vaultAddr)
	p.Branch = strings.ToLower(currentBranch)
	p.VaultToken = vaultToken
	p.Cause = strings.ToLower(cause)
	p.SpecsPath = specsPath
	p.OutputPath = outputPath
	p.OutputFormat = outputFormat

	return &p
}

// Validate does a basic validation of values received from the environment
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

	// validate output format
	if c.OutputFormat != "helm" && c.OutputFormat != "dotenv" {
		errors = append(errors, "empty output format")
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

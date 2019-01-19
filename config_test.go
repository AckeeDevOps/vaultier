package main

import (
	"os"
	"testing"

	"github.com/vranystepan/vaultier/config"
)

func TestValidConfiguration(t *testing.T) {
	os.Setenv("PLUGIN_VAULT_ADDR", "http://vault.co.uk")
	os.Setenv("PLUGIN_VAULT_TOKEN", "abcdefg")
	os.Setenv("PLUGIN_BRANCH", "master")
	os.Setenv("PLUGIN_RUN_CAUSE", "delivery")
	os.Setenv("PLUGIN_OUTPUT_FORMAT", "helm")
	os.Setenv("PLUGIN_SECRET_SPECS_PATH", "/tmp/input")
	os.Setenv("PLUGIN_SECRET_OUTPUT_PATH", "/tmp/output")

	cfg := config.Create()
	err := cfg.Validate()
	if err != nil {
		t.Errorf("Validator should be happy but it return %s", err)
	}
}

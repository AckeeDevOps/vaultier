package config

import (
	"os"
	"testing"
)

func setEnv() {
	os.Setenv("PLUGIN_VAULT_ADDR", "http://vault.co.uk")
	os.Setenv("PLUGIN_VAULT_TOKEN", "abcdefg")
	os.Setenv("PLUGIN_BRANCH", "master")
	os.Setenv("PLUGIN_RUN_CAUSE", "delivery")
	os.Setenv("PLUGIN_OUTPUT_FORMAT", "helm")
	os.Setenv("PLUGIN_SECRET_SPECS_PATH", "/tmp/input")
	os.Setenv("PLUGIN_SECRET_OUTPUT_PATH", "/tmp/output")
}

func TestValidConfiguration(t *testing.T) {
	setEnv()
	cfg := Create()
	err := cfg.Validate()
	if err != nil {
		t.Errorf("Validator should be happy but it return %s", err)
	}
}

func TestValidConfigurationWithoutInput(t *testing.T) {
	setEnv()
	os.Setenv("PLUGIN_SECRET_SPECS_PATH", "")

	cfg := Create()
	err := cfg.Validate()
	if err != nil {
		t.Errorf("Validator should be happy but it return %s", err)
	}
}

func TestInvalidConfigurationWithoutAddress(t *testing.T) {
	setEnv()
	os.Setenv("PLUGIN_VAULT_ADDR", "")

	cfg := Create()
	err := cfg.Validate()
	if err == nil {
		t.Errorf("Validator should return error but it does not")
	}
}

func TestInvalidConfigurationWithoutToken(t *testing.T) {
	setEnv()
	os.Setenv("PLUGIN_VAULT_TOKEN", "") // <= here

	cfg := Create()
	err := cfg.Validate()
	if err == nil {
		t.Errorf("Validator should return error but it does not")
	}
}

func TestInvalidConfigurationWithoutBranch(t *testing.T) {
	setEnv()
	os.Setenv("PLUGIN_BRANCH", "") // <= here

	cfg := Create()
	err := cfg.Validate()
	if err == nil {
		t.Errorf("Validator should return error but it does not")
	}
}

func TestInvalidConfigurationWithoutRunCause(t *testing.T) {
	setEnv()
	os.Setenv("PLUGIN_RUN_CAUSE", "") // <= here

	cfg := Create()
	err := cfg.Validate()
	if err != nil {
		t.Errorf("Validator should be happy but it return %s", err)
	}
}

func TestInvalidConfigurationWithoutOutputFormat(t *testing.T) {
	setEnv()
	os.Setenv("PLUGIN_OUTPUT_FORMAT", "") // <= here

	cfg := Create()
	err := cfg.Validate()
	if err == nil {
		t.Errorf("Validator should return error but it does not")
	}
}

func TestInvalidConfigurationWithoutOutputPath(t *testing.T) {
	setEnv()
	os.Setenv("PLUGIN_SECRET_OUTPUT_PATH", "") // <= here

	cfg := Create()
	err := cfg.Validate()
	if err == nil {
		t.Errorf("Validator should return error but it does not")
	}
}

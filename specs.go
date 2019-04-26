package main

import (
	client "github.com/AckeeDevOps/vaultier/client"
)

/*
SecretPathEntry represents Vault path e.g. secret/data/keys
KeyMap allows to retreive multiple keys from the single path
*/
type SecretPathEntry struct {
	Path   string                     `yaml:"path"`
	KeyMap []client.SecretKeyMapEntry `yaml:"keyMap"`
}

/*
Branch represents configuration for the certain git branch
*/
type Environment struct {
	Name    string            `yaml:"name"`
	Secrets []SecretPathEntry `yaml:"secrets"`
}

/*
TestConfig is related to CI testing activities e.g.
secrets for database for integration testing
*/
type TestConfig struct {
	Secrets []SecretPathEntry `yaml:"secrets"`
}

/*
Specs is representation of input YAML specification
*/
type Specs struct {
	Environments []Environment `yaml:"environments"`
}

// perform validation of required fields
//func (s Specs) validate() error {
//}

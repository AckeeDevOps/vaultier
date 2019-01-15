package main

import (
	"fmt"
	"strings"
)

/*
SecretKeyMapEntry represents details about one KEY:VALUE secret pair
VaultKey represent name of the key under data are stored in Vault
LocalKey represents name of the local key e.g. KEY becomes results['KEY']
*/
type SecretKeyMapEntry struct {
	VaultKey string `yaml:"vaultKey"`
	LocalKey string `yaml:"localKey"`
}

/*
SecretPathEntry represents Vault path e.g. secret/data/keys
KeyMap allows to retreive multiple keys from the single path
*/
type SecretPathEntry struct {
	Path   string              `yaml:"path"`   // required
	KeyMap []SecretKeyMapEntry `yaml:"keymap"` // required
}

/*
SpecsEntry represents configuration for the certain git branch
*/
type SpecsEntry struct {
	Branch  string            `yaml:"branch"`  // required
	Secrets []SecretPathEntry `yaml:"secrets"` // required
}

/*
Specs is representation of input YAML specification
example:
---
vaultAddr: ''
token: ''
specs:
- branch: master
  secrets:
  - path: secret/data/test
    keyMap:
      - vaultKey: key
        localKey: KEY
*/
type Specs struct {
	Specs []SpecsEntry `yaml:"specs"` // required
}

func (s Specs) validate() error {
	var errors []string

	/*
		Specs can't be emty, if you don't need any secrets -
		just don't use this plugin.
	*/
	if cap(s.Specs) == 0 {
		errors = append(errors, "specs can't be empty")
	}

	// validate branches
	for _, branch := range s.Specs {

		// check branch name
		if branch.Branch == "" {
			errors = append(errors, "specs/[].branch can't be empty")
		}

		// check if secrets specification exist
		if cap(branch.Secrets) == 0 {
			errors = append(errors, "specs/[].secrets can't be empty")
		}
	}

	if cap(errors) > 0 {
		return fmt.Errorf(strings.Join(errors[:], "\n"))
	}

	// everything's alright
	return nil
}

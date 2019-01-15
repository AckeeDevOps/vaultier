package main

import (
	validator "gopkg.in/go-playground/validator.v9"
)

/*
SecretKeyMapEntry represents details about one KEY:VALUE secret pair
VaultKey represent name of the key under data are stored in Vault
LocalKey represents name of the local key e.g. KEY becomes results['KEY']
*/
type SecretKeyMapEntry struct {
	VaultKey string `yaml:"vaultKey" validate:"required"`
	LocalKey string `yaml:"localKey" validate:"required"`
}

/*
SecretPathEntry represents Vault path e.g. secret/data/keys
KeyMap allows to retreive multiple keys from the single path
*/
type SecretPathEntry struct {
	Path   string              `yaml:"path" validate:"required"`                 // required
	KeyMap []SecretKeyMapEntry `yaml:"keyMap" validate:"required,dive,required"` // required
}

/*
SpecsEntry represents configuration for the certain git branch
*/
type SpecsEntry struct {
	Branch  string             `yaml:"branch" validate:"required"`
	Secrets []*SecretPathEntry `yaml:"secrets" validate:"required,dive,required"`
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
	Specs []SpecsEntry `yaml:"specs" validate:"required,dive,required"`
}

// perform validation of required fields
func (s Specs) validate() error {
	validate := validator.New()
	err := validate.Struct(s)
	return err
}

package main

import (
	client "github.com/vranystepan/vaultier/client"
	validator "gopkg.in/go-playground/validator.v9"
)

/*
SecretPathEntry represents Vault path e.g. secret/data/keys
KeyMap allows to retreive multiple keys from the single path
*/
type SecretPathEntry struct {
	Path   string                     `yaml:"path" validate:"required"`
	KeyMap []client.SecretKeyMapEntry `yaml:"keyMap" validate:"required,dive,required"`
}

/*
SpecsEntry represents configuration for the certain git branch
*/
type SpecsEntry struct {
	Branch  string            `yaml:"branch" validate:"required"`
	Secrets []SecretPathEntry `yaml:"secrets" validate:"required,dive,required"`
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
	return validate.Struct(s)
}

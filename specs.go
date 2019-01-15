package main

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
	Path   string              `yaml:"path"`
	KeyMap []SecretKeyMapEntry `yaml:"keymap"`
}

/*
SpecsEntry represents configuration for the certain git branch
*/
type SpecsEntry struct {
	Branch  string            `yaml:"branch"`
	Secrets []SecretPathEntry `yaml:"secrets"`
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
	VaultAddr string       `yaml:"vaultAddr"` // optional
	Token     string       `yaml:"token"`     // optional, don't do that
	Specs     []SpecsEntry `yaml:"specs"`     // required
}

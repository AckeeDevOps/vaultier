package main

type secretKeyMapEntry struct {
	VaultKey string `yaml:"vaultKey"`
	LocalKey string `yaml:"localKey"`
}

type secretPapthEntry struct {
	Path   string              `yaml:"path"`
	KeyMap []secretKeyMapEntry `yaml:"keymap"`
}

type specsEntry struct {
	Branch  string             `yaml:"branch"`
	Secrets []secretPapthEntry `yaml:"secrets"`
}

// Specs is representation of input YAML specification
/*
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
	Specs     []specsEntry `yaml:"specs"`     // required
}

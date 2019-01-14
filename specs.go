package main

type secretPapthEntry struct {
	Path     string `json:"path"`
	VaultKey string `json:"vaultKey"`
	LocalKey string `json:"localKey"`
}

type specsEntry struct {
	Branch  string             `json:"branch"`
	Secrets []secretPapthEntry `json:"secrets"`
}

type specs struct {
	VaultAddr string       `json:"vaultAddr"`
	Token     string       `json:"token"` // don't do that
	Specs     []specsEntry `json:"specs"`
}

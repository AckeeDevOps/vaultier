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

// Specs is representation of input JSON specification
/*
example:
{
	"vaultAddr": "https://vault.co.uk/",
	"token": "1234567",
	"specs": [
		{
			"branch": "master",
			"secrets": [
				{"path": "secret/data/key", vaultKey: "key", "localKey": "KEY"}
			]
		}
	]
}
*/
type Specs struct {
	VaultAddr string       `json:"vaultAddr"` // optional
	Token     string       `json:"token"`     // optional, don't do that
	Specs     []specsEntry `json:"specs"`     // required
}

package client

// VaultResponseData represents the actual secrets stored in Vault
type VaultResponseData struct {
	Data map[string]interface{} `json:"data"`
}

// VaultResponseMetadata represens secret's metadata
type VaultResponseMetadata struct {
	CreatedTime  string `json:"created_time"`
	DeletionTime string `json:"deletion_time"`
	Destroyed    bool   `json:"destroyed"`
	Version      int    `json:"version"`
}

// VaultResponse represents the actual Vault response
type VaultResponse struct {
	RequestID     string            `json:"request_id"`
	LeaseID       string            `json:"lease_id"`
	Renewable     bool              `json:"renewable"`
	LeaseDuration int               `json:"lease_duration"`
	Data          VaultResponseData `json:"data"`
	Errors        []string          `json:"errors"`
}

package client

type VaultResponseData struct {
	Data map[string]string `json:"data"`
}

type VaultResponseMetadata struct {
	CreatedTime  string `json:"created_time"`
	DeletionTime string `json:"deletion_time"`
	Destroyed    bool   `json:"destroyed"`
	Version      int    `json:"version"`
}

type VaultResponse struct {
	RequestID     string            `json:"request_id"`
	LeaseID       string            `json:"lease_id"`
	Renewable     bool              `json:"renewable"`
	LeaseDuration int               `json:"lease_duration"`
	Data          VaultResponseData `json:"data"`
}

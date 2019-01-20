package client

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty"
)

// VaultFetcher implements VaultResponseFetcher
type VaultFetcher struct{}

// Fetch wrapt resty GET request
func (VaultFetcher) Fetch(token string, url string) (*VaultResponse, error) {
	resp, err := resty.R().
		SetHeader("Accept", "application/json").
		SetHeader("X-Vault-Token", token).
		Get(url)
	if err != nil {
		return nil, fmt.Errorf("Could not get secrets from vault:\n%s", err)
	}

	respJSON := VaultResponse{}
	err = json.Unmarshal(resp.Body(), &respJSON)
	if err != nil {
		return nil, fmt.Errorf("Could not unmarshal JSON\n%s", err)
	}
	return &respJSON, nil
}

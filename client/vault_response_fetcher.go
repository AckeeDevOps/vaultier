package client

// VaultResponseFetcher is an interface which wraps http library
type VaultResponseFetcher interface {
	Fetch(token string, url string) (*VaultResponse, error)
}

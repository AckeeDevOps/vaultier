package client

import (
	"crypto/tls"
	"fmt"
	"net/url"

	"github.com/go-resty/resty"
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

// Client communicates with Vault
type Client struct {
	addr     string
	token    string
	insecure bool
}

// New creates a new Client
func New(addr string, token string, insecure bool) *Client {
	return &Client{
		addr:     addr,
		token:    token,
		insecure: insecure,
	}
}

// Get gets required
func (c Client) Get(path string, keyMap []SecretKeyMapEntry) (map[string]string, error) {
	// prepare output map
	secrets := make(map[string]string)

	// parse url
	u, err := url.Parse(c.addr)
	if err != nil {
		return nil, fmt.Errorf("could not parse vault addr")
	}

	// get base url
	scheme := u.Scheme
	host := u.Host

	// assemble the base url
	baseURL := fmt.Sprintf("%s://%s/v1/", scheme, host)

	// configure resty
	resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: c.insecure})

	// get secrets
	resp, err := resty.R().
		SetHeader("Accept", "application/json").
		SetHeader("X-Vault-Token", c.token).
		Get(baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("Could not get secrets from vault:\n%s", err)
	}

	// process results
	fmt.Println(baseURL + path)
	fmt.Println(string(resp.Body()))

	return secrets, nil
}

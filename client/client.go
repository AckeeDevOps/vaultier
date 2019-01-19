package client

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/go-resty/resty"
)

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
	secrets := map[string]string{}

	// sanitize path
	path = strings.Trim(path, " ")
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}

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
	respJSON := VaultResponse{}
	json.Unmarshal(resp.Body(), &respJSON)

	// validate length of response object
	if cap(keyMap) != 0 && len(respJSON.Data.Data) == 0 {
		return nil, fmt.Errorf("response does not contain any secrets")
	}

	for _, m := range keyMap {
		secrets[m.LocalKey] = respJSON.Data.Data[m.VaultKey]
	}

	return secrets, nil
}

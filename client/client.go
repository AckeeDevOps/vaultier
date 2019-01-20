package client

import (
	"crypto/tls"
	"fmt"
	"log"
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
func (c Client) Get(path string, keyMap []SecretKeyMapEntry, fetcher VaultResponseFetcher) (map[string]string, error) {
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

	// process results
	respJSON, err := fetcher.Fetch(c.token, (baseURL + path))
	if err != nil {
		return nil, fmt.Errorf("could not ret vault response: %s", err)
	}

	// validate length of response object
	if cap(keyMap) != 0 && len(respJSON.Data.Data) == 0 {
		if cap(respJSON.Errors) > 0 {
			log.Printf("response contains following errors: %s", strings.Join(respJSON.Errors[:], "; "))
			return nil, fmt.Errorf("response contains errors")
		}
		log.Printf("response from %s is empty", path)
		return nil, fmt.Errorf("response does not contain any secrets nor errors")
	}

	for _, m := range keyMap {
		secrets[m.LocalKey] = respJSON.Data.Data[m.VaultKey]
	}

	return secrets, nil
}

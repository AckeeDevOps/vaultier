package client

import "testing"

type mockFetcherSuccess struct{}
type mockFetcherFailure struct{}
type mockFetcherEmpty struct{}

func (mockFetcherSuccess) Fetch(token string, url string) (*VaultResponse, error) {
	r := VaultResponse{
		RequestID:     "id",
		LeaseID:       "lease id",
		Renewable:     true,
		LeaseDuration: 1,
		Data: VaultResponseData{
			Data: map[string]string{
				"vaultVar1": "secret1",
				"vaultVar2": "secret2",
			},
		},
	}
	return &r, nil
}

func (mockFetcherFailure) Fetch(token string, url string) (*VaultResponse, error) {
	r := VaultResponse{
		Errors: []string{
			"this is utterly wrong",
			"you should not be here",
		},
	}
	return &r, nil
}

func (mockFetcherEmpty) Fetch(token string, url string) (*VaultResponse, error) {
	r := VaultResponse{}
	return &r, nil
}

var c = Client{
	addr:     "https://vault.co.uk",
	token:    "123456789",
	insecure: false,
}

var keyMap = []SecretKeyMapEntry{
	SecretKeyMapEntry{
		LocalKey: "localVar1",
		VaultKey: "vaultVar1",
	},
	SecretKeyMapEntry{
		LocalKey: "localVar2",
		VaultKey: "vaultVar2",
	},
}

func TestValidResponse(t *testing.T) {
	r, err := c.Get("/path/to/secrets", keyMap, mockFetcherSuccess{})

	if err != nil {
		t.Errorf("client returns error even with valid Vault reponse")
	}

	if r["localVar1"] != "secret1" {
		t.Errorf("client does not return expected remapped data, it returns %s", r)
	}

	if r["localVar2"] != "secret2" {
		t.Errorf("client does not return expected remapped data, it returns %s", r)
	}

}

func TestErrorResponse(t *testing.T) {
	_, err := c.Get("/path/to/secrets", keyMap, mockFetcherFailure{})

	if err == nil {
		t.Errorf("client should return error but it does not")
	}
}

func TestEmptyResponse(t *testing.T) {
	_, err := c.Get("/path/to/secrets", keyMap, mockFetcherEmpty{})

	if err == nil {
		t.Errorf("client should return error but it does not")
	}
}

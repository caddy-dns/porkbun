package porkbun

import (
	"fmt"
	"testing"

	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	porkbun "github.com/libdns/porkbun"
)

func TestUnmarshalCaddyFile(t *testing.T) {
	tests := []string{
		`porkbun {
			api_key thekey
			api_secret_key itsasecret
		}`}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			// given
			dispenser := caddyfile.NewTestDispenser(tc)
			p := Provider{&porkbun.Provider{}}
			// when
			err := p.UnmarshalCaddyfile(dispenser)
			// then
			if err != nil {
				t.Errorf("UnmarshalCaddyfile failed with %v", err)
				return
			}

			expectedAPIKey := "thekey"
			actualAPIKey := p.Provider.APIKey
			if expectedAPIKey != actualAPIKey {
				t.Errorf("Expected APIKey to be '%s' but got '%s'", expectedAPIKey, actualAPIKey)
			}

			expectedAPISecretKey := "itsasecret"
			actualApiSecretKey := p.Provider.APISecretKey
			if expectedAPISecretKey != actualApiSecretKey {
				t.Errorf("Expected ApiSecretKey to be '%s' but got '%s'", expectedAPISecretKey, actualApiSecretKey)
			}
		})
	}
}

package porkbun

import (
	"fmt"
	"testing"

	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	porkbun "github.com/libdns/porkbun"
)

func TestUnmarshalCaddyFile(t *testing.T) {
	fmt.Println("Testing valid config parses")

	config := `porkbun {
			api_key thekey
			api_secret_key itsasecret
		}`

	// given
	dispenser := caddyfile.NewTestDispenser(config)
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
}

func TestEmptyConfig(t *testing.T) {
	fmt.Println("Testing empty config fails to parse... ")
	config := "porkbun"

	dispenser := caddyfile.NewTestDispenser(config)
	p := Provider{&porkbun.Provider{}}

	err := p.UnmarshalCaddyfile(dispenser)
	if err == nil {
		t.Errorf(
			"UnmarshalCaddyfile should have provided an error, but none was received. api_token = %s, api_secret_key = %s",
			p.Provider.APIKey,
			p.Provider.APISecretKey,
		)
	}
}

func TestPartialConfig(t *testing.T) {
	fmt.Println("Testing partial config fails to parse... ")
	config := `
	porkbun {
		api_secret_key itsasecret
	}`

	dispenser := caddyfile.NewTestDispenser(config)
	p := Provider{&porkbun.Provider{}}

	err := p.UnmarshalCaddyfile(dispenser)
	if err == nil {
		t.Errorf(
			"UnmarshalCaddyfile should have provided an error, but none was received. api_token = %s, api_secret_key = %s",
			p.Provider.APIKey,
			p.Provider.APISecretKey,
		)
	}
}

func TestTooManyArgs(t *testing.T) {
	fmt.Println("Testing too many args... ")
	config := `porkbun {
			api_key thekey
			api_secret_key itsasecret
			something_else fail
		}`

	dispenser := caddyfile.NewTestDispenser(config)
	p := Provider{&porkbun.Provider{}}

	err := p.UnmarshalCaddyfile(dispenser)
	if err == nil {
		t.Errorf(
			"UnmarshalCaddyfile should have provided an error, but none was received. api_token = %s, api_secret_key = %s",
			p.Provider.APIKey,
			p.Provider.APISecretKey,
		)
	}
}

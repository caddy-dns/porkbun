package porkbun

import (
	"strings"
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	porkbun "github.com/libdns/porkbun"
)

// Provider lets Caddy read and manipulate DNS records hosted by this DNS provider.
type Provider struct{ *porkbun.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.porkbun",
		New: func() caddy.Module { return &Provider{new(porkbun.Provider)} },
	}
}

// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	p.Provider.APIKey = strings.TrimSpace(caddy.NewReplacer().ReplaceAll(p.Provider.APIKey, ""))
	p.Provider.APISecretKey = strings.TrimSpace(caddy.NewReplacer().ReplaceAll(p.Provider.APISecretKey, ""))
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
// porkbun {
//     api_key <api_key>
//	   api_secret_key <api_secret_key>
// }
//
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "api_key":
				if d.NextArg() {
					p.Provider.APIKey = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "api_secret_key":
				if d.NextArg() {
					p.Provider.APISecretKey = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.ArgErr()
			}
		}
	}
	if p.Provider.APIKey == "" {
		return d.Err("No api_key set")
	}
	if p.Provider.APISecretKey == "" {
		return d.Err("No api_secret_key set")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)

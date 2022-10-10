
Porkbun module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with Porkbun.

## Caddy module name

```
dns.providers.porkbun
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "porkbun",
				"api_key": "{env.PORKBUN_API_KEY}",
        		"api_secret_key": "{env.PORKBUN_API_SECRET_KEY}"
			}
		}
	}
}
```

or with the Caddyfile:

```
# globally
{
	acme_dns porkbun {
			"api_key": "{env.PORKBUN_API_KEY}",
			"api_secret_key": "{env.PORKBUN_API_SECRET_KEY}"
	}
}
```

```
# one site
tls {
	dns porkbun {
			"api_key": "{env.PORKBUN_API_KEY}",
			"api_secret_key": "{env.PORKBUN_API_PASSWORD}"
	}
}
```

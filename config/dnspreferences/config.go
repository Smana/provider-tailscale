package dnspreferences

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tailscale_dns_preferences", func(r *config.Resource) {
		r.ShortGroup = "dns"
		r.Kind = "Preferences"
	})
}

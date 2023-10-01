package tailnetkey

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tailscale_tailnet_key", func(r *config.Resource) {
		r.ShortGroup = "tailnet"
		r.Kind = "Key"
	})
}

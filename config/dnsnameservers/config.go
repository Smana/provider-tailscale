package dnsnameservers

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tailscale_dns_search_paths", func(r *config.Resource) {
		r.ShortGroup = "dns"
		r.Kind = "SearchPaths"
	})
}

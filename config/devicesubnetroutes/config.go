package devicesubnetroutes

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tailscale_device_subnet_routes", func(r *config.Resource) {
		r.ShortGroup = "device"
		r.Kind = "SubnetRoutes"
	})
}

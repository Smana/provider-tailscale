/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/Smana/provider-tailscale/config/acl"
	"github.com/Smana/provider-tailscale/config/deviceauthorization"
	"github.com/Smana/provider-tailscale/config/devicekey"
	"github.com/Smana/provider-tailscale/config/devicesubnetroutes"
	"github.com/Smana/provider-tailscale/config/devicetags"
	"github.com/Smana/provider-tailscale/config/dnsnameservers"
	"github.com/Smana/provider-tailscale/config/dnspreferences"
	"github.com/Smana/provider-tailscale/config/dnssearchpaths"
	"github.com/Smana/provider-tailscale/config/tailnetkey"
)

const (
	resourcePrefix = "tailscale"
	modulePath     = "github.com/Smana/provider-tailscale"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		acl.Configure,
		deviceauthorization.Configure,
		devicekey.Configure,
		devicesubnetroutes.Configure,
		devicetags.Configure,
		dnsnameservers.Configure,
		dnspreferences.Configure,
		dnssearchpaths.Configure,
		tailnetkey.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

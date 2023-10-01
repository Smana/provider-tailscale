/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	acl "github.com/Smana/provider-tailscale/internal/controller/acl/acl"
	autorization "github.com/Smana/provider-tailscale/internal/controller/device/autorization"
	key "github.com/Smana/provider-tailscale/internal/controller/device/key"
	subnetroutes "github.com/Smana/provider-tailscale/internal/controller/device/subnetroutes"
	tags "github.com/Smana/provider-tailscale/internal/controller/device/tags"
	nameservers "github.com/Smana/provider-tailscale/internal/controller/dns/nameservers"
	preferences "github.com/Smana/provider-tailscale/internal/controller/dns/preferences"
	searchpaths "github.com/Smana/provider-tailscale/internal/controller/dns/searchpaths"
	providerconfig "github.com/Smana/provider-tailscale/internal/controller/providerconfig"
	keytailnet "github.com/Smana/provider-tailscale/internal/controller/tailnet/key"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		acl.Setup,
		autorization.Setup,
		key.Setup,
		subnetroutes.Setup,
		tags.Setup,
		nameservers.Setup,
		preferences.Setup,
		searchpaths.Setup,
		providerconfig.Setup,
		keytailnet.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/upbound/upjet/pkg/terraform"

	"github.com/Smana/provider-tailscale/apis/v1beta1"
)

const (
	keyApiKey            = "api_key"
	keyBaseUrl           = "base_url"
	keyOauthClientID     = "oauth_client_id"
	keyOauthClientSecret = "oauth_client_secret"
	keyTailnet           = "tailnet"
	keyUserAgent         = "user_agent"
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal tailscale credentials as JSON"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// Set credentials in Terraform provider configuration.
		/*ps.Configuration = map[string]any{
			"username": creds["username"],
			"password": creds["password"],
		}*/
		ps.Configuration = map[string]any{}
		if v, ok := creds[keyApiKey]; ok {
			ps.Configuration[keyApiKey] = v
		}
		if v, ok := creds[keyBaseUrl]; ok {
			ps.Configuration[keyBaseUrl] = v
		}
		if v, ok := creds[keyOauthClientID]; ok {
			ps.Configuration[keyOauthClientID] = v
		}
		if v, ok := creds[keyOauthClientSecret]; ok {
			ps.Configuration[keyOauthClientSecret] = v
		}
		if v, ok := creds[keyTailnet]; ok {
			ps.Configuration[keyTailnet] = v
		}
		if v, ok := creds[keyUserAgent]; ok {
			ps.Configuration[keyUserAgent] = v
		}
		return ps, nil
	}
}
/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/crossplane-contrib/provider-mongodb-atlas/apis/v1beta1"
)

const (
	keyPublicKey  = "public_key"
	keyPrivateKey = "private_key"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal mongodbatlas credentials as JSON"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
//
//nolint:gocyclo // Breaking this up doesn't seem worth yet more layers of abstraction.
func TerraformSetupBuilder(tfProvider *schema.Provider) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{}
		cnf := terraform.ProviderConfiguration{}
		// baseURL := "https://api.github.com"

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

		mongodbCreds := map[string]string{}
		if err := json.Unmarshal(data, &mongodbCreds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		if v, ok := mongodbCreds[keyPublicKey]; ok {
			cnf[keyPublicKey] = v
		}
		if v, ok := mongodbCreds[keyPrivateKey]; ok {
			cnf[keyPrivateKey] = v
		}

		ps.Configuration = cnf

		return ps, errors.Wrap(configureNoForkMongoDBClient(ctx, &ps, *tfProvider), "failed to configure the no-fork MongoDB Atlas client")
	}

}

func configureNoForkMongoDBClient(ctx context.Context, ps *terraform.Setup, p schema.Provider) error {
	diag := p.Configure(context.WithoutCancel(ctx), &tfsdk.ResourceConfig{
		Config: ps.Configuration,
	})
	if diag != nil && diag.HasError() {
		return errors.Errorf("failed to configure the provider: %v", diag)
	}
	ps.Meta = p.Meta()

	return nil
}

// // TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// // returns Terraform provider setup configuration
// func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
// 	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
// 		ps := terraform.Setup{
// 			Version: version,
// 			Requirement: terraform.ProviderRequirement{
// 				Source:  providerSource,
// 				Version: providerVersion,
// 			},
// 		}

// 		configRef := mg.GetProviderConfigReference()
// 		if configRef == nil {
// 			return ps, errors.New(errNoProviderConfig)
// 		}
// 		pc := &v1beta1.ProviderConfig{}
// 		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
// 			return ps, errors.Wrap(err, errGetProviderConfig)
// 		}

// 		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
// 		if err := t.Track(ctx, mg); err != nil {
// 			return ps, errors.Wrap(err, errTrackUsage)
// 		}

// 		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
// 		if err != nil {
// 			return ps, errors.Wrap(err, errExtractCredentials)
// 		}
// 		creds := map[string]string{}
// 		if err := json.Unmarshal(data, &creds); err != nil {
// 			return ps, errors.Wrap(err, errUnmarshalCredentials)
// 		}

// 		// Set credentials in Terraform provider configuration.
// 		ps.Configuration = map[string]any{}

// 		if v, ok := creds[keyPublicKey]; ok {
// 			ps.Configuration[keyPublicKey] = v
// 		}
// 		if v, ok := creds[keyPrivateKey]; ok {
// 			ps.Configuration[keyPrivateKey] = v
// 		}
// 		return ps, nil
// 	}
// }

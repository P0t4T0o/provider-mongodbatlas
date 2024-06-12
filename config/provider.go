/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	"context"
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/registry/reference"
	conversiontfjson "github.com/crossplane/upjet/pkg/types/conversion/tfjson"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/crossplane-contrib/provider-mongodbatlas/config/database"
	"github.com/crossplane-contrib/provider-mongodbatlas/config/mongodbatlas"
	"github.com/crossplane-contrib/provider-mongodbatlas/config/privatelink"
	"github.com/crossplane-contrib/provider-mongodbatlas/config/project"
)

const (
	resourcePrefix = "mongodbatlas"
	modulePath     = "github.com/crossplane-contrib/provider-mongodbatlas"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// workaround for the TF AWS v4.67.0-based no-fork release: We would like to
// keep the types in the generated CRDs intact
// (prevent number->int type replacements).
func getProviderSchema(s string) (*schema.Provider, error) {
	ps := tfjson.ProviderSchemas{}
	if err := ps.UnmarshalJSON([]byte(s)); err != nil {
		panic(err)
	}
	if len(ps.Schemas) != 1 {
		return nil, errors.Errorf("there should exactly be 1 provider schema but there are %d", len(ps.Schemas))
	}
	var rs map[string]*tfjson.Schema
	for _, v := range ps.Schemas {
		rs = v.ResourceSchemas
		break
	}
	return &schema.Provider{
		ResourcesMap: conversiontfjson.GetV2ResourceMap(rs),
	}, nil
}

// GetProvider returns provider configuration
func GetProvider(_ context.Context, generationProvider bool) (*ujconfig.Provider, error) {
	var p *schema.Provider
	var err error
	if generationProvider {
		p, err = getProviderSchema(providerSchema)
	} else {
		p, err = getProviderSchema(providerSchema)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get the Terraform provider schema with generation mode set to %t", generationProvider)
	}

	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithShortName("mongodbatlas"),
		ujconfig.WithRootGroup("mongodbatlas.crossplane.io"),
		ujconfig.WithIncludeList(CLIReconciledResourceList()),
		ujconfig.WithTerraformPluginSDKIncludeList(TerraformPluginSDKResourceList()),
		ujconfig.WithDefaultResourceOptions(ResourceConfigurator()),
		ujconfig.WithReferenceInjectors([]ujconfig.ReferenceInjector{reference.NewInjector(modulePath)}),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithTerraformProvider(p),
		ujconfig.WithDefaultResourceOptions(
			// API group overrides from Terraform import statements
			groupKindOverride(),
			commonReferences(),
		),
	)

	// add custom config functions
	for _, configure := range ProviderConfiguration {
		configure(pc)
	}

	pc.ConfigureResources()

	// // This function runs after the special configurations were applied. However, if some references were configured in
	// // the configuration files, the reference generator does not override them.
	// for _, r := range pc.Resources {
	// 	delete(r.References, "resource_group_name")
	// 	if err := common.AddCommonReferences(r); err != nil {
	// 		panic(err)
	// 	}
	// }
	return pc, err
}

// CLIReconciledResourceList returns the list of resources that have external
// name configured in ExternalNameConfigs table and to be reconciled under
// the TF CLI based architecture.
func CLIReconciledResourceList() []string {
	l := make([]string, len(CLIReconciledExternalNameConfigs))
	i := 0
	for name := range CLIReconciledExternalNameConfigs {
		// Expected format is regex and we'd like to have exact matches.
		l[i] = name + "$"
		i++
	}
	return l
}

// TerraformPluginSDKResourceList returns the list of resources that have external
// name configured in ExternalNameConfigs table and to be reconciled under
// the no-fork architecture.
func TerraformPluginSDKResourceList() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// Expected format is regex and we'd like to have exact matches.
		l[i] = name + "$"
		i++
	}
	return l
}

// Configure configures the specified Provider.
type Configure func(provider *ujconfig.Provider)

// Configurator is a registry for provider Configs.
type Configurator []Configure

// AddConfig adds a Config to the Configurator registry.
func (c *Configurator) AddConfig(conf Configure) {
	*c = append(*c, conf)
}

// ProviderConfiguration is a global registry to be used by
// the resource providers to register their Config functions.
var ProviderConfiguration = Configurator{}

func init() {
	ProviderConfiguration.AddConfig(database.Configure)
	ProviderConfiguration.AddConfig(mongodbatlas.Configure)
	ProviderConfiguration.AddConfig(project.Configure)
	ProviderConfiguration.AddConfig(privatelink.Configure)
}

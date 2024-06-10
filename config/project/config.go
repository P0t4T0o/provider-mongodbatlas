/*
Copyright 2022 The Crossplane Authors.

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

package project

import (
	"context"
	"fmt"
	"strings"

	"github.com/crossplane/upjet/pkg/config"

	"github.com/crossplane-contrib/provider-mongodb-atlas/config/common"
)

// Configure configures the root group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("mongodbatlas_project", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.SetIdentifierArgumentFn = common.SetIdentifierFunc
		r.ExternalName.GetExternalNameFn = getExternalNameFunc
		r.ExternalName.GetIDFn = func(_ context.Context, externalName string, parameters map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
			parts := strings.Split(externalName, ":")
			if len(parts) != 2 {
				return "", nil
			}
			return parts[1], nil
		}
		// https://github.com/crossplane/upjet/issues/372#issuecomment-1999277392
		r.ExternalName.OmittedFields = []string{"ip_addresses"}
	})

	p.AddResourceConfigurator("mongodbatlas_project_ip_access_list", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"ip_address", "cidr_block"},
		}
		r.References["project_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-mongodb-atlas/apis/mongodbatlas/v1alpha1.Project",
		}
	})

	p.AddResourceConfigurator("mongodbatlas_alert_configuration", func(r *config.Resource) {
		r.References["project_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-mongodb-atlas/apis/mongodbatlas/v1alpha1.Project",
		}
	})

	p.AddResourceConfigurator("mongodbatlas_project_api_key", func(r *config.Resource) {
		r.References["project_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-mongodb-atlas/apis/mongodbatlas/v1alpha1.Project",
		}
	})

	p.AddResourceConfigurator("mongodbatlas_project_invitation", func(r *config.Resource) {
		r.References["project_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-mongodb-atlas/apis/mongodbatlas/v1alpha1.Project",
		}
	})

	p.AddResourceConfigurator("mongodbatlas_auditing", func(r *config.Resource) {
		r.References["project_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-mongodb-atlas/apis/mongodbatlas/v1alpha1.Project",
		}
	})

	p.AddResourceConfigurator("mongodbatlas_maintenance_window", func(r *config.Resource) {
		r.ExternalName.OmittedFields = []string{"start_asap"}
		r.References["project_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-mongodb-atlas/apis/mongodbatlas/v1alpha1.Project",
		}
	})
}

func getExternalNameFunc(tfstate map[string]interface{}) (string, error) {
	return fmt.Sprintf("%s:%s", tfstate["name"], tfstate["id"]), nil
}

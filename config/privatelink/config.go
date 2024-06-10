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

package privatelink

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the root group
func Configure(p *config.Provider) {

	p.AddResourceConfigurator("mongodbatlas_privatelink_endpoint", func(r *config.Resource) {
		r.References["project_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-mongodb-atlas/apis/mongodbatlas/v1alpha1.Project",
		}
	})

	p.AddResourceConfigurator("mongodbatlas_privatelink_endpoint_service", func(r *config.Resource) {
		r.References["project_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-mongodb-atlas/apis/mongodbatlas/v1alpha1.Project",
		}
	})
}

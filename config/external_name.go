/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"github.com/crossplane/upjet/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"mongodbatlas_cluster":                config.ExternalName{},
	"mongodbatlas_advanced_cluster":       config.ExternalName{},
	"mongodbatlas_project":                config.ExternalName{},
	"mongodbatlas_project_ip_access_list": config.ExternalName{},
	"mongodbatlas_database_user":          config.ExternalName{},

	"mongodbatlas_federated_query_limit":                                       config.ExternalName{},
	"mongodbatlas_ldap_verify":                                                 config.ExternalName{},
	"mongodbatlas_online_archive":                                              config.ExternalName{},
	"mongodbatlas_third_party_integration":                                     config.ExternalName{},
	"mongodbatlas_privatelink_endpoint_service_data_federation_online_archive": config.ExternalName{},
	// "mongodbatlas_search_deployment":                                           config.ExternalName{},
	"mongodbatlas_api_key":                                 config.ExternalName{},
	"mongodbatlas_search_index":                            config.ExternalName{},
	"mongodbatlas_organization":                            config.ExternalName{},
	"mongodbatlas_cluster_outage_simulation":               config.ExternalName{},
	"mongodbatlas_federated_settings_org_config":           config.ExternalName{},
	"mongodbatlas_privatelink_endpoint":                    config.ExternalName{},
	"mongodbatlas_privatelink_endpoint_service_serverless": config.ExternalName{},
	// "mongodbatlas_teams":                                   config.ExternalName{}, deprecated
	"mongodbatlas_auditing":                             config.ExternalName{},
	"mongodbatlas_project_invitation":                   config.ExternalName{},
	"mongodbatlas_cloud_backup_schedule":                config.ExternalName{},
	"mongodbatlas_privatelink_endpoint_serverless":      config.ExternalName{},
	"mongodbatlas_alert_configuration":                  config.ExternalName{},
	"mongodbatlas_custom_dns_configuration_cluster_aws": config.ExternalName{},
	"mongodbatlas_cloud_backup_snapshot_restore_job":    config.ExternalName{},
	// "mongodbatlas_push_based_log_export":                   config.ExternalName{},   // (optional) field timeouts: invalid schema type TypeInvalid
	"mongodbatlas_team":                  config.ExternalName{},
	"mongodbatlas_encryption_at_rest":    config.ExternalName{},
	"mongodbatlas_access_list_api_key":   config.ExternalName{},
	"mongodbatlas_global_cluster_config": config.ExternalName{},
	// "mongodbatlas_stream_instance":                config.ExternalName{}, // (required) field data_process_region: invalid schema type TypeInvalid
	"mongodbatlas_cloud_provider_access_setup":         config.ExternalName{},
	"mongodbatlas_project_api_key":                     config.ExternalName{},
	"mongodbatlas_private_endpoint_regional_mode":      config.ExternalName{},
	"mongodbatlas_org_invitation":                      config.ExternalName{},
	"mongodbatlas_cloud_provider_access_authorization": config.ExternalName{},
	"mongodbatlas_ldap_configuration":                  config.ExternalName{},
	"mongodbatlas_data_lake_pipeline":                  config.ExternalName{},
	"mongodbatlas_privatelink_endpoint_service":        config.ExternalName{},
	"mongodbatlas_custom_db_role":                      config.ExternalName{},
	"mongodbatlas_backup_compliance_policy":            config.ExternalName{},
	"mongodbatlas_cloud_backup_snapshot_export_bucket": config.ExternalName{},
	"mongodbatlas_cloud_backup_snapshot_export_job":    config.ExternalName{},
	// "mongodbatlas_stream_connection":                   config.ExternalName{}, // field authentication: invalid schema type TypeInvalid
	"mongodbatlas_event_trigger":                        config.ExternalName{},
	"mongodbatlas_cloud_backup_snapshot":                config.ExternalName{},
	"mongodbatlas_maintenance_window":                   config.ExternalName{},
	"mongodbatlas_network_container":                    config.ExternalName{},
	"mongodbatlas_federated_database_instance":          config.ExternalName{},
	"mongodbatlas_network_peering":                      config.ExternalName{},
	"mongodbatlas_x509_authentication_database_user":    config.ExternalName{},
	"mongodbatlas_federated_settings_org_role_mapping":  config.ExternalName{},
	"mongodbatlas_federated_settings_identity_provider": config.ExternalName{},
	"mongodbatlas_serverless_instance":                  config.ExternalName{},
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}

// CLIReconciledExternalNameConfigs contains all external name configurations which use the classical CLI based architecture.
var CLIReconciledExternalNameConfigs = map[string]config.ExternalName{}

// ResourceConfigurator applies all external name configs
// listed in the table TerraformPluginSDKExternalNameConfigs and
// CLIReconciledExternalNameConfigs and sets the version
// of those resources to v1beta1. For those resource in
// TerraformPluginSDKExternalNameConfigs, it also sets
// config.Resource.UseNoForkClient to `true`.
func ResourceConfigurator() config.ResourceOption {
	return func(r *config.Resource) {
		// if configured both for the no-fork and CLI based architectures,
		// no-fork configuration prevails
		e, configured := ExternalNameConfigs[r.Name]
		if !configured {
			e, configured = CLIReconciledExternalNameConfigs[r.Name]
		}
		if !configured {
			return
		}
		r.Version = "v1alpha1"
		r.ExternalName = e
	}
}

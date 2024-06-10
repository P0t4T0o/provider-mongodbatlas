/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BackupScheduleInitParameters struct {

	// Flag that indicates whether automatic export of cloud backup snapshots to the AWS bucket is enabled. Value can be one of the following:
	AutoExportEnabled *bool `json:"autoExportEnabled,omitempty" tf:"auto_export_enabled,omitempty"`

	// The name of the Atlas cluster that contains the snapshot backup policy you want to retrieve.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-mongodbatlas/apis/mongodbatlas/v1alpha1.Cluster
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",false)
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Reference to a Cluster in mongodbatlas to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameRef *v1.Reference `json:"clusterNameRef,omitempty" tf:"-"`

	// Selector for a Cluster in mongodbatlas to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameSelector *v1.Selector `json:"clusterNameSelector,omitempty" tf:"-"`

	CopySettings []CopySettingsInitParameters `json:"copySettings,omitempty" tf:"copy_settings,omitempty"`

	Export []ExportInitParameters `json:"export,omitempty" tf:"export,omitempty"`

	// Daily policy item
	PolicyItemDaily []PolicyItemDailyInitParameters `json:"policyItemDaily,omitempty" tf:"policy_item_daily,omitempty"`

	// Hourly policy item
	PolicyItemHourly []PolicyItemHourlyInitParameters `json:"policyItemHourly,omitempty" tf:"policy_item_hourly,omitempty"`

	// Monthly policy item
	PolicyItemMonthly []PolicyItemMonthlyInitParameters `json:"policyItemMonthly,omitempty" tf:"policy_item_monthly,omitempty"`

	// Weekly policy item
	PolicyItemWeekly []PolicyItemWeeklyInitParameters `json:"policyItemWeekly,omitempty" tf:"policy_item_weekly,omitempty"`

	// Yearly policy item
	PolicyItemYearly []PolicyItemYearlyInitParameters `json:"policyItemYearly,omitempty" tf:"policy_item_yearly,omitempty"`

	// The unique identifier of the project for the Atlas cluster.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-mongodbatlas/apis/mongodbatlas/v1alpha1.Cluster
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("project_id",false)
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Cluster in mongodbatlas to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Cluster in mongodbatlas to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// UTC Hour of day between 0 and 23, inclusive, representing which hour of the day that Atlas takes snapshots for backup policy items.
	ReferenceHourOfDay *float64 `json:"referenceHourOfDay,omitempty" tf:"reference_hour_of_day,omitempty"`

	// UTC Minutes after reference_hour_of_day that Atlas takes snapshots for backup policy items. Must be between 0 and 59, inclusive.
	ReferenceMinuteOfHour *float64 `json:"referenceMinuteOfHour,omitempty" tf:"reference_minute_of_hour,omitempty"`

	// Number of days back in time you can restore to with point-in-time accuracy. Must be a positive, non-zero integer.
	RestoreWindowDays *float64 `json:"restoreWindowDays,omitempty" tf:"restore_window_days,omitempty"`

	// Specify true to apply the retention changes in the updated backup policy to snapshots that Atlas took previously.
	UpdateSnapshots *bool `json:"updateSnapshots,omitempty" tf:"update_snapshots,omitempty"`

	// Specify true to use organization and project names instead of organization and project UUIDs in the path for the metadata files that Atlas uploads to your S3 bucket after it finishes exporting the snapshots. To learn more about the metadata files that Atlas uploads, see Export Cloud Backup Snapshot.
	UseOrgAndGroupNamesInExportPrefix *bool `json:"useOrgAndGroupNamesInExportPrefix,omitempty" tf:"use_org_and_group_names_in_export_prefix,omitempty"`
}

type BackupScheduleObservation struct {

	// Flag that indicates whether automatic export of cloud backup snapshots to the AWS bucket is enabled. Value can be one of the following:
	AutoExportEnabled *bool `json:"autoExportEnabled,omitempty" tf:"auto_export_enabled,omitempty"`

	// Unique identifier of the Atlas cluster.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// The name of the Atlas cluster that contains the snapshot backup policy you want to retrieve.
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	CopySettings []CopySettingsObservation `json:"copySettings,omitempty" tf:"copy_settings,omitempty"`

	Export []ExportObservation `json:"export,omitempty" tf:"export,omitempty"`

	// Unique identifier of the backup policy item.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Unique identifier of the backup policy.
	IDPolicy *string `json:"idPolicy,omitempty" tf:"id_policy,omitempty"`

	// Timestamp in the number of seconds that have elapsed since the UNIX epoch when Atlas takes the next snapshot.
	NextSnapshot *string `json:"nextSnapshot,omitempty" tf:"next_snapshot,omitempty"`

	// Daily policy item
	PolicyItemDaily []PolicyItemDailyObservation `json:"policyItemDaily,omitempty" tf:"policy_item_daily,omitempty"`

	// Hourly policy item
	PolicyItemHourly []PolicyItemHourlyObservation `json:"policyItemHourly,omitempty" tf:"policy_item_hourly,omitempty"`

	// Monthly policy item
	PolicyItemMonthly []PolicyItemMonthlyObservation `json:"policyItemMonthly,omitempty" tf:"policy_item_monthly,omitempty"`

	// Weekly policy item
	PolicyItemWeekly []PolicyItemWeeklyObservation `json:"policyItemWeekly,omitempty" tf:"policy_item_weekly,omitempty"`

	// Yearly policy item
	PolicyItemYearly []PolicyItemYearlyObservation `json:"policyItemYearly,omitempty" tf:"policy_item_yearly,omitempty"`

	// The unique identifier of the project for the Atlas cluster.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// UTC Hour of day between 0 and 23, inclusive, representing which hour of the day that Atlas takes snapshots for backup policy items.
	ReferenceHourOfDay *float64 `json:"referenceHourOfDay,omitempty" tf:"reference_hour_of_day,omitempty"`

	// UTC Minutes after reference_hour_of_day that Atlas takes snapshots for backup policy items. Must be between 0 and 59, inclusive.
	ReferenceMinuteOfHour *float64 `json:"referenceMinuteOfHour,omitempty" tf:"reference_minute_of_hour,omitempty"`

	// Number of days back in time you can restore to with point-in-time accuracy. Must be a positive, non-zero integer.
	RestoreWindowDays *float64 `json:"restoreWindowDays,omitempty" tf:"restore_window_days,omitempty"`

	// Specify true to apply the retention changes in the updated backup policy to snapshots that Atlas took previously.
	UpdateSnapshots *bool `json:"updateSnapshots,omitempty" tf:"update_snapshots,omitempty"`

	// Specify true to use organization and project names instead of organization and project UUIDs in the path for the metadata files that Atlas uploads to your S3 bucket after it finishes exporting the snapshots. To learn more about the metadata files that Atlas uploads, see Export Cloud Backup Snapshot.
	UseOrgAndGroupNamesInExportPrefix *bool `json:"useOrgAndGroupNamesInExportPrefix,omitempty" tf:"use_org_and_group_names_in_export_prefix,omitempty"`
}

type BackupScheduleParameters struct {

	// Flag that indicates whether automatic export of cloud backup snapshots to the AWS bucket is enabled. Value can be one of the following:
	// +kubebuilder:validation:Optional
	AutoExportEnabled *bool `json:"autoExportEnabled,omitempty" tf:"auto_export_enabled,omitempty"`

	// The name of the Atlas cluster that contains the snapshot backup policy you want to retrieve.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-mongodbatlas/apis/mongodbatlas/v1alpha1.Cluster
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Reference to a Cluster in mongodbatlas to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameRef *v1.Reference `json:"clusterNameRef,omitempty" tf:"-"`

	// Selector for a Cluster in mongodbatlas to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameSelector *v1.Selector `json:"clusterNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	CopySettings []CopySettingsParameters `json:"copySettings,omitempty" tf:"copy_settings,omitempty"`

	// +kubebuilder:validation:Optional
	Export []ExportParameters `json:"export,omitempty" tf:"export,omitempty"`

	// Daily policy item
	// +kubebuilder:validation:Optional
	PolicyItemDaily []PolicyItemDailyParameters `json:"policyItemDaily,omitempty" tf:"policy_item_daily,omitempty"`

	// Hourly policy item
	// +kubebuilder:validation:Optional
	PolicyItemHourly []PolicyItemHourlyParameters `json:"policyItemHourly,omitempty" tf:"policy_item_hourly,omitempty"`

	// Monthly policy item
	// +kubebuilder:validation:Optional
	PolicyItemMonthly []PolicyItemMonthlyParameters `json:"policyItemMonthly,omitempty" tf:"policy_item_monthly,omitempty"`

	// Weekly policy item
	// +kubebuilder:validation:Optional
	PolicyItemWeekly []PolicyItemWeeklyParameters `json:"policyItemWeekly,omitempty" tf:"policy_item_weekly,omitempty"`

	// Yearly policy item
	// +kubebuilder:validation:Optional
	PolicyItemYearly []PolicyItemYearlyParameters `json:"policyItemYearly,omitempty" tf:"policy_item_yearly,omitempty"`

	// The unique identifier of the project for the Atlas cluster.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-mongodbatlas/apis/mongodbatlas/v1alpha1.Cluster
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("project_id",false)
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Cluster in mongodbatlas to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Cluster in mongodbatlas to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// UTC Hour of day between 0 and 23, inclusive, representing which hour of the day that Atlas takes snapshots for backup policy items.
	// +kubebuilder:validation:Optional
	ReferenceHourOfDay *float64 `json:"referenceHourOfDay,omitempty" tf:"reference_hour_of_day,omitempty"`

	// UTC Minutes after reference_hour_of_day that Atlas takes snapshots for backup policy items. Must be between 0 and 59, inclusive.
	// +kubebuilder:validation:Optional
	ReferenceMinuteOfHour *float64 `json:"referenceMinuteOfHour,omitempty" tf:"reference_minute_of_hour,omitempty"`

	// Number of days back in time you can restore to with point-in-time accuracy. Must be a positive, non-zero integer.
	// +kubebuilder:validation:Optional
	RestoreWindowDays *float64 `json:"restoreWindowDays,omitempty" tf:"restore_window_days,omitempty"`

	// Specify true to apply the retention changes in the updated backup policy to snapshots that Atlas took previously.
	// +kubebuilder:validation:Optional
	UpdateSnapshots *bool `json:"updateSnapshots,omitempty" tf:"update_snapshots,omitempty"`

	// Specify true to use organization and project names instead of organization and project UUIDs in the path for the metadata files that Atlas uploads to your S3 bucket after it finishes exporting the snapshots. To learn more about the metadata files that Atlas uploads, see Export Cloud Backup Snapshot.
	// +kubebuilder:validation:Optional
	UseOrgAndGroupNamesInExportPrefix *bool `json:"useOrgAndGroupNamesInExportPrefix,omitempty" tf:"use_org_and_group_names_in_export_prefix,omitempty"`
}

type CopySettingsInitParameters struct {

	// Human-readable label that identifies the cloud provider that stores the snapshot copy. i.e. "AWS" "AZURE" "GCP"
	CloudProvider *string `json:"cloudProvider,omitempty" tf:"cloud_provider,omitempty"`

	// List that describes which types of snapshots to copy. i.e. "HOURLY" "DAILY" "WEEKLY" "MONTHLY" "ON_DEMAND"
	// +listType=set
	Frequencies []*string `json:"frequencies,omitempty" tf:"frequencies,omitempty"`

	// Target region to copy snapshots belonging to replicationSpecId to. Please supply the 'Atlas Region' which can be found under https://www.mongodb.com/docs/atlas/reference/cloud-providers/ 'regions' link
	RegionName *string `json:"regionName,omitempty" tf:"region_name,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the replication object for a zone in a cluster. For global clusters, there can be multiple zones to choose from. For sharded clusters and replica set clusters, there is only one zone in the cluster. To find the Replication Spec Id, consult the replicationSpecs array returned from Return One Multi-Cloud Cluster in One Project.
	ReplicationSpecID *string `json:"replicationSpecId,omitempty" tf:"replication_spec_id,omitempty"`

	// Flag that indicates whether to copy the oplogs to the target region. You can use the oplogs to perform point-in-time restores.
	ShouldCopyOplogs *bool `json:"shouldCopyOplogs,omitempty" tf:"should_copy_oplogs,omitempty"`
}

type CopySettingsObservation struct {

	// Human-readable label that identifies the cloud provider that stores the snapshot copy. i.e. "AWS" "AZURE" "GCP"
	CloudProvider *string `json:"cloudProvider,omitempty" tf:"cloud_provider,omitempty"`

	// List that describes which types of snapshots to copy. i.e. "HOURLY" "DAILY" "WEEKLY" "MONTHLY" "ON_DEMAND"
	// +listType=set
	Frequencies []*string `json:"frequencies,omitempty" tf:"frequencies,omitempty"`

	// Target region to copy snapshots belonging to replicationSpecId to. Please supply the 'Atlas Region' which can be found under https://www.mongodb.com/docs/atlas/reference/cloud-providers/ 'regions' link
	RegionName *string `json:"regionName,omitempty" tf:"region_name,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the replication object for a zone in a cluster. For global clusters, there can be multiple zones to choose from. For sharded clusters and replica set clusters, there is only one zone in the cluster. To find the Replication Spec Id, consult the replicationSpecs array returned from Return One Multi-Cloud Cluster in One Project.
	ReplicationSpecID *string `json:"replicationSpecId,omitempty" tf:"replication_spec_id,omitempty"`

	// Flag that indicates whether to copy the oplogs to the target region. You can use the oplogs to perform point-in-time restores.
	ShouldCopyOplogs *bool `json:"shouldCopyOplogs,omitempty" tf:"should_copy_oplogs,omitempty"`
}

type CopySettingsParameters struct {

	// Human-readable label that identifies the cloud provider that stores the snapshot copy. i.e. "AWS" "AZURE" "GCP"
	// +kubebuilder:validation:Optional
	CloudProvider *string `json:"cloudProvider,omitempty" tf:"cloud_provider,omitempty"`

	// List that describes which types of snapshots to copy. i.e. "HOURLY" "DAILY" "WEEKLY" "MONTHLY" "ON_DEMAND"
	// +kubebuilder:validation:Optional
	// +listType=set
	Frequencies []*string `json:"frequencies,omitempty" tf:"frequencies,omitempty"`

	// Target region to copy snapshots belonging to replicationSpecId to. Please supply the 'Atlas Region' which can be found under https://www.mongodb.com/docs/atlas/reference/cloud-providers/ 'regions' link
	// +kubebuilder:validation:Optional
	RegionName *string `json:"regionName,omitempty" tf:"region_name,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the replication object for a zone in a cluster. For global clusters, there can be multiple zones to choose from. For sharded clusters and replica set clusters, there is only one zone in the cluster. To find the Replication Spec Id, consult the replicationSpecs array returned from Return One Multi-Cloud Cluster in One Project.
	// +kubebuilder:validation:Optional
	ReplicationSpecID *string `json:"replicationSpecId,omitempty" tf:"replication_spec_id,omitempty"`

	// Flag that indicates whether to copy the oplogs to the target region. You can use the oplogs to perform point-in-time restores.
	// +kubebuilder:validation:Optional
	ShouldCopyOplogs *bool `json:"shouldCopyOplogs,omitempty" tf:"should_copy_oplogs,omitempty"`
}

type ExportInitParameters struct {

	// Unique identifier of the mongodbatlas_cloud_backup_snapshot_export_bucket export_bucket_id value.
	ExportBucketID *string `json:"exportBucketId,omitempty" tf:"export_bucket_id,omitempty"`

	// Frequency associated with the export snapshot item.
	FrequencyType *string `json:"frequencyType,omitempty" tf:"frequency_type,omitempty"`
}

type ExportObservation struct {

	// Unique identifier of the mongodbatlas_cloud_backup_snapshot_export_bucket export_bucket_id value.
	ExportBucketID *string `json:"exportBucketId,omitempty" tf:"export_bucket_id,omitempty"`

	// Frequency associated with the export snapshot item.
	FrequencyType *string `json:"frequencyType,omitempty" tf:"frequency_type,omitempty"`
}

type ExportParameters struct {

	// Unique identifier of the mongodbatlas_cloud_backup_snapshot_export_bucket export_bucket_id value.
	// +kubebuilder:validation:Optional
	ExportBucketID *string `json:"exportBucketId,omitempty" tf:"export_bucket_id,omitempty"`

	// Frequency associated with the export snapshot item.
	// +kubebuilder:validation:Optional
	FrequencyType *string `json:"frequencyType,omitempty" tf:"frequency_type,omitempty"`
}

type PolicyItemDailyInitParameters struct {

	// Desired frequency of the new backup policy item specified by frequency_type (hourly in this case). The supported values for hourly policies are 1, 2, 4, 6, 8 or 12 hours. Note that 12 hours is the only accepted value for NVMe clusters.
	FrequencyInterval *float64 `json:"frequencyInterval,omitempty" tf:"frequency_interval,omitempty"`

	// Scope of the backup policy item: days, weeks, months, or years.
	RetentionUnit *string `json:"retentionUnit,omitempty" tf:"retention_unit,omitempty"`

	// Value to associate with retention_unit.
	RetentionValue *float64 `json:"retentionValue,omitempty" tf:"retention_value,omitempty"`
}

type PolicyItemDailyObservation struct {

	// Desired frequency of the new backup policy item specified by frequency_type (hourly in this case). The supported values for hourly policies are 1, 2, 4, 6, 8 or 12 hours. Note that 12 hours is the only accepted value for NVMe clusters.
	FrequencyInterval *float64 `json:"frequencyInterval,omitempty" tf:"frequency_interval,omitempty"`

	// Frequency associated with the export snapshot item.
	FrequencyType *string `json:"frequencyType,omitempty" tf:"frequency_type,omitempty"`

	// Unique identifier of the backup policy item.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Scope of the backup policy item: days, weeks, months, or years.
	RetentionUnit *string `json:"retentionUnit,omitempty" tf:"retention_unit,omitempty"`

	// Value to associate with retention_unit.
	RetentionValue *float64 `json:"retentionValue,omitempty" tf:"retention_value,omitempty"`
}

type PolicyItemDailyParameters struct {

	// Desired frequency of the new backup policy item specified by frequency_type (hourly in this case). The supported values for hourly policies are 1, 2, 4, 6, 8 or 12 hours. Note that 12 hours is the only accepted value for NVMe clusters.
	// +kubebuilder:validation:Optional
	FrequencyInterval *float64 `json:"frequencyInterval" tf:"frequency_interval,omitempty"`

	// Scope of the backup policy item: days, weeks, months, or years.
	// +kubebuilder:validation:Optional
	RetentionUnit *string `json:"retentionUnit" tf:"retention_unit,omitempty"`

	// Value to associate with retention_unit.
	// +kubebuilder:validation:Optional
	RetentionValue *float64 `json:"retentionValue" tf:"retention_value,omitempty"`
}

type PolicyItemHourlyInitParameters struct {

	// Desired frequency of the new backup policy item specified by frequency_type (hourly in this case). The supported values for hourly policies are 1, 2, 4, 6, 8 or 12 hours. Note that 12 hours is the only accepted value for NVMe clusters.
	FrequencyInterval *float64 `json:"frequencyInterval,omitempty" tf:"frequency_interval,omitempty"`

	// Scope of the backup policy item: days, weeks, months, or years.
	RetentionUnit *string `json:"retentionUnit,omitempty" tf:"retention_unit,omitempty"`

	// Value to associate with retention_unit.
	RetentionValue *float64 `json:"retentionValue,omitempty" tf:"retention_value,omitempty"`
}

type PolicyItemHourlyObservation struct {

	// Desired frequency of the new backup policy item specified by frequency_type (hourly in this case). The supported values for hourly policies are 1, 2, 4, 6, 8 or 12 hours. Note that 12 hours is the only accepted value for NVMe clusters.
	FrequencyInterval *float64 `json:"frequencyInterval,omitempty" tf:"frequency_interval,omitempty"`

	// Frequency associated with the export snapshot item.
	FrequencyType *string `json:"frequencyType,omitempty" tf:"frequency_type,omitempty"`

	// Unique identifier of the backup policy item.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Scope of the backup policy item: days, weeks, months, or years.
	RetentionUnit *string `json:"retentionUnit,omitempty" tf:"retention_unit,omitempty"`

	// Value to associate with retention_unit.
	RetentionValue *float64 `json:"retentionValue,omitempty" tf:"retention_value,omitempty"`
}

type PolicyItemHourlyParameters struct {

	// Desired frequency of the new backup policy item specified by frequency_type (hourly in this case). The supported values for hourly policies are 1, 2, 4, 6, 8 or 12 hours. Note that 12 hours is the only accepted value for NVMe clusters.
	// +kubebuilder:validation:Optional
	FrequencyInterval *float64 `json:"frequencyInterval" tf:"frequency_interval,omitempty"`

	// Scope of the backup policy item: days, weeks, months, or years.
	// +kubebuilder:validation:Optional
	RetentionUnit *string `json:"retentionUnit" tf:"retention_unit,omitempty"`

	// Value to associate with retention_unit.
	// +kubebuilder:validation:Optional
	RetentionValue *float64 `json:"retentionValue" tf:"retention_value,omitempty"`
}

type PolicyItemMonthlyInitParameters struct {

	// Desired frequency of the new backup policy item specified by frequency_type (hourly in this case). The supported values for hourly policies are 1, 2, 4, 6, 8 or 12 hours. Note that 12 hours is the only accepted value for NVMe clusters.
	FrequencyInterval *float64 `json:"frequencyInterval,omitempty" tf:"frequency_interval,omitempty"`

	// Scope of the backup policy item: days, weeks, months, or years.
	RetentionUnit *string `json:"retentionUnit,omitempty" tf:"retention_unit,omitempty"`

	// Value to associate with retention_unit.
	RetentionValue *float64 `json:"retentionValue,omitempty" tf:"retention_value,omitempty"`
}

type PolicyItemMonthlyObservation struct {

	// Desired frequency of the new backup policy item specified by frequency_type (hourly in this case). The supported values for hourly policies are 1, 2, 4, 6, 8 or 12 hours. Note that 12 hours is the only accepted value for NVMe clusters.
	FrequencyInterval *float64 `json:"frequencyInterval,omitempty" tf:"frequency_interval,omitempty"`

	// Frequency associated with the export snapshot item.
	FrequencyType *string `json:"frequencyType,omitempty" tf:"frequency_type,omitempty"`

	// Unique identifier of the backup policy item.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Scope of the backup policy item: days, weeks, months, or years.
	RetentionUnit *string `json:"retentionUnit,omitempty" tf:"retention_unit,omitempty"`

	// Value to associate with retention_unit.
	RetentionValue *float64 `json:"retentionValue,omitempty" tf:"retention_value,omitempty"`
}

type PolicyItemMonthlyParameters struct {

	// Desired frequency of the new backup policy item specified by frequency_type (hourly in this case). The supported values for hourly policies are 1, 2, 4, 6, 8 or 12 hours. Note that 12 hours is the only accepted value for NVMe clusters.
	// +kubebuilder:validation:Optional
	FrequencyInterval *float64 `json:"frequencyInterval" tf:"frequency_interval,omitempty"`

	// Scope of the backup policy item: days, weeks, months, or years.
	// +kubebuilder:validation:Optional
	RetentionUnit *string `json:"retentionUnit" tf:"retention_unit,omitempty"`

	// Value to associate with retention_unit.
	// +kubebuilder:validation:Optional
	RetentionValue *float64 `json:"retentionValue" tf:"retention_value,omitempty"`
}

type PolicyItemWeeklyInitParameters struct {

	// Desired frequency of the new backup policy item specified by frequency_type (hourly in this case). The supported values for hourly policies are 1, 2, 4, 6, 8 or 12 hours. Note that 12 hours is the only accepted value for NVMe clusters.
	FrequencyInterval *float64 `json:"frequencyInterval,omitempty" tf:"frequency_interval,omitempty"`

	// Scope of the backup policy item: days, weeks, months, or years.
	RetentionUnit *string `json:"retentionUnit,omitempty" tf:"retention_unit,omitempty"`

	// Value to associate with retention_unit.
	RetentionValue *float64 `json:"retentionValue,omitempty" tf:"retention_value,omitempty"`
}

type PolicyItemWeeklyObservation struct {

	// Desired frequency of the new backup policy item specified by frequency_type (hourly in this case). The supported values for hourly policies are 1, 2, 4, 6, 8 or 12 hours. Note that 12 hours is the only accepted value for NVMe clusters.
	FrequencyInterval *float64 `json:"frequencyInterval,omitempty" tf:"frequency_interval,omitempty"`

	// Frequency associated with the export snapshot item.
	FrequencyType *string `json:"frequencyType,omitempty" tf:"frequency_type,omitempty"`

	// Unique identifier of the backup policy item.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Scope of the backup policy item: days, weeks, months, or years.
	RetentionUnit *string `json:"retentionUnit,omitempty" tf:"retention_unit,omitempty"`

	// Value to associate with retention_unit.
	RetentionValue *float64 `json:"retentionValue,omitempty" tf:"retention_value,omitempty"`
}

type PolicyItemWeeklyParameters struct {

	// Desired frequency of the new backup policy item specified by frequency_type (hourly in this case). The supported values for hourly policies are 1, 2, 4, 6, 8 or 12 hours. Note that 12 hours is the only accepted value for NVMe clusters.
	// +kubebuilder:validation:Optional
	FrequencyInterval *float64 `json:"frequencyInterval" tf:"frequency_interval,omitempty"`

	// Scope of the backup policy item: days, weeks, months, or years.
	// +kubebuilder:validation:Optional
	RetentionUnit *string `json:"retentionUnit" tf:"retention_unit,omitempty"`

	// Value to associate with retention_unit.
	// +kubebuilder:validation:Optional
	RetentionValue *float64 `json:"retentionValue" tf:"retention_value,omitempty"`
}

type PolicyItemYearlyInitParameters struct {

	// Desired frequency of the new backup policy item specified by frequency_type (hourly in this case). The supported values for hourly policies are 1, 2, 4, 6, 8 or 12 hours. Note that 12 hours is the only accepted value for NVMe clusters.
	FrequencyInterval *float64 `json:"frequencyInterval,omitempty" tf:"frequency_interval,omitempty"`

	// Scope of the backup policy item: days, weeks, months, or years.
	RetentionUnit *string `json:"retentionUnit,omitempty" tf:"retention_unit,omitempty"`

	// Value to associate with retention_unit.
	RetentionValue *float64 `json:"retentionValue,omitempty" tf:"retention_value,omitempty"`
}

type PolicyItemYearlyObservation struct {

	// Desired frequency of the new backup policy item specified by frequency_type (hourly in this case). The supported values for hourly policies are 1, 2, 4, 6, 8 or 12 hours. Note that 12 hours is the only accepted value for NVMe clusters.
	FrequencyInterval *float64 `json:"frequencyInterval,omitempty" tf:"frequency_interval,omitempty"`

	// Frequency associated with the export snapshot item.
	FrequencyType *string `json:"frequencyType,omitempty" tf:"frequency_type,omitempty"`

	// Unique identifier of the backup policy item.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Scope of the backup policy item: days, weeks, months, or years.
	RetentionUnit *string `json:"retentionUnit,omitempty" tf:"retention_unit,omitempty"`

	// Value to associate with retention_unit.
	RetentionValue *float64 `json:"retentionValue,omitempty" tf:"retention_value,omitempty"`
}

type PolicyItemYearlyParameters struct {

	// Desired frequency of the new backup policy item specified by frequency_type (hourly in this case). The supported values for hourly policies are 1, 2, 4, 6, 8 or 12 hours. Note that 12 hours is the only accepted value for NVMe clusters.
	// +kubebuilder:validation:Optional
	FrequencyInterval *float64 `json:"frequencyInterval" tf:"frequency_interval,omitempty"`

	// Scope of the backup policy item: days, weeks, months, or years.
	// +kubebuilder:validation:Optional
	RetentionUnit *string `json:"retentionUnit" tf:"retention_unit,omitempty"`

	// Value to associate with retention_unit.
	// +kubebuilder:validation:Optional
	RetentionValue *float64 `json:"retentionValue" tf:"retention_value,omitempty"`
}

// BackupScheduleSpec defines the desired state of BackupSchedule
type BackupScheduleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupScheduleParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider BackupScheduleInitParameters `json:"initProvider,omitempty"`
}

// BackupScheduleStatus defines the observed state of BackupSchedule.
type BackupScheduleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupScheduleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BackupSchedule is the Schema for the BackupSchedules API. Provides a Cloud Backup Schedule resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlas}
type BackupSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupScheduleSpec   `json:"spec"`
	Status            BackupScheduleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupScheduleList contains a list of BackupSchedules
type BackupScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupSchedule `json:"items"`
}

// Repository type metadata.
var (
	BackupSchedule_Kind             = "BackupSchedule"
	BackupSchedule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackupSchedule_Kind}.String()
	BackupSchedule_KindAPIVersion   = BackupSchedule_Kind + "." + CRDGroupVersion.String()
	BackupSchedule_GroupVersionKind = CRDGroupVersion.WithKind(BackupSchedule_Kind)
)

func init() {
	SchemeBuilder.Register(&BackupSchedule{}, &BackupScheduleList{})
}

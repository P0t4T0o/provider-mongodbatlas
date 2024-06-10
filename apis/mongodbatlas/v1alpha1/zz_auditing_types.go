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

type AuditingInitParameters struct {

	// Indicates whether the auditing system captures successful authentication attempts for audit filters using the "atype" : "authCheck" auditing event. For more information, see auditAuthorizationSuccess.  Warning! Enabling Audit authorization successes can severely impact cluster performance. Enable this option with caution.
	AuditAuthorizationSuccess *bool `json:"auditAuthorizationSuccess,omitempty" tf:"audit_authorization_success,omitempty"`

	// JSON-formatted audit filter. For complete documentation on custom auditing filters, see Configure Audit Filters.
	AuditFilter *string `json:"auditFilter,omitempty" tf:"audit_filter,omitempty"`

	// Denotes whether or not the project associated with the {project_id} has database auditing enabled.  Defaults to false.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The unique ID for the project to configure auditing. Note: When changing this value to a different project_id it will delete the current audit settings for the original project that was assigned to.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-mongodbatlas/apis/mongodbatlas/v1alpha1.Project
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in mongodbatlas to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in mongodbatlas to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`
}

type AuditingObservation struct {

	// Indicates whether the auditing system captures successful authentication attempts for audit filters using the "atype" : "authCheck" auditing event. For more information, see auditAuthorizationSuccess.  Warning! Enabling Audit authorization successes can severely impact cluster performance. Enable this option with caution.
	AuditAuthorizationSuccess *bool `json:"auditAuthorizationSuccess,omitempty" tf:"audit_authorization_success,omitempty"`

	// JSON-formatted audit filter. For complete documentation on custom auditing filters, see Configure Audit Filters.
	AuditFilter *string `json:"auditFilter,omitempty" tf:"audit_filter,omitempty"`

	// Denotes the configuration method for the audit filter. Possible values are:
	ConfigurationType *string `json:"configurationType,omitempty" tf:"configuration_type,omitempty"`

	// Denotes whether or not the project associated with the {project_id} has database auditing enabled.  Defaults to false.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The unique ID for the project to configure auditing. Note: When changing this value to a different project_id it will delete the current audit settings for the original project that was assigned to.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type AuditingParameters struct {

	// Indicates whether the auditing system captures successful authentication attempts for audit filters using the "atype" : "authCheck" auditing event. For more information, see auditAuthorizationSuccess.  Warning! Enabling Audit authorization successes can severely impact cluster performance. Enable this option with caution.
	// +kubebuilder:validation:Optional
	AuditAuthorizationSuccess *bool `json:"auditAuthorizationSuccess,omitempty" tf:"audit_authorization_success,omitempty"`

	// JSON-formatted audit filter. For complete documentation on custom auditing filters, see Configure Audit Filters.
	// +kubebuilder:validation:Optional
	AuditFilter *string `json:"auditFilter,omitempty" tf:"audit_filter,omitempty"`

	// Denotes whether or not the project associated with the {project_id} has database auditing enabled.  Defaults to false.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The unique ID for the project to configure auditing. Note: When changing this value to a different project_id it will delete the current audit settings for the original project that was assigned to.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-mongodbatlas/apis/mongodbatlas/v1alpha1.Project
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in mongodbatlas to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in mongodbatlas to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`
}

// AuditingSpec defines the desired state of Auditing
type AuditingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuditingParameters `json:"forProvider"`
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
	InitProvider AuditingInitParameters `json:"initProvider,omitempty"`
}

// AuditingStatus defines the observed state of Auditing.
type AuditingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuditingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Auditing is the Schema for the Auditings API. Provides a Auditing resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlas}
type Auditing struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AuditingSpec   `json:"spec"`
	Status            AuditingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuditingList contains a list of Auditings
type AuditingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Auditing `json:"items"`
}

// Repository type metadata.
var (
	Auditing_Kind             = "Auditing"
	Auditing_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Auditing_Kind}.String()
	Auditing_KindAPIVersion   = Auditing_Kind + "." + CRDGroupVersion.String()
	Auditing_GroupVersionKind = CRDGroupVersion.WithKind(Auditing_Kind)
)

func init() {
	SchemeBuilder.Register(&Auditing{}, &AuditingList{})
}

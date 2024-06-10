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

type WindowInitParameters struct {

	// Defer any scheduled maintenance for the given project for one week.
	AutoDefer *bool `json:"autoDefer,omitempty" tf:"auto_defer,omitempty"`

	// Flag that indicates whether you want to defer all maintenance windows one week they would be triggered.
	AutoDeferOnceEnabled *bool `json:"autoDeferOnceEnabled,omitempty" tf:"auto_defer_once_enabled,omitempty"`

	// Day of the week when you would like the maintenance window to start as a 1-based integer: Su=1, M=2, T=3, W=4, T=5, F=6, Sa=7.
	DayOfWeek *float64 `json:"dayOfWeek,omitempty" tf:"day_of_week,omitempty"`

	// Defer the next scheduled maintenance for the given project for one week.
	Defer *bool `json:"defer,omitempty" tf:"defer,omitempty"`

	// Hour of the day when you would like the maintenance window to start. This parameter uses the 24-hour clock, where midnight is 0, noon is 12 (Time zone is UTC).
	HourOfDay *float64 `json:"hourOfDay,omitempty" tf:"hour_of_day,omitempty"`

	// The unique identifier of the project for the Maintenance Window.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-mongodbatlas/apis/mongodbatlas/v1alpha1.Project
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in mongodbatlas to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in mongodbatlas to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`
}

type WindowObservation struct {

	// Defer any scheduled maintenance for the given project for one week.
	AutoDefer *bool `json:"autoDefer,omitempty" tf:"auto_defer,omitempty"`

	// Flag that indicates whether you want to defer all maintenance windows one week they would be triggered.
	AutoDeferOnceEnabled *bool `json:"autoDeferOnceEnabled,omitempty" tf:"auto_defer_once_enabled,omitempty"`

	// Day of the week when you would like the maintenance window to start as a 1-based integer: Su=1, M=2, T=3, W=4, T=5, F=6, Sa=7.
	DayOfWeek *float64 `json:"dayOfWeek,omitempty" tf:"day_of_week,omitempty"`

	// Defer the next scheduled maintenance for the given project for one week.
	Defer *bool `json:"defer,omitempty" tf:"defer,omitempty"`

	// Hour of the day when you would like the maintenance window to start. This parameter uses the 24-hour clock, where midnight is 0, noon is 12 (Time zone is UTC).
	HourOfDay *float64 `json:"hourOfDay,omitempty" tf:"hour_of_day,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Number of times the current maintenance event for this project has been deferred, there can be a maximum of 2 deferrals.
	NumberOfDeferrals *float64 `json:"numberOfDeferrals,omitempty" tf:"number_of_deferrals,omitempty"`

	// The unique identifier of the project for the Maintenance Window.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type WindowParameters struct {

	// Defer any scheduled maintenance for the given project for one week.
	// +kubebuilder:validation:Optional
	AutoDefer *bool `json:"autoDefer,omitempty" tf:"auto_defer,omitempty"`

	// Flag that indicates whether you want to defer all maintenance windows one week they would be triggered.
	// +kubebuilder:validation:Optional
	AutoDeferOnceEnabled *bool `json:"autoDeferOnceEnabled,omitempty" tf:"auto_defer_once_enabled,omitempty"`

	// Day of the week when you would like the maintenance window to start as a 1-based integer: Su=1, M=2, T=3, W=4, T=5, F=6, Sa=7.
	// +kubebuilder:validation:Optional
	DayOfWeek *float64 `json:"dayOfWeek,omitempty" tf:"day_of_week,omitempty"`

	// Defer the next scheduled maintenance for the given project for one week.
	// +kubebuilder:validation:Optional
	Defer *bool `json:"defer,omitempty" tf:"defer,omitempty"`

	// Hour of the day when you would like the maintenance window to start. This parameter uses the 24-hour clock, where midnight is 0, noon is 12 (Time zone is UTC).
	// +kubebuilder:validation:Optional
	HourOfDay *float64 `json:"hourOfDay,omitempty" tf:"hour_of_day,omitempty"`

	// The unique identifier of the project for the Maintenance Window.
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

// WindowSpec defines the desired state of Window
type WindowSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WindowParameters `json:"forProvider"`
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
	InitProvider WindowInitParameters `json:"initProvider,omitempty"`
}

// WindowStatus defines the observed state of Window.
type WindowStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WindowObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Window is the Schema for the Windows API. Provides an Maintenance Window resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlas}
type Window struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dayOfWeek) || (has(self.initProvider) && has(self.initProvider.dayOfWeek))",message="spec.forProvider.dayOfWeek is a required parameter"
	Spec   WindowSpec   `json:"spec"`
	Status WindowStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WindowList contains a list of Windows
type WindowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Window `json:"items"`
}

// Repository type metadata.
var (
	Window_Kind             = "Window"
	Window_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Window_Kind}.String()
	Window_KindAPIVersion   = Window_Kind + "." + CRDGroupVersion.String()
	Window_GroupVersionKind = CRDGroupVersion.WithKind(Window_Kind)
)

func init() {
	SchemeBuilder.Register(&Window{}, &WindowList{})
}

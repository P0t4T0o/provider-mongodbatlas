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

type EndpointServiceDataFederationOnlineArchiveInitParameters struct {

	// Human-readable string to associate with this private endpoint.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Human-readable label to identify VPC endpoint DNS name.
	CustomerEndpointDNSName *string `json:"customerEndpointDnsName,omitempty" tf:"customer_endpoint_dns_name,omitempty"`

	// Unique 22-character alphanumeric string that identifies the private endpoint. See Atlas Data Lake supports Amazon Web Services private endpoints using the AWS PrivateLink feature.
	EndpointID *string `json:"endpointId,omitempty" tf:"endpoint_id,omitempty"`

	// Unique 24-hexadecimal digit string that identifies your project.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-mongodb-atlas/apis/mongodbatlas/v1alpha1.Project
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in mongodbatlas to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in mongodbatlas to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// Human-readable label that identifies the cloud service provider.
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// Human-readable label to identify the region of VPC endpoint.  Requires the Atlas region name, see the reference list for AWS, GCP, Azure.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type EndpointServiceDataFederationOnlineArchiveObservation struct {

	// Human-readable string to associate with this private endpoint.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Human-readable label to identify VPC endpoint DNS name.
	CustomerEndpointDNSName *string `json:"customerEndpointDnsName,omitempty" tf:"customer_endpoint_dns_name,omitempty"`

	// Unique 22-character alphanumeric string that identifies the private endpoint. See Atlas Data Lake supports Amazon Web Services private endpoints using the AWS PrivateLink feature.
	EndpointID *string `json:"endpointId,omitempty" tf:"endpoint_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Unique 24-hexadecimal digit string that identifies your project.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Human-readable label that identifies the cloud service provider.
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// Human-readable label to identify the region of VPC endpoint.  Requires the Atlas region name, see the reference list for AWS, GCP, Azure.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Human-readable label that identifies the resource type associated with this private endpoint.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type EndpointServiceDataFederationOnlineArchiveParameters struct {

	// Human-readable string to associate with this private endpoint.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Human-readable label to identify VPC endpoint DNS name.
	// +kubebuilder:validation:Optional
	CustomerEndpointDNSName *string `json:"customerEndpointDnsName,omitempty" tf:"customer_endpoint_dns_name,omitempty"`

	// Unique 22-character alphanumeric string that identifies the private endpoint. See Atlas Data Lake supports Amazon Web Services private endpoints using the AWS PrivateLink feature.
	// +kubebuilder:validation:Optional
	EndpointID *string `json:"endpointId,omitempty" tf:"endpoint_id,omitempty"`

	// Unique 24-hexadecimal digit string that identifies your project.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-mongodb-atlas/apis/mongodbatlas/v1alpha1.Project
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in mongodbatlas to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in mongodbatlas to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// Human-readable label that identifies the cloud service provider.
	// +kubebuilder:validation:Optional
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// Human-readable label to identify the region of VPC endpoint.  Requires the Atlas region name, see the reference list for AWS, GCP, Azure.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// EndpointServiceDataFederationOnlineArchiveSpec defines the desired state of EndpointServiceDataFederationOnlineArchive
type EndpointServiceDataFederationOnlineArchiveSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EndpointServiceDataFederationOnlineArchiveParameters `json:"forProvider"`
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
	InitProvider EndpointServiceDataFederationOnlineArchiveInitParameters `json:"initProvider,omitempty"`
}

// EndpointServiceDataFederationOnlineArchiveStatus defines the observed state of EndpointServiceDataFederationOnlineArchive.
type EndpointServiceDataFederationOnlineArchiveStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EndpointServiceDataFederationOnlineArchiveObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// EndpointServiceDataFederationOnlineArchive is the Schema for the EndpointServiceDataFederationOnlineArchives API. Provides a Privatelink Endpoint Service Data Federation Online Archive resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlas}
type EndpointServiceDataFederationOnlineArchive struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.endpointId) || (has(self.initProvider) && has(self.initProvider.endpointId))",message="spec.forProvider.endpointId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.providerName) || (has(self.initProvider) && has(self.initProvider.providerName))",message="spec.forProvider.providerName is a required parameter"
	Spec   EndpointServiceDataFederationOnlineArchiveSpec   `json:"spec"`
	Status EndpointServiceDataFederationOnlineArchiveStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointServiceDataFederationOnlineArchiveList contains a list of EndpointServiceDataFederationOnlineArchives
type EndpointServiceDataFederationOnlineArchiveList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EndpointServiceDataFederationOnlineArchive `json:"items"`
}

// Repository type metadata.
var (
	EndpointServiceDataFederationOnlineArchive_Kind             = "EndpointServiceDataFederationOnlineArchive"
	EndpointServiceDataFederationOnlineArchive_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EndpointServiceDataFederationOnlineArchive_Kind}.String()
	EndpointServiceDataFederationOnlineArchive_KindAPIVersion   = EndpointServiceDataFederationOnlineArchive_Kind + "." + CRDGroupVersion.String()
	EndpointServiceDataFederationOnlineArchive_GroupVersionKind = CRDGroupVersion.WithKind(EndpointServiceDataFederationOnlineArchive_Kind)
)

func init() {
	SchemeBuilder.Register(&EndpointServiceDataFederationOnlineArchive{}, &EndpointServiceDataFederationOnlineArchiveList{})
}

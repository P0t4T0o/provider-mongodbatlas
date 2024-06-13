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

type InvitationInitParameters struct {

	// Unique 24-hexadecimal digit string that identifies the organization to which you want to invite a user.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Atlas roles to assign to the invited user. If the user accepts the invitation, Atlas assigns these roles to them. The MongoDB Documentation describes the roles a user can have.
	// +listType=set
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`

	// An array of unique 24-hexadecimal digit strings that identify the teams that the user was invited to join.
	// +listType=set
	TeamsIds []*string `json:"teamsIds,omitempty" tf:"teams_ids,omitempty"`

	// Email address of the invited user. This is the address to which Atlas sends the invite. If the user accepts the invitation, they log in to Atlas with this username.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type InvitationObservation struct {

	// Timestamp in ISO 8601 date and time format in UTC when Atlas sent the invitation.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Timestamp in ISO 8601 date and time format in UTC when the invitation expires. Users have 30 days to accept an invitation.
	ExpiresAt *string `json:"expiresAt,omitempty" tf:"expires_at,omitempty"`

	// Autogenerated unique string that identifies this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the invitation in Atlas.
	InvitationID *string `json:"invitationId,omitempty" tf:"invitation_id,omitempty"`

	// Atlas user who invited username to the organization.
	InviterUsername *string `json:"inviterUsername,omitempty" tf:"inviter_username,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the organization to which you want to invite a user.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Atlas roles to assign to the invited user. If the user accepts the invitation, Atlas assigns these roles to them. The MongoDB Documentation describes the roles a user can have.
	// +listType=set
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`

	// An array of unique 24-hexadecimal digit strings that identify the teams that the user was invited to join.
	// +listType=set
	TeamsIds []*string `json:"teamsIds,omitempty" tf:"teams_ids,omitempty"`

	// Email address of the invited user. This is the address to which Atlas sends the invite. If the user accepts the invitation, they log in to Atlas with this username.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type InvitationParameters struct {

	// Unique 24-hexadecimal digit string that identifies the organization to which you want to invite a user.
	// +kubebuilder:validation:Optional
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Atlas roles to assign to the invited user. If the user accepts the invitation, Atlas assigns these roles to them. The MongoDB Documentation describes the roles a user can have.
	// +kubebuilder:validation:Optional
	// +listType=set
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`

	// An array of unique 24-hexadecimal digit strings that identify the teams that the user was invited to join.
	// +kubebuilder:validation:Optional
	// +listType=set
	TeamsIds []*string `json:"teamsIds,omitempty" tf:"teams_ids,omitempty"`

	// Email address of the invited user. This is the address to which Atlas sends the invite. If the user accepts the invitation, they log in to Atlas with this username.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// InvitationSpec defines the desired state of Invitation
type InvitationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InvitationParameters `json:"forProvider"`
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
	InitProvider InvitationInitParameters `json:"initProvider,omitempty"`
}

// InvitationStatus defines the observed state of Invitation.
type InvitationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InvitationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Invitation is the Schema for the Invitations API. Provides an Atlas Organization Invitation resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlas}
type Invitation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.orgId) || (has(self.initProvider) && has(self.initProvider.orgId))",message="spec.forProvider.orgId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.roles) || (has(self.initProvider) && has(self.initProvider.roles))",message="spec.forProvider.roles is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.username) || (has(self.initProvider) && has(self.initProvider.username))",message="spec.forProvider.username is a required parameter"
	Spec   InvitationSpec   `json:"spec"`
	Status InvitationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InvitationList contains a list of Invitations
type InvitationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Invitation `json:"items"`
}

// Repository type metadata.
var (
	Invitation_Kind             = "Invitation"
	Invitation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Invitation_Kind}.String()
	Invitation_KindAPIVersion   = Invitation_Kind + "." + CRDGroupVersion.String()
	Invitation_GroupVersionKind = CRDGroupVersion.WithKind(Invitation_Kind)
)

func init() {
	SchemeBuilder.Register(&Invitation{}, &InvitationList{})
}

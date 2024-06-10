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

type SettingsIdentityProviderInitParameters struct {

	// List that contains the domains associated with the identity provider.
	AssociatedDomains []*string `json:"associatedDomains,omitempty" tf:"associated_domains,omitempty"`

	// Identifier of the intended recipient of the token used in OIDC IdP.
	AudienceClaim []*string `json:"audienceClaim,omitempty" tf:"audience_claim,omitempty"`

	// Client identifier that is assigned to an application by the OIDC Identity Provider.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
	FederationSettingsID *string `json:"federationSettingsId,omitempty" tf:"federation_settings_id,omitempty"`

	// Identifier of the claim which contains OIDC IdP Group IDs in the token.
	GroupsClaim *string `json:"groupsClaim,omitempty" tf:"groups_claim,omitempty"`

	// Unique string that identifies the issuer of the IdP.
	IssuerURI *string `json:"issuerUri,omitempty" tf:"issuer_uri,omitempty"`

	// Human-readable label that identifies the identity provider.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The protocol of the identity provider. Either SAML or OIDC.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// SAML Authentication Request Protocol HTTP method binding (POST or REDIRECT) that Federated Authentication uses to send the authentication request. Atlas supports the following binding values:
	RequestBinding *string `json:"requestBinding,omitempty" tf:"request_binding,omitempty"`

	// Scopes that MongoDB applications will request from the authorization endpoint used for OIDC IdPs.
	RequestedScopes []*string `json:"requestedScopes,omitempty" tf:"requested_scopes,omitempty"`

	// Signature algorithm that Federated Authentication uses to encrypt the identity provider signature.  Valid values include SHA-1 and SHA-256.
	ResponseSignatureAlgorithm *string `json:"responseSignatureAlgorithm,omitempty" tf:"response_signature_algorithm,omitempty"`

	// Flag that indicates whether the identity provider has SSO debug enabled.
	SsoDebugEnabled *bool `json:"ssoDebugEnabled,omitempty" tf:"sso_debug_enabled,omitempty"`

	// Unique string that identifies the intended audience of the SAML assertion.
	SsoURL *string `json:"ssoUrl,omitempty" tf:"sso_url,omitempty"`

	// String enum that indicates whether the identity provider is active or not. Accepted values are ACTIVE or INACTIVE.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Identifier of the claim which contains the user ID in the token used for OIDC IdPs.
	UserClaim *string `json:"userClaim,omitempty" tf:"user_claim,omitempty"`
}

type SettingsIdentityProviderObservation struct {

	// List that contains the domains associated with the identity provider.
	AssociatedDomains []*string `json:"associatedDomains,omitempty" tf:"associated_domains,omitempty"`

	// Identifier of the intended recipient of the token used in OIDC IdP.
	AudienceClaim []*string `json:"audienceClaim,omitempty" tf:"audience_claim,omitempty"`

	// Client identifier that is assigned to an application by the OIDC Identity Provider.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
	FederationSettingsID *string `json:"federationSettingsId,omitempty" tf:"federation_settings_id,omitempty"`

	// Identifier of the claim which contains OIDC IdP Group IDs in the token.
	GroupsClaim *string `json:"groupsClaim,omitempty" tf:"groups_claim,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the IdP.
	IdpID *string `json:"idpId,omitempty" tf:"idp_id,omitempty"`

	// Unique string that identifies the issuer of the IdP.
	IssuerURI *string `json:"issuerUri,omitempty" tf:"issuer_uri,omitempty"`

	// Human-readable label that identifies the identity provider.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Unique 20-hexadecimal digit string that identifies the IdP.
	OktaIdpID *string `json:"oktaIdpId,omitempty" tf:"okta_idp_id,omitempty"`

	// The protocol of the identity provider. Either SAML or OIDC.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// SAML Authentication Request Protocol HTTP method binding (POST or REDIRECT) that Federated Authentication uses to send the authentication request. Atlas supports the following binding values:
	RequestBinding *string `json:"requestBinding,omitempty" tf:"request_binding,omitempty"`

	// Scopes that MongoDB applications will request from the authorization endpoint used for OIDC IdPs.
	RequestedScopes []*string `json:"requestedScopes,omitempty" tf:"requested_scopes,omitempty"`

	// Signature algorithm that Federated Authentication uses to encrypt the identity provider signature.  Valid values include SHA-1 and SHA-256.
	ResponseSignatureAlgorithm *string `json:"responseSignatureAlgorithm,omitempty" tf:"response_signature_algorithm,omitempty"`

	// Flag that indicates whether the identity provider has SSO debug enabled.
	SsoDebugEnabled *bool `json:"ssoDebugEnabled,omitempty" tf:"sso_debug_enabled,omitempty"`

	// Unique string that identifies the intended audience of the SAML assertion.
	SsoURL *string `json:"ssoUrl,omitempty" tf:"sso_url,omitempty"`

	// String enum that indicates whether the identity provider is active or not. Accepted values are ACTIVE or INACTIVE.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Identifier of the claim which contains the user ID in the token used for OIDC IdPs.
	UserClaim *string `json:"userClaim,omitempty" tf:"user_claim,omitempty"`
}

type SettingsIdentityProviderParameters struct {

	// List that contains the domains associated with the identity provider.
	// +kubebuilder:validation:Optional
	AssociatedDomains []*string `json:"associatedDomains,omitempty" tf:"associated_domains,omitempty"`

	// Identifier of the intended recipient of the token used in OIDC IdP.
	// +kubebuilder:validation:Optional
	AudienceClaim []*string `json:"audienceClaim,omitempty" tf:"audience_claim,omitempty"`

	// Client identifier that is assigned to an application by the OIDC Identity Provider.
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
	// +kubebuilder:validation:Optional
	FederationSettingsID *string `json:"federationSettingsId,omitempty" tf:"federation_settings_id,omitempty"`

	// Identifier of the claim which contains OIDC IdP Group IDs in the token.
	// +kubebuilder:validation:Optional
	GroupsClaim *string `json:"groupsClaim,omitempty" tf:"groups_claim,omitempty"`

	// Unique string that identifies the issuer of the IdP.
	// +kubebuilder:validation:Optional
	IssuerURI *string `json:"issuerUri,omitempty" tf:"issuer_uri,omitempty"`

	// Human-readable label that identifies the identity provider.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The protocol of the identity provider. Either SAML or OIDC.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// SAML Authentication Request Protocol HTTP method binding (POST or REDIRECT) that Federated Authentication uses to send the authentication request. Atlas supports the following binding values:
	// +kubebuilder:validation:Optional
	RequestBinding *string `json:"requestBinding,omitempty" tf:"request_binding,omitempty"`

	// Scopes that MongoDB applications will request from the authorization endpoint used for OIDC IdPs.
	// +kubebuilder:validation:Optional
	RequestedScopes []*string `json:"requestedScopes,omitempty" tf:"requested_scopes,omitempty"`

	// Signature algorithm that Federated Authentication uses to encrypt the identity provider signature.  Valid values include SHA-1 and SHA-256.
	// +kubebuilder:validation:Optional
	ResponseSignatureAlgorithm *string `json:"responseSignatureAlgorithm,omitempty" tf:"response_signature_algorithm,omitempty"`

	// Flag that indicates whether the identity provider has SSO debug enabled.
	// +kubebuilder:validation:Optional
	SsoDebugEnabled *bool `json:"ssoDebugEnabled,omitempty" tf:"sso_debug_enabled,omitempty"`

	// Unique string that identifies the intended audience of the SAML assertion.
	// +kubebuilder:validation:Optional
	SsoURL *string `json:"ssoUrl,omitempty" tf:"sso_url,omitempty"`

	// String enum that indicates whether the identity provider is active or not. Accepted values are ACTIVE or INACTIVE.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Identifier of the claim which contains the user ID in the token used for OIDC IdPs.
	// +kubebuilder:validation:Optional
	UserClaim *string `json:"userClaim,omitempty" tf:"user_claim,omitempty"`
}

// SettingsIdentityProviderSpec defines the desired state of SettingsIdentityProvider
type SettingsIdentityProviderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SettingsIdentityProviderParameters `json:"forProvider"`
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
	InitProvider SettingsIdentityProviderInitParameters `json:"initProvider,omitempty"`
}

// SettingsIdentityProviderStatus defines the observed state of SettingsIdentityProvider.
type SettingsIdentityProviderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SettingsIdentityProviderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SettingsIdentityProvider is the Schema for the SettingsIdentityProviders API. Provides a federated settings Identity Provider resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlas}
type SettingsIdentityProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.federationSettingsId) || (has(self.initProvider) && has(self.initProvider.federationSettingsId))",message="spec.forProvider.federationSettingsId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.issuerUri) || (has(self.initProvider) && has(self.initProvider.issuerUri))",message="spec.forProvider.issuerUri is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   SettingsIdentityProviderSpec   `json:"spec"`
	Status SettingsIdentityProviderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SettingsIdentityProviderList contains a list of SettingsIdentityProviders
type SettingsIdentityProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SettingsIdentityProvider `json:"items"`
}

// Repository type metadata.
var (
	SettingsIdentityProvider_Kind             = "SettingsIdentityProvider"
	SettingsIdentityProvider_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SettingsIdentityProvider_Kind}.String()
	SettingsIdentityProvider_KindAPIVersion   = SettingsIdentityProvider_Kind + "." + CRDGroupVersion.String()
	SettingsIdentityProvider_GroupVersionKind = CRDGroupVersion.WithKind(SettingsIdentityProvider_Kind)
)

func init() {
	SchemeBuilder.Register(&SettingsIdentityProvider{}, &SettingsIdentityProviderList{})
}

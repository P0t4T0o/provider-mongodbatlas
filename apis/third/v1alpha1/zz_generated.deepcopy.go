//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartyIntegration) DeepCopyInto(out *PartyIntegration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartyIntegration.
func (in *PartyIntegration) DeepCopy() *PartyIntegration {
	if in == nil {
		return nil
	}
	out := new(PartyIntegration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PartyIntegration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartyIntegrationInitParameters) DeepCopyInto(out *PartyIntegrationInitParameters) {
	*out = *in
	if in.APIKeySecretRef != nil {
		in, out := &in.APIKeySecretRef, &out.APIKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.ChannelName != nil {
		in, out := &in.ChannelName, &out.ChannelName
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.MicrosoftTeamsWebhookURLSecretRef != nil {
		in, out := &in.MicrosoftTeamsWebhookURLSecretRef, &out.MicrosoftTeamsWebhookURLSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.PasswordSecretRef != nil {
		in, out := &in.PasswordSecretRef, &out.PasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RoutingKeySecretRef != nil {
		in, out := &in.RoutingKeySecretRef, &out.RoutingKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.Scheme != nil {
		in, out := &in.Scheme, &out.Scheme
		*out = new(string)
		**out = **in
	}
	if in.SecretSecretRef != nil {
		in, out := &in.SecretSecretRef, &out.SecretSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.ServiceDiscoverySecretRef != nil {
		in, out := &in.ServiceDiscoverySecretRef, &out.ServiceDiscoverySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.ServiceKeySecretRef != nil {
		in, out := &in.ServiceKeySecretRef, &out.ServiceKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.TeamName != nil {
		in, out := &in.TeamName, &out.TeamName
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
	if in.UserNameSecretRef != nil {
		in, out := &in.UserNameSecretRef, &out.UserNameSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartyIntegrationInitParameters.
func (in *PartyIntegrationInitParameters) DeepCopy() *PartyIntegrationInitParameters {
	if in == nil {
		return nil
	}
	out := new(PartyIntegrationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartyIntegrationList) DeepCopyInto(out *PartyIntegrationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PartyIntegration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartyIntegrationList.
func (in *PartyIntegrationList) DeepCopy() *PartyIntegrationList {
	if in == nil {
		return nil
	}
	out := new(PartyIntegrationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PartyIntegrationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartyIntegrationObservation) DeepCopyInto(out *PartyIntegrationObservation) {
	*out = *in
	if in.ChannelName != nil {
		in, out := &in.ChannelName, &out.ChannelName
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Scheme != nil {
		in, out := &in.Scheme, &out.Scheme
		*out = new(string)
		**out = **in
	}
	if in.TeamName != nil {
		in, out := &in.TeamName, &out.TeamName
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartyIntegrationObservation.
func (in *PartyIntegrationObservation) DeepCopy() *PartyIntegrationObservation {
	if in == nil {
		return nil
	}
	out := new(PartyIntegrationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartyIntegrationParameters) DeepCopyInto(out *PartyIntegrationParameters) {
	*out = *in
	if in.APIKeySecretRef != nil {
		in, out := &in.APIKeySecretRef, &out.APIKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.ChannelName != nil {
		in, out := &in.ChannelName, &out.ChannelName
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.MicrosoftTeamsWebhookURLSecretRef != nil {
		in, out := &in.MicrosoftTeamsWebhookURLSecretRef, &out.MicrosoftTeamsWebhookURLSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.PasswordSecretRef != nil {
		in, out := &in.PasswordSecretRef, &out.PasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RoutingKeySecretRef != nil {
		in, out := &in.RoutingKeySecretRef, &out.RoutingKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.Scheme != nil {
		in, out := &in.Scheme, &out.Scheme
		*out = new(string)
		**out = **in
	}
	if in.SecretSecretRef != nil {
		in, out := &in.SecretSecretRef, &out.SecretSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.ServiceDiscoverySecretRef != nil {
		in, out := &in.ServiceDiscoverySecretRef, &out.ServiceDiscoverySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.ServiceKeySecretRef != nil {
		in, out := &in.ServiceKeySecretRef, &out.ServiceKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.TeamName != nil {
		in, out := &in.TeamName, &out.TeamName
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
	if in.UserNameSecretRef != nil {
		in, out := &in.UserNameSecretRef, &out.UserNameSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartyIntegrationParameters.
func (in *PartyIntegrationParameters) DeepCopy() *PartyIntegrationParameters {
	if in == nil {
		return nil
	}
	out := new(PartyIntegrationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartyIntegrationSpec) DeepCopyInto(out *PartyIntegrationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartyIntegrationSpec.
func (in *PartyIntegrationSpec) DeepCopy() *PartyIntegrationSpec {
	if in == nil {
		return nil
	}
	out := new(PartyIntegrationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartyIntegrationStatus) DeepCopyInto(out *PartyIntegrationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartyIntegrationStatus.
func (in *PartyIntegrationStatus) DeepCopy() *PartyIntegrationStatus {
	if in == nil {
		return nil
	}
	out := new(PartyIntegrationStatus)
	in.DeepCopyInto(out)
	return out
}

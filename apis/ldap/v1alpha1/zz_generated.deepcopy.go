//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Configuration) DeepCopyInto(out *Configuration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configuration.
func (in *Configuration) DeepCopy() *Configuration {
	if in == nil {
		return nil
	}
	out := new(Configuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Configuration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationInitParameters) DeepCopyInto(out *ConfigurationInitParameters) {
	*out = *in
	if in.AuthenticationEnabled != nil {
		in, out := &in.AuthenticationEnabled, &out.AuthenticationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AuthorizationEnabled != nil {
		in, out := &in.AuthorizationEnabled, &out.AuthorizationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AuthzQueryTemplate != nil {
		in, out := &in.AuthzQueryTemplate, &out.AuthzQueryTemplate
		*out = new(string)
		**out = **in
	}
	out.BindPasswordSecretRef = in.BindPasswordSecretRef
	if in.BindUsername != nil {
		in, out := &in.BindUsername, &out.BindUsername
		*out = new(string)
		**out = **in
	}
	if in.CACertificate != nil {
		in, out := &in.CACertificate, &out.CACertificate
		*out = new(string)
		**out = **in
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.UserToDnMapping != nil {
		in, out := &in.UserToDnMapping, &out.UserToDnMapping
		*out = make([]UserToDnMappingInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationInitParameters.
func (in *ConfigurationInitParameters) DeepCopy() *ConfigurationInitParameters {
	if in == nil {
		return nil
	}
	out := new(ConfigurationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationList) DeepCopyInto(out *ConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Configuration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationList.
func (in *ConfigurationList) DeepCopy() *ConfigurationList {
	if in == nil {
		return nil
	}
	out := new(ConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationObservation) DeepCopyInto(out *ConfigurationObservation) {
	*out = *in
	if in.AuthenticationEnabled != nil {
		in, out := &in.AuthenticationEnabled, &out.AuthenticationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AuthorizationEnabled != nil {
		in, out := &in.AuthorizationEnabled, &out.AuthorizationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AuthzQueryTemplate != nil {
		in, out := &in.AuthzQueryTemplate, &out.AuthzQueryTemplate
		*out = new(string)
		**out = **in
	}
	if in.BindUsername != nil {
		in, out := &in.BindUsername, &out.BindUsername
		*out = new(string)
		**out = **in
	}
	if in.CACertificate != nil {
		in, out := &in.CACertificate, &out.CACertificate
		*out = new(string)
		**out = **in
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.UserToDnMapping != nil {
		in, out := &in.UserToDnMapping, &out.UserToDnMapping
		*out = make([]UserToDnMappingObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationObservation.
func (in *ConfigurationObservation) DeepCopy() *ConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(ConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationParameters) DeepCopyInto(out *ConfigurationParameters) {
	*out = *in
	if in.AuthenticationEnabled != nil {
		in, out := &in.AuthenticationEnabled, &out.AuthenticationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AuthorizationEnabled != nil {
		in, out := &in.AuthorizationEnabled, &out.AuthorizationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AuthzQueryTemplate != nil {
		in, out := &in.AuthzQueryTemplate, &out.AuthzQueryTemplate
		*out = new(string)
		**out = **in
	}
	out.BindPasswordSecretRef = in.BindPasswordSecretRef
	if in.BindUsername != nil {
		in, out := &in.BindUsername, &out.BindUsername
		*out = new(string)
		**out = **in
	}
	if in.CACertificate != nil {
		in, out := &in.CACertificate, &out.CACertificate
		*out = new(string)
		**out = **in
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.UserToDnMapping != nil {
		in, out := &in.UserToDnMapping, &out.UserToDnMapping
		*out = make([]UserToDnMappingParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationParameters.
func (in *ConfigurationParameters) DeepCopy() *ConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(ConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationSpec) DeepCopyInto(out *ConfigurationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationSpec.
func (in *ConfigurationSpec) DeepCopy() *ConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationStatus) DeepCopyInto(out *ConfigurationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationStatus.
func (in *ConfigurationStatus) DeepCopy() *ConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(ConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinksInitParameters) DeepCopyInto(out *LinksInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinksInitParameters.
func (in *LinksInitParameters) DeepCopy() *LinksInitParameters {
	if in == nil {
		return nil
	}
	out := new(LinksInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinksObservation) DeepCopyInto(out *LinksObservation) {
	*out = *in
	if in.Href != nil {
		in, out := &in.Href, &out.Href
		*out = new(string)
		**out = **in
	}
	if in.Rel != nil {
		in, out := &in.Rel, &out.Rel
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinksObservation.
func (in *LinksObservation) DeepCopy() *LinksObservation {
	if in == nil {
		return nil
	}
	out := new(LinksObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinksParameters) DeepCopyInto(out *LinksParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinksParameters.
func (in *LinksParameters) DeepCopy() *LinksParameters {
	if in == nil {
		return nil
	}
	out := new(LinksParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserToDnMappingInitParameters) DeepCopyInto(out *UserToDnMappingInitParameters) {
	*out = *in
	if in.LdapQuery != nil {
		in, out := &in.LdapQuery, &out.LdapQuery
		*out = new(string)
		**out = **in
	}
	if in.Match != nil {
		in, out := &in.Match, &out.Match
		*out = new(string)
		**out = **in
	}
	if in.Substitution != nil {
		in, out := &in.Substitution, &out.Substitution
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserToDnMappingInitParameters.
func (in *UserToDnMappingInitParameters) DeepCopy() *UserToDnMappingInitParameters {
	if in == nil {
		return nil
	}
	out := new(UserToDnMappingInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserToDnMappingObservation) DeepCopyInto(out *UserToDnMappingObservation) {
	*out = *in
	if in.LdapQuery != nil {
		in, out := &in.LdapQuery, &out.LdapQuery
		*out = new(string)
		**out = **in
	}
	if in.Match != nil {
		in, out := &in.Match, &out.Match
		*out = new(string)
		**out = **in
	}
	if in.Substitution != nil {
		in, out := &in.Substitution, &out.Substitution
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserToDnMappingObservation.
func (in *UserToDnMappingObservation) DeepCopy() *UserToDnMappingObservation {
	if in == nil {
		return nil
	}
	out := new(UserToDnMappingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserToDnMappingParameters) DeepCopyInto(out *UserToDnMappingParameters) {
	*out = *in
	if in.LdapQuery != nil {
		in, out := &in.LdapQuery, &out.LdapQuery
		*out = new(string)
		**out = **in
	}
	if in.Match != nil {
		in, out := &in.Match, &out.Match
		*out = new(string)
		**out = **in
	}
	if in.Substitution != nil {
		in, out := &in.Substitution, &out.Substitution
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserToDnMappingParameters.
func (in *UserToDnMappingParameters) DeepCopy() *UserToDnMappingParameters {
	if in == nil {
		return nil
	}
	out := new(UserToDnMappingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidationsInitParameters) DeepCopyInto(out *ValidationsInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidationsInitParameters.
func (in *ValidationsInitParameters) DeepCopy() *ValidationsInitParameters {
	if in == nil {
		return nil
	}
	out := new(ValidationsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidationsObservation) DeepCopyInto(out *ValidationsObservation) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.ValidationType != nil {
		in, out := &in.ValidationType, &out.ValidationType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidationsObservation.
func (in *ValidationsObservation) DeepCopy() *ValidationsObservation {
	if in == nil {
		return nil
	}
	out := new(ValidationsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidationsParameters) DeepCopyInto(out *ValidationsParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidationsParameters.
func (in *ValidationsParameters) DeepCopy() *ValidationsParameters {
	if in == nil {
		return nil
	}
	out := new(ValidationsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Verify) DeepCopyInto(out *Verify) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Verify.
func (in *Verify) DeepCopy() *Verify {
	if in == nil {
		return nil
	}
	out := new(Verify)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Verify) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerifyInitParameters) DeepCopyInto(out *VerifyInitParameters) {
	*out = *in
	if in.AuthzQueryTemplate != nil {
		in, out := &in.AuthzQueryTemplate, &out.AuthzQueryTemplate
		*out = new(string)
		**out = **in
	}
	if in.BindPassword != nil {
		in, out := &in.BindPassword, &out.BindPassword
		*out = new(string)
		**out = **in
	}
	if in.BindUsername != nil {
		in, out := &in.BindUsername, &out.BindUsername
		*out = new(string)
		**out = **in
	}
	if in.CACertificate != nil {
		in, out := &in.CACertificate, &out.CACertificate
		*out = new(string)
		**out = **in
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerifyInitParameters.
func (in *VerifyInitParameters) DeepCopy() *VerifyInitParameters {
	if in == nil {
		return nil
	}
	out := new(VerifyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerifyList) DeepCopyInto(out *VerifyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Verify, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerifyList.
func (in *VerifyList) DeepCopy() *VerifyList {
	if in == nil {
		return nil
	}
	out := new(VerifyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VerifyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerifyObservation) DeepCopyInto(out *VerifyObservation) {
	*out = *in
	if in.AuthzQueryTemplate != nil {
		in, out := &in.AuthzQueryTemplate, &out.AuthzQueryTemplate
		*out = new(string)
		**out = **in
	}
	if in.BindPassword != nil {
		in, out := &in.BindPassword, &out.BindPassword
		*out = new(string)
		**out = **in
	}
	if in.BindUsername != nil {
		in, out := &in.BindUsername, &out.BindUsername
		*out = new(string)
		**out = **in
	}
	if in.CACertificate != nil {
		in, out := &in.CACertificate, &out.CACertificate
		*out = new(string)
		**out = **in
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = make([]LinksObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.RequestID != nil {
		in, out := &in.RequestID, &out.RequestID
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Validations != nil {
		in, out := &in.Validations, &out.Validations
		*out = make([]ValidationsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerifyObservation.
func (in *VerifyObservation) DeepCopy() *VerifyObservation {
	if in == nil {
		return nil
	}
	out := new(VerifyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerifyParameters) DeepCopyInto(out *VerifyParameters) {
	*out = *in
	if in.AuthzQueryTemplate != nil {
		in, out := &in.AuthzQueryTemplate, &out.AuthzQueryTemplate
		*out = new(string)
		**out = **in
	}
	if in.BindPassword != nil {
		in, out := &in.BindPassword, &out.BindPassword
		*out = new(string)
		**out = **in
	}
	if in.BindUsername != nil {
		in, out := &in.BindUsername, &out.BindUsername
		*out = new(string)
		**out = **in
	}
	if in.CACertificate != nil {
		in, out := &in.CACertificate, &out.CACertificate
		*out = new(string)
		**out = **in
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerifyParameters.
func (in *VerifyParameters) DeepCopy() *VerifyParameters {
	if in == nil {
		return nil
	}
	out := new(VerifyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerifySpec) DeepCopyInto(out *VerifySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerifySpec.
func (in *VerifySpec) DeepCopy() *VerifySpec {
	if in == nil {
		return nil
	}
	out := new(VerifySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerifyStatus) DeepCopyInto(out *VerifyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerifyStatus.
func (in *VerifyStatus) DeepCopy() *VerifyStatus {
	if in == nil {
		return nil
	}
	out := new(VerifyStatus)
	in.DeepCopyInto(out)
	return out
}

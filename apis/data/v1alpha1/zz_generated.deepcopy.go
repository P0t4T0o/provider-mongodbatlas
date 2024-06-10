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
func (in *IngestionSchedulesInitParameters) DeepCopyInto(out *IngestionSchedulesInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngestionSchedulesInitParameters.
func (in *IngestionSchedulesInitParameters) DeepCopy() *IngestionSchedulesInitParameters {
	if in == nil {
		return nil
	}
	out := new(IngestionSchedulesInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngestionSchedulesObservation) DeepCopyInto(out *IngestionSchedulesObservation) {
	*out = *in
	if in.FrequencyInterval != nil {
		in, out := &in.FrequencyInterval, &out.FrequencyInterval
		*out = new(float64)
		**out = **in
	}
	if in.FrequencyType != nil {
		in, out := &in.FrequencyType, &out.FrequencyType
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.RetentionUnit != nil {
		in, out := &in.RetentionUnit, &out.RetentionUnit
		*out = new(string)
		**out = **in
	}
	if in.RetentionValue != nil {
		in, out := &in.RetentionValue, &out.RetentionValue
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngestionSchedulesObservation.
func (in *IngestionSchedulesObservation) DeepCopy() *IngestionSchedulesObservation {
	if in == nil {
		return nil
	}
	out := new(IngestionSchedulesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngestionSchedulesParameters) DeepCopyInto(out *IngestionSchedulesParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngestionSchedulesParameters.
func (in *IngestionSchedulesParameters) DeepCopy() *IngestionSchedulesParameters {
	if in == nil {
		return nil
	}
	out := new(IngestionSchedulesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LakePipeline) DeepCopyInto(out *LakePipeline) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LakePipeline.
func (in *LakePipeline) DeepCopy() *LakePipeline {
	if in == nil {
		return nil
	}
	out := new(LakePipeline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LakePipeline) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LakePipelineInitParameters) DeepCopyInto(out *LakePipelineInitParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.ProjectIDRef != nil {
		in, out := &in.ProjectIDRef, &out.ProjectIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ProjectIDSelector != nil {
		in, out := &in.ProjectIDSelector, &out.ProjectIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Sink != nil {
		in, out := &in.Sink, &out.Sink
		*out = make([]SinkInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = make([]SourceInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Transformations != nil {
		in, out := &in.Transformations, &out.Transformations
		*out = make([]TransformationsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LakePipelineInitParameters.
func (in *LakePipelineInitParameters) DeepCopy() *LakePipelineInitParameters {
	if in == nil {
		return nil
	}
	out := new(LakePipelineInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LakePipelineList) DeepCopyInto(out *LakePipelineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LakePipeline, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LakePipelineList.
func (in *LakePipelineList) DeepCopy() *LakePipelineList {
	if in == nil {
		return nil
	}
	out := new(LakePipelineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LakePipelineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LakePipelineObservation) DeepCopyInto(out *LakePipelineObservation) {
	*out = *in
	if in.CreatedDate != nil {
		in, out := &in.CreatedDate, &out.CreatedDate
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IngestionSchedules != nil {
		in, out := &in.IngestionSchedules, &out.IngestionSchedules
		*out = make([]IngestionSchedulesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastUpdatedDate != nil {
		in, out := &in.LastUpdatedDate, &out.LastUpdatedDate
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.Sink != nil {
		in, out := &in.Sink, &out.Sink
		*out = make([]SinkObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Snapshots != nil {
		in, out := &in.Snapshots, &out.Snapshots
		*out = make([]SnapshotsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = make([]SourceObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.Transformations != nil {
		in, out := &in.Transformations, &out.Transformations
		*out = make([]TransformationsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LakePipelineObservation.
func (in *LakePipelineObservation) DeepCopy() *LakePipelineObservation {
	if in == nil {
		return nil
	}
	out := new(LakePipelineObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LakePipelineParameters) DeepCopyInto(out *LakePipelineParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.ProjectIDRef != nil {
		in, out := &in.ProjectIDRef, &out.ProjectIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ProjectIDSelector != nil {
		in, out := &in.ProjectIDSelector, &out.ProjectIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Sink != nil {
		in, out := &in.Sink, &out.Sink
		*out = make([]SinkParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = make([]SourceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Transformations != nil {
		in, out := &in.Transformations, &out.Transformations
		*out = make([]TransformationsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LakePipelineParameters.
func (in *LakePipelineParameters) DeepCopy() *LakePipelineParameters {
	if in == nil {
		return nil
	}
	out := new(LakePipelineParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LakePipelineSpec) DeepCopyInto(out *LakePipelineSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LakePipelineSpec.
func (in *LakePipelineSpec) DeepCopy() *LakePipelineSpec {
	if in == nil {
		return nil
	}
	out := new(LakePipelineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LakePipelineStatus) DeepCopyInto(out *LakePipelineStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LakePipelineStatus.
func (in *LakePipelineStatus) DeepCopy() *LakePipelineStatus {
	if in == nil {
		return nil
	}
	out := new(LakePipelineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionFieldsInitParameters) DeepCopyInto(out *PartitionFieldsInitParameters) {
	*out = *in
	if in.FieldName != nil {
		in, out := &in.FieldName, &out.FieldName
		*out = new(string)
		**out = **in
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionFieldsInitParameters.
func (in *PartitionFieldsInitParameters) DeepCopy() *PartitionFieldsInitParameters {
	if in == nil {
		return nil
	}
	out := new(PartitionFieldsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionFieldsObservation) DeepCopyInto(out *PartitionFieldsObservation) {
	*out = *in
	if in.FieldName != nil {
		in, out := &in.FieldName, &out.FieldName
		*out = new(string)
		**out = **in
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionFieldsObservation.
func (in *PartitionFieldsObservation) DeepCopy() *PartitionFieldsObservation {
	if in == nil {
		return nil
	}
	out := new(PartitionFieldsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionFieldsParameters) DeepCopyInto(out *PartitionFieldsParameters) {
	*out = *in
	if in.FieldName != nil {
		in, out := &in.FieldName, &out.FieldName
		*out = new(string)
		**out = **in
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionFieldsParameters.
func (in *PartitionFieldsParameters) DeepCopy() *PartitionFieldsParameters {
	if in == nil {
		return nil
	}
	out := new(PartitionFieldsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SinkInitParameters) DeepCopyInto(out *SinkInitParameters) {
	*out = *in
	if in.PartitionFields != nil {
		in, out := &in.PartitionFields, &out.PartitionFields
		*out = make([]PartitionFieldsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Provider != nil {
		in, out := &in.Provider, &out.Provider
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SinkInitParameters.
func (in *SinkInitParameters) DeepCopy() *SinkInitParameters {
	if in == nil {
		return nil
	}
	out := new(SinkInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SinkObservation) DeepCopyInto(out *SinkObservation) {
	*out = *in
	if in.PartitionFields != nil {
		in, out := &in.PartitionFields, &out.PartitionFields
		*out = make([]PartitionFieldsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Provider != nil {
		in, out := &in.Provider, &out.Provider
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SinkObservation.
func (in *SinkObservation) DeepCopy() *SinkObservation {
	if in == nil {
		return nil
	}
	out := new(SinkObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SinkParameters) DeepCopyInto(out *SinkParameters) {
	*out = *in
	if in.PartitionFields != nil {
		in, out := &in.PartitionFields, &out.PartitionFields
		*out = make([]PartitionFieldsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Provider != nil {
		in, out := &in.Provider, &out.Provider
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SinkParameters.
func (in *SinkParameters) DeepCopy() *SinkParameters {
	if in == nil {
		return nil
	}
	out := new(SinkParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotsInitParameters) DeepCopyInto(out *SnapshotsInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotsInitParameters.
func (in *SnapshotsInitParameters) DeepCopy() *SnapshotsInitParameters {
	if in == nil {
		return nil
	}
	out := new(SnapshotsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotsObservation) DeepCopyInto(out *SnapshotsObservation) {
	*out = *in
	if in.CopyRegion != nil {
		in, out := &in.CopyRegion, &out.CopyRegion
		*out = new(string)
		**out = **in
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.ExpiresAt != nil {
		in, out := &in.ExpiresAt, &out.ExpiresAt
		*out = new(string)
		**out = **in
	}
	if in.FrequencyYype != nil {
		in, out := &in.FrequencyYype, &out.FrequencyYype
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.MasterKey != nil {
		in, out := &in.MasterKey, &out.MasterKey
		*out = new(string)
		**out = **in
	}
	if in.MongodVersion != nil {
		in, out := &in.MongodVersion, &out.MongodVersion
		*out = new(string)
		**out = **in
	}
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Provider != nil {
		in, out := &in.Provider, &out.Provider
		*out = new(string)
		**out = **in
	}
	if in.ReplicaSetName != nil {
		in, out := &in.ReplicaSetName, &out.ReplicaSetName
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.SnapshotType != nil {
		in, out := &in.SnapshotType, &out.SnapshotType
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotsObservation.
func (in *SnapshotsObservation) DeepCopy() *SnapshotsObservation {
	if in == nil {
		return nil
	}
	out := new(SnapshotsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotsParameters) DeepCopyInto(out *SnapshotsParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotsParameters.
func (in *SnapshotsParameters) DeepCopy() *SnapshotsParameters {
	if in == nil {
		return nil
	}
	out := new(SnapshotsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceInitParameters) DeepCopyInto(out *SourceInitParameters) {
	*out = *in
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.ClusterNameRef != nil {
		in, out := &in.ClusterNameRef, &out.ClusterNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterNameSelector != nil {
		in, out := &in.ClusterNameSelector, &out.ClusterNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.CollectionName != nil {
		in, out := &in.CollectionName, &out.CollectionName
		*out = new(string)
		**out = **in
	}
	if in.DatabaseName != nil {
		in, out := &in.DatabaseName, &out.DatabaseName
		*out = new(string)
		**out = **in
	}
	if in.PolicyItemID != nil {
		in, out := &in.PolicyItemID, &out.PolicyItemID
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceInitParameters.
func (in *SourceInitParameters) DeepCopy() *SourceInitParameters {
	if in == nil {
		return nil
	}
	out := new(SourceInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceObservation) DeepCopyInto(out *SourceObservation) {
	*out = *in
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.CollectionName != nil {
		in, out := &in.CollectionName, &out.CollectionName
		*out = new(string)
		**out = **in
	}
	if in.DatabaseName != nil {
		in, out := &in.DatabaseName, &out.DatabaseName
		*out = new(string)
		**out = **in
	}
	if in.PolicyItemID != nil {
		in, out := &in.PolicyItemID, &out.PolicyItemID
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceObservation.
func (in *SourceObservation) DeepCopy() *SourceObservation {
	if in == nil {
		return nil
	}
	out := new(SourceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceParameters) DeepCopyInto(out *SourceParameters) {
	*out = *in
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.ClusterNameRef != nil {
		in, out := &in.ClusterNameRef, &out.ClusterNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterNameSelector != nil {
		in, out := &in.ClusterNameSelector, &out.ClusterNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.CollectionName != nil {
		in, out := &in.CollectionName, &out.CollectionName
		*out = new(string)
		**out = **in
	}
	if in.DatabaseName != nil {
		in, out := &in.DatabaseName, &out.DatabaseName
		*out = new(string)
		**out = **in
	}
	if in.PolicyItemID != nil {
		in, out := &in.PolicyItemID, &out.PolicyItemID
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceParameters.
func (in *SourceParameters) DeepCopy() *SourceParameters {
	if in == nil {
		return nil
	}
	out := new(SourceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransformationsInitParameters) DeepCopyInto(out *TransformationsInitParameters) {
	*out = *in
	if in.Field != nil {
		in, out := &in.Field, &out.Field
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransformationsInitParameters.
func (in *TransformationsInitParameters) DeepCopy() *TransformationsInitParameters {
	if in == nil {
		return nil
	}
	out := new(TransformationsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransformationsObservation) DeepCopyInto(out *TransformationsObservation) {
	*out = *in
	if in.Field != nil {
		in, out := &in.Field, &out.Field
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransformationsObservation.
func (in *TransformationsObservation) DeepCopy() *TransformationsObservation {
	if in == nil {
		return nil
	}
	out := new(TransformationsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransformationsParameters) DeepCopyInto(out *TransformationsParameters) {
	*out = *in
	if in.Field != nil {
		in, out := &in.Field, &out.Field
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransformationsParameters.
func (in *TransformationsParameters) DeepCopy() *TransformationsParameters {
	if in == nil {
		return nil
	}
	out := new(TransformationsParameters)
	in.DeepCopyInto(out)
	return out
}

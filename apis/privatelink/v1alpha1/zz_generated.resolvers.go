/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha11 "github.com/crossplane-contrib/provider-mongodb-atlas/apis/mongodbatlas/v1alpha1"
	v1alpha1 "github.com/crossplane-contrib/provider-mongodb-atlas/apis/serverless/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this EndpointServerless.
func (mg *EndpointServerless) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceName),
		Extract:      resource.ExtractParamPath("name", false),
		Reference:    mg.Spec.ForProvider.InstanceNameRef,
		Selector:     mg.Spec.ForProvider.InstanceNameSelector,
		To: reference.To{
			List:    &v1alpha1.InstanceList{},
			Managed: &v1alpha1.Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceName")
	}
	mg.Spec.ForProvider.InstanceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceName),
		Extract:      resource.ExtractParamPath("name", false),
		Reference:    mg.Spec.InitProvider.InstanceNameRef,
		Selector:     mg.Spec.InitProvider.InstanceNameSelector,
		To: reference.To{
			List:    &v1alpha1.InstanceList{},
			Managed: &v1alpha1.Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceName")
	}
	mg.Spec.InitProvider.InstanceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this EndpointServiceDataFederationOnlineArchive.
func (mg *EndpointServiceDataFederationOnlineArchive) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ProjectID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ProjectIDRef,
		Selector:     mg.Spec.ForProvider.ProjectIDSelector,
		To: reference.To{
			List:    &v1alpha11.ProjectList{},
			Managed: &v1alpha11.Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ProjectID")
	}
	mg.Spec.ForProvider.ProjectID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProjectIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ProjectID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.ProjectIDRef,
		Selector:     mg.Spec.InitProvider.ProjectIDSelector,
		To: reference.To{
			List:    &v1alpha11.ProjectList{},
			Managed: &v1alpha11.Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ProjectID")
	}
	mg.Spec.InitProvider.ProjectID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ProjectIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this EndpointServiceServerless.
func (mg *EndpointServiceServerless) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EndpointID),
		Extract:      resource.ExtractParamPath("endpoint_id", true),
		Reference:    mg.Spec.ForProvider.EndpointIDRef,
		Selector:     mg.Spec.ForProvider.EndpointIDSelector,
		To: reference.To{
			List:    &EndpointServerlessList{},
			Managed: &EndpointServerless{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EndpointID")
	}
	mg.Spec.ForProvider.EndpointID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EndpointIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceName),
		Extract:      resource.ExtractParamPath("name", false),
		Reference:    mg.Spec.ForProvider.InstanceNameRef,
		Selector:     mg.Spec.ForProvider.InstanceNameSelector,
		To: reference.To{
			List:    &v1alpha1.InstanceList{},
			Managed: &v1alpha1.Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceName")
	}
	mg.Spec.ForProvider.InstanceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ProjectID),
		Extract:      resource.ExtractParamPath("project_id", false),
		Reference:    mg.Spec.ForProvider.ProjectIDRef,
		Selector:     mg.Spec.ForProvider.ProjectIDSelector,
		To: reference.To{
			List:    &EndpointServerlessList{},
			Managed: &EndpointServerless{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ProjectID")
	}
	mg.Spec.ForProvider.ProjectID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProjectIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EndpointID),
		Extract:      resource.ExtractParamPath("endpoint_id", true),
		Reference:    mg.Spec.InitProvider.EndpointIDRef,
		Selector:     mg.Spec.InitProvider.EndpointIDSelector,
		To: reference.To{
			List:    &EndpointServerlessList{},
			Managed: &EndpointServerless{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.EndpointID")
	}
	mg.Spec.InitProvider.EndpointID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.EndpointIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceName),
		Extract:      resource.ExtractParamPath("name", false),
		Reference:    mg.Spec.InitProvider.InstanceNameRef,
		Selector:     mg.Spec.InitProvider.InstanceNameSelector,
		To: reference.To{
			List:    &v1alpha1.InstanceList{},
			Managed: &v1alpha1.Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceName")
	}
	mg.Spec.InitProvider.InstanceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ProjectID),
		Extract:      resource.ExtractParamPath("project_id", false),
		Reference:    mg.Spec.InitProvider.ProjectIDRef,
		Selector:     mg.Spec.InitProvider.ProjectIDSelector,
		To: reference.To{
			List:    &EndpointServerlessList{},
			Managed: &EndpointServerless{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ProjectID")
	}
	mg.Spec.InitProvider.ProjectID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ProjectIDRef = rsp.ResolvedReference

	return nil
}

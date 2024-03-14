// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *AlertManagerDefinition) ResolveReferences( // ResolveReferences of this AlertManagerDefinition.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("amp.aws.upbound.io", "v1beta1", "Workspace", "WorkspaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WorkspaceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.WorkspaceIDRef,
			Selector:     mg.Spec.ForProvider.WorkspaceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.WorkspaceID")
	}
	mg.Spec.ForProvider.WorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.WorkspaceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("amp.aws.upbound.io", "v1beta1", "Workspace", "WorkspaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.WorkspaceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.WorkspaceIDRef,
			Selector:     mg.Spec.InitProvider.WorkspaceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.WorkspaceID")
	}
	mg.Spec.InitProvider.WorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.WorkspaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RuleGroupNamespace.
func (mg *RuleGroupNamespace) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("amp.aws.upbound.io", "v1beta1", "Workspace", "WorkspaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WorkspaceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.WorkspaceIDRef,
			Selector:     mg.Spec.ForProvider.WorkspaceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.WorkspaceID")
	}
	mg.Spec.ForProvider.WorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.WorkspaceIDRef = rsp.ResolvedReference

	return nil
}

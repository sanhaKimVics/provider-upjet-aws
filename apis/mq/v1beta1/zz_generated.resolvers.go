// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Broker.
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func (mg *Broker) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Configuration); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("mq.aws.upbound.io", "v1beta1", "Configuration", "ConfigurationList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Configuration[i3].ID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Configuration[i3].IDRef,
				Selector:     mg.Spec.ForProvider.Configuration[i3].IDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Configuration[i3].ID")
		}
		mg.Spec.ForProvider.Configuration[i3].ID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Configuration[i3].IDRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "SecurityGroup", "SecurityGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroups),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.SecurityGroupRefs,
			Selector:      mg.Spec.ForProvider.SecurityGroupSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroups")
	}
	mg.Spec.ForProvider.SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SubnetIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.SubnetIDRefs,
			Selector:      mg.Spec.ForProvider.SubnetIDSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetIds")
	}
	mg.Spec.ForProvider.SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SubnetIDRefs = mrsp.ResolvedReferences

	for i3 := 0; i3 < len(mg.Spec.InitProvider.Configuration); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("mq.aws.upbound.io", "v1beta1", "Configuration", "ConfigurationList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Configuration[i3].ID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.Configuration[i3].IDRef,
				Selector:     mg.Spec.InitProvider.Configuration[i3].IDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Configuration[i3].ID")
		}
		mg.Spec.InitProvider.Configuration[i3].ID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Configuration[i3].IDRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "SecurityGroup", "SecurityGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.SecurityGroups),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.SecurityGroupRefs,
			Selector:      mg.Spec.InitProvider.SecurityGroupSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SecurityGroups")
	}
	mg.Spec.InitProvider.SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.SecurityGroupRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.SubnetIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.SubnetIDRefs,
			Selector:      mg.Spec.InitProvider.SubnetIDSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SubnetIds")
	}
	mg.Spec.InitProvider.SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.SubnetIDRefs = mrsp.ResolvedReferences

	return nil
}

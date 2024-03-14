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
	errors "github.com/pkg/errors"
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *CertificateValidation) ResolveReferences( // ResolveReferences of this CertificateValidation.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("acm.aws.upbound.io", "v1beta1", "Certificate", "CertificateList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CertificateArn),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.CertificateArnRef,
			Selector:     mg.Spec.ForProvider.CertificateArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CertificateArn")
	}
	mg.Spec.ForProvider.CertificateArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CertificateArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("acm.aws.upbound.io", "v1beta1", "Certificate", "CertificateList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CertificateArn),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.CertificateArnRef,
			Selector:     mg.Spec.InitProvider.CertificateArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CertificateArn")
	}
	mg.Spec.InitProvider.CertificateArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CertificateArnRef = rsp.ResolvedReference

	return nil
}

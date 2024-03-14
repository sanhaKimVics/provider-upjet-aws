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
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *RateBasedRule) ResolveReferences( // ResolveReferences of this RateBasedRule.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Predicates); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("waf.aws.upbound.io", "v1beta1", "IPSet", "IPSetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Predicates[i3].DataID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Predicates[i3].DataIDRef,
				Selector:     mg.Spec.ForProvider.Predicates[i3].DataIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Predicates[i3].DataID")
		}
		mg.Spec.ForProvider.Predicates[i3].DataID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Predicates[i3].DataIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Predicates); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("waf.aws.upbound.io", "v1beta1", "IPSet", "IPSetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Predicates[i3].DataID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.Predicates[i3].DataIDRef,
				Selector:     mg.Spec.InitProvider.Predicates[i3].DataIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Predicates[i3].DataID")
		}
		mg.Spec.InitProvider.Predicates[i3].DataID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Predicates[i3].DataIDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this RegexMatchSet.
func (mg *RegexMatchSet) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.RegexMatchTuple); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("waf.aws.upbound.io", "v1beta1", "RegexPatternSet", "RegexPatternSetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RegexMatchTuple[i3].RegexPatternSetID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.RegexMatchTuple[i3].RegexPatternSetIDRef,
				Selector:     mg.Spec.ForProvider.RegexMatchTuple[i3].RegexPatternSetIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.RegexMatchTuple[i3].RegexPatternSetID")
		}
		mg.Spec.ForProvider.RegexMatchTuple[i3].RegexPatternSetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.RegexMatchTuple[i3].RegexPatternSetIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.RegexMatchTuple); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("waf.aws.upbound.io", "v1beta1", "RegexPatternSet", "RegexPatternSetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RegexMatchTuple[i3].RegexPatternSetID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.RegexMatchTuple[i3].RegexPatternSetIDRef,
				Selector:     mg.Spec.InitProvider.RegexMatchTuple[i3].RegexPatternSetIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.RegexMatchTuple[i3].RegexPatternSetID")
		}
		mg.Spec.InitProvider.RegexMatchTuple[i3].RegexPatternSetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.RegexMatchTuple[i3].RegexPatternSetIDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Rule.
func (mg *Rule) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Predicates); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("waf.aws.upbound.io", "v1beta1", "IPSet", "IPSetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Predicates[i3].DataID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Predicates[i3].DataIDRef,
				Selector:     mg.Spec.ForProvider.Predicates[i3].DataIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Predicates[i3].DataID")
		}
		mg.Spec.ForProvider.Predicates[i3].DataID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Predicates[i3].DataIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Predicates); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("waf.aws.upbound.io", "v1beta1", "IPSet", "IPSetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Predicates[i3].DataID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.Predicates[i3].DataIDRef,
				Selector:     mg.Spec.InitProvider.Predicates[i3].DataIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Predicates[i3].DataID")
		}
		mg.Spec.InitProvider.Predicates[i3].DataID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Predicates[i3].DataIDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this WebACL.
func (mg *WebACL) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.LoggingConfiguration); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("firehose.aws.upbound.io", "v1beta1", "DeliveryStream", "DeliveryStreamList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LoggingConfiguration[i3].LogDestination),
				Extract:      resource.ExtractParamPath("arn", false),
				Reference:    mg.Spec.ForProvider.LoggingConfiguration[i3].LogDestinationRef,
				Selector:     mg.Spec.ForProvider.LoggingConfiguration[i3].LogDestinationSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.LoggingConfiguration[i3].LogDestination")
		}
		mg.Spec.ForProvider.LoggingConfiguration[i3].LogDestination = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.LoggingConfiguration[i3].LogDestinationRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Rules); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("waf.aws.upbound.io", "v1beta1", "Rule", "RuleList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Rules[i3].RuleID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Rules[i3].RuleIDRef,
				Selector:     mg.Spec.ForProvider.Rules[i3].RuleIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Rules[i3].RuleID")
		}
		mg.Spec.ForProvider.Rules[i3].RuleID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Rules[i3].RuleIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.LoggingConfiguration); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("firehose.aws.upbound.io", "v1beta1", "DeliveryStream", "DeliveryStreamList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LoggingConfiguration[i3].LogDestination),
				Extract:      resource.ExtractParamPath("arn", false),
				Reference:    mg.Spec.InitProvider.LoggingConfiguration[i3].LogDestinationRef,
				Selector:     mg.Spec.InitProvider.LoggingConfiguration[i3].LogDestinationSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.LoggingConfiguration[i3].LogDestination")
		}
		mg.Spec.InitProvider.LoggingConfiguration[i3].LogDestination = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.LoggingConfiguration[i3].LogDestinationRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Rules); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("waf.aws.upbound.io", "v1beta1", "Rule", "RuleList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Rules[i3].RuleID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.Rules[i3].RuleIDRef,
				Selector:     mg.Spec.InitProvider.Rules[i3].RuleIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Rules[i3].RuleID")
		}
		mg.Spec.InitProvider.Rules[i3].RuleID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Rules[i3].RuleIDRef = rsp.ResolvedReference

	}

	return nil
}

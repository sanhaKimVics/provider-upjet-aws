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
	common "github.com/upbound/provider-aws/config/common"
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *Activation) ResolveReferences( // ResolveReferences of this Activation.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IAMRole),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.IAMRoleRef,
			Selector:     mg.Spec.ForProvider.IAMRoleSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IAMRole")
	}
	mg.Spec.ForProvider.IAMRole = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IAMRoleRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.IAMRole),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.IAMRoleRef,
			Selector:     mg.Spec.InitProvider.IAMRoleSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.IAMRole")
	}
	mg.Spec.InitProvider.IAMRole = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.IAMRoleRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Association.
func (mg *Association) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("ssm.aws.upbound.io", "v1beta1", "Document", "DocumentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Name),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.NameRef,
			Selector:     mg.Spec.ForProvider.NameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Name")
	}
	mg.Spec.ForProvider.Name = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ssm.aws.upbound.io", "v1beta1", "Document", "DocumentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Name),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.NameRef,
			Selector:     mg.Spec.InitProvider.NameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Name")
	}
	mg.Spec.InitProvider.Name = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DefaultPatchBaseline.
func (mg *DefaultPatchBaseline) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("ssm.aws.upbound.io", "v1beta1", "PatchBaseline", "PatchBaselineList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BaselineID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.BaselineIDRef,
			Selector:     mg.Spec.ForProvider.BaselineIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.BaselineID")
	}
	mg.Spec.ForProvider.BaselineID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BaselineIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ssm.aws.upbound.io", "v1beta1", "PatchBaseline", "PatchBaselineList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OperatingSystem),
			Extract:      resource.ExtractParamPath("operating_system", false),
			Reference:    mg.Spec.ForProvider.OperatingSystemRef,
			Selector:     mg.Spec.ForProvider.OperatingSystemSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.OperatingSystem")
	}
	mg.Spec.ForProvider.OperatingSystem = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OperatingSystemRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ssm.aws.upbound.io", "v1beta1", "PatchBaseline", "PatchBaselineList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.BaselineID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.BaselineIDRef,
			Selector:     mg.Spec.InitProvider.BaselineIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.BaselineID")
	}
	mg.Spec.InitProvider.BaselineID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BaselineIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ssm.aws.upbound.io", "v1beta1", "PatchBaseline", "PatchBaselineList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.OperatingSystem),
			Extract:      resource.ExtractParamPath("operating_system", false),
			Reference:    mg.Spec.InitProvider.OperatingSystemRef,
			Selector:     mg.Spec.InitProvider.OperatingSystemSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.OperatingSystem")
	}
	mg.Spec.InitProvider.OperatingSystem = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.OperatingSystemRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this MaintenanceWindowTarget.
func (mg *MaintenanceWindowTarget) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("ssm.aws.upbound.io", "v1beta1", "MaintenanceWindow", "MaintenanceWindowList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WindowID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.WindowIDRef,
			Selector:     mg.Spec.ForProvider.WindowIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.WindowID")
	}
	mg.Spec.ForProvider.WindowID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.WindowIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ssm.aws.upbound.io", "v1beta1", "MaintenanceWindow", "MaintenanceWindowList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.WindowID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.WindowIDRef,
			Selector:     mg.Spec.InitProvider.WindowIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.WindowID")
	}
	mg.Spec.InitProvider.WindowID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.WindowIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this MaintenanceWindowTask.
func (mg *MaintenanceWindowTask) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.ServiceRoleArnRef,
			Selector:     mg.Spec.ForProvider.ServiceRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceRoleArn")
	}
	mg.Spec.ForProvider.ServiceRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceRoleArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("lambda.aws.upbound.io", "v1beta1", "Function", "FunctionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TaskArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.TaskArnRef,
			Selector:     mg.Spec.ForProvider.TaskArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TaskArn")
	}
	mg.Spec.ForProvider.TaskArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TaskArnRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.TaskInvocationParameters); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.TaskInvocationParameters[i3].RunCommandParameters); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].NotificationConfig); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("sns.aws.upbound.io", "v1beta1", "Topic", "TopicList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].NotificationConfig[i5].NotificationArn),
						Extract:      resource.ExtractParamPath("arn", true),
						Reference:    mg.Spec.ForProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].NotificationConfig[i5].NotificationArnRef,
						Selector:     mg.Spec.ForProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].NotificationConfig[i5].NotificationArnSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].NotificationConfig[i5].NotificationArn")
				}
				mg.Spec.ForProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].NotificationConfig[i5].NotificationArn = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].NotificationConfig[i5].NotificationArnRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.TaskInvocationParameters); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.TaskInvocationParameters[i3].RunCommandParameters); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Bucket", "BucketList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].OutputS3Bucket),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.ForProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].OutputS3BucketRef,
					Selector:     mg.Spec.ForProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].OutputS3BucketSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].OutputS3Bucket")
			}
			mg.Spec.ForProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].OutputS3Bucket = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].OutputS3BucketRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.TaskInvocationParameters); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.TaskInvocationParameters[i3].RunCommandParameters); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].ServiceRoleArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.ForProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].ServiceRoleArnRef,
					Selector:     mg.Spec.ForProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].ServiceRoleArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].ServiceRoleArn")
			}
			mg.Spec.ForProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].ServiceRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].ServiceRoleArnRef = rsp.ResolvedReference

		}
	}
	{
		m, l, err = apisresolver.GetManagedResource("ssm.aws.upbound.io", "v1beta1", "MaintenanceWindow", "MaintenanceWindowList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WindowID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.WindowIDRef,
			Selector:     mg.Spec.ForProvider.WindowIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.WindowID")
	}
	mg.Spec.ForProvider.WindowID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.WindowIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.ServiceRoleArnRef,
			Selector:     mg.Spec.InitProvider.ServiceRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServiceRoleArn")
	}
	mg.Spec.InitProvider.ServiceRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceRoleArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("lambda.aws.upbound.io", "v1beta1", "Function", "FunctionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.TaskArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.InitProvider.TaskArnRef,
			Selector:     mg.Spec.InitProvider.TaskArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.TaskArn")
	}
	mg.Spec.InitProvider.TaskArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TaskArnRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.TaskInvocationParameters); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.TaskInvocationParameters[i3].RunCommandParameters); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].NotificationConfig); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("sns.aws.upbound.io", "v1beta1", "Topic", "TopicList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].NotificationConfig[i5].NotificationArn),
						Extract:      resource.ExtractParamPath("arn", true),
						Reference:    mg.Spec.InitProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].NotificationConfig[i5].NotificationArnRef,
						Selector:     mg.Spec.InitProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].NotificationConfig[i5].NotificationArnSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].NotificationConfig[i5].NotificationArn")
				}
				mg.Spec.InitProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].NotificationConfig[i5].NotificationArn = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].NotificationConfig[i5].NotificationArnRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.TaskInvocationParameters); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.TaskInvocationParameters[i3].RunCommandParameters); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Bucket", "BucketList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].OutputS3Bucket),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.InitProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].OutputS3BucketRef,
					Selector:     mg.Spec.InitProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].OutputS3BucketSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].OutputS3Bucket")
			}
			mg.Spec.InitProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].OutputS3Bucket = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].OutputS3BucketRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.TaskInvocationParameters); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.TaskInvocationParameters[i3].RunCommandParameters); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].ServiceRoleArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.InitProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].ServiceRoleArnRef,
					Selector:     mg.Spec.InitProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].ServiceRoleArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].ServiceRoleArn")
			}
			mg.Spec.InitProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].ServiceRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.TaskInvocationParameters[i3].RunCommandParameters[i4].ServiceRoleArnRef = rsp.ResolvedReference

		}
	}
	{
		m, l, err = apisresolver.GetManagedResource("ssm.aws.upbound.io", "v1beta1", "MaintenanceWindow", "MaintenanceWindowList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.WindowID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.WindowIDRef,
			Selector:     mg.Spec.InitProvider.WindowIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.WindowID")
	}
	mg.Spec.InitProvider.WindowID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.WindowIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this PatchGroup.
func (mg *PatchGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("ssm.aws.upbound.io", "v1beta1", "PatchBaseline", "PatchBaselineList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BaselineID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.BaselineIDRef,
			Selector:     mg.Spec.ForProvider.BaselineIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.BaselineID")
	}
	mg.Spec.ForProvider.BaselineID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BaselineIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ssm.aws.upbound.io", "v1beta1", "PatchBaseline", "PatchBaselineList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.BaselineID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.BaselineIDRef,
			Selector:     mg.Spec.InitProvider.BaselineIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.BaselineID")
	}
	mg.Spec.InitProvider.BaselineID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BaselineIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ResourceDataSync.
func (mg *ResourceDataSync) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.S3Destination); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Bucket", "BucketList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.S3Destination[i3].BucketName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.S3Destination[i3].BucketNameRef,
				Selector:     mg.Spec.ForProvider.S3Destination[i3].BucketNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.S3Destination[i3].BucketName")
		}
		mg.Spec.ForProvider.S3Destination[i3].BucketName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.S3Destination[i3].BucketNameRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.S3Destination); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Bucket", "BucketList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.S3Destination[i3].Region),
				Extract:      resource.ExtractParamPath("region", false),
				Reference:    mg.Spec.ForProvider.S3Destination[i3].RegionRef,
				Selector:     mg.Spec.ForProvider.S3Destination[i3].RegionSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.S3Destination[i3].Region")
		}
		mg.Spec.ForProvider.S3Destination[i3].Region = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.S3Destination[i3].RegionRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.S3Destination); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Bucket", "BucketList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.S3Destination[i3].BucketName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.S3Destination[i3].BucketNameRef,
				Selector:     mg.Spec.InitProvider.S3Destination[i3].BucketNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.S3Destination[i3].BucketName")
		}
		mg.Spec.InitProvider.S3Destination[i3].BucketName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.S3Destination[i3].BucketNameRef = rsp.ResolvedReference

	}

	return nil
}

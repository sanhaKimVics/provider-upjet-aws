// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"dario.cat/mergo"
	"github.com/pkg/errors"

	"github.com/crossplane/upjet/pkg/resource"
	"github.com/crossplane/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this TransitGatewayVPCAttachmentAccepter
func (mg *TransitGatewayVPCAttachmentAccepter) GetTerraformResourceType() string {
	return "aws_ec2_transit_gateway_vpc_attachment_accepter"
}

// GetConnectionDetailsMapping for this TransitGatewayVPCAttachmentAccepter
func (tr *TransitGatewayVPCAttachmentAccepter) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this TransitGatewayVPCAttachmentAccepter
func (tr *TransitGatewayVPCAttachmentAccepter) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this TransitGatewayVPCAttachmentAccepter
func (tr *TransitGatewayVPCAttachmentAccepter) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this TransitGatewayVPCAttachmentAccepter
func (tr *TransitGatewayVPCAttachmentAccepter) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this TransitGatewayVPCAttachmentAccepter
func (tr *TransitGatewayVPCAttachmentAccepter) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this TransitGatewayVPCAttachmentAccepter
func (tr *TransitGatewayVPCAttachmentAccepter) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this TransitGatewayVPCAttachmentAccepter
func (tr *TransitGatewayVPCAttachmentAccepter) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// GetInitParameters of this TransitGatewayVPCAttachmentAccepter
func (tr *TransitGatewayVPCAttachmentAccepter) GetMergedParameters(shouldMergeInitProvider bool) (map[string]any, error) {
	params, err := tr.GetParameters()
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get parameters for resource '%q'", tr.GetName())
	}
	if !shouldMergeInitProvider {
		return params, nil
	}

	initParams, err := tr.GetInitParameters()
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get init parameters for resource '%q'", tr.GetName())
	}

	// Note(lsviben): mergo.WithSliceDeepCopy is needed to merge the
	// slices from the initProvider to forProvider. As it also sets
	// overwrite to true, we need to set it back to false, we don't
	// want to overwrite the forProvider fields with the initProvider
	// fields.
	err = mergo.Merge(&params, initParams, mergo.WithSliceDeepCopy, func(c *mergo.Config) {
		c.Overwrite = false
	})
	if err != nil {
		return nil, errors.Wrapf(err, "cannot merge spec.initProvider and spec.forProvider parameters for resource '%q'", tr.GetName())
	}

	return params, nil
}

// LateInitialize this TransitGatewayVPCAttachmentAccepter using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *TransitGatewayVPCAttachmentAccepter) LateInitialize(attrs []byte) (bool, error) {
	params := &TransitGatewayVPCAttachmentAccepterParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *TransitGatewayVPCAttachmentAccepter) GetTerraformSchemaVersion() int {
	return 0
}

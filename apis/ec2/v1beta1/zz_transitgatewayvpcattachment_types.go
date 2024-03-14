// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TransitGatewayVPCAttachmentInitParameters struct {

	// Whether Appliance Mode support is enabled. If enabled, a traffic flow between a source and destination uses the same Availability Zone for the VPC attachment for the lifetime of that flow. Valid values: disable, enable. Default value: disable.
	ApplianceModeSupport *string `json:"applianceModeSupport,omitempty" tf:"appliance_mode_support,omitempty"`

	// Whether DNS support is enabled. Valid values: disable, enable. Default value: enable.
	DNSSupport *string `json:"dnsSupport,omitempty" tf:"dns_support,omitempty"`

	// Whether IPv6 support is enabled. Valid values: disable, enable. Default value: disable.
	IPv6Support *string `json:"ipv6Support,omitempty" tf:"ipv6_support,omitempty"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Identifiers of EC2 Subnets.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Boolean whether the VPC Attachment should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: true.
	TransitGatewayDefaultRouteTableAssociation *bool `json:"transitGatewayDefaultRouteTableAssociation,omitempty" tf:"transit_gateway_default_route_table_association,omitempty"`

	// Boolean whether the VPC Attachment should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: true.
	TransitGatewayDefaultRouteTablePropagation *bool `json:"transitGatewayDefaultRouteTablePropagation,omitempty" tf:"transit_gateway_default_route_table_propagation,omitempty"`

	// Identifier of EC2 Transit Gateway.
	// +crossplane:generate:reference:type=TransitGateway
	TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id,omitempty"`

	// Reference to a TransitGateway to populate transitGatewayId.
	// +kubebuilder:validation:Optional
	TransitGatewayIDRef *v1.Reference `json:"transitGatewayIdRef,omitempty" tf:"-"`

	// Selector for a TransitGateway to populate transitGatewayId.
	// +kubebuilder:validation:Optional
	TransitGatewayIDSelector *v1.Selector `json:"transitGatewayIdSelector,omitempty" tf:"-"`

	// Identifier of EC2 VPC.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type TransitGatewayVPCAttachmentObservation struct {

	// Whether Appliance Mode support is enabled. If enabled, a traffic flow between a source and destination uses the same Availability Zone for the VPC attachment for the lifetime of that flow. Valid values: disable, enable. Default value: disable.
	ApplianceModeSupport *string `json:"applianceModeSupport,omitempty" tf:"appliance_mode_support,omitempty"`

	// Whether DNS support is enabled. Valid values: disable, enable. Default value: enable.
	DNSSupport *string `json:"dnsSupport,omitempty" tf:"dns_support,omitempty"`

	// EC2 Transit Gateway Attachment identifier
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether IPv6 support is enabled. Valid values: disable, enable. Default value: disable.
	IPv6Support *string `json:"ipv6Support,omitempty" tf:"ipv6_support,omitempty"`

	// Identifiers of EC2 Subnets.
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Boolean whether the VPC Attachment should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: true.
	TransitGatewayDefaultRouteTableAssociation *bool `json:"transitGatewayDefaultRouteTableAssociation,omitempty" tf:"transit_gateway_default_route_table_association,omitempty"`

	// Boolean whether the VPC Attachment should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: true.
	TransitGatewayDefaultRouteTablePropagation *bool `json:"transitGatewayDefaultRouteTablePropagation,omitempty" tf:"transit_gateway_default_route_table_propagation,omitempty"`

	// Identifier of EC2 Transit Gateway.
	TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id,omitempty"`

	// Identifier of EC2 VPC.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Identifier of the AWS account that owns the EC2 VPC.
	VPCOwnerID *string `json:"vpcOwnerId,omitempty" tf:"vpc_owner_id,omitempty"`
}

type TransitGatewayVPCAttachmentParameters struct {

	// Whether Appliance Mode support is enabled. If enabled, a traffic flow between a source and destination uses the same Availability Zone for the VPC attachment for the lifetime of that flow. Valid values: disable, enable. Default value: disable.
	// +kubebuilder:validation:Optional
	ApplianceModeSupport *string `json:"applianceModeSupport,omitempty" tf:"appliance_mode_support,omitempty"`

	// Whether DNS support is enabled. Valid values: disable, enable. Default value: enable.
	// +kubebuilder:validation:Optional
	DNSSupport *string `json:"dnsSupport,omitempty" tf:"dns_support,omitempty"`

	// Whether IPv6 support is enabled. Valid values: disable, enable. Default value: disable.
	// +kubebuilder:validation:Optional
	IPv6Support *string `json:"ipv6Support,omitempty" tf:"ipv6_support,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Identifiers of EC2 Subnets.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +kubebuilder:validation:Optional
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Boolean whether the VPC Attachment should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: true.
	// +kubebuilder:validation:Optional
	TransitGatewayDefaultRouteTableAssociation *bool `json:"transitGatewayDefaultRouteTableAssociation,omitempty" tf:"transit_gateway_default_route_table_association,omitempty"`

	// Boolean whether the VPC Attachment should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: true.
	// +kubebuilder:validation:Optional
	TransitGatewayDefaultRouteTablePropagation *bool `json:"transitGatewayDefaultRouteTablePropagation,omitempty" tf:"transit_gateway_default_route_table_propagation,omitempty"`

	// Identifier of EC2 Transit Gateway.
	// +crossplane:generate:reference:type=TransitGateway
	// +kubebuilder:validation:Optional
	TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id,omitempty"`

	// Reference to a TransitGateway to populate transitGatewayId.
	// +kubebuilder:validation:Optional
	TransitGatewayIDRef *v1.Reference `json:"transitGatewayIdRef,omitempty" tf:"-"`

	// Selector for a TransitGateway to populate transitGatewayId.
	// +kubebuilder:validation:Optional
	TransitGatewayIDSelector *v1.Selector `json:"transitGatewayIdSelector,omitempty" tf:"-"`

	// Identifier of EC2 VPC.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// TransitGatewayVPCAttachmentSpec defines the desired state of TransitGatewayVPCAttachment
type TransitGatewayVPCAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitGatewayVPCAttachmentParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider TransitGatewayVPCAttachmentInitParameters `json:"initProvider,omitempty"`
}

// TransitGatewayVPCAttachmentStatus defines the observed state of TransitGatewayVPCAttachment.
type TransitGatewayVPCAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitGatewayVPCAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TransitGatewayVPCAttachment is the Schema for the TransitGatewayVPCAttachments API. Manages an EC2 Transit Gateway VPC Attachment
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TransitGatewayVPCAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransitGatewayVPCAttachmentSpec   `json:"spec"`
	Status            TransitGatewayVPCAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayVPCAttachmentList contains a list of TransitGatewayVPCAttachments
type TransitGatewayVPCAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitGatewayVPCAttachment `json:"items"`
}

// Repository type metadata.
var (
	TransitGatewayVPCAttachment_Kind             = "TransitGatewayVPCAttachment"
	TransitGatewayVPCAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitGatewayVPCAttachment_Kind}.String()
	TransitGatewayVPCAttachment_KindAPIVersion   = TransitGatewayVPCAttachment_Kind + "." + CRDGroupVersion.String()
	TransitGatewayVPCAttachment_GroupVersionKind = CRDGroupVersion.WithKind(TransitGatewayVPCAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&TransitGatewayVPCAttachment{}, &TransitGatewayVPCAttachmentList{})
}

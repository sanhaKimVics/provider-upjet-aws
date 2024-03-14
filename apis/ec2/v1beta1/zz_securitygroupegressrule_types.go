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

type SecurityGroupEgressRuleInitParameters struct {

	// The destination IPv4 CIDR range.
	CidrIPv4 *string `json:"cidrIpv4,omitempty" tf:"cidr_ipv4,omitempty"`

	// The destination IPv6 CIDR range.
	CidrIPv6 *string `json:"cidrIpv6,omitempty" tf:"cidr_ipv6,omitempty"`

	// The security group rule description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type.
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// The IP protocol name or number. Use -1 to specify all protocols. Note that if ip_protocol is set to -1, it translates to all protocols, all port ranges, and from_port and to_port values should not be defined.
	IPProtocol *string `json:"ipProtocol,omitempty" tf:"ip_protocol,omitempty"`

	// The ID of the destination prefix list.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.ManagedPrefixList
	PrefixListID *string `json:"prefixListId,omitempty" tf:"prefix_list_id,omitempty"`

	// Reference to a ManagedPrefixList in ec2 to populate prefixListId.
	// +kubebuilder:validation:Optional
	PrefixListIDRef *v1.Reference `json:"prefixListIdRef,omitempty" tf:"-"`

	// Selector for a ManagedPrefixList in ec2 to populate prefixListId.
	// +kubebuilder:validation:Optional
	PrefixListIDSelector *v1.Selector `json:"prefixListIdSelector,omitempty" tf:"-"`

	// The destination security group that is referenced in the rule.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	ReferencedSecurityGroupID *string `json:"referencedSecurityGroupId,omitempty" tf:"referenced_security_group_id,omitempty"`

	// Reference to a SecurityGroup in ec2 to populate referencedSecurityGroupId.
	// +kubebuilder:validation:Optional
	ReferencedSecurityGroupIDRef *v1.Reference `json:"referencedSecurityGroupIdRef,omitempty" tf:"-"`

	// Selector for a SecurityGroup in ec2 to populate referencedSecurityGroupId.
	// +kubebuilder:validation:Optional
	ReferencedSecurityGroupIDSelector *v1.Selector `json:"referencedSecurityGroupIdSelector,omitempty" tf:"-"`

	// The ID of the security group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Reference to a SecurityGroup in ec2 to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRef *v1.Reference `json:"securityGroupIdRef,omitempty" tf:"-"`

	// Selector for a SecurityGroup in ec2 to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code.
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type SecurityGroupEgressRuleObservation struct {

	// The Amazon Resource Name (ARN) of the security group rule.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The destination IPv4 CIDR range.
	CidrIPv4 *string `json:"cidrIpv4,omitempty" tf:"cidr_ipv4,omitempty"`

	// The destination IPv6 CIDR range.
	CidrIPv6 *string `json:"cidrIpv6,omitempty" tf:"cidr_ipv6,omitempty"`

	// The security group rule description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type.
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IP protocol name or number. Use -1 to specify all protocols. Note that if ip_protocol is set to -1, it translates to all protocols, all port ranges, and from_port and to_port values should not be defined.
	IPProtocol *string `json:"ipProtocol,omitempty" tf:"ip_protocol,omitempty"`

	// The ID of the destination prefix list.
	PrefixListID *string `json:"prefixListId,omitempty" tf:"prefix_list_id,omitempty"`

	// The destination security group that is referenced in the rule.
	ReferencedSecurityGroupID *string `json:"referencedSecurityGroupId,omitempty" tf:"referenced_security_group_id,omitempty"`

	// The ID of the security group.
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// The ID of the security group rule.
	SecurityGroupRuleID *string `json:"securityGroupRuleId,omitempty" tf:"security_group_rule_id,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code.
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type SecurityGroupEgressRuleParameters struct {

	// The destination IPv4 CIDR range.
	// +kubebuilder:validation:Optional
	CidrIPv4 *string `json:"cidrIpv4,omitempty" tf:"cidr_ipv4,omitempty"`

	// The destination IPv6 CIDR range.
	// +kubebuilder:validation:Optional
	CidrIPv6 *string `json:"cidrIpv6,omitempty" tf:"cidr_ipv6,omitempty"`

	// The security group rule description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type.
	// +kubebuilder:validation:Optional
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// The IP protocol name or number. Use -1 to specify all protocols. Note that if ip_protocol is set to -1, it translates to all protocols, all port ranges, and from_port and to_port values should not be defined.
	// +kubebuilder:validation:Optional
	IPProtocol *string `json:"ipProtocol,omitempty" tf:"ip_protocol,omitempty"`

	// The ID of the destination prefix list.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.ManagedPrefixList
	// +kubebuilder:validation:Optional
	PrefixListID *string `json:"prefixListId,omitempty" tf:"prefix_list_id,omitempty"`

	// Reference to a ManagedPrefixList in ec2 to populate prefixListId.
	// +kubebuilder:validation:Optional
	PrefixListIDRef *v1.Reference `json:"prefixListIdRef,omitempty" tf:"-"`

	// Selector for a ManagedPrefixList in ec2 to populate prefixListId.
	// +kubebuilder:validation:Optional
	PrefixListIDSelector *v1.Selector `json:"prefixListIdSelector,omitempty" tf:"-"`

	// The destination security group that is referenced in the rule.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +kubebuilder:validation:Optional
	ReferencedSecurityGroupID *string `json:"referencedSecurityGroupId,omitempty" tf:"referenced_security_group_id,omitempty"`

	// Reference to a SecurityGroup in ec2 to populate referencedSecurityGroupId.
	// +kubebuilder:validation:Optional
	ReferencedSecurityGroupIDRef *v1.Reference `json:"referencedSecurityGroupIdRef,omitempty" tf:"-"`

	// Selector for a SecurityGroup in ec2 to populate referencedSecurityGroupId.
	// +kubebuilder:validation:Optional
	ReferencedSecurityGroupIDSelector *v1.Selector `json:"referencedSecurityGroupIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ID of the security group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +kubebuilder:validation:Optional
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Reference to a SecurityGroup in ec2 to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRef *v1.Reference `json:"securityGroupIdRef,omitempty" tf:"-"`

	// Selector for a SecurityGroup in ec2 to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code.
	// +kubebuilder:validation:Optional
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

// SecurityGroupEgressRuleSpec defines the desired state of SecurityGroupEgressRule
type SecurityGroupEgressRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityGroupEgressRuleParameters `json:"forProvider"`
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
	InitProvider SecurityGroupEgressRuleInitParameters `json:"initProvider,omitempty"`
}

// SecurityGroupEgressRuleStatus defines the observed state of SecurityGroupEgressRule.
type SecurityGroupEgressRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityGroupEgressRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SecurityGroupEgressRule is the Schema for the SecurityGroupEgressRules API. Provides a VPC security group egress rule resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SecurityGroupEgressRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ipProtocol) || (has(self.initProvider) && has(self.initProvider.ipProtocol))",message="spec.forProvider.ipProtocol is a required parameter"
	Spec   SecurityGroupEgressRuleSpec   `json:"spec"`
	Status SecurityGroupEgressRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroupEgressRuleList contains a list of SecurityGroupEgressRules
type SecurityGroupEgressRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityGroupEgressRule `json:"items"`
}

// Repository type metadata.
var (
	SecurityGroupEgressRule_Kind             = "SecurityGroupEgressRule"
	SecurityGroupEgressRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityGroupEgressRule_Kind}.String()
	SecurityGroupEgressRule_KindAPIVersion   = SecurityGroupEgressRule_Kind + "." + CRDGroupVersion.String()
	SecurityGroupEgressRule_GroupVersionKind = CRDGroupVersion.WithKind(SecurityGroupEgressRule_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityGroupEgressRule{}, &SecurityGroupEgressRuleList{})
}

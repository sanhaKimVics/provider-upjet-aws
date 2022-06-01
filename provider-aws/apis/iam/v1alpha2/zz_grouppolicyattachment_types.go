/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type GroupPolicyAttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type GroupPolicyAttachmentParameters struct {

	// +crossplane:generate:reference:type=Group
	// +kubebuilder:validation:Optional
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// +kubebuilder:validation:Optional
	GroupRef *v1.Reference `json:"groupRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	GroupSelector *v1.Selector `json:"groupSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=Policy
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	PolicyArn *string `json:"policyArn,omitempty" tf:"policy_arn,omitempty"`

	// +kubebuilder:validation:Optional
	PolicyArnRef *v1.Reference `json:"policyArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PolicyArnSelector *v1.Selector `json:"policyArnSelector,omitempty" tf:"-"`
}

// GroupPolicyAttachmentSpec defines the desired state of GroupPolicyAttachment
type GroupPolicyAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupPolicyAttachmentParameters `json:"forProvider"`
}

// GroupPolicyAttachmentStatus defines the observed state of GroupPolicyAttachment.
type GroupPolicyAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupPolicyAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GroupPolicyAttachment is the Schema for the GroupPolicyAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type GroupPolicyAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GroupPolicyAttachmentSpec   `json:"spec"`
	Status            GroupPolicyAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupPolicyAttachmentList contains a list of GroupPolicyAttachments
type GroupPolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GroupPolicyAttachment `json:"items"`
}

// Repository type metadata.
var (
	GroupPolicyAttachment_Kind             = "GroupPolicyAttachment"
	GroupPolicyAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GroupPolicyAttachment_Kind}.String()
	GroupPolicyAttachment_KindAPIVersion   = GroupPolicyAttachment_Kind + "." + CRDGroupVersion.String()
	GroupPolicyAttachment_GroupVersionKind = CRDGroupVersion.WithKind(GroupPolicyAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&GroupPolicyAttachment{}, &GroupPolicyAttachmentList{})
}

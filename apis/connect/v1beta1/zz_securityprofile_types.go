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

type SecurityProfileInitParameters struct {

	// Specifies the description of the Security Profile.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the identifier of the hosting Amazon Connect Instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/connect/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in connect to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in connect to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Specifies the name of the Security Profile.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies a list of permissions assigned to the security profile.
	// +listType=set
	Permissions []*string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SecurityProfileObservation struct {

	// The Amazon Resource Name (ARN) of the Security Profile.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Specifies the description of the Security Profile.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The identifier of the hosting Amazon Connect Instance and identifier of the Security Profile separated by a colon (:).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the identifier of the hosting Amazon Connect Instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Specifies the name of the Security Profile.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The organization resource identifier for the security profile.
	OrganizationResourceID *string `json:"organizationResourceId,omitempty" tf:"organization_resource_id,omitempty"`

	// Specifies a list of permissions assigned to the security profile.
	// +listType=set
	Permissions []*string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// The identifier for the Security Profile.
	SecurityProfileID *string `json:"securityProfileId,omitempty" tf:"security_profile_id,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type SecurityProfileParameters struct {

	// Specifies the description of the Security Profile.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the identifier of the hosting Amazon Connect Instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/connect/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in connect to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in connect to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Specifies the name of the Security Profile.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies a list of permissions assigned to the security profile.
	// +kubebuilder:validation:Optional
	// +listType=set
	Permissions []*string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// SecurityProfileSpec defines the desired state of SecurityProfile
type SecurityProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityProfileParameters `json:"forProvider"`
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
	InitProvider SecurityProfileInitParameters `json:"initProvider,omitempty"`
}

// SecurityProfileStatus defines the observed state of SecurityProfile.
type SecurityProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SecurityProfile is the Schema for the SecurityProfiles API. Provides details about a specific Amazon Connect Security Profile.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SecurityProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   SecurityProfileSpec   `json:"spec"`
	Status SecurityProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityProfileList contains a list of SecurityProfiles
type SecurityProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityProfile `json:"items"`
}

// Repository type metadata.
var (
	SecurityProfile_Kind             = "SecurityProfile"
	SecurityProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityProfile_Kind}.String()
	SecurityProfile_KindAPIVersion   = SecurityProfile_Kind + "." + CRDGroupVersion.String()
	SecurityProfile_GroupVersionKind = CRDGroupVersion.WithKind(SecurityProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityProfile{}, &SecurityProfileList{})
}

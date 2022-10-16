/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SecretVersionObservation struct {

	// The ARN of the secret.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A pipe delimited combination of secret ID and version ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The unique identifier of the version of the secret.
	VersionID *string `json:"versionId,omitempty" tf:"version_id,omitempty"`
}

type SecretVersionParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Specifies binary data that you want to encrypt and store in this version of the secret. This is required if secret_string is not set. Needs to be encoded to base64.
	// +kubebuilder:validation:Optional
	SecretBinarySecretRef *v1.SecretKeySelector `json:"secretBinarySecretRef,omitempty" tf:"-"`

	// Specifies the secret to which you want to add a new version. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret. The secret must already exist.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/secretsmanager/v1beta1.Secret
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SecretID *string `json:"secretId,omitempty" tf:"secret_id,omitempty"`

	// Reference to a Secret in secretsmanager to populate secretId.
	// +kubebuilder:validation:Optional
	SecretIDRef *v1.Reference `json:"secretIdRef,omitempty" tf:"-"`

	// Selector for a Secret in secretsmanager to populate secretId.
	// +kubebuilder:validation:Optional
	SecretIDSelector *v1.Selector `json:"secretIdSelector,omitempty" tf:"-"`

	// Specifies text data that you want to encrypt and store in this version of the secret. This is required if secret_binary is not set.
	// +kubebuilder:validation:Optional
	SecretStringSecretRef *v1.SecretKeySelector `json:"secretStringSecretRef,omitempty" tf:"-"`

	// Specifies a list of staging labels that are attached to this version of the secret. A staging label must be unique to a single version of the secret. If you specify a staging label that's already associated with a different version of the same secret then that staging label is automatically removed from the other version and attached to this version. If you do not specify a value, then AWS Secrets Manager automatically moves the staging label AWSCURRENT to this new version on creation.
	// +kubebuilder:validation:Optional
	VersionStages []*string `json:"versionStages,omitempty" tf:"version_stages,omitempty"`
}

// SecretVersionSpec defines the desired state of SecretVersion
type SecretVersionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretVersionParameters `json:"forProvider"`
}

// SecretVersionStatus defines the observed state of SecretVersion.
type SecretVersionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretVersionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecretVersion is the Schema for the SecretVersions API. Provides a resource to manage AWS Secrets Manager secret version including its secret value
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SecretVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecretVersionSpec   `json:"spec"`
	Status            SecretVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretVersionList contains a list of SecretVersions
type SecretVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretVersion `json:"items"`
}

// Repository type metadata.
var (
	SecretVersion_Kind             = "SecretVersion"
	SecretVersion_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretVersion_Kind}.String()
	SecretVersion_KindAPIVersion   = SecretVersion_Kind + "." + CRDGroupVersion.String()
	SecretVersion_GroupVersionKind = CRDGroupVersion.WithKind(SecretVersion_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretVersion{}, &SecretVersionList{})
}

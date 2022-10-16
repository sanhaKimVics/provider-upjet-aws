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

type BucketObjectLockConfigurationObservation struct {

	// The bucket or bucket and expected_bucket_owner separated by a comma (,) if the latter is provided.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BucketObjectLockConfigurationParameters struct {

	// The name of the bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// The account ID of the expected bucket owner.
	// +kubebuilder:validation:Optional
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// Indicates whether this bucket has an Object Lock configuration enabled. Defaults to Enabled. Valid values: Enabled.
	// +kubebuilder:validation:Optional
	ObjectLockEnabled *string `json:"objectLockEnabled,omitempty" tf:"object_lock_enabled,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Configuration block for specifying the Object Lock rule for the specified object detailed below.
	// +kubebuilder:validation:Required
	Rule []BucketObjectLockConfigurationRuleParameters `json:"rule" tf:"rule,omitempty"`

	// A token to allow Object Lock to be enabled for an existing bucket. You must contact AWS support for the bucket's "Object Lock token".
	// The token is generated in the back-end when versioning is enabled on a bucket. For more details on versioning, see the aws_s3_bucket_versioning resource.
	// +kubebuilder:validation:Optional
	TokenSecretRef *v1.SecretKeySelector `json:"tokenSecretRef,omitempty" tf:"-"`
}

type BucketObjectLockConfigurationRuleObservation struct {
}

type BucketObjectLockConfigurationRuleParameters struct {

	// A configuration block for specifying the default Object Lock retention settings for new objects placed in the specified bucket detailed below.
	// +kubebuilder:validation:Required
	DefaultRetention []RuleDefaultRetentionParameters `json:"defaultRetention" tf:"default_retention,omitempty"`
}

type RuleDefaultRetentionObservation struct {
}

type RuleDefaultRetentionParameters struct {

	// The number of days that you want to specify for the default retention period.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// The default Object Lock retention mode you want to apply to new objects placed in the specified bucket. Valid values: COMPLIANCE, GOVERNANCE.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The number of years that you want to specify for the default retention period.
	// +kubebuilder:validation:Optional
	Years *float64 `json:"years,omitempty" tf:"years,omitempty"`
}

// BucketObjectLockConfigurationSpec defines the desired state of BucketObjectLockConfiguration
type BucketObjectLockConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketObjectLockConfigurationParameters `json:"forProvider"`
}

// BucketObjectLockConfigurationStatus defines the observed state of BucketObjectLockConfiguration.
type BucketObjectLockConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketObjectLockConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketObjectLockConfiguration is the Schema for the BucketObjectLockConfigurations API. Provides an S3 bucket Object Lock configuration resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BucketObjectLockConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketObjectLockConfigurationSpec   `json:"spec"`
	Status            BucketObjectLockConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketObjectLockConfigurationList contains a list of BucketObjectLockConfigurations
type BucketObjectLockConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketObjectLockConfiguration `json:"items"`
}

// Repository type metadata.
var (
	BucketObjectLockConfiguration_Kind             = "BucketObjectLockConfiguration"
	BucketObjectLockConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketObjectLockConfiguration_Kind}.String()
	BucketObjectLockConfiguration_KindAPIVersion   = BucketObjectLockConfiguration_Kind + "." + CRDGroupVersion.String()
	BucketObjectLockConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(BucketObjectLockConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketObjectLockConfiguration{}, &BucketObjectLockConfigurationList{})
}

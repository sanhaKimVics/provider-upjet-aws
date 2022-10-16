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

type ACLConfigurationObservation struct {
}

type ACLConfigurationParameters struct {

	// The Amazon S3 canned ACL that Athena should specify when storing query results. Valid value is BUCKET_OWNER_FULL_CONTROL.
	// +kubebuilder:validation:Required
	S3ACLOption *string `json:"s3AclOption" tf:"s3_acl_option,omitempty"`
}

type DatabaseObservation struct {

	// The database name
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DatabaseParameters struct {

	// Indicates that an Amazon S3 canned ACL should be set to control ownership of stored query results. See ACL Configuration below.
	// +kubebuilder:validation:Optional
	ACLConfiguration []ACLConfigurationParameters `json:"aclConfiguration,omitempty" tf:"acl_configuration,omitempty"`

	// Name of S3 bucket to save the results of the query execution.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Description of the database.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// The encryption key block AWS Athena uses to decrypt the data in S3, such as an AWS Key Management Service (AWS KMS) key. See Encryption Configuration below.
	// +kubebuilder:validation:Optional
	EncryptionConfiguration []EncryptionConfigurationParameters `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration,omitempty"`

	// The AWS account ID that you expect to be the owner of the Amazon S3 bucket.
	// +kubebuilder:validation:Optional
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// A boolean that indicates all tables should be deleted from the database so that the database can be destroyed without error. The tables are not recoverable.
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// A key-value map of custom metadata properties for the database definition.
	// +kubebuilder:validation:Optional
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type EncryptionConfigurationObservation struct {
}

type EncryptionConfigurationParameters struct {

	// The type of key; one of SSE_S3, SSE_KMS, CSE_KMS
	// +kubebuilder:validation:Required
	EncryptionOption *string `json:"encryptionOption" tf:"encryption_option,omitempty"`

	// The KMS key ARN or ID; required for key types SSE_KMS and CSE_KMS.
	// +kubebuilder:validation:Optional
	KMSKey *string `json:"kmsKey,omitempty" tf:"kms_key,omitempty"`
}

// DatabaseSpec defines the desired state of Database
type DatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatabaseParameters `json:"forProvider"`
}

// DatabaseStatus defines the observed state of Database.
type DatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Database is the Schema for the Databases API. Provides an Athena database.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Database struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseSpec   `json:"spec"`
	Status            DatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseList contains a list of Databases
type DatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Database `json:"items"`
}

// Repository type metadata.
var (
	Database_Kind             = "Database"
	Database_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Database_Kind}.String()
	Database_KindAPIVersion   = Database_Kind + "." + CRDGroupVersion.String()
	Database_GroupVersionKind = CRDGroupVersion.WithKind(Database_Kind)
)

func init() {
	SchemeBuilder.Register(&Database{}, &DatabaseList{})
}

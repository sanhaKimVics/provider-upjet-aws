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

type DestinationInitParameters struct {

	// A configuration block describing the S3 Source object: See S3 Source below for details.
	S3 []S3InitParameters `json:"s3,omitempty" tf:"s3,omitempty"`
}

type DestinationObservation struct {

	// A configuration block describing the S3 Source object: See S3 Source below for details.
	S3 []S3Observation `json:"s3,omitempty" tf:"s3,omitempty"`
}

type DestinationParameters struct {

	// A configuration block describing the S3 Source object: See S3 Source below for details.
	// +kubebuilder:validation:Optional
	S3 []S3Parameters `json:"s3" tf:"s3,omitempty"`
}

type RevocationRecordInitParameters struct {
}

type RevocationRecordObservation struct {
	Reason *string `json:"reason,omitempty" tf:"reason,omitempty"`

	RevokedAt *string `json:"revokedAt,omitempty" tf:"revoked_at,omitempty"`

	RevokedBy *string `json:"revokedBy,omitempty" tf:"revoked_by,omitempty"`
}

type RevocationRecordParameters struct {
}

type S3InitParameters struct {

	// Name of the S3 bucket.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// An Amazon S3 object key prefix that you can use to limit signed objects keys to begin with the specified prefix.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type S3Observation struct {

	// Name of the S3 bucket.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// An Amazon S3 object key prefix that you can use to limit signed objects keys to begin with the specified prefix.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type S3Parameters struct {

	// Name of the S3 bucket.
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// An Amazon S3 object key prefix that you can use to limit signed objects keys to begin with the specified prefix.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type SignedObjectInitParameters struct {
}

type SignedObjectObservation struct {

	// A configuration block describing the S3 Source object: See S3 Source below for details.
	S3 []SignedObjectS3Observation `json:"s3,omitempty" tf:"s3,omitempty"`
}

type SignedObjectParameters struct {
}

type SignedObjectS3InitParameters struct {
}

type SignedObjectS3Observation struct {

	// Name of the S3 bucket.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Key name of the object that contains your unsigned code.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type SignedObjectS3Parameters struct {
}

type SigningJobInitParameters struct {

	// The S3 bucket in which to save your signed object. See Destination below for details.
	Destination []DestinationInitParameters `json:"destination,omitempty" tf:"destination,omitempty"`

	// Set this argument to true to ignore signing job failures and retrieve failed status and reason. Default false.
	IgnoreSigningJobFailure *bool `json:"ignoreSigningJobFailure,omitempty" tf:"ignore_signing_job_failure,omitempty"`

	// The name of the profile to initiate the signing operation.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/signer/v1beta1.SigningProfile
	ProfileName *string `json:"profileName,omitempty" tf:"profile_name,omitempty"`

	// Reference to a SigningProfile in signer to populate profileName.
	// +kubebuilder:validation:Optional
	ProfileNameRef *v1.Reference `json:"profileNameRef,omitempty" tf:"-"`

	// Selector for a SigningProfile in signer to populate profileName.
	// +kubebuilder:validation:Optional
	ProfileNameSelector *v1.Selector `json:"profileNameSelector,omitempty" tf:"-"`

	// The S3 bucket that contains the object to sign. See Source below for details.
	Source []SourceInitParameters `json:"source,omitempty" tf:"source,omitempty"`
}

type SigningJobObservation struct {

	// Date and time in RFC3339 format that the signing job was completed.
	CompletedAt *string `json:"completedAt,omitempty" tf:"completed_at,omitempty"`

	// Date and time in RFC3339 format that the signing job was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The S3 bucket in which to save your signed object. See Destination below for details.
	Destination []DestinationObservation `json:"destination,omitempty" tf:"destination,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Set this argument to true to ignore signing job failures and retrieve failed status and reason. Default false.
	IgnoreSigningJobFailure *bool `json:"ignoreSigningJobFailure,omitempty" tf:"ignore_signing_job_failure,omitempty"`

	// The ID of the signing job on output.
	JobID *string `json:"jobId,omitempty" tf:"job_id,omitempty"`

	// The IAM entity that initiated the signing job.
	JobInvoker *string `json:"jobInvoker,omitempty" tf:"job_invoker,omitempty"`

	// The AWS account ID of the job owner.
	JobOwner *string `json:"jobOwner,omitempty" tf:"job_owner,omitempty"`

	// A human-readable name for the signing platform associated with the signing job.
	PlatformDisplayName *string `json:"platformDisplayName,omitempty" tf:"platform_display_name,omitempty"`

	// The platform to which your signed code image will be distributed.
	PlatformID *string `json:"platformId,omitempty" tf:"platform_id,omitempty"`

	// The name of the profile to initiate the signing operation.
	ProfileName *string `json:"profileName,omitempty" tf:"profile_name,omitempty"`

	// The version of the signing profile used to initiate the signing job.
	ProfileVersion *string `json:"profileVersion,omitempty" tf:"profile_version,omitempty"`

	// The IAM principal that requested the signing job.
	RequestedBy *string `json:"requestedBy,omitempty" tf:"requested_by,omitempty"`

	// A revocation record if the signature generated by the signing job has been revoked. Contains a timestamp and the ID of the IAM entity that revoked the signature.
	RevocationRecord []RevocationRecordObservation `json:"revocationRecord,omitempty" tf:"revocation_record,omitempty"`

	// The time when the signature of a signing job expires.
	SignatureExpiresAt *string `json:"signatureExpiresAt,omitempty" tf:"signature_expires_at,omitempty"`

	// Name of the S3 bucket where the signed code image is saved by code signing.
	SignedObject []SignedObjectObservation `json:"signedObject,omitempty" tf:"signed_object,omitempty"`

	// The S3 bucket that contains the object to sign. See Source below for details.
	Source []SourceObservation `json:"source,omitempty" tf:"source,omitempty"`

	// Status of the signing job.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// String value that contains the status reason.
	StatusReason *string `json:"statusReason,omitempty" tf:"status_reason,omitempty"`
}

type SigningJobParameters struct {

	// The S3 bucket in which to save your signed object. See Destination below for details.
	// +kubebuilder:validation:Optional
	Destination []DestinationParameters `json:"destination,omitempty" tf:"destination,omitempty"`

	// Set this argument to true to ignore signing job failures and retrieve failed status and reason. Default false.
	// +kubebuilder:validation:Optional
	IgnoreSigningJobFailure *bool `json:"ignoreSigningJobFailure,omitempty" tf:"ignore_signing_job_failure,omitempty"`

	// The name of the profile to initiate the signing operation.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/signer/v1beta1.SigningProfile
	// +kubebuilder:validation:Optional
	ProfileName *string `json:"profileName,omitempty" tf:"profile_name,omitempty"`

	// Reference to a SigningProfile in signer to populate profileName.
	// +kubebuilder:validation:Optional
	ProfileNameRef *v1.Reference `json:"profileNameRef,omitempty" tf:"-"`

	// Selector for a SigningProfile in signer to populate profileName.
	// +kubebuilder:validation:Optional
	ProfileNameSelector *v1.Selector `json:"profileNameSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The S3 bucket that contains the object to sign. See Source below for details.
	// +kubebuilder:validation:Optional
	Source []SourceParameters `json:"source,omitempty" tf:"source,omitempty"`
}

type SourceInitParameters struct {

	// A configuration block describing the S3 Source object: See S3 Source below for details.
	S3 []SourceS3InitParameters `json:"s3,omitempty" tf:"s3,omitempty"`
}

type SourceObservation struct {

	// A configuration block describing the S3 Source object: See S3 Source below for details.
	S3 []SourceS3Observation `json:"s3,omitempty" tf:"s3,omitempty"`
}

type SourceParameters struct {

	// A configuration block describing the S3 Source object: See S3 Source below for details.
	// +kubebuilder:validation:Optional
	S3 []SourceS3Parameters `json:"s3" tf:"s3,omitempty"`
}

type SourceS3InitParameters struct {

	// Name of the S3 bucket.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Key name of the object that contains your unsigned code.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Version of your source image in your version enabled S3 bucket.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type SourceS3Observation struct {

	// Name of the S3 bucket.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Key name of the object that contains your unsigned code.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Version of your source image in your version enabled S3 bucket.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type SourceS3Parameters struct {

	// Name of the S3 bucket.
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// Key name of the object that contains your unsigned code.
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// Version of your source image in your version enabled S3 bucket.
	// +kubebuilder:validation:Optional
	Version *string `json:"version" tf:"version,omitempty"`
}

// SigningJobSpec defines the desired state of SigningJob
type SigningJobSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SigningJobParameters `json:"forProvider"`
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
	InitProvider SigningJobInitParameters `json:"initProvider,omitempty"`
}

// SigningJobStatus defines the observed state of SigningJob.
type SigningJobStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SigningJobObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SigningJob is the Schema for the SigningJobs API. Creates a Signer Signing Job.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SigningJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.destination) || (has(self.initProvider) && has(self.initProvider.destination))",message="spec.forProvider.destination is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.source) || (has(self.initProvider) && has(self.initProvider.source))",message="spec.forProvider.source is a required parameter"
	Spec   SigningJobSpec   `json:"spec"`
	Status SigningJobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SigningJobList contains a list of SigningJobs
type SigningJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SigningJob `json:"items"`
}

// Repository type metadata.
var (
	SigningJob_Kind             = "SigningJob"
	SigningJob_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SigningJob_Kind}.String()
	SigningJob_KindAPIVersion   = SigningJob_Kind + "." + CRDGroupVersion.String()
	SigningJob_GroupVersionKind = CRDGroupVersion.WithKind(SigningJob_Kind)
)

func init() {
	SchemeBuilder.Register(&SigningJob{}, &SigningJobList{})
}

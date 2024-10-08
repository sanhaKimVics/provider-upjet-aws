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

type DBSnapshotCopyInitParameters struct {

	// Whether to copy existing tags. Defaults to false.
	CopyTags *bool `json:"copyTags,omitempty" tf:"copy_tags,omitempty"`

	// The Destination region to place snapshot copy.
	DestinationRegion *string `json:"destinationRegion,omitempty" tf:"destination_region,omitempty"`

	// KMS key ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// The name of an option group to associate with the copy of the snapshot.
	OptionGroupName *string `json:"optionGroupName,omitempty" tf:"option_group_name,omitempty"`

	// he URL that contains a Signature Version 4 signed request.
	PresignedURL *string `json:"presignedUrl,omitempty" tf:"presigned_url,omitempty"`

	// List of AWS Account ids to share snapshot with, use all to make snaphot public.
	// +listType=set
	SharedAccounts []*string `json:"sharedAccounts,omitempty" tf:"shared_accounts,omitempty"`

	// Snapshot identifier of the source snapshot.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/rds/v1beta1.Snapshot
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("db_snapshot_arn",true)
	SourceDBSnapshotIdentifier *string `json:"sourceDbSnapshotIdentifier,omitempty" tf:"source_db_snapshot_identifier,omitempty"`

	// Reference to a Snapshot in rds to populate sourceDbSnapshotIdentifier.
	// +kubebuilder:validation:Optional
	SourceDBSnapshotIdentifierRef *v1.Reference `json:"sourceDbSnapshotIdentifierRef,omitempty" tf:"-"`

	// Selector for a Snapshot in rds to populate sourceDbSnapshotIdentifier.
	// +kubebuilder:validation:Optional
	SourceDBSnapshotIdentifierSelector *v1.Selector `json:"sourceDbSnapshotIdentifierSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The external custom Availability Zone.
	TargetCustomAvailabilityZone *string `json:"targetCustomAvailabilityZone,omitempty" tf:"target_custom_availability_zone,omitempty"`

	// The Identifier for the snapshot.
	TargetDBSnapshotIdentifier *string `json:"targetDbSnapshotIdentifier,omitempty" tf:"target_db_snapshot_identifier,omitempty"`
}

type DBSnapshotCopyObservation struct {

	// Specifies the allocated storage size in gigabytes (GB).
	AllocatedStorage *float64 `json:"allocatedStorage,omitempty" tf:"allocated_storage,omitempty"`

	// Specifies the name of the Availability Zone the DB instance was located in at the time of the DB snapshot.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Whether to copy existing tags. Defaults to false.
	CopyTags *bool `json:"copyTags,omitempty" tf:"copy_tags,omitempty"`

	// The Amazon Resource Name (ARN) for the DB snapshot.
	DBSnapshotArn *string `json:"dbSnapshotArn,omitempty" tf:"db_snapshot_arn,omitempty"`

	// The Destination region to place snapshot copy.
	DestinationRegion *string `json:"destinationRegion,omitempty" tf:"destination_region,omitempty"`

	// Specifies whether the DB snapshot is encrypted.
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// Specifies the name of the database engine.
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// Specifies the version of the database engine.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// Snapshot Identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the Provisioned IOPS (I/O operations per second) value of the DB instance at the time of the snapshot.
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// KMS key ID.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// License model information for the restored DB instance.
	LicenseModel *string `json:"licenseModel,omitempty" tf:"license_model,omitempty"`

	// The name of an option group to associate with the copy of the snapshot.
	OptionGroupName *string `json:"optionGroupName,omitempty" tf:"option_group_name,omitempty"`

	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// he URL that contains a Signature Version 4 signed request.
	PresignedURL *string `json:"presignedUrl,omitempty" tf:"presigned_url,omitempty"`

	// List of AWS Account ids to share snapshot with, use all to make snaphot public.
	// +listType=set
	SharedAccounts []*string `json:"sharedAccounts,omitempty" tf:"shared_accounts,omitempty"`

	SnapshotType *string `json:"snapshotType,omitempty" tf:"snapshot_type,omitempty"`

	// Snapshot identifier of the source snapshot.
	SourceDBSnapshotIdentifier *string `json:"sourceDbSnapshotIdentifier,omitempty" tf:"source_db_snapshot_identifier,omitempty"`

	// The region that the DB snapshot was created in or copied from.
	SourceRegion *string `json:"sourceRegion,omitempty" tf:"source_region,omitempty"`

	// Specifies the storage type associated with DB snapshot.
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The external custom Availability Zone.
	TargetCustomAvailabilityZone *string `json:"targetCustomAvailabilityZone,omitempty" tf:"target_custom_availability_zone,omitempty"`

	// The Identifier for the snapshot.
	TargetDBSnapshotIdentifier *string `json:"targetDbSnapshotIdentifier,omitempty" tf:"target_db_snapshot_identifier,omitempty"`

	// Provides the VPC ID associated with the DB snapshot.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type DBSnapshotCopyParameters struct {

	// Whether to copy existing tags. Defaults to false.
	// +kubebuilder:validation:Optional
	CopyTags *bool `json:"copyTags,omitempty" tf:"copy_tags,omitempty"`

	// The Destination region to place snapshot copy.
	// +kubebuilder:validation:Optional
	DestinationRegion *string `json:"destinationRegion,omitempty" tf:"destination_region,omitempty"`

	// KMS key ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// The name of an option group to associate with the copy of the snapshot.
	// +kubebuilder:validation:Optional
	OptionGroupName *string `json:"optionGroupName,omitempty" tf:"option_group_name,omitempty"`

	// he URL that contains a Signature Version 4 signed request.
	// +kubebuilder:validation:Optional
	PresignedURL *string `json:"presignedUrl,omitempty" tf:"presigned_url,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// List of AWS Account ids to share snapshot with, use all to make snaphot public.
	// +kubebuilder:validation:Optional
	// +listType=set
	SharedAccounts []*string `json:"sharedAccounts,omitempty" tf:"shared_accounts,omitempty"`

	// Snapshot identifier of the source snapshot.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/rds/v1beta1.Snapshot
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("db_snapshot_arn",true)
	// +kubebuilder:validation:Optional
	SourceDBSnapshotIdentifier *string `json:"sourceDbSnapshotIdentifier,omitempty" tf:"source_db_snapshot_identifier,omitempty"`

	// Reference to a Snapshot in rds to populate sourceDbSnapshotIdentifier.
	// +kubebuilder:validation:Optional
	SourceDBSnapshotIdentifierRef *v1.Reference `json:"sourceDbSnapshotIdentifierRef,omitempty" tf:"-"`

	// Selector for a Snapshot in rds to populate sourceDbSnapshotIdentifier.
	// +kubebuilder:validation:Optional
	SourceDBSnapshotIdentifierSelector *v1.Selector `json:"sourceDbSnapshotIdentifierSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The external custom Availability Zone.
	// +kubebuilder:validation:Optional
	TargetCustomAvailabilityZone *string `json:"targetCustomAvailabilityZone,omitempty" tf:"target_custom_availability_zone,omitempty"`

	// The Identifier for the snapshot.
	// +kubebuilder:validation:Optional
	TargetDBSnapshotIdentifier *string `json:"targetDbSnapshotIdentifier,omitempty" tf:"target_db_snapshot_identifier,omitempty"`
}

// DBSnapshotCopySpec defines the desired state of DBSnapshotCopy
type DBSnapshotCopySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DBSnapshotCopyParameters `json:"forProvider"`
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
	InitProvider DBSnapshotCopyInitParameters `json:"initProvider,omitempty"`
}

// DBSnapshotCopyStatus defines the observed state of DBSnapshotCopy.
type DBSnapshotCopyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DBSnapshotCopyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DBSnapshotCopy is the Schema for the DBSnapshotCopys API. Manages an RDS database instance snapshot copy.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DBSnapshotCopy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.targetDbSnapshotIdentifier) || (has(self.initProvider) && has(self.initProvider.targetDbSnapshotIdentifier))",message="spec.forProvider.targetDbSnapshotIdentifier is a required parameter"
	Spec   DBSnapshotCopySpec   `json:"spec"`
	Status DBSnapshotCopyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DBSnapshotCopyList contains a list of DBSnapshotCopys
type DBSnapshotCopyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DBSnapshotCopy `json:"items"`
}

// Repository type metadata.
var (
	DBSnapshotCopy_Kind             = "DBSnapshotCopy"
	DBSnapshotCopy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DBSnapshotCopy_Kind}.String()
	DBSnapshotCopy_KindAPIVersion   = DBSnapshotCopy_Kind + "." + CRDGroupVersion.String()
	DBSnapshotCopy_GroupVersionKind = CRDGroupVersion.WithKind(DBSnapshotCopy_Kind)
)

func init() {
	SchemeBuilder.Register(&DBSnapshotCopy{}, &DBSnapshotCopyList{})
}

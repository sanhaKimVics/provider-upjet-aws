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

type AccessControlTranslationObservation struct {
}

type AccessControlTranslationParameters struct {

	// +kubebuilder:validation:Required
	Owner *string `json:"owner" tf:"owner,omitempty"`
}

type ApplyServerSideEncryptionByDefaultObservation struct {
}

type ApplyServerSideEncryptionByDefaultParameters struct {

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/kms/v1alpha2.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	KMSMasterKeyID *string `json:"kmsMasterKeyId,omitempty" tf:"kms_master_key_id,omitempty"`

	// +kubebuilder:validation:Optional
	KMSMasterKeyIDRef *v1.Reference `json:"kmsMasterKeyIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	KMSMasterKeyIDSelector *v1.Selector `json:"kmsMasterKeyIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	SseAlgorithm *string `json:"sseAlgorithm" tf:"sse_algorithm,omitempty"`
}

type BucketObservation struct {
	BucketDomainName *string `json:"bucketDomainName,omitempty" tf:"bucket_domain_name,omitempty"`

	BucketRegionalDomainName *string `json:"bucketRegionalDomainName,omitempty" tf:"bucket_regional_domain_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	WebsiteDomain *string `json:"websiteDomain,omitempty" tf:"website_domain,omitempty"`

	WebsiteEndpoint *string `json:"websiteEndpoint,omitempty" tf:"website_endpoint,omitempty"`
}

type BucketParameters struct {

	// +kubebuilder:validation:Optional
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// +kubebuilder:validation:Optional
	AccelerationStatus *string `json:"accelerationStatus,omitempty" tf:"acceleration_status,omitempty"`

	// +kubebuilder:validation:Optional
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// +kubebuilder:validation:Optional
	CorsRule []CorsRuleParameters `json:"corsRule,omitempty" tf:"cors_rule,omitempty"`

	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// +kubebuilder:validation:Optional
	Grant []GrantParameters `json:"grant,omitempty" tf:"grant,omitempty"`

	// +kubebuilder:validation:Optional
	HostedZoneID *string `json:"hostedZoneId,omitempty" tf:"hosted_zone_id,omitempty"`

	// +kubebuilder:validation:Optional
	LifecycleRule []LifecycleRuleParameters `json:"lifecycleRule,omitempty" tf:"lifecycle_rule,omitempty"`

	// +kubebuilder:validation:Optional
	Logging []LoggingParameters `json:"logging,omitempty" tf:"logging,omitempty"`

	// +kubebuilder:validation:Optional
	ObjectLockConfiguration []ObjectLockConfigurationParameters `json:"objectLockConfiguration,omitempty" tf:"object_lock_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	ObjectLockEnabled *bool `json:"objectLockEnabled,omitempty" tf:"object_lock_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	ReplicationConfiguration []ReplicationConfigurationParameters `json:"replicationConfiguration,omitempty" tf:"replication_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	RequestPayer *string `json:"requestPayer,omitempty" tf:"request_payer,omitempty"`

	// +kubebuilder:validation:Optional
	ServerSideEncryptionConfiguration []ServerSideEncryptionConfigurationParameters `json:"serverSideEncryptionConfiguration,omitempty" tf:"server_side_encryption_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Versioning []VersioningParameters `json:"versioning,omitempty" tf:"versioning,omitempty"`

	// +kubebuilder:validation:Optional
	Website []WebsiteParameters `json:"website,omitempty" tf:"website,omitempty"`
}

type CorsRuleObservation struct {
}

type CorsRuleParameters struct {

	// +kubebuilder:validation:Optional
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// +kubebuilder:validation:Required
	AllowedMethods []*string `json:"allowedMethods" tf:"allowed_methods,omitempty"`

	// +kubebuilder:validation:Required
	AllowedOrigins []*string `json:"allowedOrigins" tf:"allowed_origins,omitempty"`

	// +kubebuilder:validation:Optional
	ExposeHeaders []*string `json:"exposeHeaders,omitempty" tf:"expose_headers,omitempty"`

	// +kubebuilder:validation:Optional
	MaxAgeSeconds *float64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`
}

type DefaultRetentionObservation struct {
}

type DefaultRetentionParameters struct {

	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// +kubebuilder:validation:Optional
	Years *float64 `json:"years,omitempty" tf:"years,omitempty"`
}

type DestinationObservation struct {
}

type DestinationParameters struct {

	// +kubebuilder:validation:Optional
	AccessControlTranslation []AccessControlTranslationParameters `json:"accessControlTranslation,omitempty" tf:"access_control_translation,omitempty"`

	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// +kubebuilder:validation:Required
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// +kubebuilder:validation:Optional
	Metrics []MetricsParameters `json:"metrics,omitempty" tf:"metrics,omitempty"`

	// +kubebuilder:validation:Optional
	ReplicaKMSKeyID *string `json:"replicaKmsKeyId,omitempty" tf:"replica_kms_key_id,omitempty"`

	// +kubebuilder:validation:Optional
	ReplicationTime []ReplicationTimeParameters `json:"replicationTime,omitempty" tf:"replication_time,omitempty"`

	// +kubebuilder:validation:Optional
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type ExpirationObservation struct {
}

type ExpirationParameters struct {

	// +kubebuilder:validation:Optional
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// +kubebuilder:validation:Optional
	ExpiredObjectDeleteMarker *bool `json:"expiredObjectDeleteMarker,omitempty" tf:"expired_object_delete_marker,omitempty"`
}

type FilterObservation struct {
}

type FilterParameters struct {

	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type GrantObservation struct {
}

type GrantParameters struct {

	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:Required
	Permissions []*string `json:"permissions" tf:"permissions,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type LifecycleRuleObservation struct {
}

type LifecycleRuleParameters struct {

	// +kubebuilder:validation:Optional
	AbortIncompleteMultipartUploadDays *float64 `json:"abortIncompleteMultipartUploadDays,omitempty" tf:"abort_incomplete_multipart_upload_days,omitempty"`

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Expiration []ExpirationParameters `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	NoncurrentVersionExpiration []NoncurrentVersionExpirationParameters `json:"noncurrentVersionExpiration,omitempty" tf:"noncurrent_version_expiration,omitempty"`

	// +kubebuilder:validation:Optional
	NoncurrentVersionTransition []NoncurrentVersionTransitionParameters `json:"noncurrentVersionTransition,omitempty" tf:"noncurrent_version_transition,omitempty"`

	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Transition []TransitionParameters `json:"transition,omitempty" tf:"transition,omitempty"`
}

type LoggingObservation struct {
}

type LoggingParameters struct {

	// +kubebuilder:validation:Required
	TargetBucket *string `json:"targetBucket" tf:"target_bucket,omitempty"`

	// +kubebuilder:validation:Optional
	TargetPrefix *string `json:"targetPrefix,omitempty" tf:"target_prefix,omitempty"`
}

type MetricsObservation struct {
}

type MetricsParameters struct {

	// +kubebuilder:validation:Optional
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type NoncurrentVersionExpirationObservation struct {
}

type NoncurrentVersionExpirationParameters struct {

	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`
}

type NoncurrentVersionTransitionObservation struct {
}

type NoncurrentVersionTransitionParameters struct {

	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// +kubebuilder:validation:Required
	StorageClass *string `json:"storageClass" tf:"storage_class,omitempty"`
}

type ObjectLockConfigurationObservation struct {
}

type ObjectLockConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	ObjectLockEnabled *string `json:"objectLockEnabled,omitempty" tf:"object_lock_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Rule []RuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type ReplicationConfigurationObservation struct {
}

type ReplicationConfigurationParameters struct {

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// +kubebuilder:validation:Required
	Rules []RulesParameters `json:"rules" tf:"rules,omitempty"`
}

type ReplicationTimeObservation struct {
}

type ReplicationTimeParameters struct {

	// +kubebuilder:validation:Optional
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type RuleObservation struct {
}

type RuleParameters struct {

	// +kubebuilder:validation:Required
	DefaultRetention []DefaultRetentionParameters `json:"defaultRetention" tf:"default_retention,omitempty"`
}

type RulesObservation struct {
}

type RulesParameters struct {

	// +kubebuilder:validation:Optional
	DeleteMarkerReplicationStatus *string `json:"deleteMarkerReplicationStatus,omitempty" tf:"delete_marker_replication_status,omitempty"`

	// +kubebuilder:validation:Required
	Destination []DestinationParameters `json:"destination" tf:"destination,omitempty"`

	// +kubebuilder:validation:Optional
	Filter []FilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// +kubebuilder:validation:Optional
	SourceSelectionCriteria []SourceSelectionCriteriaParameters `json:"sourceSelectionCriteria,omitempty" tf:"source_selection_criteria,omitempty"`

	// +kubebuilder:validation:Required
	Status *string `json:"status" tf:"status,omitempty"`
}

type ServerSideEncryptionConfigurationObservation struct {
}

type ServerSideEncryptionConfigurationParameters struct {

	// +kubebuilder:validation:Required
	Rule []ServerSideEncryptionConfigurationRuleParameters `json:"rule" tf:"rule,omitempty"`
}

type ServerSideEncryptionConfigurationRuleObservation struct {
}

type ServerSideEncryptionConfigurationRuleParameters struct {

	// +kubebuilder:validation:Required
	ApplyServerSideEncryptionByDefault []ApplyServerSideEncryptionByDefaultParameters `json:"applyServerSideEncryptionByDefault" tf:"apply_server_side_encryption_by_default,omitempty"`

	// +kubebuilder:validation:Optional
	BucketKeyEnabled *bool `json:"bucketKeyEnabled,omitempty" tf:"bucket_key_enabled,omitempty"`
}

type SourceSelectionCriteriaObservation struct {
}

type SourceSelectionCriteriaParameters struct {

	// +kubebuilder:validation:Optional
	SseKMSEncryptedObjects []SseKMSEncryptedObjectsParameters `json:"sseKmsEncryptedObjects,omitempty" tf:"sse_kms_encrypted_objects,omitempty"`
}

type SseKMSEncryptedObjectsObservation struct {
}

type SseKMSEncryptedObjectsParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type TransitionObservation struct {
}

type TransitionParameters struct {

	// +kubebuilder:validation:Optional
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// +kubebuilder:validation:Required
	StorageClass *string `json:"storageClass" tf:"storage_class,omitempty"`
}

type VersioningObservation struct {
}

type VersioningParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	MfaDelete *bool `json:"mfaDelete,omitempty" tf:"mfa_delete,omitempty"`
}

type WebsiteObservation struct {
}

type WebsiteParameters struct {

	// +kubebuilder:validation:Optional
	ErrorDocument *string `json:"errorDocument,omitempty" tf:"error_document,omitempty"`

	// +kubebuilder:validation:Optional
	IndexDocument *string `json:"indexDocument,omitempty" tf:"index_document,omitempty"`

	// +kubebuilder:validation:Optional
	RedirectAllRequestsTo *string `json:"redirectAllRequestsTo,omitempty" tf:"redirect_all_requests_to,omitempty"`

	// +kubebuilder:validation:Optional
	RoutingRules *string `json:"routingRules,omitempty" tf:"routing_rules,omitempty"`
}

// BucketSpec defines the desired state of Bucket
type BucketSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketParameters `json:"forProvider"`
}

// BucketStatus defines the observed state of Bucket.
type BucketStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Bucket is the Schema for the Buckets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Bucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketSpec   `json:"spec"`
	Status            BucketStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketList contains a list of Buckets
type BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bucket `json:"items"`
}

// Repository type metadata.
var (
	Bucket_Kind             = "Bucket"
	Bucket_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Bucket_Kind}.String()
	Bucket_KindAPIVersion   = Bucket_Kind + "." + CRDGroupVersion.String()
	Bucket_GroupVersionKind = CRDGroupVersion.WithKind(Bucket_Kind)
)

func init() {
	SchemeBuilder.Register(&Bucket{}, &BucketList{})
}

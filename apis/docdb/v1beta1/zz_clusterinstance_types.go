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

type ClusterInstanceObservation struct {

	// Amazon Resource Name (ARN) of cluster instance
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The DB subnet group to associate with this DB instance.
	DBSubnetGroupName *string `json:"dbSubnetGroupName,omitempty" tf:"db_subnet_group_name,omitempty"`

	// The region-unique, immutable identifier for the DB instance.
	DbiResourceID *string `json:"dbiResourceId,omitempty" tf:"dbi_resource_id,omitempty"`

	// The DNS address for this instance. May not be writable
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The database engine version
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ARN for the KMS encryption key if one is set to the cluster.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// The database port
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The daily time range during which automated backups are created if automated backups are enabled.
	PreferredBackupWindow *string `json:"preferredBackupWindow,omitempty" tf:"preferred_backup_window,omitempty"`

	PubliclyAccessible *bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`

	// Specifies whether the DB cluster is encrypted.
	StorageEncrypted *bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// – Boolean indicating if this instance is writable. False indicates this instance is a read replica.
	Writer *bool `json:"writer,omitempty" tf:"writer,omitempty"`
}

type ClusterInstanceParameters struct {

	// Specifies whether any database modifications
	// are applied immediately, or during the next maintenance window. Default isfalse.
	// +kubebuilder:validation:Optional
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default true.
	// +kubebuilder:validation:Optional
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`

	// The EC2 Availability Zone that the DB instance is created in. See docs about the details.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The identifier of the CA certificate for the DB instance.
	// +kubebuilder:validation:Optional
	CACertIdentifier *string `json:"caCertIdentifier,omitempty" tf:"ca_cert_identifier,omitempty"`

	// The identifier of the aws_docdb_cluster in which to launch this instance.
	// +crossplane:generate:reference:type=Cluster
	// +kubebuilder:validation:Optional
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier,omitempty"`

	// Reference to a Cluster to populate clusterIdentifier.
	// +kubebuilder:validation:Optional
	ClusterIdentifierRef *v1.Reference `json:"clusterIdentifierRef,omitempty" tf:"-"`

	// Selector for a Cluster to populate clusterIdentifier.
	// +kubebuilder:validation:Optional
	ClusterIdentifierSelector *v1.Selector `json:"clusterIdentifierSelector,omitempty" tf:"-"`

	// The name of the database engine to be used for the DocDB instance. Defaults to docdb. Valid Values: docdb.
	// +kubebuilder:validation:Optional
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// The instance class to use. For details on CPU and memory, see Scaling for DocDB Instances. DocDB currently
	// supports the below instance classes. Please see AWS Documentation for complete details.
	// +kubebuilder:validation:Required
	InstanceClass *string `json:"instanceClass" tf:"instance_class,omitempty"`

	// The window to perform maintenance in.
	// Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
	// +kubebuilder:validation:Optional
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window,omitempty"`

	// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
	// +kubebuilder:validation:Optional
	PromotionTier *float64 `json:"promotionTier,omitempty" tf:"promotion_tier,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A map of tags to assign to the instance. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ClusterInstanceSpec defines the desired state of ClusterInstance
type ClusterInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterInstanceParameters `json:"forProvider"`
}

// ClusterInstanceStatus defines the observed state of ClusterInstance.
type ClusterInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterInstance is the Schema for the ClusterInstances API. Provides an DocDB Cluster Resource Instance
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ClusterInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterInstanceSpec   `json:"spec"`
	Status            ClusterInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterInstanceList contains a list of ClusterInstances
type ClusterInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterInstance `json:"items"`
}

// Repository type metadata.
var (
	ClusterInstance_Kind             = "ClusterInstance"
	ClusterInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClusterInstance_Kind}.String()
	ClusterInstance_KindAPIVersion   = ClusterInstance_Kind + "." + CRDGroupVersion.String()
	ClusterInstance_GroupVersionKind = CRDGroupVersion.WithKind(ClusterInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&ClusterInstance{}, &ClusterInstanceList{})
}

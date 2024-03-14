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

type AccountInitParameters struct {

	// ARN of an IAM role for CloudWatch (to allow logging & monitoring). See more in AWS Docs. Logging & monitoring can be enabled/disabled and otherwise tuned on the API Gateway Stage level.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	CloudwatchRoleArn *string `json:"cloudwatchRoleArn,omitempty" tf:"cloudwatch_role_arn,omitempty"`

	// Reference to a Role in iam to populate cloudwatchRoleArn.
	// +kubebuilder:validation:Optional
	CloudwatchRoleArnRef *v1.Reference `json:"cloudwatchRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate cloudwatchRoleArn.
	// +kubebuilder:validation:Optional
	CloudwatchRoleArnSelector *v1.Selector `json:"cloudwatchRoleArnSelector,omitempty" tf:"-"`
}

type AccountObservation struct {

	// The version of the API keys used for the account.
	APIKeyVersion *string `json:"apiKeyVersion,omitempty" tf:"api_key_version,omitempty"`

	// ARN of an IAM role for CloudWatch (to allow logging & monitoring). See more in AWS Docs. Logging & monitoring can be enabled/disabled and otherwise tuned on the API Gateway Stage level.
	CloudwatchRoleArn *string `json:"cloudwatchRoleArn,omitempty" tf:"cloudwatch_role_arn,omitempty"`

	// A list of features supported for the account.
	// +listType=set
	Features []*string `json:"features,omitempty" tf:"features,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Account-Level throttle settings. See exported fields below.
	ThrottleSettings []ThrottleSettingsObservation `json:"throttleSettings,omitempty" tf:"throttle_settings,omitempty"`
}

type AccountParameters struct {

	// ARN of an IAM role for CloudWatch (to allow logging & monitoring). See more in AWS Docs. Logging & monitoring can be enabled/disabled and otherwise tuned on the API Gateway Stage level.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	CloudwatchRoleArn *string `json:"cloudwatchRoleArn,omitempty" tf:"cloudwatch_role_arn,omitempty"`

	// Reference to a Role in iam to populate cloudwatchRoleArn.
	// +kubebuilder:validation:Optional
	CloudwatchRoleArnRef *v1.Reference `json:"cloudwatchRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate cloudwatchRoleArn.
	// +kubebuilder:validation:Optional
	CloudwatchRoleArnSelector *v1.Selector `json:"cloudwatchRoleArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type ThrottleSettingsInitParameters struct {
}

type ThrottleSettingsObservation struct {

	// Absolute maximum number of times API Gateway allows the API to be called per second (RPS).
	BurstLimit *float64 `json:"burstLimit,omitempty" tf:"burst_limit,omitempty"`

	// Number of times API Gateway allows the API to be called per second on average (RPS).
	RateLimit *float64 `json:"rateLimit,omitempty" tf:"rate_limit,omitempty"`
}

type ThrottleSettingsParameters struct {
}

// AccountSpec defines the desired state of Account
type AccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccountParameters `json:"forProvider"`
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
	InitProvider AccountInitParameters `json:"initProvider,omitempty"`
}

// AccountStatus defines the observed state of Account.
type AccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Account is the Schema for the Accounts API. Provides a settings of an API Gateway Account.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Account struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccountSpec   `json:"spec"`
	Status            AccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccountList contains a list of Accounts
type AccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Account `json:"items"`
}

// Repository type metadata.
var (
	Account_Kind             = "Account"
	Account_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Account_Kind}.String()
	Account_KindAPIVersion   = Account_Kind + "." + CRDGroupVersion.String()
	Account_GroupVersionKind = CRDGroupVersion.WithKind(Account_Kind)
)

func init() {
	SchemeBuilder.Register(&Account{}, &AccountList{})
}

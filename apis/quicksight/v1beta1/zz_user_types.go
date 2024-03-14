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

type UserInitParameters struct {

	// The ID for the AWS account that the user is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
	AwsAccountID *string `json:"awsAccountId,omitempty" tf:"aws_account_id,omitempty"`

	// The email address of the user that you want to register.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The ARN of the IAM user or role that you are registering with Amazon QuickSight.
	IAMArn *string `json:"iamArn,omitempty" tf:"iam_arn,omitempty"`

	// Amazon QuickSight supports several ways of managing the identity of users. This parameter accepts either  IAM or QUICKSIGHT. If IAM is specified, the iam_arn must also be specified.
	IdentityType *string `json:"identityType,omitempty" tf:"identity_type,omitempty"`

	// The Amazon Quicksight namespace to create the user in. Defaults to default.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The name of the IAM session to use when assuming roles that can embed QuickSight dashboards. Only valid for registering users using an assumed IAM role. Additionally, if registering multiple users using the same IAM role, each user needs to have a unique session name.
	SessionName *string `json:"sessionName,omitempty" tf:"session_name,omitempty"`

	// The Amazon QuickSight user name that you want to create for the user you are registering. Only valid for registering a user with identity_type set to QUICKSIGHT.
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`

	// The Amazon QuickSight role of the user. The user role can be one of the following: READER, AUTHOR, or ADMIN
	UserRole *string `json:"userRole,omitempty" tf:"user_role,omitempty"`
}

type UserObservation struct {

	// Amazon Resource Name (ARN) of the user
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ID for the AWS account that the user is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
	AwsAccountID *string `json:"awsAccountId,omitempty" tf:"aws_account_id,omitempty"`

	// The email address of the user that you want to register.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The ARN of the IAM user or role that you are registering with Amazon QuickSight.
	IAMArn *string `json:"iamArn,omitempty" tf:"iam_arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Amazon QuickSight supports several ways of managing the identity of users. This parameter accepts either  IAM or QUICKSIGHT. If IAM is specified, the iam_arn must also be specified.
	IdentityType *string `json:"identityType,omitempty" tf:"identity_type,omitempty"`

	// The Amazon Quicksight namespace to create the user in. Defaults to default.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The name of the IAM session to use when assuming roles that can embed QuickSight dashboards. Only valid for registering users using an assumed IAM role. Additionally, if registering multiple users using the same IAM role, each user needs to have a unique session name.
	SessionName *string `json:"sessionName,omitempty" tf:"session_name,omitempty"`

	// The Amazon QuickSight user name that you want to create for the user you are registering. Only valid for registering a user with identity_type set to QUICKSIGHT.
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`

	// The Amazon QuickSight role of the user. The user role can be one of the following: READER, AUTHOR, or ADMIN
	UserRole *string `json:"userRole,omitempty" tf:"user_role,omitempty"`
}

type UserParameters struct {

	// The ID for the AWS account that the user is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
	// +kubebuilder:validation:Optional
	AwsAccountID *string `json:"awsAccountId,omitempty" tf:"aws_account_id,omitempty"`

	// The email address of the user that you want to register.
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The ARN of the IAM user or role that you are registering with Amazon QuickSight.
	// +kubebuilder:validation:Optional
	IAMArn *string `json:"iamArn,omitempty" tf:"iam_arn,omitempty"`

	// Amazon QuickSight supports several ways of managing the identity of users. This parameter accepts either  IAM or QUICKSIGHT. If IAM is specified, the iam_arn must also be specified.
	// +kubebuilder:validation:Optional
	IdentityType *string `json:"identityType,omitempty" tf:"identity_type,omitempty"`

	// The Amazon Quicksight namespace to create the user in. Defaults to default.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The name of the IAM session to use when assuming roles that can embed QuickSight dashboards. Only valid for registering users using an assumed IAM role. Additionally, if registering multiple users using the same IAM role, each user needs to have a unique session name.
	// +kubebuilder:validation:Optional
	SessionName *string `json:"sessionName,omitempty" tf:"session_name,omitempty"`

	// The Amazon QuickSight user name that you want to create for the user you are registering. Only valid for registering a user with identity_type set to QUICKSIGHT.
	// +kubebuilder:validation:Optional
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`

	// The Amazon QuickSight role of the user. The user role can be one of the following: READER, AUTHOR, or ADMIN
	// +kubebuilder:validation:Optional
	UserRole *string `json:"userRole,omitempty" tf:"user_role,omitempty"`
}

// UserSpec defines the desired state of User
type UserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserParameters `json:"forProvider"`
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
	InitProvider UserInitParameters `json:"initProvider,omitempty"`
}

// UserStatus defines the observed state of User.
type UserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// User is the Schema for the Users API. Manages a Resource QuickSight User.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.email) || (has(self.initProvider) && has(self.initProvider.email))",message="spec.forProvider.email is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.identityType) || (has(self.initProvider) && has(self.initProvider.identityType))",message="spec.forProvider.identityType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.userRole) || (has(self.initProvider) && has(self.initProvider.userRole))",message="spec.forProvider.userRole is a required parameter"
	Spec   UserSpec   `json:"spec"`
	Status UserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserList contains a list of Users
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User `json:"items"`
}

// Repository type metadata.
var (
	User_Kind             = "User"
	User_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: User_Kind}.String()
	User_KindAPIVersion   = User_Kind + "." + CRDGroupVersion.String()
	User_GroupVersionKind = CRDGroupVersion.WithKind(User_Kind)
)

func init() {
	SchemeBuilder.Register(&User{}, &UserList{})
}

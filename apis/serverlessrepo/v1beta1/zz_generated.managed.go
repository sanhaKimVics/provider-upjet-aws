/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this CloudFormationStack.
func (mg *CloudFormationStack) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this CloudFormationStack.
func (mg *CloudFormationStack) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this CloudFormationStack.
func (mg *CloudFormationStack) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this CloudFormationStack.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *CloudFormationStack) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this CloudFormationStack.
func (mg *CloudFormationStack) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this CloudFormationStack.
func (mg *CloudFormationStack) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this CloudFormationStack.
func (mg *CloudFormationStack) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this CloudFormationStack.
func (mg *CloudFormationStack) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this CloudFormationStack.
func (mg *CloudFormationStack) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this CloudFormationStack.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *CloudFormationStack) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this CloudFormationStack.
func (mg *CloudFormationStack) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this CloudFormationStack.
func (mg *CloudFormationStack) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
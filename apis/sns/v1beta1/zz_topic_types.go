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

type TopicInitParameters struct {

	// IAM role for failure feedback
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	ApplicationFailureFeedbackRoleArn *string `json:"applicationFailureFeedbackRoleArn,omitempty" tf:"application_failure_feedback_role_arn,omitempty"`

	// Reference to a Role in iam to populate applicationFailureFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	ApplicationFailureFeedbackRoleArnRef *v1.Reference `json:"applicationFailureFeedbackRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate applicationFailureFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	ApplicationFailureFeedbackRoleArnSelector *v1.Selector `json:"applicationFailureFeedbackRoleArnSelector,omitempty" tf:"-"`

	// The IAM role permitted to receive success feedback for this topic
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	ApplicationSuccessFeedbackRoleArn *string `json:"applicationSuccessFeedbackRoleArn,omitempty" tf:"application_success_feedback_role_arn,omitempty"`

	// Reference to a Role in iam to populate applicationSuccessFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	ApplicationSuccessFeedbackRoleArnRef *v1.Reference `json:"applicationSuccessFeedbackRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate applicationSuccessFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	ApplicationSuccessFeedbackRoleArnSelector *v1.Selector `json:"applicationSuccessFeedbackRoleArnSelector,omitempty" tf:"-"`

	// Percentage of success to sample
	ApplicationSuccessFeedbackSampleRate *float64 `json:"applicationSuccessFeedbackSampleRate,omitempty" tf:"application_success_feedback_sample_rate,omitempty"`

	// The message archive policy for FIFO topics. More details in the AWS documentation.
	ArchivePolicy *string `json:"archivePolicy,omitempty" tf:"archive_policy,omitempty"`

	// Enables content-based deduplication for FIFO topics. For more information, see the related documentation
	ContentBasedDeduplication *bool `json:"contentBasedDeduplication,omitempty" tf:"content_based_deduplication,omitempty"`

	// The SNS delivery policy. More details in the AWS documentation.
	DeliveryPolicy *string `json:"deliveryPolicy,omitempty" tf:"delivery_policy,omitempty"`

	// The display name for the topic
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Boolean indicating whether or not to create a FIFO (first-in-first-out) topic (default is false).
	FifoTopic *bool `json:"fifoTopic,omitempty" tf:"fifo_topic,omitempty"`

	// IAM role for failure feedback
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	FirehoseFailureFeedbackRoleArn *string `json:"firehoseFailureFeedbackRoleArn,omitempty" tf:"firehose_failure_feedback_role_arn,omitempty"`

	// Reference to a Role in iam to populate firehoseFailureFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	FirehoseFailureFeedbackRoleArnRef *v1.Reference `json:"firehoseFailureFeedbackRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate firehoseFailureFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	FirehoseFailureFeedbackRoleArnSelector *v1.Selector `json:"firehoseFailureFeedbackRoleArnSelector,omitempty" tf:"-"`

	// The IAM role permitted to receive success feedback for this topic
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	FirehoseSuccessFeedbackRoleArn *string `json:"firehoseSuccessFeedbackRoleArn,omitempty" tf:"firehose_success_feedback_role_arn,omitempty"`

	// Reference to a Role in iam to populate firehoseSuccessFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	FirehoseSuccessFeedbackRoleArnRef *v1.Reference `json:"firehoseSuccessFeedbackRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate firehoseSuccessFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	FirehoseSuccessFeedbackRoleArnSelector *v1.Selector `json:"firehoseSuccessFeedbackRoleArnSelector,omitempty" tf:"-"`

	// Percentage of success to sample
	FirehoseSuccessFeedbackSampleRate *float64 `json:"firehoseSuccessFeedbackSampleRate,omitempty" tf:"firehose_success_feedback_sample_rate,omitempty"`

	// IAM role for failure feedback
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	HTTPFailureFeedbackRoleArn *string `json:"httpFailureFeedbackRoleArn,omitempty" tf:"http_failure_feedback_role_arn,omitempty"`

	// Reference to a Role in iam to populate httpFailureFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	HTTPFailureFeedbackRoleArnRef *v1.Reference `json:"httpFailureFeedbackRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate httpFailureFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	HTTPFailureFeedbackRoleArnSelector *v1.Selector `json:"httpFailureFeedbackRoleArnSelector,omitempty" tf:"-"`

	// The IAM role permitted to receive success feedback for this topic
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	HTTPSuccessFeedbackRoleArn *string `json:"httpSuccessFeedbackRoleArn,omitempty" tf:"http_success_feedback_role_arn,omitempty"`

	// Reference to a Role in iam to populate httpSuccessFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	HTTPSuccessFeedbackRoleArnRef *v1.Reference `json:"httpSuccessFeedbackRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate httpSuccessFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	HTTPSuccessFeedbackRoleArnSelector *v1.Selector `json:"httpSuccessFeedbackRoleArnSelector,omitempty" tf:"-"`

	// Percentage of success to sample
	HTTPSuccessFeedbackSampleRate *float64 `json:"httpSuccessFeedbackSampleRate,omitempty" tf:"http_success_feedback_sample_rate,omitempty"`

	// The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see Key Terms
	KMSMasterKeyID *string `json:"kmsMasterKeyId,omitempty" tf:"kms_master_key_id,omitempty"`

	// IAM role for failure feedback
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	LambdaFailureFeedbackRoleArn *string `json:"lambdaFailureFeedbackRoleArn,omitempty" tf:"lambda_failure_feedback_role_arn,omitempty"`

	// Reference to a Role in iam to populate lambdaFailureFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	LambdaFailureFeedbackRoleArnRef *v1.Reference `json:"lambdaFailureFeedbackRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate lambdaFailureFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	LambdaFailureFeedbackRoleArnSelector *v1.Selector `json:"lambdaFailureFeedbackRoleArnSelector,omitempty" tf:"-"`

	// The IAM role permitted to receive success feedback for this topic
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	LambdaSuccessFeedbackRoleArn *string `json:"lambdaSuccessFeedbackRoleArn,omitempty" tf:"lambda_success_feedback_role_arn,omitempty"`

	// Reference to a Role in iam to populate lambdaSuccessFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	LambdaSuccessFeedbackRoleArnRef *v1.Reference `json:"lambdaSuccessFeedbackRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate lambdaSuccessFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	LambdaSuccessFeedbackRoleArnSelector *v1.Selector `json:"lambdaSuccessFeedbackRoleArnSelector,omitempty" tf:"-"`

	// Percentage of success to sample
	LambdaSuccessFeedbackSampleRate *float64 `json:"lambdaSuccessFeedbackSampleRate,omitempty" tf:"lambda_success_feedback_sample_rate,omitempty"`

	// The fully-formed AWS policy as JSON.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// If SignatureVersion should be 1 (SHA1) or 2 (SHA256). The signature version corresponds to the hashing algorithm used while creating the signature of the notifications, subscription confirmations, or unsubscribe confirmation messages sent by Amazon SNS.
	SignatureVersion *float64 `json:"signatureVersion,omitempty" tf:"signature_version,omitempty"`

	// IAM role for failure feedback
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	SqsFailureFeedbackRoleArn *string `json:"sqsFailureFeedbackRoleArn,omitempty" tf:"sqs_failure_feedback_role_arn,omitempty"`

	// Reference to a Role in iam to populate sqsFailureFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	SqsFailureFeedbackRoleArnRef *v1.Reference `json:"sqsFailureFeedbackRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate sqsFailureFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	SqsFailureFeedbackRoleArnSelector *v1.Selector `json:"sqsFailureFeedbackRoleArnSelector,omitempty" tf:"-"`

	// The IAM role permitted to receive success feedback for this topic
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	SqsSuccessFeedbackRoleArn *string `json:"sqsSuccessFeedbackRoleArn,omitempty" tf:"sqs_success_feedback_role_arn,omitempty"`

	// Reference to a Role in iam to populate sqsSuccessFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	SqsSuccessFeedbackRoleArnRef *v1.Reference `json:"sqsSuccessFeedbackRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate sqsSuccessFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	SqsSuccessFeedbackRoleArnSelector *v1.Selector `json:"sqsSuccessFeedbackRoleArnSelector,omitempty" tf:"-"`

	// Percentage of success to sample
	SqsSuccessFeedbackSampleRate *float64 `json:"sqsSuccessFeedbackSampleRate,omitempty" tf:"sqs_success_feedback_sample_rate,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Tracing mode of an Amazon SNS topic. Valid values: "PassThrough", "Active".
	TracingConfig *string `json:"tracingConfig,omitempty" tf:"tracing_config,omitempty"`
}

type TopicObservation struct {

	// IAM role for failure feedback
	ApplicationFailureFeedbackRoleArn *string `json:"applicationFailureFeedbackRoleArn,omitempty" tf:"application_failure_feedback_role_arn,omitempty"`

	// The IAM role permitted to receive success feedback for this topic
	ApplicationSuccessFeedbackRoleArn *string `json:"applicationSuccessFeedbackRoleArn,omitempty" tf:"application_success_feedback_role_arn,omitempty"`

	// Percentage of success to sample
	ApplicationSuccessFeedbackSampleRate *float64 `json:"applicationSuccessFeedbackSampleRate,omitempty" tf:"application_success_feedback_sample_rate,omitempty"`

	// The message archive policy for FIFO topics. More details in the AWS documentation.
	ArchivePolicy *string `json:"archivePolicy,omitempty" tf:"archive_policy,omitempty"`

	// The ARN of the SNS topic, as a more obvious property (clone of id)
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The oldest timestamp at which a FIFO topic subscriber can start a replay.
	BeginningArchiveTime *string `json:"beginningArchiveTime,omitempty" tf:"beginning_archive_time,omitempty"`

	// Enables content-based deduplication for FIFO topics. For more information, see the related documentation
	ContentBasedDeduplication *bool `json:"contentBasedDeduplication,omitempty" tf:"content_based_deduplication,omitempty"`

	// The SNS delivery policy. More details in the AWS documentation.
	DeliveryPolicy *string `json:"deliveryPolicy,omitempty" tf:"delivery_policy,omitempty"`

	// The display name for the topic
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Boolean indicating whether or not to create a FIFO (first-in-first-out) topic (default is false).
	FifoTopic *bool `json:"fifoTopic,omitempty" tf:"fifo_topic,omitempty"`

	// IAM role for failure feedback
	FirehoseFailureFeedbackRoleArn *string `json:"firehoseFailureFeedbackRoleArn,omitempty" tf:"firehose_failure_feedback_role_arn,omitempty"`

	// The IAM role permitted to receive success feedback for this topic
	FirehoseSuccessFeedbackRoleArn *string `json:"firehoseSuccessFeedbackRoleArn,omitempty" tf:"firehose_success_feedback_role_arn,omitempty"`

	// Percentage of success to sample
	FirehoseSuccessFeedbackSampleRate *float64 `json:"firehoseSuccessFeedbackSampleRate,omitempty" tf:"firehose_success_feedback_sample_rate,omitempty"`

	// IAM role for failure feedback
	HTTPFailureFeedbackRoleArn *string `json:"httpFailureFeedbackRoleArn,omitempty" tf:"http_failure_feedback_role_arn,omitempty"`

	// The IAM role permitted to receive success feedback for this topic
	HTTPSuccessFeedbackRoleArn *string `json:"httpSuccessFeedbackRoleArn,omitempty" tf:"http_success_feedback_role_arn,omitempty"`

	// Percentage of success to sample
	HTTPSuccessFeedbackSampleRate *float64 `json:"httpSuccessFeedbackSampleRate,omitempty" tf:"http_success_feedback_sample_rate,omitempty"`

	// The ARN of the SNS topic
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see Key Terms
	KMSMasterKeyID *string `json:"kmsMasterKeyId,omitempty" tf:"kms_master_key_id,omitempty"`

	// IAM role for failure feedback
	LambdaFailureFeedbackRoleArn *string `json:"lambdaFailureFeedbackRoleArn,omitempty" tf:"lambda_failure_feedback_role_arn,omitempty"`

	// The IAM role permitted to receive success feedback for this topic
	LambdaSuccessFeedbackRoleArn *string `json:"lambdaSuccessFeedbackRoleArn,omitempty" tf:"lambda_success_feedback_role_arn,omitempty"`

	// Percentage of success to sample
	LambdaSuccessFeedbackSampleRate *float64 `json:"lambdaSuccessFeedbackSampleRate,omitempty" tf:"lambda_success_feedback_sample_rate,omitempty"`

	// The AWS Account ID of the SNS topic owner
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// The fully-formed AWS policy as JSON.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// If SignatureVersion should be 1 (SHA1) or 2 (SHA256). The signature version corresponds to the hashing algorithm used while creating the signature of the notifications, subscription confirmations, or unsubscribe confirmation messages sent by Amazon SNS.
	SignatureVersion *float64 `json:"signatureVersion,omitempty" tf:"signature_version,omitempty"`

	// IAM role for failure feedback
	SqsFailureFeedbackRoleArn *string `json:"sqsFailureFeedbackRoleArn,omitempty" tf:"sqs_failure_feedback_role_arn,omitempty"`

	// The IAM role permitted to receive success feedback for this topic
	SqsSuccessFeedbackRoleArn *string `json:"sqsSuccessFeedbackRoleArn,omitempty" tf:"sqs_success_feedback_role_arn,omitempty"`

	// Percentage of success to sample
	SqsSuccessFeedbackSampleRate *float64 `json:"sqsSuccessFeedbackSampleRate,omitempty" tf:"sqs_success_feedback_sample_rate,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Tracing mode of an Amazon SNS topic. Valid values: "PassThrough", "Active".
	TracingConfig *string `json:"tracingConfig,omitempty" tf:"tracing_config,omitempty"`
}

type TopicParameters struct {

	// IAM role for failure feedback
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	ApplicationFailureFeedbackRoleArn *string `json:"applicationFailureFeedbackRoleArn,omitempty" tf:"application_failure_feedback_role_arn,omitempty"`

	// Reference to a Role in iam to populate applicationFailureFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	ApplicationFailureFeedbackRoleArnRef *v1.Reference `json:"applicationFailureFeedbackRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate applicationFailureFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	ApplicationFailureFeedbackRoleArnSelector *v1.Selector `json:"applicationFailureFeedbackRoleArnSelector,omitempty" tf:"-"`

	// The IAM role permitted to receive success feedback for this topic
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	ApplicationSuccessFeedbackRoleArn *string `json:"applicationSuccessFeedbackRoleArn,omitempty" tf:"application_success_feedback_role_arn,omitempty"`

	// Reference to a Role in iam to populate applicationSuccessFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	ApplicationSuccessFeedbackRoleArnRef *v1.Reference `json:"applicationSuccessFeedbackRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate applicationSuccessFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	ApplicationSuccessFeedbackRoleArnSelector *v1.Selector `json:"applicationSuccessFeedbackRoleArnSelector,omitempty" tf:"-"`

	// Percentage of success to sample
	// +kubebuilder:validation:Optional
	ApplicationSuccessFeedbackSampleRate *float64 `json:"applicationSuccessFeedbackSampleRate,omitempty" tf:"application_success_feedback_sample_rate,omitempty"`

	// The message archive policy for FIFO topics. More details in the AWS documentation.
	// +kubebuilder:validation:Optional
	ArchivePolicy *string `json:"archivePolicy,omitempty" tf:"archive_policy,omitempty"`

	// Enables content-based deduplication for FIFO topics. For more information, see the related documentation
	// +kubebuilder:validation:Optional
	ContentBasedDeduplication *bool `json:"contentBasedDeduplication,omitempty" tf:"content_based_deduplication,omitempty"`

	// The SNS delivery policy. More details in the AWS documentation.
	// +kubebuilder:validation:Optional
	DeliveryPolicy *string `json:"deliveryPolicy,omitempty" tf:"delivery_policy,omitempty"`

	// The display name for the topic
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Boolean indicating whether or not to create a FIFO (first-in-first-out) topic (default is false).
	// +kubebuilder:validation:Optional
	FifoTopic *bool `json:"fifoTopic,omitempty" tf:"fifo_topic,omitempty"`

	// IAM role for failure feedback
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	FirehoseFailureFeedbackRoleArn *string `json:"firehoseFailureFeedbackRoleArn,omitempty" tf:"firehose_failure_feedback_role_arn,omitempty"`

	// Reference to a Role in iam to populate firehoseFailureFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	FirehoseFailureFeedbackRoleArnRef *v1.Reference `json:"firehoseFailureFeedbackRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate firehoseFailureFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	FirehoseFailureFeedbackRoleArnSelector *v1.Selector `json:"firehoseFailureFeedbackRoleArnSelector,omitempty" tf:"-"`

	// The IAM role permitted to receive success feedback for this topic
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	FirehoseSuccessFeedbackRoleArn *string `json:"firehoseSuccessFeedbackRoleArn,omitempty" tf:"firehose_success_feedback_role_arn,omitempty"`

	// Reference to a Role in iam to populate firehoseSuccessFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	FirehoseSuccessFeedbackRoleArnRef *v1.Reference `json:"firehoseSuccessFeedbackRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate firehoseSuccessFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	FirehoseSuccessFeedbackRoleArnSelector *v1.Selector `json:"firehoseSuccessFeedbackRoleArnSelector,omitempty" tf:"-"`

	// Percentage of success to sample
	// +kubebuilder:validation:Optional
	FirehoseSuccessFeedbackSampleRate *float64 `json:"firehoseSuccessFeedbackSampleRate,omitempty" tf:"firehose_success_feedback_sample_rate,omitempty"`

	// IAM role for failure feedback
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	HTTPFailureFeedbackRoleArn *string `json:"httpFailureFeedbackRoleArn,omitempty" tf:"http_failure_feedback_role_arn,omitempty"`

	// Reference to a Role in iam to populate httpFailureFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	HTTPFailureFeedbackRoleArnRef *v1.Reference `json:"httpFailureFeedbackRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate httpFailureFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	HTTPFailureFeedbackRoleArnSelector *v1.Selector `json:"httpFailureFeedbackRoleArnSelector,omitempty" tf:"-"`

	// The IAM role permitted to receive success feedback for this topic
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	HTTPSuccessFeedbackRoleArn *string `json:"httpSuccessFeedbackRoleArn,omitempty" tf:"http_success_feedback_role_arn,omitempty"`

	// Reference to a Role in iam to populate httpSuccessFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	HTTPSuccessFeedbackRoleArnRef *v1.Reference `json:"httpSuccessFeedbackRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate httpSuccessFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	HTTPSuccessFeedbackRoleArnSelector *v1.Selector `json:"httpSuccessFeedbackRoleArnSelector,omitempty" tf:"-"`

	// Percentage of success to sample
	// +kubebuilder:validation:Optional
	HTTPSuccessFeedbackSampleRate *float64 `json:"httpSuccessFeedbackSampleRate,omitempty" tf:"http_success_feedback_sample_rate,omitempty"`

	// The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see Key Terms
	// +kubebuilder:validation:Optional
	KMSMasterKeyID *string `json:"kmsMasterKeyId,omitempty" tf:"kms_master_key_id,omitempty"`

	// IAM role for failure feedback
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	LambdaFailureFeedbackRoleArn *string `json:"lambdaFailureFeedbackRoleArn,omitempty" tf:"lambda_failure_feedback_role_arn,omitempty"`

	// Reference to a Role in iam to populate lambdaFailureFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	LambdaFailureFeedbackRoleArnRef *v1.Reference `json:"lambdaFailureFeedbackRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate lambdaFailureFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	LambdaFailureFeedbackRoleArnSelector *v1.Selector `json:"lambdaFailureFeedbackRoleArnSelector,omitempty" tf:"-"`

	// The IAM role permitted to receive success feedback for this topic
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	LambdaSuccessFeedbackRoleArn *string `json:"lambdaSuccessFeedbackRoleArn,omitempty" tf:"lambda_success_feedback_role_arn,omitempty"`

	// Reference to a Role in iam to populate lambdaSuccessFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	LambdaSuccessFeedbackRoleArnRef *v1.Reference `json:"lambdaSuccessFeedbackRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate lambdaSuccessFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	LambdaSuccessFeedbackRoleArnSelector *v1.Selector `json:"lambdaSuccessFeedbackRoleArnSelector,omitempty" tf:"-"`

	// Percentage of success to sample
	// +kubebuilder:validation:Optional
	LambdaSuccessFeedbackSampleRate *float64 `json:"lambdaSuccessFeedbackSampleRate,omitempty" tf:"lambda_success_feedback_sample_rate,omitempty"`

	// The fully-formed AWS policy as JSON.
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// If SignatureVersion should be 1 (SHA1) or 2 (SHA256). The signature version corresponds to the hashing algorithm used while creating the signature of the notifications, subscription confirmations, or unsubscribe confirmation messages sent by Amazon SNS.
	// +kubebuilder:validation:Optional
	SignatureVersion *float64 `json:"signatureVersion,omitempty" tf:"signature_version,omitempty"`

	// IAM role for failure feedback
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	SqsFailureFeedbackRoleArn *string `json:"sqsFailureFeedbackRoleArn,omitempty" tf:"sqs_failure_feedback_role_arn,omitempty"`

	// Reference to a Role in iam to populate sqsFailureFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	SqsFailureFeedbackRoleArnRef *v1.Reference `json:"sqsFailureFeedbackRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate sqsFailureFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	SqsFailureFeedbackRoleArnSelector *v1.Selector `json:"sqsFailureFeedbackRoleArnSelector,omitempty" tf:"-"`

	// The IAM role permitted to receive success feedback for this topic
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	SqsSuccessFeedbackRoleArn *string `json:"sqsSuccessFeedbackRoleArn,omitempty" tf:"sqs_success_feedback_role_arn,omitempty"`

	// Reference to a Role in iam to populate sqsSuccessFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	SqsSuccessFeedbackRoleArnRef *v1.Reference `json:"sqsSuccessFeedbackRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate sqsSuccessFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	SqsSuccessFeedbackRoleArnSelector *v1.Selector `json:"sqsSuccessFeedbackRoleArnSelector,omitempty" tf:"-"`

	// Percentage of success to sample
	// +kubebuilder:validation:Optional
	SqsSuccessFeedbackSampleRate *float64 `json:"sqsSuccessFeedbackSampleRate,omitempty" tf:"sqs_success_feedback_sample_rate,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Tracing mode of an Amazon SNS topic. Valid values: "PassThrough", "Active".
	// +kubebuilder:validation:Optional
	TracingConfig *string `json:"tracingConfig,omitempty" tf:"tracing_config,omitempty"`
}

// TopicSpec defines the desired state of Topic
type TopicSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TopicParameters `json:"forProvider"`
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
	InitProvider TopicInitParameters `json:"initProvider,omitempty"`
}

// TopicStatus defines the observed state of Topic.
type TopicStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TopicObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Topic is the Schema for the Topics API. Provides an SNS topic resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Topic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TopicSpec   `json:"spec"`
	Status            TopicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TopicList contains a list of Topics
type TopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Topic `json:"items"`
}

// Repository type metadata.
var (
	Topic_Kind             = "Topic"
	Topic_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Topic_Kind}.String()
	Topic_KindAPIVersion   = Topic_Kind + "." + CRDGroupVersion.String()
	Topic_GroupVersionKind = CRDGroupVersion.WithKind(Topic_Kind)
)

func init() {
	SchemeBuilder.Register(&Topic{}, &TopicList{})
}

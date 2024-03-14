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

type RuleInitParameters struct {

	// The description of the rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name or ARN of the event bus to associate with this rule.
	// If you omit this, the default event bus is used.
	// +crossplane:generate:reference:type=Bus
	EventBusName *string `json:"eventBusName,omitempty" tf:"event_bus_name,omitempty"`

	// Reference to a Bus to populate eventBusName.
	// +kubebuilder:validation:Optional
	EventBusNameRef *v1.Reference `json:"eventBusNameRef,omitempty" tf:"-"`

	// Selector for a Bus to populate eventBusName.
	// +kubebuilder:validation:Optional
	EventBusNameSelector *v1.Selector `json:"eventBusNameSelector,omitempty" tf:"-"`

	// The event pattern described a JSON object. At least one of schedule_expression or event_pattern is required. See full documentation of Events and Event Patterns in EventBridge for details. Note: The event pattern size is 2048 by default but it is adjustable up to 4096 characters by submitting a service quota increase request. See Amazon EventBridge quotas for details.
	EventPattern *string `json:"eventPattern,omitempty" tf:"event_pattern,omitempty"`

	// Whether the rule should be enabled.
	// Defaults to true.
	// Conflicts with state.
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`

	// The Amazon Resource Name (ARN) associated with the role that is used for target invocation.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// The scheduling expression. For example, cron(0 20 * * ? *) or rate(5 minutes). At least one of schedule_expression or event_pattern is required. Can only be used on the default event bus. For more information, refer to the AWS documentation Schedule Expressions for Rules.
	ScheduleExpression *string `json:"scheduleExpression,omitempty" tf:"schedule_expression,omitempty"`

	// State of the rule.
	// Valid values are DISABLED, ENABLED, and ENABLED_WITH_ALL_CLOUDTRAIL_MANAGEMENT_EVENTS.
	// When state is ENABLED, the rule is enabled for all events except those delivered by CloudTrail.
	// To also enable the rule for events delivered by CloudTrail, set state to ENABLED_WITH_ALL_CLOUDTRAIL_MANAGEMENT_EVENTS.
	// Defaults to ENABLED.
	// Conflicts with is_enabled.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RuleObservation struct {

	// The Amazon Resource Name (ARN) of the rule.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The description of the rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name or ARN of the event bus to associate with this rule.
	// If you omit this, the default event bus is used.
	EventBusName *string `json:"eventBusName,omitempty" tf:"event_bus_name,omitempty"`

	// The event pattern described a JSON object. At least one of schedule_expression or event_pattern is required. See full documentation of Events and Event Patterns in EventBridge for details. Note: The event pattern size is 2048 by default but it is adjustable up to 4096 characters by submitting a service quota increase request. See Amazon EventBridge quotas for details.
	EventPattern *string `json:"eventPattern,omitempty" tf:"event_pattern,omitempty"`

	// The name of the rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether the rule should be enabled.
	// Defaults to true.
	// Conflicts with state.
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`

	// The Amazon Resource Name (ARN) associated with the role that is used for target invocation.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// The scheduling expression. For example, cron(0 20 * * ? *) or rate(5 minutes). At least one of schedule_expression or event_pattern is required. Can only be used on the default event bus. For more information, refer to the AWS documentation Schedule Expressions for Rules.
	ScheduleExpression *string `json:"scheduleExpression,omitempty" tf:"schedule_expression,omitempty"`

	// State of the rule.
	// Valid values are DISABLED, ENABLED, and ENABLED_WITH_ALL_CLOUDTRAIL_MANAGEMENT_EVENTS.
	// When state is ENABLED, the rule is enabled for all events except those delivered by CloudTrail.
	// To also enable the rule for events delivered by CloudTrail, set state to ENABLED_WITH_ALL_CLOUDTRAIL_MANAGEMENT_EVENTS.
	// Defaults to ENABLED.
	// Conflicts with is_enabled.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type RuleParameters struct {

	// The description of the rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name or ARN of the event bus to associate with this rule.
	// If you omit this, the default event bus is used.
	// +crossplane:generate:reference:type=Bus
	// +kubebuilder:validation:Optional
	EventBusName *string `json:"eventBusName,omitempty" tf:"event_bus_name,omitempty"`

	// Reference to a Bus to populate eventBusName.
	// +kubebuilder:validation:Optional
	EventBusNameRef *v1.Reference `json:"eventBusNameRef,omitempty" tf:"-"`

	// Selector for a Bus to populate eventBusName.
	// +kubebuilder:validation:Optional
	EventBusNameSelector *v1.Selector `json:"eventBusNameSelector,omitempty" tf:"-"`

	// The event pattern described a JSON object. At least one of schedule_expression or event_pattern is required. See full documentation of Events and Event Patterns in EventBridge for details. Note: The event pattern size is 2048 by default but it is adjustable up to 4096 characters by submitting a service quota increase request. See Amazon EventBridge quotas for details.
	// +kubebuilder:validation:Optional
	EventPattern *string `json:"eventPattern,omitempty" tf:"event_pattern,omitempty"`

	// Whether the rule should be enabled.
	// Defaults to true.
	// Conflicts with state.
	// +kubebuilder:validation:Optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The Amazon Resource Name (ARN) associated with the role that is used for target invocation.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// The scheduling expression. For example, cron(0 20 * * ? *) or rate(5 minutes). At least one of schedule_expression or event_pattern is required. Can only be used on the default event bus. For more information, refer to the AWS documentation Schedule Expressions for Rules.
	// +kubebuilder:validation:Optional
	ScheduleExpression *string `json:"scheduleExpression,omitempty" tf:"schedule_expression,omitempty"`

	// State of the rule.
	// Valid values are DISABLED, ENABLED, and ENABLED_WITH_ALL_CLOUDTRAIL_MANAGEMENT_EVENTS.
	// When state is ENABLED, the rule is enabled for all events except those delivered by CloudTrail.
	// To also enable the rule for events delivered by CloudTrail, set state to ENABLED_WITH_ALL_CLOUDTRAIL_MANAGEMENT_EVENTS.
	// Defaults to ENABLED.
	// Conflicts with is_enabled.
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// RuleSpec defines the desired state of Rule
type RuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RuleParameters `json:"forProvider"`
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
	InitProvider RuleInitParameters `json:"initProvider,omitempty"`
}

// RuleStatus defines the observed state of Rule.
type RuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Rule is the Schema for the Rules API. Provides an EventBridge Rule resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Rule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RuleSpec   `json:"spec"`
	Status            RuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RuleList contains a list of Rules
type RuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Rule `json:"items"`
}

// Repository type metadata.
var (
	Rule_Kind             = "Rule"
	Rule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Rule_Kind}.String()
	Rule_KindAPIVersion   = Rule_Kind + "." + CRDGroupVersion.String()
	Rule_GroupVersionKind = CRDGroupVersion.WithKind(Rule_Kind)
)

func init() {
	SchemeBuilder.Register(&Rule{}, &RuleList{})
}

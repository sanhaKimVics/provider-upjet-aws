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

type WorkflowInitParameters struct {

	// –  A map of default run properties for this workflow. These properties are passed to all jobs associated to the workflow.
	// +mapType=granular
	DefaultRunProperties map[string]*string `json:"defaultRunProperties,omitempty" tf:"default_run_properties,omitempty"`

	// –  Description of the workflow.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Prevents exceeding the maximum number of concurrent runs of any of the component jobs. If you leave this parameter blank, there is no limit to the number of concurrent workflow runs.
	MaxConcurrentRuns *float64 `json:"maxConcurrentRuns,omitempty" tf:"max_concurrent_runs,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type WorkflowObservation struct {

	// Amazon Resource Name (ARN) of Glue Workflow
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// –  A map of default run properties for this workflow. These properties are passed to all jobs associated to the workflow.
	// +mapType=granular
	DefaultRunProperties map[string]*string `json:"defaultRunProperties,omitempty" tf:"default_run_properties,omitempty"`

	// –  Description of the workflow.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Workflow name
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Prevents exceeding the maximum number of concurrent runs of any of the component jobs. If you leave this parameter blank, there is no limit to the number of concurrent workflow runs.
	MaxConcurrentRuns *float64 `json:"maxConcurrentRuns,omitempty" tf:"max_concurrent_runs,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type WorkflowParameters struct {

	// –  A map of default run properties for this workflow. These properties are passed to all jobs associated to the workflow.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	DefaultRunProperties map[string]*string `json:"defaultRunProperties,omitempty" tf:"default_run_properties,omitempty"`

	// –  Description of the workflow.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Prevents exceeding the maximum number of concurrent runs of any of the component jobs. If you leave this parameter blank, there is no limit to the number of concurrent workflow runs.
	// +kubebuilder:validation:Optional
	MaxConcurrentRuns *float64 `json:"maxConcurrentRuns,omitempty" tf:"max_concurrent_runs,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// WorkflowSpec defines the desired state of Workflow
type WorkflowSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkflowParameters `json:"forProvider"`
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
	InitProvider WorkflowInitParameters `json:"initProvider,omitempty"`
}

// WorkflowStatus defines the observed state of Workflow.
type WorkflowStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkflowObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Workflow is the Schema for the Workflows API. Provides a Glue Workflow resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Workflow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkflowSpec   `json:"spec"`
	Status            WorkflowStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkflowList contains a list of Workflows
type WorkflowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workflow `json:"items"`
}

// Repository type metadata.
var (
	Workflow_Kind             = "Workflow"
	Workflow_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Workflow_Kind}.String()
	Workflow_KindAPIVersion   = Workflow_Kind + "." + CRDGroupVersion.String()
	Workflow_GroupVersionKind = CRDGroupVersion.WithKind(Workflow_Kind)
)

func init() {
	SchemeBuilder.Register(&Workflow{}, &WorkflowList{})
}

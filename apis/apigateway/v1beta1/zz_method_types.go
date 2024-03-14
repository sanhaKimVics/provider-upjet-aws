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

type MethodInitParameters struct {

	// Specify if the method requires an API key
	APIKeyRequired *bool `json:"apiKeyRequired,omitempty" tf:"api_key_required,omitempty"`

	// Type of authorization used for the method (NONE, CUSTOM, AWS_IAM, COGNITO_USER_POOLS)
	Authorization *string `json:"authorization,omitempty" tf:"authorization,omitempty"`

	// Authorization scopes used when the authorization is COGNITO_USER_POOLS
	// +listType=set
	AuthorizationScopes []*string `json:"authorizationScopes,omitempty" tf:"authorization_scopes,omitempty"`

	// Authorizer id to be used when the authorization is CUSTOM or COGNITO_USER_POOLS
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.Authorizer
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	AuthorizerID *string `json:"authorizerId,omitempty" tf:"authorizer_id,omitempty"`

	// Reference to a Authorizer in apigateway to populate authorizerId.
	// +kubebuilder:validation:Optional
	AuthorizerIDRef *v1.Reference `json:"authorizerIdRef,omitempty" tf:"-"`

	// Selector for a Authorizer in apigateway to populate authorizerId.
	// +kubebuilder:validation:Optional
	AuthorizerIDSelector *v1.Selector `json:"authorizerIdSelector,omitempty" tf:"-"`

	// HTTP Method (GET, POST, PUT, DELETE, HEAD, OPTIONS, ANY)
	HTTPMethod *string `json:"httpMethod,omitempty" tf:"http_method,omitempty"`

	// Function name that will be given to the method when generating an SDK through API Gateway. If omitted, API Gateway will generate a function name based on the resource path and HTTP verb.
	OperationName *string `json:"operationName,omitempty" tf:"operation_name,omitempty"`

	// Map of the API models used for the request's content type
	// where key is the content type (e.g., application/json)
	// and value is either Error, Empty (built-in models) or aws_api_gateway_model's name.
	// +mapType=granular
	RequestModels map[string]*string `json:"requestModels,omitempty" tf:"request_models,omitempty"`

	// Map of request parameters (from the path, query string and headers) that should be passed to the integration. The boolean value indicates whether the parameter is required (true) or optional (false).
	// For example: request_parameters = {"method.request.header.X-Some-Header" = true "method.request.querystring.some-query-param" = true} would define that the header X-Some-Header and the query string some-query-param must be provided in the request.
	// +mapType=granular
	RequestParameters map[string]*bool `json:"requestParameters,omitempty" tf:"request_parameters,omitempty"`

	// ID of a aws_api_gateway_request_validator
	RequestValidatorID *string `json:"requestValidatorId,omitempty" tf:"request_validator_id,omitempty"`

	// API resource ID
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.Resource
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Reference to a Resource in apigateway to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDRef *v1.Reference `json:"resourceIdRef,omitempty" tf:"-"`

	// Selector for a Resource in apigateway to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDSelector *v1.Selector `json:"resourceIdSelector,omitempty" tf:"-"`

	// ID of the associated REST API
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.RestAPI
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	RestAPIID *string `json:"restApiId,omitempty" tf:"rest_api_id,omitempty"`

	// Reference to a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDRef *v1.Reference `json:"restApiIdRef,omitempty" tf:"-"`

	// Selector for a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDSelector *v1.Selector `json:"restApiIdSelector,omitempty" tf:"-"`
}

type MethodObservation struct {

	// Specify if the method requires an API key
	APIKeyRequired *bool `json:"apiKeyRequired,omitempty" tf:"api_key_required,omitempty"`

	// Type of authorization used for the method (NONE, CUSTOM, AWS_IAM, COGNITO_USER_POOLS)
	Authorization *string `json:"authorization,omitempty" tf:"authorization,omitempty"`

	// Authorization scopes used when the authorization is COGNITO_USER_POOLS
	// +listType=set
	AuthorizationScopes []*string `json:"authorizationScopes,omitempty" tf:"authorization_scopes,omitempty"`

	// Authorizer id to be used when the authorization is CUSTOM or COGNITO_USER_POOLS
	AuthorizerID *string `json:"authorizerId,omitempty" tf:"authorizer_id,omitempty"`

	// HTTP Method (GET, POST, PUT, DELETE, HEAD, OPTIONS, ANY)
	HTTPMethod *string `json:"httpMethod,omitempty" tf:"http_method,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Function name that will be given to the method when generating an SDK through API Gateway. If omitted, API Gateway will generate a function name based on the resource path and HTTP verb.
	OperationName *string `json:"operationName,omitempty" tf:"operation_name,omitempty"`

	// Map of the API models used for the request's content type
	// where key is the content type (e.g., application/json)
	// and value is either Error, Empty (built-in models) or aws_api_gateway_model's name.
	// +mapType=granular
	RequestModels map[string]*string `json:"requestModels,omitempty" tf:"request_models,omitempty"`

	// Map of request parameters (from the path, query string and headers) that should be passed to the integration. The boolean value indicates whether the parameter is required (true) or optional (false).
	// For example: request_parameters = {"method.request.header.X-Some-Header" = true "method.request.querystring.some-query-param" = true} would define that the header X-Some-Header and the query string some-query-param must be provided in the request.
	// +mapType=granular
	RequestParameters map[string]*bool `json:"requestParameters,omitempty" tf:"request_parameters,omitempty"`

	// ID of a aws_api_gateway_request_validator
	RequestValidatorID *string `json:"requestValidatorId,omitempty" tf:"request_validator_id,omitempty"`

	// API resource ID
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// ID of the associated REST API
	RestAPIID *string `json:"restApiId,omitempty" tf:"rest_api_id,omitempty"`
}

type MethodParameters struct {

	// Specify if the method requires an API key
	// +kubebuilder:validation:Optional
	APIKeyRequired *bool `json:"apiKeyRequired,omitempty" tf:"api_key_required,omitempty"`

	// Type of authorization used for the method (NONE, CUSTOM, AWS_IAM, COGNITO_USER_POOLS)
	// +kubebuilder:validation:Optional
	Authorization *string `json:"authorization,omitempty" tf:"authorization,omitempty"`

	// Authorization scopes used when the authorization is COGNITO_USER_POOLS
	// +kubebuilder:validation:Optional
	// +listType=set
	AuthorizationScopes []*string `json:"authorizationScopes,omitempty" tf:"authorization_scopes,omitempty"`

	// Authorizer id to be used when the authorization is CUSTOM or COGNITO_USER_POOLS
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.Authorizer
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	AuthorizerID *string `json:"authorizerId,omitempty" tf:"authorizer_id,omitempty"`

	// Reference to a Authorizer in apigateway to populate authorizerId.
	// +kubebuilder:validation:Optional
	AuthorizerIDRef *v1.Reference `json:"authorizerIdRef,omitempty" tf:"-"`

	// Selector for a Authorizer in apigateway to populate authorizerId.
	// +kubebuilder:validation:Optional
	AuthorizerIDSelector *v1.Selector `json:"authorizerIdSelector,omitempty" tf:"-"`

	// HTTP Method (GET, POST, PUT, DELETE, HEAD, OPTIONS, ANY)
	// +kubebuilder:validation:Optional
	HTTPMethod *string `json:"httpMethod,omitempty" tf:"http_method,omitempty"`

	// Function name that will be given to the method when generating an SDK through API Gateway. If omitted, API Gateway will generate a function name based on the resource path and HTTP verb.
	// +kubebuilder:validation:Optional
	OperationName *string `json:"operationName,omitempty" tf:"operation_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Map of the API models used for the request's content type
	// where key is the content type (e.g., application/json)
	// and value is either Error, Empty (built-in models) or aws_api_gateway_model's name.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	RequestModels map[string]*string `json:"requestModels,omitempty" tf:"request_models,omitempty"`

	// Map of request parameters (from the path, query string and headers) that should be passed to the integration. The boolean value indicates whether the parameter is required (true) or optional (false).
	// For example: request_parameters = {"method.request.header.X-Some-Header" = true "method.request.querystring.some-query-param" = true} would define that the header X-Some-Header and the query string some-query-param must be provided in the request.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	RequestParameters map[string]*bool `json:"requestParameters,omitempty" tf:"request_parameters,omitempty"`

	// ID of a aws_api_gateway_request_validator
	// +kubebuilder:validation:Optional
	RequestValidatorID *string `json:"requestValidatorId,omitempty" tf:"request_validator_id,omitempty"`

	// API resource ID
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.Resource
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Reference to a Resource in apigateway to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDRef *v1.Reference `json:"resourceIdRef,omitempty" tf:"-"`

	// Selector for a Resource in apigateway to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDSelector *v1.Selector `json:"resourceIdSelector,omitempty" tf:"-"`

	// ID of the associated REST API
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.RestAPI
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RestAPIID *string `json:"restApiId,omitempty" tf:"rest_api_id,omitempty"`

	// Reference to a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDRef *v1.Reference `json:"restApiIdRef,omitempty" tf:"-"`

	// Selector for a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDSelector *v1.Selector `json:"restApiIdSelector,omitempty" tf:"-"`
}

// MethodSpec defines the desired state of Method
type MethodSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MethodParameters `json:"forProvider"`
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
	InitProvider MethodInitParameters `json:"initProvider,omitempty"`
}

// MethodStatus defines the observed state of Method.
type MethodStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MethodObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Method is the Schema for the Methods API. Provides a HTTP Method for an API Gateway Resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Method struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.authorization) || (has(self.initProvider) && has(self.initProvider.authorization))",message="spec.forProvider.authorization is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.httpMethod) || (has(self.initProvider) && has(self.initProvider.httpMethod))",message="spec.forProvider.httpMethod is a required parameter"
	Spec   MethodSpec   `json:"spec"`
	Status MethodStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MethodList contains a list of Methods
type MethodList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Method `json:"items"`
}

// Repository type metadata.
var (
	Method_Kind             = "Method"
	Method_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Method_Kind}.String()
	Method_KindAPIVersion   = Method_Kind + "." + CRDGroupVersion.String()
	Method_GroupVersionKind = CRDGroupVersion.WithKind(Method_Kind)
)

func init() {
	SchemeBuilder.Register(&Method{}, &MethodList{})
}

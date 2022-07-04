/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1beta1 "github.com/upbound/official-providers/provider-aws/apis/acm/v1beta1"
	v1beta1acmpca "github.com/upbound/official-providers/provider-aws/apis/acmpca/v1beta1"
	v1beta1apigatewayv2 "github.com/upbound/official-providers/provider-aws/apis/apigatewayv2/v1beta1"
	v1beta1autoscaling "github.com/upbound/official-providers/provider-aws/apis/autoscaling/v1beta1"
	v1beta1backup "github.com/upbound/official-providers/provider-aws/apis/backup/v1beta1"
	v1beta1cloudfront "github.com/upbound/official-providers/provider-aws/apis/cloudfront/v1beta1"
	v1beta1docdb "github.com/upbound/official-providers/provider-aws/apis/docdb/v1beta1"
	v1beta1dynamodb "github.com/upbound/official-providers/provider-aws/apis/dynamodb/v1beta1"
	v1beta1ec2 "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1"
	v1beta1ecr "github.com/upbound/official-providers/provider-aws/apis/ecr/v1beta1"
	v1beta1ecrpublic "github.com/upbound/official-providers/provider-aws/apis/ecrpublic/v1beta1"
	v1beta1ecs "github.com/upbound/official-providers/provider-aws/apis/ecs/v1beta1"
	v1beta1efs "github.com/upbound/official-providers/provider-aws/apis/efs/v1beta1"
	v1beta1eks "github.com/upbound/official-providers/provider-aws/apis/eks/v1beta1"
	v1beta1elasticache "github.com/upbound/official-providers/provider-aws/apis/elasticache/v1beta1"
	v1beta1elbv2 "github.com/upbound/official-providers/provider-aws/apis/elbv2/v1beta1"
	v1beta1globalaccelerator "github.com/upbound/official-providers/provider-aws/apis/globalaccelerator/v1beta1"
	v1beta1glue "github.com/upbound/official-providers/provider-aws/apis/glue/v1beta1"
	v1beta1iam "github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1"
	v1beta1kms "github.com/upbound/official-providers/provider-aws/apis/kms/v1beta1"
	v1beta1mq "github.com/upbound/official-providers/provider-aws/apis/mq/v1beta1"
	v1beta1neptune "github.com/upbound/official-providers/provider-aws/apis/neptune/v1beta1"
	v1beta1rds "github.com/upbound/official-providers/provider-aws/apis/rds/v1beta1"
	v1beta1resourcegroups "github.com/upbound/official-providers/provider-aws/apis/resourcegroups/v1beta1"
	v1beta1route53 "github.com/upbound/official-providers/provider-aws/apis/route53/v1beta1"
	v1beta1route53resolver "github.com/upbound/official-providers/provider-aws/apis/route53resolver/v1beta1"
	v1beta1s3 "github.com/upbound/official-providers/provider-aws/apis/s3/v1beta1"
	v1beta1secretsmanager "github.com/upbound/official-providers/provider-aws/apis/secretsmanager/v1beta1"
	v1beta1servicediscovery "github.com/upbound/official-providers/provider-aws/apis/servicediscovery/v1beta1"
	v1beta1sns "github.com/upbound/official-providers/provider-aws/apis/sns/v1beta1"
	v1beta1sqs "github.com/upbound/official-providers/provider-aws/apis/sqs/v1beta1"
	v1beta1transfer "github.com/upbound/official-providers/provider-aws/apis/transfer/v1beta1"
	v1alpha1 "github.com/upbound/official-providers/provider-aws/apis/v1alpha1"
	v1beta1apis "github.com/upbound/official-providers/provider-aws/apis/v1beta1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1beta1.SchemeBuilder.AddToScheme,
		v1beta1acmpca.SchemeBuilder.AddToScheme,
		v1beta1apigatewayv2.SchemeBuilder.AddToScheme,
		v1beta1autoscaling.SchemeBuilder.AddToScheme,
		v1beta1backup.SchemeBuilder.AddToScheme,
		v1beta1cloudfront.SchemeBuilder.AddToScheme,
		v1beta1docdb.SchemeBuilder.AddToScheme,
		v1beta1dynamodb.SchemeBuilder.AddToScheme,
		v1beta1ec2.SchemeBuilder.AddToScheme,
		v1beta1ecr.SchemeBuilder.AddToScheme,
		v1beta1ecrpublic.SchemeBuilder.AddToScheme,
		v1beta1ecs.SchemeBuilder.AddToScheme,
		v1beta1efs.SchemeBuilder.AddToScheme,
		v1beta1eks.SchemeBuilder.AddToScheme,
		v1beta1elasticache.SchemeBuilder.AddToScheme,
		v1beta1elbv2.SchemeBuilder.AddToScheme,
		v1beta1globalaccelerator.SchemeBuilder.AddToScheme,
		v1beta1glue.SchemeBuilder.AddToScheme,
		v1beta1iam.SchemeBuilder.AddToScheme,
		v1beta1kms.SchemeBuilder.AddToScheme,
		v1beta1mq.SchemeBuilder.AddToScheme,
		v1beta1neptune.SchemeBuilder.AddToScheme,
		v1beta1rds.SchemeBuilder.AddToScheme,
		v1beta1resourcegroups.SchemeBuilder.AddToScheme,
		v1beta1route53.SchemeBuilder.AddToScheme,
		v1beta1route53resolver.SchemeBuilder.AddToScheme,
		v1beta1s3.SchemeBuilder.AddToScheme,
		v1beta1secretsmanager.SchemeBuilder.AddToScheme,
		v1beta1servicediscovery.SchemeBuilder.AddToScheme,
		v1beta1sns.SchemeBuilder.AddToScheme,
		v1beta1sqs.SchemeBuilder.AddToScheme,
		v1beta1transfer.SchemeBuilder.AddToScheme,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1beta1apis.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}

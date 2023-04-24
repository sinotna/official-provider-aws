/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1beta1 "github.com/dkb-bank/official-provider-aws/apis/amp/v1beta1"
	v1beta1cloudwatchlogs "github.com/dkb-bank/official-provider-aws/apis/cloudwatchlogs/v1beta1"
	v1beta1datasync "github.com/dkb-bank/official-provider-aws/apis/datasync/v1beta1"
	v1beta1directconnect "github.com/dkb-bank/official-provider-aws/apis/directconnect/v1beta1"
	v1beta1dms "github.com/dkb-bank/official-provider-aws/apis/dms/v1beta1"
	v1beta1dynamodb "github.com/dkb-bank/official-provider-aws/apis/dynamodb/v1beta1"
	v1beta1ec2 "github.com/dkb-bank/official-provider-aws/apis/ec2/v1beta1"
	v1beta1ecr "github.com/dkb-bank/official-provider-aws/apis/ecr/v1beta1"
	v1beta1elbv2 "github.com/dkb-bank/official-provider-aws/apis/elbv2/v1beta1"
	v1beta1firehose "github.com/dkb-bank/official-provider-aws/apis/firehose/v1beta1"
	v1beta1grafana "github.com/dkb-bank/official-provider-aws/apis/grafana/v1beta1"
	v1beta1iam "github.com/dkb-bank/official-provider-aws/apis/iam/v1beta1"
	v1beta1kinesis "github.com/dkb-bank/official-provider-aws/apis/kinesis/v1beta1"
	v1beta1kinesisanalytics "github.com/dkb-bank/official-provider-aws/apis/kinesisanalytics/v1beta1"
	v1beta1kinesisanalyticsv2 "github.com/dkb-bank/official-provider-aws/apis/kinesisanalyticsv2/v1beta1"
	v1beta1kinesisvideo "github.com/dkb-bank/official-provider-aws/apis/kinesisvideo/v1beta1"
	v1beta1kms "github.com/dkb-bank/official-provider-aws/apis/kms/v1beta1"
	v1beta1networkfirewall "github.com/dkb-bank/official-provider-aws/apis/networkfirewall/v1beta1"
	v1beta1s3 "github.com/dkb-bank/official-provider-aws/apis/s3/v1beta1"
	v1beta1ses "github.com/dkb-bank/official-provider-aws/apis/ses/v1beta1"
	v1beta1ssm "github.com/dkb-bank/official-provider-aws/apis/ssm/v1beta1"
	v1alpha1 "github.com/dkb-bank/official-provider-aws/apis/v1alpha1"
	v1beta1apis "github.com/dkb-bank/official-provider-aws/apis/v1beta1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1beta1.SchemeBuilder.AddToScheme,
		v1beta1cloudwatchlogs.SchemeBuilder.AddToScheme,
		v1beta1datasync.SchemeBuilder.AddToScheme,
		v1beta1directconnect.SchemeBuilder.AddToScheme,
		v1beta1dms.SchemeBuilder.AddToScheme,
		v1beta1dynamodb.SchemeBuilder.AddToScheme,
		v1beta1ec2.SchemeBuilder.AddToScheme,
		v1beta1ecr.SchemeBuilder.AddToScheme,
		v1beta1elbv2.SchemeBuilder.AddToScheme,
		v1beta1firehose.SchemeBuilder.AddToScheme,
		v1beta1grafana.SchemeBuilder.AddToScheme,
		v1beta1iam.SchemeBuilder.AddToScheme,
		v1beta1kinesis.SchemeBuilder.AddToScheme,
		v1beta1kinesisanalytics.SchemeBuilder.AddToScheme,
		v1beta1kinesisanalyticsv2.SchemeBuilder.AddToScheme,
		v1beta1kinesisvideo.SchemeBuilder.AddToScheme,
		v1beta1kms.SchemeBuilder.AddToScheme,
		v1beta1networkfirewall.SchemeBuilder.AddToScheme,
		v1beta1s3.SchemeBuilder.AddToScheme,
		v1beta1ses.SchemeBuilder.AddToScheme,
		v1beta1ssm.SchemeBuilder.AddToScheme,
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

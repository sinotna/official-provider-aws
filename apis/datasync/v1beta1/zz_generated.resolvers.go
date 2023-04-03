/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1beta11 "github.com/dkb-bank/official-provider-aws/apis/iam/v1beta1"
	v1beta1 "github.com/dkb-bank/official-provider-aws/apis/s3/v1beta1"
	errors "github.com/pkg/errors"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this LocationS3.
func (mg *LocationS3) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.S3BucketArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.S3BucketArnRef,
		Selector:     mg.Spec.ForProvider.S3BucketArnSelector,
		To: reference.To{
			List:    &v1beta1.BucketList{},
			Managed: &v1beta1.Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.S3BucketArn")
	}
	mg.Spec.ForProvider.S3BucketArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.S3BucketArnRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.S3Config); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.S3Config[i3].BucketAccessRoleArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.S3Config[i3].BucketAccessRoleArnRef,
			Selector:     mg.Spec.ForProvider.S3Config[i3].BucketAccessRoleArnSelector,
			To: reference.To{
				List:    &v1beta11.RoleList{},
				Managed: &v1beta11.Role{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.S3Config[i3].BucketAccessRoleArn")
		}
		mg.Spec.ForProvider.S3Config[i3].BucketAccessRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.S3Config[i3].BucketAccessRoleArnRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Task.
func (mg *Task) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DestinationLocationArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.DestinationLocationArnRef,
		Selector:     mg.Spec.ForProvider.DestinationLocationArnSelector,
		To: reference.To{
			List:    &LocationS3List{},
			Managed: &LocationS3{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DestinationLocationArn")
	}
	mg.Spec.ForProvider.DestinationLocationArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DestinationLocationArnRef = rsp.ResolvedReference

	return nil
}

/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1beta12 "github.com/dkb-bank/official-provider-aws/apis/ec2/v1beta1"
	v1beta11 "github.com/dkb-bank/official-provider-aws/apis/iam/v1beta1"
	v1beta1 "github.com/dkb-bank/official-provider-aws/apis/kms/v1beta1"
	errors "github.com/pkg/errors"
	common "github.com/upbound/provider-aws/config/common"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Endpoint.
func (mg *Endpoint) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyArn),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KMSKeyArnRef,
		Selector:     mg.Spec.ForProvider.KMSKeyArnSelector,
		To: reference.To{
			List:    &v1beta1.KeyList{},
			Managed: &v1beta1.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyArn")
	}
	mg.Spec.ForProvider.KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SecretsManagerAccessRoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.SecretsManagerAccessRoleArnRef,
		Selector:     mg.Spec.ForProvider.SecretsManagerAccessRoleArnSelector,
		To: reference.To{
			List:    &v1beta11.RoleList{},
			Managed: &v1beta11.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecretsManagerAccessRoleArn")
	}
	mg.Spec.ForProvider.SecretsManagerAccessRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SecretsManagerAccessRoleArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ReplicationInstance.
func (mg *ReplicationInstance) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyArn),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KMSKeyArnRef,
		Selector:     mg.Spec.ForProvider.KMSKeyArnSelector,
		To: reference.To{
			List:    &v1beta1.KeyList{},
			Managed: &v1beta1.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyArn")
	}
	mg.Spec.ForProvider.KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ReplicationSubnetGroupID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ReplicationSubnetGroupIDRef,
		Selector:     mg.Spec.ForProvider.ReplicationSubnetGroupIDSelector,
		To: reference.To{
			List:    &ReplicationSubnetGroupList{},
			Managed: &ReplicationSubnetGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ReplicationSubnetGroupID")
	}
	mg.Spec.ForProvider.ReplicationSubnetGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ReplicationSubnetGroupIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCSecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.VPCSecurityGroupIDRefs,
		Selector:      mg.Spec.ForProvider.VPCSecurityGroupIDSelector,
		To: reference.To{
			List:    &v1beta12.SecurityGroupList{},
			Managed: &v1beta12.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCSecurityGroupIds")
	}
	mg.Spec.ForProvider.VPCSecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.VPCSecurityGroupIDRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this ReplicationSubnetGroup.
func (mg *ReplicationSubnetGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SubnetIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SubnetIDRefs,
		Selector:      mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &v1beta12.SubnetList{},
			Managed: &v1beta12.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetIds")
	}
	mg.Spec.ForProvider.SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SubnetIDRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this ReplicationTask.
func (mg *ReplicationTask) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ReplicationInstanceArn),
		Extract:      resource.ExtractParamPath("replication_instance_arn", true),
		Reference:    mg.Spec.ForProvider.ReplicationInstanceArnRef,
		Selector:     mg.Spec.ForProvider.ReplicationInstanceArnSelector,
		To: reference.To{
			List:    &ReplicationInstanceList{},
			Managed: &ReplicationInstance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ReplicationInstanceArn")
	}
	mg.Spec.ForProvider.ReplicationInstanceArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ReplicationInstanceArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SourceEndpointArn),
		Extract:      resource.ExtractParamPath("endpoint_arn", true),
		Reference:    mg.Spec.ForProvider.SourceEndpointArnRef,
		Selector:     mg.Spec.ForProvider.SourceEndpointArnSelector,
		To: reference.To{
			List:    &EndpointList{},
			Managed: &Endpoint{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SourceEndpointArn")
	}
	mg.Spec.ForProvider.SourceEndpointArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SourceEndpointArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TargetEndpointArn),
		Extract:      resource.ExtractParamPath("endpoint_arn", true),
		Reference:    mg.Spec.ForProvider.TargetEndpointArnRef,
		Selector:     mg.Spec.ForProvider.TargetEndpointArnSelector,
		To: reference.To{
			List:    &EndpointList{},
			Managed: &Endpoint{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TargetEndpointArn")
	}
	mg.Spec.ForProvider.TargetEndpointArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TargetEndpointArnRef = rsp.ResolvedReference

	return nil
}

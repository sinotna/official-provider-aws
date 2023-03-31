/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1beta11 "github.com/dkb-bank/official-provider-aws/apis/firehose/v1beta1"
	v1beta1 "github.com/dkb-bank/official-provider-aws/apis/iam/v1beta1"
	errors "github.com/pkg/errors"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this DomainMailFrom.
func (mg *DomainMailFrom) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Domain),
		Extract:      resource.ExtractParamPath("domain", false),
		Reference:    mg.Spec.ForProvider.DomainRef,
		Selector:     mg.Spec.ForProvider.DomainSelector,
		To: reference.To{
			List:    &DomainIdentityList{},
			Managed: &DomainIdentity{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Domain")
	}
	mg.Spec.ForProvider.Domain = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DomainRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this EventDestination.
func (mg *EventDestination) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConfigurationSetName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ConfigurationSetNameRef,
		Selector:     mg.Spec.ForProvider.ConfigurationSetNameSelector,
		To: reference.To{
			List:    &ConfigurationSetList{},
			Managed: &ConfigurationSet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ConfigurationSetName")
	}
	mg.Spec.ForProvider.ConfigurationSetName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ConfigurationSetNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.KinesisDestination); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KinesisDestination[i3].RoleArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.KinesisDestination[i3].RoleArnRef,
			Selector:     mg.Spec.ForProvider.KinesisDestination[i3].RoleArnSelector,
			To: reference.To{
				List:    &v1beta1.RoleList{},
				Managed: &v1beta1.Role{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.KinesisDestination[i3].RoleArn")
		}
		mg.Spec.ForProvider.KinesisDestination[i3].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.KinesisDestination[i3].RoleArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.KinesisDestination); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KinesisDestination[i3].StreamArn),
			Extract:      resource.ExtractParamPath("arn", false),
			Reference:    mg.Spec.ForProvider.KinesisDestination[i3].StreamArnRef,
			Selector:     mg.Spec.ForProvider.KinesisDestination[i3].StreamArnSelector,
			To: reference.To{
				List:    &v1beta11.DeliveryStreamList{},
				Managed: &v1beta11.DeliveryStream{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.KinesisDestination[i3].StreamArn")
		}
		mg.Spec.ForProvider.KinesisDestination[i3].StreamArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.KinesisDestination[i3].StreamArnRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this IdentityNotificationTopic.
func (mg *IdentityNotificationTopic) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Identity),
		Extract:      resource.ExtractParamPath("domain", false),
		Reference:    mg.Spec.ForProvider.IdentityRef,
		Selector:     mg.Spec.ForProvider.IdentitySelector,
		To: reference.To{
			List:    &DomainIdentityList{},
			Managed: &DomainIdentity{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Identity")
	}
	mg.Spec.ForProvider.Identity = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IdentityRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this IdentityPolicy.
func (mg *IdentityPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Identity),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.IdentityRef,
		Selector:     mg.Spec.ForProvider.IdentitySelector,
		To: reference.To{
			List:    &DomainIdentityList{},
			Managed: &DomainIdentity{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Identity")
	}
	mg.Spec.ForProvider.Identity = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IdentityRef = rsp.ResolvedReference

	return nil
}

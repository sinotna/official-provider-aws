/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this BGPPeer.
func (mg *BGPPeer) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualInterfaceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VirtualInterfaceIDRef,
		Selector:     mg.Spec.ForProvider.VirtualInterfaceIDSelector,
		To: reference.To{
			List:    &PrivateVirtualInterfaceList{},
			Managed: &PrivateVirtualInterface{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualInterfaceID")
	}
	mg.Spec.ForProvider.VirtualInterfaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualInterfaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ConnectionAssociation.
func (mg *ConnectionAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConnectionID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ConnectionIDRef,
		Selector:     mg.Spec.ForProvider.ConnectionIDSelector,
		To: reference.To{
			List:    &ConnectionList{},
			Managed: &Connection{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ConnectionID")
	}
	mg.Spec.ForProvider.ConnectionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ConnectionIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LagID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.LagIDRef,
		Selector:     mg.Spec.ForProvider.LagIDSelector,
		To: reference.To{
			List:    &LagList{},
			Managed: &Lag{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LagID")
	}
	mg.Spec.ForProvider.LagID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LagIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this GatewayAssociation.
func (mg *GatewayAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DxGatewayID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.DxGatewayIDRef,
		Selector:     mg.Spec.ForProvider.DxGatewayIDSelector,
		To: reference.To{
			List:    &GatewayList{},
			Managed: &Gateway{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DxGatewayID")
	}
	mg.Spec.ForProvider.DxGatewayID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DxGatewayIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this GatewayAssociationProposal.
func (mg *GatewayAssociationProposal) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DxGatewayID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.DxGatewayIDRef,
		Selector:     mg.Spec.ForProvider.DxGatewayIDSelector,
		To: reference.To{
			List:    &GatewayList{},
			Managed: &Gateway{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DxGatewayID")
	}
	mg.Spec.ForProvider.DxGatewayID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DxGatewayIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DxGatewayOwnerAccountID),
		Extract:      resource.ExtractParamPath("owner_account_id", true),
		Reference:    mg.Spec.ForProvider.DxGatewayOwnerAccountIDRef,
		Selector:     mg.Spec.ForProvider.DxGatewayOwnerAccountIDSelector,
		To: reference.To{
			List:    &GatewayList{},
			Managed: &Gateway{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DxGatewayOwnerAccountID")
	}
	mg.Spec.ForProvider.DxGatewayOwnerAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DxGatewayOwnerAccountIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this HostedPrivateVirtualInterfaceAccepter.
func (mg *HostedPrivateVirtualInterfaceAccepter) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualInterfaceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VirtualInterfaceIDRef,
		Selector:     mg.Spec.ForProvider.VirtualInterfaceIDSelector,
		To: reference.To{
			List:    &HostedPrivateVirtualInterfaceList{},
			Managed: &HostedPrivateVirtualInterface{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualInterfaceID")
	}
	mg.Spec.ForProvider.VirtualInterfaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualInterfaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this HostedPublicVirtualInterfaceAccepter.
func (mg *HostedPublicVirtualInterfaceAccepter) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualInterfaceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VirtualInterfaceIDRef,
		Selector:     mg.Spec.ForProvider.VirtualInterfaceIDSelector,
		To: reference.To{
			List:    &HostedPublicVirtualInterfaceList{},
			Managed: &HostedPublicVirtualInterface{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualInterfaceID")
	}
	mg.Spec.ForProvider.VirtualInterfaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualInterfaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this HostedTransitVirtualInterface.
func (mg *HostedTransitVirtualInterface) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConnectionID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ConnectionIDRef,
		Selector:     mg.Spec.ForProvider.ConnectionIDSelector,
		To: reference.To{
			List:    &ConnectionList{},
			Managed: &Connection{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ConnectionID")
	}
	mg.Spec.ForProvider.ConnectionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ConnectionIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this HostedTransitVirtualInterfaceAccepter.
func (mg *HostedTransitVirtualInterfaceAccepter) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DxGatewayID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.DxGatewayIDRef,
		Selector:     mg.Spec.ForProvider.DxGatewayIDSelector,
		To: reference.To{
			List:    &GatewayList{},
			Managed: &Gateway{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DxGatewayID")
	}
	mg.Spec.ForProvider.DxGatewayID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DxGatewayIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualInterfaceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VirtualInterfaceIDRef,
		Selector:     mg.Spec.ForProvider.VirtualInterfaceIDSelector,
		To: reference.To{
			List:    &HostedTransitVirtualInterfaceList{},
			Managed: &HostedTransitVirtualInterface{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualInterfaceID")
	}
	mg.Spec.ForProvider.VirtualInterfaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualInterfaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TransitVirtualInterface.
func (mg *TransitVirtualInterface) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConnectionID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ConnectionIDRef,
		Selector:     mg.Spec.ForProvider.ConnectionIDSelector,
		To: reference.To{
			List:    &ConnectionList{},
			Managed: &Connection{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ConnectionID")
	}
	mg.Spec.ForProvider.ConnectionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ConnectionIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DxGatewayID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.DxGatewayIDRef,
		Selector:     mg.Spec.ForProvider.DxGatewayIDSelector,
		To: reference.To{
			List:    &GatewayList{},
			Managed: &Gateway{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DxGatewayID")
	}
	mg.Spec.ForProvider.DxGatewayID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DxGatewayIDRef = rsp.ResolvedReference

	return nil
}

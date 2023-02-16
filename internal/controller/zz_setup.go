/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	bgppeer "github.com/dkb-bank/official-provider-aws/internal/controller/directconnect/bgppeer"
	connection "github.com/dkb-bank/official-provider-aws/internal/controller/directconnect/connection"
	connectionassociation "github.com/dkb-bank/official-provider-aws/internal/controller/directconnect/connectionassociation"
	gateway "github.com/dkb-bank/official-provider-aws/internal/controller/directconnect/gateway"
	gatewayassociation "github.com/dkb-bank/official-provider-aws/internal/controller/directconnect/gatewayassociation"
	gatewayassociationproposal "github.com/dkb-bank/official-provider-aws/internal/controller/directconnect/gatewayassociationproposal"
	hostedprivatevirtualinterface "github.com/dkb-bank/official-provider-aws/internal/controller/directconnect/hostedprivatevirtualinterface"
	hostedprivatevirtualinterfaceaccepter "github.com/dkb-bank/official-provider-aws/internal/controller/directconnect/hostedprivatevirtualinterfaceaccepter"
	hostedpublicvirtualinterface "github.com/dkb-bank/official-provider-aws/internal/controller/directconnect/hostedpublicvirtualinterface"
	hostedpublicvirtualinterfaceaccepter "github.com/dkb-bank/official-provider-aws/internal/controller/directconnect/hostedpublicvirtualinterfaceaccepter"
	hostedtransitvirtualinterface "github.com/dkb-bank/official-provider-aws/internal/controller/directconnect/hostedtransitvirtualinterface"
	hostedtransitvirtualinterfaceaccepter "github.com/dkb-bank/official-provider-aws/internal/controller/directconnect/hostedtransitvirtualinterfaceaccepter"
	lag "github.com/dkb-bank/official-provider-aws/internal/controller/directconnect/lag"
	privatevirtualinterface "github.com/dkb-bank/official-provider-aws/internal/controller/directconnect/privatevirtualinterface"
	publicvirtualinterface "github.com/dkb-bank/official-provider-aws/internal/controller/directconnect/publicvirtualinterface"
	transitvirtualinterface "github.com/dkb-bank/official-provider-aws/internal/controller/directconnect/transitvirtualinterface"
	providerconfig "github.com/dkb-bank/official-provider-aws/internal/controller/providerconfig"
	activereceiptruleset "github.com/dkb-bank/official-provider-aws/internal/controller/ses/activereceiptruleset"
	configurationset "github.com/dkb-bank/official-provider-aws/internal/controller/ses/configurationset"
	domaindkim "github.com/dkb-bank/official-provider-aws/internal/controller/ses/domaindkim"
	domainidentity "github.com/dkb-bank/official-provider-aws/internal/controller/ses/domainidentity"
	domainmailfrom "github.com/dkb-bank/official-provider-aws/internal/controller/ses/domainmailfrom"
	emailidentity "github.com/dkb-bank/official-provider-aws/internal/controller/ses/emailidentity"
	eventdestination "github.com/dkb-bank/official-provider-aws/internal/controller/ses/eventdestination"
	identitynotificationtopic "github.com/dkb-bank/official-provider-aws/internal/controller/ses/identitynotificationtopic"
	identitypolicy "github.com/dkb-bank/official-provider-aws/internal/controller/ses/identitypolicy"
	receiptfilter "github.com/dkb-bank/official-provider-aws/internal/controller/ses/receiptfilter"
	receiptrule "github.com/dkb-bank/official-provider-aws/internal/controller/ses/receiptrule"
	receiptruleset "github.com/dkb-bank/official-provider-aws/internal/controller/ses/receiptruleset"
	template "github.com/dkb-bank/official-provider-aws/internal/controller/ses/template"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bgppeer.Setup,
		connection.Setup,
		connectionassociation.Setup,
		gateway.Setup,
		gatewayassociation.Setup,
		gatewayassociationproposal.Setup,
		hostedprivatevirtualinterface.Setup,
		hostedprivatevirtualinterfaceaccepter.Setup,
		hostedpublicvirtualinterface.Setup,
		hostedpublicvirtualinterfaceaccepter.Setup,
		hostedtransitvirtualinterface.Setup,
		hostedtransitvirtualinterfaceaccepter.Setup,
		lag.Setup,
		privatevirtualinterface.Setup,
		publicvirtualinterface.Setup,
		transitvirtualinterface.Setup,
		providerconfig.Setup,
		activereceiptruleset.Setup,
		configurationset.Setup,
		domaindkim.Setup,
		domainidentity.Setup,
		domainmailfrom.Setup,
		emailidentity.Setup,
		eventdestination.Setup,
		identitynotificationtopic.Setup,
		identitypolicy.Setup,
		receiptfilter.Setup,
		receiptrule.Setup,
		receiptruleset.Setup,
		template.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

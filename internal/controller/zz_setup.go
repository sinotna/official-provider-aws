/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	definition "github.com/dkb-bank/official-provider-aws/internal/controller/cloudwatchlogs/definition"
	destination "github.com/dkb-bank/official-provider-aws/internal/controller/cloudwatchlogs/destination"
	destinationpolicy "github.com/dkb-bank/official-provider-aws/internal/controller/cloudwatchlogs/destinationpolicy"
	group "github.com/dkb-bank/official-provider-aws/internal/controller/cloudwatchlogs/group"
	metricfilter "github.com/dkb-bank/official-provider-aws/internal/controller/cloudwatchlogs/metricfilter"
	resourcepolicy "github.com/dkb-bank/official-provider-aws/internal/controller/cloudwatchlogs/resourcepolicy"
	stream "github.com/dkb-bank/official-provider-aws/internal/controller/cloudwatchlogs/stream"
	subscriptionfilter "github.com/dkb-bank/official-provider-aws/internal/controller/cloudwatchlogs/subscriptionfilter"
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
	accesskey "github.com/dkb-bank/official-provider-aws/internal/controller/iam/accesskey"
	accountalias "github.com/dkb-bank/official-provider-aws/internal/controller/iam/accountalias"
	accountpasswordpolicy "github.com/dkb-bank/official-provider-aws/internal/controller/iam/accountpasswordpolicy"
	groupiam "github.com/dkb-bank/official-provider-aws/internal/controller/iam/group"
	groupmembership "github.com/dkb-bank/official-provider-aws/internal/controller/iam/groupmembership"
	grouppolicyattachment "github.com/dkb-bank/official-provider-aws/internal/controller/iam/grouppolicyattachment"
	instanceprofile "github.com/dkb-bank/official-provider-aws/internal/controller/iam/instanceprofile"
	openidconnectprovider "github.com/dkb-bank/official-provider-aws/internal/controller/iam/openidconnectprovider"
	policy "github.com/dkb-bank/official-provider-aws/internal/controller/iam/policy"
	role "github.com/dkb-bank/official-provider-aws/internal/controller/iam/role"
	rolepolicyattachment "github.com/dkb-bank/official-provider-aws/internal/controller/iam/rolepolicyattachment"
	samlprovider "github.com/dkb-bank/official-provider-aws/internal/controller/iam/samlprovider"
	servercertificate "github.com/dkb-bank/official-provider-aws/internal/controller/iam/servercertificate"
	servicelinkedrole "github.com/dkb-bank/official-provider-aws/internal/controller/iam/servicelinkedrole"
	servicespecificcredential "github.com/dkb-bank/official-provider-aws/internal/controller/iam/servicespecificcredential"
	signingcertificate "github.com/dkb-bank/official-provider-aws/internal/controller/iam/signingcertificate"
	user "github.com/dkb-bank/official-provider-aws/internal/controller/iam/user"
	usergroupmembership "github.com/dkb-bank/official-provider-aws/internal/controller/iam/usergroupmembership"
	userloginprofile "github.com/dkb-bank/official-provider-aws/internal/controller/iam/userloginprofile"
	userpolicyattachment "github.com/dkb-bank/official-provider-aws/internal/controller/iam/userpolicyattachment"
	usersshkey "github.com/dkb-bank/official-provider-aws/internal/controller/iam/usersshkey"
	virtualmfadevice "github.com/dkb-bank/official-provider-aws/internal/controller/iam/virtualmfadevice"
	alias "github.com/dkb-bank/official-provider-aws/internal/controller/kms/alias"
	ciphertext "github.com/dkb-bank/official-provider-aws/internal/controller/kms/ciphertext"
	externalkey "github.com/dkb-bank/official-provider-aws/internal/controller/kms/externalkey"
	grant "github.com/dkb-bank/official-provider-aws/internal/controller/kms/grant"
	key "github.com/dkb-bank/official-provider-aws/internal/controller/kms/key"
	replicaexternalkey "github.com/dkb-bank/official-provider-aws/internal/controller/kms/replicaexternalkey"
	replicakey "github.com/dkb-bank/official-provider-aws/internal/controller/kms/replicakey"
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
		definition.Setup,
		destination.Setup,
		destinationpolicy.Setup,
		group.Setup,
		metricfilter.Setup,
		resourcepolicy.Setup,
		stream.Setup,
		subscriptionfilter.Setup,
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
		accesskey.Setup,
		accountalias.Setup,
		accountpasswordpolicy.Setup,
		groupiam.Setup,
		groupmembership.Setup,
		grouppolicyattachment.Setup,
		instanceprofile.Setup,
		openidconnectprovider.Setup,
		policy.Setup,
		role.Setup,
		rolepolicyattachment.Setup,
		samlprovider.Setup,
		servercertificate.Setup,
		servicelinkedrole.Setup,
		servicespecificcredential.Setup,
		signingcertificate.Setup,
		user.Setup,
		usergroupmembership.Setup,
		userloginprofile.Setup,
		userpolicyattachment.Setup,
		usersshkey.Setup,
		virtualmfadevice.Setup,
		alias.Setup,
		ciphertext.Setup,
		externalkey.Setup,
		grant.Setup,
		key.Setup,
		replicaexternalkey.Setup,
		replicakey.Setup,
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

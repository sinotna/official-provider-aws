/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	alertmanagerdefinition "github.com/dkb-bank/official-provider-aws/internal/controller/amp/alertmanagerdefinition"
	rulegroupnamespace "github.com/dkb-bank/official-provider-aws/internal/controller/amp/rulegroupnamespace"
	workspace "github.com/dkb-bank/official-provider-aws/internal/controller/amp/workspace"
	definition "github.com/dkb-bank/official-provider-aws/internal/controller/cloudwatchlogs/definition"
	destination "github.com/dkb-bank/official-provider-aws/internal/controller/cloudwatchlogs/destination"
	destinationpolicy "github.com/dkb-bank/official-provider-aws/internal/controller/cloudwatchlogs/destinationpolicy"
	group "github.com/dkb-bank/official-provider-aws/internal/controller/cloudwatchlogs/group"
	metricfilter "github.com/dkb-bank/official-provider-aws/internal/controller/cloudwatchlogs/metricfilter"
	resourcepolicy "github.com/dkb-bank/official-provider-aws/internal/controller/cloudwatchlogs/resourcepolicy"
	stream "github.com/dkb-bank/official-provider-aws/internal/controller/cloudwatchlogs/stream"
	subscriptionfilter "github.com/dkb-bank/official-provider-aws/internal/controller/cloudwatchlogs/subscriptionfilter"
	locations3 "github.com/dkb-bank/official-provider-aws/internal/controller/datasync/locations3"
	task "github.com/dkb-bank/official-provider-aws/internal/controller/datasync/task"
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
	certificate "github.com/dkb-bank/official-provider-aws/internal/controller/dms/certificate"
	endpoint "github.com/dkb-bank/official-provider-aws/internal/controller/dms/endpoint"
	eventsubscription "github.com/dkb-bank/official-provider-aws/internal/controller/dms/eventsubscription"
	replicationinstance "github.com/dkb-bank/official-provider-aws/internal/controller/dms/replicationinstance"
	replicationsubnetgroup "github.com/dkb-bank/official-provider-aws/internal/controller/dms/replicationsubnetgroup"
	replicationtask "github.com/dkb-bank/official-provider-aws/internal/controller/dms/replicationtask"
	globaltable "github.com/dkb-bank/official-provider-aws/internal/controller/dynamodb/globaltable"
	table "github.com/dkb-bank/official-provider-aws/internal/controller/dynamodb/table"
	tag "github.com/dkb-bank/official-provider-aws/internal/controller/dynamodb/tag"
	instance "github.com/dkb-bank/official-provider-aws/internal/controller/ec2/instance"
	internetgateway "github.com/dkb-bank/official-provider-aws/internal/controller/ec2/internetgateway"
	keypair "github.com/dkb-bank/official-provider-aws/internal/controller/ec2/keypair"
	natgateway "github.com/dkb-bank/official-provider-aws/internal/controller/ec2/natgateway"
	networkinterface "github.com/dkb-bank/official-provider-aws/internal/controller/ec2/networkinterface"
	route "github.com/dkb-bank/official-provider-aws/internal/controller/ec2/route"
	routetable "github.com/dkb-bank/official-provider-aws/internal/controller/ec2/routetable"
	securitygroup "github.com/dkb-bank/official-provider-aws/internal/controller/ec2/securitygroup"
	securitygrouprule "github.com/dkb-bank/official-provider-aws/internal/controller/ec2/securitygrouprule"
	subnet "github.com/dkb-bank/official-provider-aws/internal/controller/ec2/subnet"
	tagec2 "github.com/dkb-bank/official-provider-aws/internal/controller/ec2/tag"
	transitgateway "github.com/dkb-bank/official-provider-aws/internal/controller/ec2/transitgateway"
	vpc "github.com/dkb-bank/official-provider-aws/internal/controller/ec2/vpc"
	vpcendpoint "github.com/dkb-bank/official-provider-aws/internal/controller/ec2/vpcendpoint"
	vpcpeeringconnection "github.com/dkb-bank/official-provider-aws/internal/controller/ec2/vpcpeeringconnection"
	lifecyclepolicy "github.com/dkb-bank/official-provider-aws/internal/controller/ecr/lifecyclepolicy"
	pullthroughcacherule "github.com/dkb-bank/official-provider-aws/internal/controller/ecr/pullthroughcacherule"
	registrypolicy "github.com/dkb-bank/official-provider-aws/internal/controller/ecr/registrypolicy"
	registryscanningconfiguration "github.com/dkb-bank/official-provider-aws/internal/controller/ecr/registryscanningconfiguration"
	replicationconfiguration "github.com/dkb-bank/official-provider-aws/internal/controller/ecr/replicationconfiguration"
	repository "github.com/dkb-bank/official-provider-aws/internal/controller/ecr/repository"
	repositorypolicy "github.com/dkb-bank/official-provider-aws/internal/controller/ecr/repositorypolicy"
	lb "github.com/dkb-bank/official-provider-aws/internal/controller/elbv2/lb"
	lblistener "github.com/dkb-bank/official-provider-aws/internal/controller/elbv2/lblistener"
	lblistenerrule "github.com/dkb-bank/official-provider-aws/internal/controller/elbv2/lblistenerrule"
	lbtargetgroup "github.com/dkb-bank/official-provider-aws/internal/controller/elbv2/lbtargetgroup"
	lbtargetgroupattachment "github.com/dkb-bank/official-provider-aws/internal/controller/elbv2/lbtargetgroupattachment"
	deliverystream "github.com/dkb-bank/official-provider-aws/internal/controller/firehose/deliverystream"
	roleassociation "github.com/dkb-bank/official-provider-aws/internal/controller/grafana/roleassociation"
	workspacegrafana "github.com/dkb-bank/official-provider-aws/internal/controller/grafana/workspace"
	workspacesamlconfiguration "github.com/dkb-bank/official-provider-aws/internal/controller/grafana/workspacesamlconfiguration"
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
	streamkinesis "github.com/dkb-bank/official-provider-aws/internal/controller/kinesis/stream"
	streamconsumer "github.com/dkb-bank/official-provider-aws/internal/controller/kinesis/streamconsumer"
	application "github.com/dkb-bank/official-provider-aws/internal/controller/kinesisanalytics/application"
	applicationkinesisanalyticsv2 "github.com/dkb-bank/official-provider-aws/internal/controller/kinesisanalyticsv2/application"
	applicationsnapshot "github.com/dkb-bank/official-provider-aws/internal/controller/kinesisanalyticsv2/applicationsnapshot"
	streamkinesisvideo "github.com/dkb-bank/official-provider-aws/internal/controller/kinesisvideo/stream"
	alias "github.com/dkb-bank/official-provider-aws/internal/controller/kms/alias"
	ciphertext "github.com/dkb-bank/official-provider-aws/internal/controller/kms/ciphertext"
	externalkey "github.com/dkb-bank/official-provider-aws/internal/controller/kms/externalkey"
	grant "github.com/dkb-bank/official-provider-aws/internal/controller/kms/grant"
	key "github.com/dkb-bank/official-provider-aws/internal/controller/kms/key"
	replicaexternalkey "github.com/dkb-bank/official-provider-aws/internal/controller/kms/replicaexternalkey"
	replicakey "github.com/dkb-bank/official-provider-aws/internal/controller/kms/replicakey"
	firewall "github.com/dkb-bank/official-provider-aws/internal/controller/networkfirewall/firewall"
	firewallpolicy "github.com/dkb-bank/official-provider-aws/internal/controller/networkfirewall/firewallpolicy"
	loggingconfiguration "github.com/dkb-bank/official-provider-aws/internal/controller/networkfirewall/loggingconfiguration"
	rulegroup "github.com/dkb-bank/official-provider-aws/internal/controller/networkfirewall/rulegroup"
	providerconfig "github.com/dkb-bank/official-provider-aws/internal/controller/providerconfig"
	bucket "github.com/dkb-bank/official-provider-aws/internal/controller/s3/bucket"
	bucketaccelerateconfiguration "github.com/dkb-bank/official-provider-aws/internal/controller/s3/bucketaccelerateconfiguration"
	bucketacl "github.com/dkb-bank/official-provider-aws/internal/controller/s3/bucketacl"
	bucketanalyticsconfiguration "github.com/dkb-bank/official-provider-aws/internal/controller/s3/bucketanalyticsconfiguration"
	bucketcorsconfiguration "github.com/dkb-bank/official-provider-aws/internal/controller/s3/bucketcorsconfiguration"
	bucketintelligenttieringconfiguration "github.com/dkb-bank/official-provider-aws/internal/controller/s3/bucketintelligenttieringconfiguration"
	bucketinventory "github.com/dkb-bank/official-provider-aws/internal/controller/s3/bucketinventory"
	bucketlifecycleconfiguration "github.com/dkb-bank/official-provider-aws/internal/controller/s3/bucketlifecycleconfiguration"
	bucketlogging "github.com/dkb-bank/official-provider-aws/internal/controller/s3/bucketlogging"
	bucketmetric "github.com/dkb-bank/official-provider-aws/internal/controller/s3/bucketmetric"
	bucketnotification "github.com/dkb-bank/official-provider-aws/internal/controller/s3/bucketnotification"
	bucketobject "github.com/dkb-bank/official-provider-aws/internal/controller/s3/bucketobject"
	bucketobjectlockconfiguration "github.com/dkb-bank/official-provider-aws/internal/controller/s3/bucketobjectlockconfiguration"
	bucketownershipcontrols "github.com/dkb-bank/official-provider-aws/internal/controller/s3/bucketownershipcontrols"
	bucketpolicy "github.com/dkb-bank/official-provider-aws/internal/controller/s3/bucketpolicy"
	bucketpublicaccessblock "github.com/dkb-bank/official-provider-aws/internal/controller/s3/bucketpublicaccessblock"
	bucketreplicationconfiguration "github.com/dkb-bank/official-provider-aws/internal/controller/s3/bucketreplicationconfiguration"
	bucketrequestpaymentconfiguration "github.com/dkb-bank/official-provider-aws/internal/controller/s3/bucketrequestpaymentconfiguration"
	bucketserversideencryptionconfiguration "github.com/dkb-bank/official-provider-aws/internal/controller/s3/bucketserversideencryptionconfiguration"
	bucketversioning "github.com/dkb-bank/official-provider-aws/internal/controller/s3/bucketversioning"
	bucketwebsiteconfiguration "github.com/dkb-bank/official-provider-aws/internal/controller/s3/bucketwebsiteconfiguration"
	object "github.com/dkb-bank/official-provider-aws/internal/controller/s3/object"
	objectcopy "github.com/dkb-bank/official-provider-aws/internal/controller/s3/objectcopy"
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
	resourcedatasync "github.com/dkb-bank/official-provider-aws/internal/controller/ssm/resourcedatasync"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alertmanagerdefinition.Setup,
		rulegroupnamespace.Setup,
		workspace.Setup,
		definition.Setup,
		destination.Setup,
		destinationpolicy.Setup,
		group.Setup,
		metricfilter.Setup,
		resourcepolicy.Setup,
		stream.Setup,
		subscriptionfilter.Setup,
		locations3.Setup,
		task.Setup,
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
		certificate.Setup,
		endpoint.Setup,
		eventsubscription.Setup,
		replicationinstance.Setup,
		replicationsubnetgroup.Setup,
		replicationtask.Setup,
		globaltable.Setup,
		table.Setup,
		tag.Setup,
		instance.Setup,
		internetgateway.Setup,
		keypair.Setup,
		natgateway.Setup,
		networkinterface.Setup,
		route.Setup,
		routetable.Setup,
		securitygroup.Setup,
		securitygrouprule.Setup,
		subnet.Setup,
		tagec2.Setup,
		transitgateway.Setup,
		vpc.Setup,
		vpcendpoint.Setup,
		vpcpeeringconnection.Setup,
		lifecyclepolicy.Setup,
		pullthroughcacherule.Setup,
		registrypolicy.Setup,
		registryscanningconfiguration.Setup,
		replicationconfiguration.Setup,
		repository.Setup,
		repositorypolicy.Setup,
		lb.Setup,
		lblistener.Setup,
		lblistenerrule.Setup,
		lbtargetgroup.Setup,
		lbtargetgroupattachment.Setup,
		deliverystream.Setup,
		roleassociation.Setup,
		workspacegrafana.Setup,
		workspacesamlconfiguration.Setup,
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
		streamkinesis.Setup,
		streamconsumer.Setup,
		application.Setup,
		applicationkinesisanalyticsv2.Setup,
		applicationsnapshot.Setup,
		streamkinesisvideo.Setup,
		alias.Setup,
		ciphertext.Setup,
		externalkey.Setup,
		grant.Setup,
		key.Setup,
		replicaexternalkey.Setup,
		replicakey.Setup,
		firewall.Setup,
		firewallpolicy.Setup,
		loggingconfiguration.Setup,
		rulegroup.Setup,
		providerconfig.Setup,
		bucket.Setup,
		bucketaccelerateconfiguration.Setup,
		bucketacl.Setup,
		bucketanalyticsconfiguration.Setup,
		bucketcorsconfiguration.Setup,
		bucketintelligenttieringconfiguration.Setup,
		bucketinventory.Setup,
		bucketlifecycleconfiguration.Setup,
		bucketlogging.Setup,
		bucketmetric.Setup,
		bucketnotification.Setup,
		bucketobject.Setup,
		bucketobjectlockconfiguration.Setup,
		bucketownershipcontrols.Setup,
		bucketpolicy.Setup,
		bucketpublicaccessblock.Setup,
		bucketreplicationconfiguration.Setup,
		bucketrequestpaymentconfiguration.Setup,
		bucketserversideencryptionconfiguration.Setup,
		bucketversioning.Setup,
		bucketwebsiteconfiguration.Setup,
		object.Setup,
		objectcopy.Setup,
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
		resourcedatasync.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	listapikey "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/access/listapikey"
	configuration "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/alert/configuration"
	key "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/api/key"
	compliancepolicy "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/backup/compliancepolicy"
	backupschedule "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/cloud/backupschedule"
	backupsnapshot "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/cloud/backupsnapshot"
	backupsnapshotexportbucket "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/cloud/backupsnapshotexportbucket"
	backupsnapshotexportjob "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/cloud/backupsnapshotexportjob"
	backupsnapshotrestorejob "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/cloud/backupsnapshotrestorejob"
	provideraccessauthorization "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/cloud/provideraccessauthorization"
	provideraccesssetup "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/cloud/provideraccesssetup"
	outagesimulation "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/cluster/outagesimulation"
	dbrole "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/custom/dbrole"
	dnsconfigurationclusteraws "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/custom/dnsconfigurationclusteraws"
	lakepipeline "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/data/lakepipeline"
	user "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/database/user"
	atrest "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/encryption/atrest"
	trigger "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/event/trigger"
	databaseinstance "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/federated/databaseinstance"
	querylimit "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/federated/querylimit"
	settingsidentityprovider "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/federated/settingsidentityprovider"
	settingsorgconfig "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/federated/settingsorgconfig"
	settingsorgrolemapping "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/federated/settingsorgrolemapping"
	clusterconfig "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/global/clusterconfig"
	configurationldap "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/ldap/configuration"
	verify "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/ldap/verify"
	window "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/maintenance/window"
	advancedcluster "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/mongodbatlas/advancedcluster"
	auditing "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/mongodbatlas/auditing"
	cluster "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/mongodbatlas/cluster"
	organization "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/mongodbatlas/organization"
	project "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/mongodbatlas/project"
	team "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/mongodbatlas/team"
	teams "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/mongodbatlas/teams"
	container "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/network/container"
	peering "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/network/peering"
	archive "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/online/archive"
	invitation "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/org/invitation"
	endpointregionalmode "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/private/endpointregionalmode"
	endpoint "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/privatelink/endpoint"
	endpointserverless "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/privatelink/endpointserverless"
	endpointservice "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/privatelink/endpointservice"
	endpointservicedatafederationonlinearchive "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/privatelink/endpointservicedatafederationonlinearchive"
	endpointserviceserverless "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/privatelink/endpointserviceserverless"
	apikey "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/project/apikey"
	invitationproject "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/project/invitation"
	ipaccesslist "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/project/ipaccesslist"
	providerconfig "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/providerconfig"
	index "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/search/index"
	instance "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/serverless/instance"
	partyintegration "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/third/partyintegration"
	authenticationdatabaseuser "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/x509/authenticationdatabaseuser"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		listapikey.Setup,
		configuration.Setup,
		key.Setup,
		compliancepolicy.Setup,
		backupschedule.Setup,
		backupsnapshot.Setup,
		backupsnapshotexportbucket.Setup,
		backupsnapshotexportjob.Setup,
		backupsnapshotrestorejob.Setup,
		provideraccessauthorization.Setup,
		provideraccesssetup.Setup,
		outagesimulation.Setup,
		dbrole.Setup,
		dnsconfigurationclusteraws.Setup,
		lakepipeline.Setup,
		user.Setup,
		atrest.Setup,
		trigger.Setup,
		databaseinstance.Setup,
		querylimit.Setup,
		settingsidentityprovider.Setup,
		settingsorgconfig.Setup,
		settingsorgrolemapping.Setup,
		clusterconfig.Setup,
		configurationldap.Setup,
		verify.Setup,
		window.Setup,
		advancedcluster.Setup,
		auditing.Setup,
		cluster.Setup,
		organization.Setup,
		project.Setup,
		team.Setup,
		teams.Setup,
		container.Setup,
		peering.Setup,
		archive.Setup,
		invitation.Setup,
		endpointregionalmode.Setup,
		endpoint.Setup,
		endpointserverless.Setup,
		endpointservice.Setup,
		endpointservicedatafederationonlinearchive.Setup,
		endpointserviceserverless.Setup,
		apikey.Setup,
		invitationproject.Setup,
		ipaccesslist.Setup,
		providerconfig.Setup,
		index.Setup,
		instance.Setup,
		partyintegration.Setup,
		authenticationdatabaseuser.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

module github.com/googleapis/repo-automation-bots/packages/bazelbot2

go 1.23.2

require (
	cloud.google.com/go/iam v1.2.2
	cloud.google.com/go/secretmanager v1.14.2
	github.com/go-git/go-git/v5 v5.12.0
	github.com/google/go-cmp v0.6.0
	github.com/google/go-github/v59 v59.0.0
	github.com/googleapis/gax-go/v2 v2.13.0
	github.com/julieqiu/derrors v0.0.0-20210614022941-f601489ffd41
	google.golang.org/api v0.205.0
	google.golang.org/genproto v0.0.0-20241104194629-dd2ea8efbc28
	google.golang.org/genproto/googleapis/api v0.0.0-20241104194629-dd2ea8efbc28
	google.golang.org/grpc v1.68.0
	google.golang.org/protobuf v1.35.1
	gopkg.in/yaml.v3 v3.0.1
)

require (
	cloud.google.com/go/auth v0.10.1 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.5 // indirect
	cloud.google.com/go/compute/metadata v0.5.2 // indirect
	dario.cat/mergo v1.0.0 // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/ProtonMail/go-crypto v1.0.0 // indirect
	github.com/cloudflare/circl v1.3.7 // indirect
	github.com/cyphar/filepath-securejoin v0.2.4 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-git/gcfg v1.5.1-0.20230307220236-3a3c6141e376 // indirect
	github.com/go-git/go-billy/v5 v5.5.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/s2a-go v0.1.8 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.4 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/kevinburke/ssh_config v1.2.0 // indirect
	github.com/pjbgf/sha1cd v0.3.0 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	github.com/sergi/go-diff v1.3.2-0.20230802210424-5b0b94c5c0d3 // indirect
	github.com/skeema/knownhosts v1.2.2 // indirect
	github.com/xanzy/ssh-agent v0.3.3 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.54.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.54.0 // indirect
	go.opentelemetry.io/otel v1.29.0 // indirect
	go.opentelemetry.io/otel/metric v1.29.0 // indirect
	go.opentelemetry.io/otel/trace v1.29.0 // indirect
	golang.org/x/crypto v0.28.0 // indirect
	golang.org/x/mod v0.20.0 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/oauth2 v0.23.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	golang.org/x/time v0.7.0 // indirect
	golang.org/x/tools v0.24.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241021214115-324edc3d5d38 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
)

replace cloud.google.com/go/secretmanager => ../../../secretmanager

replace cloud.google.com/go/accessapproval => ../../../accessapproval

replace cloud.google.com/go/accesscontextmanager => ../../../accesscontextmanager

replace cloud.google.com/go/ai => ../../../ai

replace cloud.google.com/go/aiplatform => ../../../aiplatform

replace cloud.google.com/go/advisorynotifications => ../../../advisorynotifications

replace cloud.google.com/go/alloydb => ../../../alloydb

replace cloud.google.com/go/analytics => ../../../analytics

replace cloud.google.com/go/apigateway => ../../../apigateway

replace cloud.google.com/go/apigeeconnect => ../../../apigeeconnect

replace cloud.google.com/go/apigeeregistry => ../../../apigeeregistry

replace cloud.google.com/go/apihub => ../../../apihub

replace cloud.google.com/go/apikeys => ../../../apikeys

replace cloud.google.com/go/appengine => ../../../appengine

replace cloud.google.com/go/apphub => ../../../apphub

replace cloud.google.com/go/apps => ../../../apps

replace cloud.google.com/go/area120 => ../../../area120

replace cloud.google.com/go/artifactregistry => ../../../artifactregistry

replace cloud.google.com/go/asset => ../../../asset

replace cloud.google.com/go/assuredworkloads => ../../../assuredworkloads

replace cloud.google.com/go/automl => ../../../automl

replace cloud.google.com/go/backupdr => ../../../backupdr

replace cloud.google.com/go/baremetalsolution => ../../../baremetalsolution

replace cloud.google.com/go/batch => ../../../batch

replace cloud.google.com/go/beyondcorp => ../../../beyondcorp

replace cloud.google.com/go/bigquery => ../../../bigquery

replace cloud.google.com/go/bigtable => ../../../bigtable

replace cloud.google.com/go/billing => ../../../billing

replace cloud.google.com/go/binaryauthorization => ../../../binaryauthorization

replace cloud.google.com/go/certificatemanager => ../../../certificatemanager

replace cloud.google.com/go/channel => ../../../channel

replace cloud.google.com/go/chat => ../../../chat

replace cloud.google.com/go/cloudbuild => ../../../cloudbuild

replace cloud.google.com/go/cloudcontrolspartner => ../../../cloudcontrolspartner

replace cloud.google.com/go/clouddms => ../../../clouddms

replace cloud.google.com/go/cloudprofiler => ../../../cloudprofiler

replace cloud.google.com/go/cloudquotas => ../../../cloudquotas

replace cloud.google.com/go/cloudtasks => ../../../cloudtasks

replace cloud.google.com/go/commerce => ../../../commerce

replace cloud.google.com/go/compute => ../../../compute

replace cloud.google.com/go/confidentialcomputing => ../../../confidentialcomputing

replace cloud.google.com/go/config => ../../../config

replace cloud.google.com/go/contactcenterinsights => ../../../contactcenterinsights

replace cloud.google.com/go/container => ../../../container

replace cloud.google.com/go/containeranalysis => ../../../containeranalysis

replace cloud.google.com/go/datacatalog => ../../../datacatalog

replace cloud.google.com/go/dataflow => ../../../dataflow

replace cloud.google.com/go/dataform => ../../../dataform

replace cloud.google.com/go/datafusion => ../../../datafusion

replace cloud.google.com/go/datalabeling => ../../../datalabeling

replace cloud.google.com/go/dataplex => ../../../dataplex

replace cloud.google.com/go/dataproc => ../../../dataproc

replace cloud.google.com/go/dataqna => ../../../dataqna

replace cloud.google.com/go/datastore => ../../../datastore

replace cloud.google.com/go/datastream => ../../../datastream

replace cloud.google.com/go/deploy => ../../../deploy

replace cloud.google.com/go/developerconnect => ../../../developerconnect

replace cloud.google.com/go/dialogflow => ../../../dialogflow

replace cloud.google.com/go/discoveryengine => ../../../discoveryengine

replace cloud.google.com/go/dlp => ../../../dlp

replace cloud.google.com/go/documentai => ../../../documentai

replace cloud.google.com/go/domains => ../../../domains

replace cloud.google.com/go/edgecontainer => ../../../edgecontainer

replace cloud.google.com/go/edgenetwork => ../../../edgenetwork

replace cloud.google.com/go/errorreporting => ../../../errorreporting

replace cloud.google.com/go/essentialcontacts => ../../../essentialcontacts

replace cloud.google.com/go/eventarc => ../../../eventarc

replace cloud.google.com/go/filestore => ../../../filestore

replace cloud.google.com/go/firestore => ../../../firestore

replace cloud.google.com/go/functions => ../../../functions

replace cloud.google.com/go/gkebackup => ../../../gkebackup

replace cloud.google.com/go/gkeconnect => ../../../gkeconnect

replace cloud.google.com/go/gkehub => ../../../gkehub

replace cloud.google.com/go/gkemulticloud => ../../../gkemulticloud

replace cloud.google.com/go/gsuiteaddons => ../../../gsuiteaddons

replace cloud.google.com/go/iam => ../../../iam

replace cloud.google.com/go/iap => ../../../iap

replace cloud.google.com/go/identitytoolkit => ../../../identitytoolkit

replace cloud.google.com/go/ids => ../../../ids

replace cloud.google.com/go/iot => ../../../iot

replace cloud.google.com/go/kms => ../../../kms

replace cloud.google.com/go/language => ../../../language

replace cloud.google.com/go/lifesciences => ../../../lifesciences

replace cloud.google.com/go/logging => ../../../logging

replace cloud.google.com/go/longrunning => ../../../longrunning

replace cloud.google.com/go/managedidentities => ../../../managedidentities

replace cloud.google.com/go/managedkafka => ../../../managedkafka

replace cloud.google.com/go/maps => ../../../maps

replace cloud.google.com/go/mediatranslation => ../../../mediatranslation

replace cloud.google.com/go/memcache => ../../../memcache

replace cloud.google.com/go/metastore => ../../../metastore

replace cloud.google.com/go/migrationcenter => ../../../migrationcenter

replace cloud.google.com/go/monitoring => ../../../monitoring

replace cloud.google.com/go/netapp => ../../../netapp

replace cloud.google.com/go/networkconnectivity => ../../../networkconnectivity

replace cloud.google.com/go/networkmanagement => ../../../networkmanagement

replace cloud.google.com/go/networksecurity => ../../../networksecurity

replace cloud.google.com/go/networkservices => ../../../networkservices

replace cloud.google.com/go/notebooks => ../../../notebooks

replace cloud.google.com/go/optimization => ../../../optimization

replace cloud.google.com/go/oracledatabase => ../../../oracledatabase

replace cloud.google.com/go/orchestration => ../../../orchestration

replace cloud.google.com/go/orgpolicy => ../../../orgpolicy

replace cloud.google.com/go/osconfig => ../../../osconfig

replace cloud.google.com/go/oslogin => ../../../oslogin

replace cloud.google.com/go/parallelstore => ../../../parallelstore

replace cloud.google.com/go/phishingprotection => ../../../phishingprotection

replace cloud.google.com/go/policysimulator => ../../../policysimulator

replace cloud.google.com/go/policytroubleshooter => ../../../policytroubleshooter

replace cloud.google.com/go/privatecatalog => ../../../privatecatalog

replace cloud.google.com/go/privilegedaccessmanager => ../../../privilegedaccessmanager

replace cloud.google.com/go/pubsub => ../../../pubsub

replace cloud.google.com/go/pubsublite => ../../../pubsublite

replace cloud.google.com/go/rapidmigrationassessment => ../../../rapidmigrationassessment

replace cloud.google.com/go/recaptchaenterprise => ../../../recaptchaenterprise

replace cloud.google.com/go/recommendationengine => ../../../recommendationengine

replace cloud.google.com/go/recommender => ../../../recommender

replace cloud.google.com/go/redis => ../../../redis

replace cloud.google.com/go/resourcemanager => ../../../resourcemanager

replace cloud.google.com/go/resourcesettings => ../../../resourcesettings

replace cloud.google.com/go/retail => ../../../retail

replace cloud.google.com/go/run => ../../../run

replace cloud.google.com/go/scheduler => ../../../scheduler

replace cloud.google.com/go/securesourcemanager => ../../../securesourcemanager

replace cloud.google.com/go/security => ../../../security

replace cloud.google.com/go/securitycenter => ../../../securitycenter

replace cloud.google.com/go/securitycentermanagement => ../../../securitycentermanagement

replace cloud.google.com/go/securityposture => ../../../securityposture

replace cloud.google.com/go/servicecontrol => ../../../servicecontrol

replace cloud.google.com/go/servicedirectory => ../../../servicedirectory

replace cloud.google.com/go/servicehealth => ../../../servicehealth

replace cloud.google.com/go/servicemanagement => ../../../servicemanagement

replace cloud.google.com/go/serviceusage => ../../../serviceusage

replace cloud.google.com/go/shell => ../../../shell

replace cloud.google.com/go/shopping => ../../../shopping

replace cloud.google.com/go/spanner => ../../../spanner

replace cloud.google.com/go/speech => ../../../speech

replace cloud.google.com/go/storage => ../../../storage

replace cloud.google.com/go/storageinsights => ../../../storageinsights

replace cloud.google.com/go/storagetransfer => ../../../storagetransfer

replace cloud.google.com/go/streetview => ../../../streetview

replace cloud.google.com/go/support => ../../../support

replace cloud.google.com/go/talent => ../../../talent

replace cloud.google.com/go/telcoautomation => ../../../telcoautomation

replace cloud.google.com/go/texttospeech => ../../../texttospeech

replace cloud.google.com/go/tpu => ../../../tpu

replace cloud.google.com/go/trace => ../../../trace

replace cloud.google.com/go/translate => ../../../translate

replace cloud.google.com/go/video => ../../../video

replace cloud.google.com/go/videointelligence => ../../../videointelligence

replace cloud.google.com/go/vision => ../../../vision

replace cloud.google.com/go/visionai => ../../../visionai

replace cloud.google.com/go/vmmigration => ../../../vmmigration

replace cloud.google.com/go/vmwareengine => ../../../vmwareengine

replace cloud.google.com/go/vpcaccess => ../../../vpcaccess

replace cloud.google.com/go/webrisk => ../../../webrisk

replace cloud.google.com/go/websecurityscanner => ../../../websecurityscanner

replace cloud.google.com/go/workflows => ../../../workflows

replace cloud.google.com/go/workstations => ../../../workstations

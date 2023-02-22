package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/accessanalyzer"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/account"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/acm"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/amp"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/amplify"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/apigateway"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/apigatewayv2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/applicationautoscaling"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/apprunner"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/appstream"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/appsync"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/athena"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/autoscaling"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/backup"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/cloudformation"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/cloudfront"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/cloudhsmv2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/cloudtrail"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/cloudwatch"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/cloudwatchlogs"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/codebuild"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/codepipeline"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/cognito"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/config"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/dax"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/directconnect"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/dms"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/docdb"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/dynamodb"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ec2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ecr"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ecrpublic"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ecs"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/efs"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/eks"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/elasticache"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/elasticbeanstalk"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/elasticsearch"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/elastictranscoder"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/elbv1"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/elbv2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/emr"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/eventbridge"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/firehose"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/frauddetector"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/fsx"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/glacier"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/glue"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/guardduty"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/iam"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/identitystore"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/inspector"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/inspector2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/iot"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/kafka"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/kinesis"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/kms"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/lambda"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/lightsail"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/mq"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/mwaa"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/neptune"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/organizations"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/qldb"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/quicksight"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ram"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/rds"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/redshift"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/resiliencehub"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/resourcegroups"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/route53"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/s3"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/sagemaker"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/savingsplans"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/scheduler"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/secretsmanager"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/securityhub"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/servicecatalog"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/servicequotas"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ses"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/shield"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/sns"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/sqs"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ssm"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ssoadmin"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/stepfunctions"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/support"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/timestream"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/transfer"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/waf"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/wafregional"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/wafv2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/workspaces"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/xray"
	"github.com/cloudquery/plugin-sdk/schema"
)

func tables() []*schema.Table {
	return []*schema.Table{
		accessanalyzer.Analyzers(),
		account.AlternateContacts(),
		account.Contacts(),
		acm.Certificates(),
		amp.Workspaces(),
		amplify.Apps(),
		apigateway.ApiKeys(),
		apigateway.ClientCertificates(),
		apigateway.DomainNames(),
		apigateway.RestApis(),
		apigateway.UsagePlans(),
		apigateway.VpcLinks(),
		apigatewayv2.Apis(),
		apigatewayv2.DomainNames(),
		apigatewayv2.VpcLinks(),
		applicationautoscaling.Policies(),
		apprunner.AutoScalingConfigurations(),
		apprunner.Connections(),
		apprunner.ObservabilityConfigurations(),
		apprunner.Services(),
		apprunner.VpcConnectors(),
		apprunner.VpcIngressConnections(),
		appstream.AppBlocks(),
		appstream.Applications(),
		appstream.DirectoryConfigs(),
		appstream.Fleets(),
		appstream.ImageBuilders(),
		appstream.Images(),
		appstream.Stacks(),
		appstream.UsageReportSubscriptions(),
		appstream.Users(),
		appsync.GraphqlApis(),
		athena.DataCatalogs(),
		athena.WorkGroups(),
		autoscaling.Groups(),
		autoscaling.LaunchConfigurations(),
		autoscaling.ScheduledActions(),
		backup.GlobalSettings(),
		backup.Plans(),
		backup.RegionSettings(),
		backup.Vaults(),
		cloudformation.Stacks(),
		cloudfront.CachePolicies(),
		cloudfront.Distributions(),
		cloudhsmv2.Backups(),
		cloudhsmv2.Clusters(),
		cloudtrail.Trails(),
		cloudwatch.Alarms(),
		cloudwatchlogs.LogGroups(),
		cloudwatchlogs.MetricFilters(),
		cloudwatchlogs.ResourcePolicies(),
		codebuild.Projects(),
		codepipeline.Pipelines(),
		codepipeline.Webhooks(),
		cognito.IdentityPools(),
		cognito.UserPools(),
		config.ConfigRules(),
		config.ConfigurationRecorders(),
		config.ConformancePacks(),
		dax.Clusters(),
		directconnect.Connections(),
		directconnect.Gateways(),
		directconnect.Lags(),
		directconnect.Locations(),
		directconnect.VirtualGateways(),
		directconnect.VirtualInterfaces(),
		dms.ReplicationInstances(),
		docdb.Certificates(),
		docdb.ClusterParameterGroups(),
		docdb.Clusters(),
		docdb.EngineVersions(),
		docdb.EventCategories(),
		docdb.Events(),
		docdb.EventSubscriptions(),
		docdb.GlobalClusters(),
		docdb.PendingMaintenanceActions(),
		docdb.SubnetGroups(),
		dynamodb.Tables(),
		ec2.AvailabilityZones(),
		ec2.ByoipCidrs(),
		ec2.CustomerGateways(),
		ec2.EbsSnapshots(),
		ec2.EbsVolumes(),
		ec2.EbsVolumesStatuses(),
		ec2.EgressOnlyInternetGateways(),
		ec2.Eips(),
		ec2.FlowLogs(),
		ec2.Hosts(),
		ec2.Images(),
		ec2.Instances(),
		ec2.InstanceStatuses(),
		ec2.InstanceTypes(),
		ec2.InternetGateways(),
		ec2.KeyPairs(),
		ec2.ManagedPrefixLists(),
		ec2.NatGateways(),
		ec2.NetworkAcls(),
		ec2.NetworkInterfaces(),
		ec2.RegionalConfigs(),
		ec2.Regions(),
		ec2.ReservedInstances(),
		ec2.RouteTables(),
		ec2.SecurityGroups(),
		ec2.Subnets(),
		ec2.TransitGateways(),
		ec2.VpcEndpoints(),
		ec2.VpcEndpointServiceConfigurations(),
		ec2.VpcEndpointServices(),
		ec2.VpcPeeringConnections(),
		ec2.Vpcs(),
		ec2.VpnGateways(),
		ecr.Registries(),
		ecr.RegistryPolicies(),
		ecr.Repositories(),
		ecrpublic.Repositories(),
		ecs.Clusters(),
		ecs.TaskDefinitions(),
		efs.Filesystems(),
		eks.Clusters(),
		elasticache.Clusters(),
		elasticache.EngineVersions(),
		elasticache.Events(),
		elasticache.GlobalReplicationGroups(),
		elasticache.ParameterGroups(),
		elasticache.ReplicationGroups(),
		elasticache.ReservedCacheNodes(),
		elasticache.ReservedCacheNodesOfferings(),
		elasticache.ServiceUpdates(),
		elasticache.Snapshots(),
		elasticache.SubnetGroups(),
		elasticache.UpdateActions(),
		elasticache.UserGroups(),
		elasticache.Users(),
		elasticbeanstalk.Applications(),
		elasticbeanstalk.ApplicationVersions(),
		elasticbeanstalk.Environments(),
		elasticsearch.Domains(),
		elasticsearch.Packages(),
		elasticsearch.Versions(),
		elasticsearch.VpcEndpoints(),
		elastictranscoder.Pipelines(),
		elastictranscoder.Presets(),
		elbv1.LoadBalancers(),
		elbv2.LoadBalancers(),
		elbv2.TargetGroups(),
		emr.BlockPublicAccessConfigs(),
		emr.Clusters(),
		emr.SecurityConfigurations(),
		eventbridge.ApiDestinations(),
		eventbridge.Archives(),
		eventbridge.Connections(),
		eventbridge.Endpoints(),
		eventbridge.EventBuses(),
		eventbridge.EventSources(),
		eventbridge.Replays(),
		firehose.DeliveryStreams(),
		frauddetector.BatchImports(),
		frauddetector.BatchPredictions(),
		frauddetector.Detectors(),
		frauddetector.EntityTypes(),
		frauddetector.EventTypes(),
		frauddetector.ExternalModels(),
		frauddetector.Labels(),
		frauddetector.Models(),
		frauddetector.Outcomes(),
		frauddetector.Variables(),
		fsx.Backups(),
		fsx.DataRepositoryAssociations(),
		fsx.DataRepositoryTasks(),
		fsx.FileCaches(),
		fsx.FileSystems(),
		fsx.Snapshots(),
		fsx.StorageVirtualMachines(),
		fsx.Volumes(),
		glacier.DataRetrievalPolicies(),
		glacier.Vaults(),
		glue.Classifiers(),
		glue.Connections(),
		glue.Crawlers(),
		glue.Databases(),
		glue.DatacatalogEncryptionSettings(),
		glue.DevEndpoints(),
		glue.Jobs(),
		glue.MlTransforms(),
		glue.Registries(),
		glue.SecurityConfigurations(),
		glue.Triggers(),
		glue.Workflows(),
		guardduty.Detectors(),
		iam.Accounts(),
		iam.CredentialReports(),
		iam.Groups(),
		iam.InstanceProfiles(),
		iam.OpenidConnectIdentityProviders(),
		iam.PasswordPolicies(),
		iam.Policies(),
		iam.Roles(),
		iam.SamlIdentityProviders(),
		iam.ServerCertificates(),
		iam.Users(),
		iam.VirtualMfaDevices(),
		identitystore.Groups(),
		identitystore.Users(),
		inspector.Findings(),
		inspector2.Findings(),
		iot.BillingGroups(),
		iot.CaCertificates(),
		iot.Certificates(),
		iot.Jobs(),
		iot.Policies(),
		iot.SecurityProfiles(),
		iot.Streams(),
		iot.ThingGroups(),
		iot.Things(),
		iot.ThingTypes(),
		iot.TopicRules(),
		kafka.Clusters(),
		kafka.Configurations(),
		kinesis.Streams(),
		kms.Aliases(),
		kms.Keys(),
		lambda.Functions(),
		lambda.Layers(),
		lambda.Runtimes(),
		lightsail.Alarms(),
		lightsail.Buckets(),
		lightsail.Certificates(),
		lightsail.ContainerServices(),
		lightsail.Databases(),
		lightsail.DatabaseSnapshots(),
		lightsail.Disks(),
		lightsail.Distributions(),
		lightsail.Instances(),
		lightsail.InstanceSnapshots(),
		lightsail.LoadBalancers(),
		lightsail.StaticIps(),
		mq.Brokers(),
		mwaa.Environments(),
		neptune.ClusterParameterGroups(),
		neptune.Clusters(),
		neptune.ClusterSnapshots(),
		neptune.DbParameterGroups(),
		neptune.EventSubscriptions(),
		neptune.GlobalClusters(),
		neptune.Instances(),
		neptune.SubnetGroups(),
		organizations.Accounts(),
		organizations.Organizations(),
		organizations.Policies(),
		organizations.ResourcePolicies(),
		organizations.Roots(),
		organizations.OrganizationalUnits(),
		qldb.Ledgers(),
		quicksight.Analyses(),
		quicksight.Dashboards(),
		quicksight.DataSets(),
		quicksight.DataSources(),
		quicksight.Folders(),
		quicksight.Groups(),
		quicksight.Templates(),
		quicksight.Users(),
		ram.Principals(),
		ram.Resources(),
		ram.ResourceShareAssociations(),
		ram.ResourceShareInvitations(),
		ram.ResourceShares(),
		ram.ResourceTypes(),
		rds.Certificates(),
		rds.ClusterParameterGroups(),
		rds.Clusters(),
		rds.ClusterSnapshots(),
		rds.DbParameterGroups(),
		rds.DbProxies(),
		rds.DbSecurityGroups(),
		rds.DbSnapshots(),
		rds.EngineVersions(),
		rds.EventSubscriptions(),
		rds.Instances(),
		rds.ReservedInstances(),
		rds.SubnetGroups(),
		redshift.Clusters(),
		redshift.EventSubscriptions(),
		redshift.SubnetGroups(),
		resiliencehub.Apps(),
		resiliencehub.ResiliencyPolicies(),
		resiliencehub.SuggestedResiliencyPolicies(),
		resourcegroups.ResourceGroups(),
		route53.DelegationSets(),
		route53.Domains(),
		route53.HealthChecks(),
		route53.HostedZones(),
		route53.Operations(),
		route53.TrafficPolicies(),
		s3.AccessPoints(),
		s3.Accounts(),
		s3.Buckets(),
		sagemaker.EndpointConfigurations(),
		sagemaker.Models(),
		sagemaker.NotebookInstances(),
		sagemaker.TrainingJobs(),
		savingsplans.Plans(),
		scheduler.ScheduleGroups(),
		scheduler.Schedules(),
		secretsmanager.Secrets(),
		securityhub.Findings(),
		servicecatalog.Portfolios(),
		servicecatalog.Products(),
		servicecatalog.ProvisionedProducts(),
		servicequotas.Services(),
		ses.ActiveReceiptRuleSets(),
		ses.ConfigurationSets(),
		ses.ContactLists(),
		ses.CustomVerificationEmailTemplates(),
		ses.Identities(),
		ses.Templates(),
		shield.Attacks(),
		shield.ProtectionGroups(),
		shield.Protections(),
		shield.Subscriptions(),
		sns.Subscriptions(),
		sns.Topics(),
		sqs.Queues(),
		ssm.Associations(),
		ssm.ComplianceSummaryItems(),
		ssm.Documents(),
		ssm.Instances(),
		ssm.Inventories(),
		ssm.InventorySchemas(),
		ssm.Parameters(),
		ssm.PatchBaselines(),
		ssoadmin.Instances(),
		stepfunctions.Activities(),
		stepfunctions.StateMachines(),
		support.Cases(),
		support.Services(),
		support.SeverityLevels(),
		support.TrustedAdvisorChecks(),
		timestream.Databases(),
		transfer.Servers(),
		waf.RuleGroups(),
		waf.Rules(),
		waf.SubscribedRuleGroups(),
		waf.WebAcls(),
		wafregional.RateBasedRules(),
		wafregional.RuleGroups(),
		wafregional.Rules(),
		wafregional.WebAcls(),
		wafv2.Ipsets(),
		wafv2.ManagedRuleGroups(),
		wafv2.RegexPatternSets(),
		wafv2.RuleGroups(),
		wafv2.WebAcls(),
		workspaces.Directories(),
		workspaces.Workspaces(),
		xray.EncryptionConfigs(),
		xray.Groups(),
		xray.ResourcePolicies(),
		xray.SamplingRules(),
	}
}

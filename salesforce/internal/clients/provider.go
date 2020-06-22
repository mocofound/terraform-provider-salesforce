
package clients

import (
	"context"

	users "github.com/mocofound/terraform-provider-salesforce/salesforce/internal/services/users/client"
)

type Client struct {
	// StopContext is used for propagating control from Terraform Core (e.g. Ctrl/Cmd+C)
	StopContext context.Context

	Account  *ResourceManagerAccount
	Features features.UserFeatures

	Advisor               *advisor.Client
	AnalysisServices      *analysisServices.Client
	ApiManagement         *apiManagement.Client
	AppConfiguration      *appConfiguration.Client
	AppInsights           *applicationInsights.Client
	AppPlatform           *appPlatform.Client
	Authorization         *authorization.Client
	Automation            *automation.Client
	Batch                 *batch.Client
	Blueprints            *blueprints.Client
	Bot                   *bot.Client
	Cdn                   *cdn.Client
	Cognitive             *cognitiveServices.Client
	Compute               *compute.Client
	Containers            *containerServices.Client
	Cosmos                *cosmosdb.Client
	CostManagement        *costmanagement.Client
	CustomProviders       *customproviders.Client
	DatabaseMigration     *datamigration.Client
	DataBricks            *databricks.Client
	DataFactory           *datafactory.Client
	Datalake              *datalake.Client
	DataShare             *datashare.Client
	DevSpace              *devspace.Client
	DevTestLabs           *devtestlabs.Client
	Dns                   *dns.Client
	EventGrid             *eventgrid.Client
	Eventhub              *eventhub.Client
	Frontdoor             *frontdoor.Client
	HDInsight             *hdinsight.Client
	HealthCare            *healthcare.Client
	IoTCentral            *iotcentral.Client
	IoTHub                *iothub.Client
	IoTTimeSeriesInsights *timeseriesinsights.Client
	KeyVault              *keyvault.Client
	Kusto                 *kusto.Client
	LogAnalytics          *loganalytics.Client
	Logic                 *logic.Client
	MachineLearning       *machinelearning.Client
	Maintenance           *maintenance.Client
	ManagedApplication    *managedapplication.Client
	ManagementGroups      *managementgroup.Client
	Maps                  *maps.Client
	MariaDB               *mariadb.Client
	Media                 *media.Client
	MixedReality          *mixedreality.Client
	Monitor               *monitor.Client
	MSI                   *msi.Client
	MSSQL                 *mssql.Client
	MySQL                 *mysql.Client
	NetApp                *netapp.Client
	Network               *network.Client
	NotificationHubs      *notificationhub.Client
	Policy                *policy.Client
	Portal                *portal.Client
	Postgres              *postgres.Client
	PowerBI               *powerBI.Client
	PrivateDns            *privatedns.Client
	RecoveryServices      *recoveryServices.Client
	Redis                 *redis.Client
	Relay                 *relay.Client
	Resource              *resource.Client
	Search                *search.Client
	SecurityCenter        *securityCenter.Client
	Sentinel              *sentinel.Client
	ServiceBus            *serviceBus.Client
	ServiceFabric         *serviceFabric.Client
	SignalR               *signalr.Client
	Storage               *storage.Client
	StreamAnalytics       *streamAnalytics.Client
	Subscription          *subscription.Client
	Sql                   *sql.Client
	TrafficManager        *trafficManager.Client
	Web                   *web.Client
}

// NOTE: it should be possible for this method to become Private once the top level Client's removed

func (client *Client) Build(ctx context.Context, o *common.ClientOptions) error {
	autorest.Count429AsRetry = false

	client.Features = o.Features
	client.StopContext = ctx

	client.Advisor = advisor.NewClient(o)
	client.AnalysisServices = analysisServices.NewClient(o)
	client.ApiManagement = apiManagement.NewClient(o)
	client.AppConfiguration = appConfiguration.NewClient(o)
	client.AppInsights = applicationInsights.NewClient(o)
	client.AppPlatform = appPlatform.NewClient(o)
	client.Authorization = authorization.NewClient(o)
	client.Automation = automation.NewClient(o)
	client.Batch = batch.NewClient(o)
	client.Blueprints = blueprints.NewClient(o)
	client.Bot = bot.NewClient(o)
	client.Cdn = cdn.NewClient(o)
	client.Cognitive = cognitiveServices.NewClient(o)
	client.Compute = compute.NewClient(o)
	client.Containers = containerServices.NewClient(o)
	client.Cosmos = cosmosdb.NewClient(o)
	client.CostManagement = costmanagement.NewClient(o)
	client.CustomProviders = customproviders.NewClient(o)
	client.DatabaseMigration = datamigration.NewClient(o)
	client.DataBricks = databricks.NewClient(o)
	client.DataFactory = datafactory.NewClient(o)
	client.Datalake = datalake.NewClient(o)
	client.DataShare = datashare.NewClient(o)
	client.DevSpace = devspace.NewClient(o)
	client.DevTestLabs = devtestlabs.NewClient(o)
	client.Dns = dns.NewClient(o)
	client.EventGrid = eventgrid.NewClient(o)
	client.Eventhub = eventhub.NewClient(o)
	client.Frontdoor = frontdoor.NewClient(o)
	client.HDInsight = hdinsight.NewClient(o)
	client.HealthCare = healthcare.NewClient(o)
	client.IoTCentral = iotcentral.NewClient(o)
	client.IoTHub = iothub.NewClient(o)
	client.IoTTimeSeriesInsights = timeseriesinsights.NewClient(o)
	client.KeyVault = keyvault.NewClient(o)
	client.Kusto = kusto.NewClient(o)
	client.LogAnalytics = loganalytics.NewClient(o)
	client.Logic = logic.NewClient(o)
	client.MachineLearning = machinelearning.NewClient(o)
	client.Maintenance = maintenance.NewClient(o)
	client.ManagedApplication = managedapplication.NewClient(o)
	client.ManagementGroups = managementgroup.NewClient(o)
	client.Maps = maps.NewClient(o)
	client.MariaDB = mariadb.NewClient(o)
	client.Media = media.NewClient(o)
	client.MixedReality = mixedreality.NewClient(o)
	client.Monitor = monitor.NewClient(o)
	client.MSI = msi.NewClient(o)
	client.MSSQL = mssql.NewClient(o)
	client.MySQL = mysql.NewClient(o)
	client.NetApp = netapp.NewClient(o)
	client.Network = network.NewClient(o)
	client.NotificationHubs = notificationhub.NewClient(o)
	client.Policy = policy.NewClient(o)
	client.Portal = portal.NewClient(o)
	client.Postgres = postgres.NewClient(o)
	client.PowerBI = powerBI.NewClient(o)
	client.PrivateDns = privatedns.NewClient(o)
	client.RecoveryServices = recoveryServices.NewClient(o)
	client.Redis = redis.NewClient(o)
	client.Relay = relay.NewClient(o)
	client.Resource = resource.NewClient(o)
	client.Search = search.NewClient(o)
	client.SecurityCenter = securityCenter.NewClient(o)
	client.Sentinel = sentinel.NewClient(o)
	client.ServiceBus = serviceBus.NewClient(o)
	client.ServiceFabric = serviceFabric.NewClient(o)
	client.SignalR = signalr.NewClient(o)
	client.Sql = sql.NewClient(o)
	client.Storage = storage.NewClient(o)
	client.StreamAnalytics = streamAnalytics.NewClient(o)
	client.Subscription = subscription.NewClient(o)
	client.TrafficManager = trafficManager.NewClient(o)
	client.Web = web.NewClient(o)

	return nil
}
© 2020 GitHub, Inc.
Terms
Privacy
Security
Status
Help
Contact GitHub
Pricing
API
Training
Blog
About

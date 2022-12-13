package recipes

import (
	clouddms "cloud.google.com/go/clouddms/apiv1"
	clouddmspb "cloud.google.com/go/clouddms/apiv1/clouddmspb"
)


func init() {
	resources := []*Resource{
		{
			Service: "clouddms",
			SubService: "connection_profiles",
			Struct: &clouddmspb.ConnectionProfile{},
			NewFunction: clouddms.NewDataMigrationClient,
			ListFunction: (&clouddms.DataMigrationClient{}).ListConnectionProfiles,
			RequestStruct: &clouddmspb.ListConnectionProfilesRequest{},
			ResponseStruct: &clouddmspb.ListConnectionProfilesResponse{},
			UnimplementedServer: &clouddmspb.UnimplementedDataMigrationServiceServer{},
			RegisterServer: clouddmspb.RegisterDataMigrationServiceServer,
			Template: "newapi_list",
			MockTemplate: "newapi_list_grpc_mock",
			ProtobufImport: "cloud.google.com/go/clouddms/apiv1/clouddmspb",
			MockImports: []string{"cloud.google.com/go/clouddms/apiv1"},
		},
		{
			Service: "clouddms",
			SubService: "migration_jobs",
			Struct: &clouddmspb.MigrationJob{},
			NewFunction: clouddms.NewDataMigrationClient,
			ListFunction: (&clouddms.DataMigrationClient{}).ListMigrationJobs,
			RequestStruct: &clouddmspb.ListMigrationJobsRequest{},
			ResponseStruct: &clouddmspb.ListMigrationJobsResponse{},
			UnimplementedServer: &clouddmspb.UnimplementedDataMigrationServiceServer{},
			RegisterServer: clouddmspb.RegisterDataMigrationServiceServer,
			Template: "newapi_list",
			MockTemplate: "newapi_list_grpc_mock",
			ProtobufImport: "cloud.google.com/go/clouddms/apiv1/clouddmspb",
			MockImports: []string{"cloud.google.com/go/clouddms/apiv1"},
		},
	}

	Resources = append(Resources, resources...)
}

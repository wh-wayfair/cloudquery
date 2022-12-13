package recipes

import (
	assuredworkloads "cloud.google.com/go/assuredworkloads/apiv1"
	assuredworkloadspb "cloud.google.com/go/assuredworkloads/apiv1/assuredworkloadspb"
)


func init() {
	resources := []*Resource{
		{
			Service: "assuredworkloads",
			SubService: "workloads",
			Struct: &assuredworkloadspb.Workload{},
			NewFunction: assuredworkloads.NewClient,
			ListFunction: (&assuredworkloads.Client{}).ListWorkloads,
			RequestStruct: &assuredworkloadspb.ListWorkloadsRequest{},
			ResponseStruct: &assuredworkloadspb.ListWorkloadsResponse{},
			UnimplementedServer: &assuredworkloadspb.UnimplementedAssuredWorkloadsServiceServer{},
			RegisterServer: assuredworkloadspb.RegisterAssuredWorkloadsServiceServer,
			Template: "newapi_list",
			MockTemplate: "newapi_list_grpc_mock",
			ProtobufImport: "cloud.google.com/go/assuredworkloads/apiv1/assuredworkloadspb",
			MockImports: []string{"cloud.google.com/go/assuredworkloads/apiv1"},
		},
	}

	Resources = append(Resources, resources...)
}

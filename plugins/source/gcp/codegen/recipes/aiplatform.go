package recipes

import (
	aiplatform "cloud.google.com/go/aiplatform/apiv1"
	aiplatformpb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
)


func init() {
	resources := []*Resource{
		{
			Service: "aiplatform",
			SubService: "requests",
			Struct: &aiplatformpb.Dataset{},
			NewFunction: aiplatform.NewDatasetClient,
			ListFunction: (&aiplatform.DatasetClient{}).ListDatasets,
			RequestStruct: &aiplatformpb.ListDatasetsRequest{},
			ResponseStruct: &aiplatformpb.ListDatasetsResponse{},
			UnimplementedServer: &aiplatformpb.UnimplementedDatasetServiceServer{},
			RegisterServer: aiplatformpb.RegisterDatasetServiceServer,
			Template: "newapi_list",
			MockTemplate: "newapi_list_grpc_mock",
			ProtobufImport: "cloud.google.com/go/aiplatform/apiv1/aiplatformpb",
			MockImports: []string{"cloud.google.com/go/aiplatform/apiv1"},
		},
	}

	Resources = append(Resources, resources...)
}

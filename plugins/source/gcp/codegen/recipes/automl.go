package recipes

import (
	automl "cloud.google.com/go/automl/apiv1"
	automlpb "cloud.google.com/go/automl/apiv1/automlpb"
)


func init() {
	resources := []*Resource{
		{
			Service: "automl",
			SubService: "datasets",
			Struct: &automlpb.Dataset{},
			NewFunction: automl.NewClient,
			ListFunction: (&automl.Client{}).ListDatasets,
			RequestStruct: &automlpb.ListDatasetsRequest{},
			ResponseStruct: &automlpb.ListDatasetsResponse{},
			UnimplementedServer: &automlpb.UnimplementedAutoMlServer{},
			RegisterServer: automlpb.RegisterAutoMlServer,
			Template: "newapi_list",
			MockTemplate: "newapi_list_grpc_mock",
			ProtobufImport: "cloud.google.com/go/automl/apiv1/automlpb",
			MockImports: []string{"cloud.google.com/go/automl/apiv1"},
		},
		{
			Service: "automl",
			SubService: "models",
			Struct: &automlpb.Model{},
			NewFunction: automl.NewClient,
			ListFunction: (&automl.Client{}).ListModels,
			RequestStruct: &automlpb.ListModelsRequest{},
			ResponseStruct: &automlpb.ListModelsResponse{},
			UnimplementedServer: &automlpb.UnimplementedAutoMlServer{},
			RegisterServer: automlpb.RegisterAutoMlServer,
			Template: "newapi_list",
			MockTemplate: "newapi_list_grpc_mock",
			ProtobufImport: "cloud.google.com/go/automl/apiv1/automlpb",
			MockImports: []string{"cloud.google.com/go/automl/apiv1"},
		},
	}

	Resources = append(Resources, resources...)
}

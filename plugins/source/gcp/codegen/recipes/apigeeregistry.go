package recipes

import (
	apigeeregistry "cloud.google.com/go/apigeeregistry/apiv1"
	apigeeregistrypb "google.golang.org/genproto/googleapis/cloud/apigeeregistry/v1"
)


func init() {
	resources := []*Resource{
		{
			Service: "apigeeregistry",
			SubService: "apis",
			Struct: &apigeeregistrypb.Api{},
			NewFunction: apigeeregistry.NewRegistryClient,
			ListFunction: (&apigeeregistry.RegistryClient{}).ListApis,
			RequestStruct: &apigeeregistrypb.ListApisRequest{},
			ResponseStruct: &apigeeregistrypb.ListApisResponse{},
			UnimplementedServer: &apigeeregistrypb.UnimplementedRegistryServer{},
			RegisterServer: apigeeregistrypb.RegisterRegistryServer,
			Template: "newapi_list",
			MockTemplate: "newapi_list_grpc_mock",
			ProtobufImport: "google.golang.org/genproto/googleapis/cloud/apigeeregistry/v1",
			MockImports: []string{"cloud.google.com/go/apigeeregistry/apiv1"},
		},
	}

	Resources = append(Resources, resources...)
}

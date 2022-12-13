package recipes

import (
	binaryauthorization "cloud.google.com/go/binaryauthorization/apiv1"
	binaryauthorizationpb "cloud.google.com/go/binaryauthorization/apiv1/binaryauthorizationpb"
)


func init() {
	resources := []*Resource{
		{
			Service: "binaryauthorization",
			SubService: "attestors",
			Struct: &binaryauthorizationpb.Attestor{},
			NewFunction: binaryauthorization.NewBinauthzManagementClient,
			ListFunction: (&binaryauthorization.BinauthzManagementClient{}).ListAttestors,
			RequestStruct: &binaryauthorizationpb.ListAttestorsRequest{},
			ResponseStruct: &binaryauthorizationpb.ListAttestorsResponse{},
			UnimplementedServer: &binaryauthorizationpb.UnimplementedBinauthzManagementServiceV1Server{},
			RegisterServer: binaryauthorizationpb.RegisterBinauthzManagementServiceV1Server,
			Template: "newapi_list",
			MockTemplate: "newapi_list_grpc_mock",
			ProtobufImport: "cloud.google.com/go/binaryauthorization/apiv1/binaryauthorizationpb",
			MockImports: []string{"cloud.google.com/go/binaryauthorization/apiv1"},
		},
	}

	Resources = append(Resources, resources...)
}

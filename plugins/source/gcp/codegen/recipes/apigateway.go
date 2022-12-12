package recipes

import (
	apigateway "cloud.google.com/go/apigateway/apiv1"
	"cloud.google.com/go/apigateway/apiv1/apigatewaypb"
)


func init() {
	resources := []*Resource{
		{
			Service: "apigateway",
			SubService: "gateways",
			Struct: &apigatewaypb.Gateway{},
			NewFunction: apigateway.NewClient,
			ListFunction: (&apigateway.Client{}).ListGateways,
			RequestStruct: &apigatewaypb.ListGatewaysRequest{},
			ResponseStruct: &apigatewaypb.ListGatewaysResponse{},
			UnimplementedServer: &apigatewaypb.UnimplementedApiGatewayServiceServer{},
			RegisterServer: apigatewaypb.RegisterApiGatewayServiceServer,
			Template: "newapi_list",
			MockTemplate: "newapi_list_grpc_mock",
			ProtobufImport: "cloud.google.com/go/apigateway/apiv1/apigatewaypb",
			MockImports: []string{"cloud.google.com/go/apigateway/apiv1"},
		},
		{
			Service: "apigateway",
			SubService: "apis",
			Struct: &apigatewaypb.Api{},
			NewFunction: apigateway.NewClient,
			ListFunction: (&apigateway.Client{}).ListApis,
			RequestStruct: &apigatewaypb.ListApisRequest{},
			ResponseStruct: &apigatewaypb.ListApisResponse{},
			UnimplementedServer: &apigatewaypb.UnimplementedApiGatewayServiceServer{},
			RegisterServer: apigatewaypb.RegisterApiGatewayServiceServer,
			Template: "newapi_list",
			MockTemplate: "newapi_list_grpc_mock",
			ProtobufImport: "cloud.google.com/go/apigateway/apiv1/apigatewaypb",
			MockImports: []string{"cloud.google.com/go/apigateway/apiv1"},
		},
	}

	Resources = append(Resources, resources...)
}

package recipes

import (
	accesscontextmanager "cloud.google.com/go/accesscontextmanager/apiv1"
	accesscontextmanagerpb "cloud.google.com/go/accesscontextmanager/apiv1/accesscontextmanagerpb"
)


func init() {
	resources := []*Resource{
		{
			Service: "accesscontextmanager",
			SubService: "user_access_bindings",
			Struct: &accesscontextmanagerpb.GcpUserAccessBinding{},
			NewFunction: accesscontextmanager.NewClient,
			ListFunction: (&accesscontextmanager.Client{}).ListGcpUserAccessBindings,
			RequestStruct: &accesscontextmanagerpb.ListGcpUserAccessBindingsRequest{},
			ResponseStruct: &accesscontextmanagerpb.ListGcpUserAccessBindingsResponse{},
			UnimplementedServer: &accesscontextmanagerpb.UnimplementedAccessContextManagerServer{},
			RegisterServer: accesscontextmanagerpb.RegisterAccessContextManagerServer,
			Template: "newapi_list",
			MockTemplate: "newapi_list_grpc_mock",
			ProtobufImport: "cloud.google.com/go/accesscontextmanager/apiv1/accesscontextmanagerpb",
			MockImports: []string{"cloud.google.com/go/accesscontextmanager/apiv1"},
		},
		{
			Service: "accesscontextmanager",
			SubService: "access_levels",
			Struct: &accesscontextmanagerpb.AccessPolicy{},
			NewFunction: accesscontextmanager.NewClient,
			ListFunction: (&accesscontextmanager.Client{}).ListAccessPolicies,
			RequestStruct: &accesscontextmanagerpb.ListAccessPoliciesRequest{},
			ResponseStruct: &accesscontextmanagerpb.ListAccessPoliciesResponse{},
			UnimplementedServer: &accesscontextmanagerpb.UnimplementedAccessContextManagerServer{},
			RegisterServer: accesscontextmanagerpb.RegisterAccessContextManagerServer,
			Template: "newapi_list",
			MockTemplate: "newapi_list_grpc_mock",
			ProtobufImport: "cloud.google.com/go/accesscontextmanager/apiv1/accesscontextmanagerpb",
			MockImports: []string{"cloud.google.com/go/accesscontextmanager/apiv1"},
		},
	}

	Resources = append(Resources, resources...)
}

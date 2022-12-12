package recipes

import (
	appengine "cloud.google.com/go/appengine/apiv1"
	appenginepb "cloud.google.com/go/appengine/apiv1/appenginepb"
)


func init() {
	resources := []*Resource{
		{
			Service: "appengine",
			SubService: "services",
			Struct: &appenginepb.Service{},
			NewFunction: appengine.NewServicesClient,
			ListFunction: (&appengine.ServicesClient{}).ListServices,
			RequestStruct: &appenginepb.ListServicesRequest{},
			ResponseStruct: &appenginepb.ListServicesResponse{},
			UnimplementedServer: &appenginepb.UnimplementedServicesServer{},
			RegisterServer: appenginepb.RegisterServicesServer,
			Template: "newapi_list",
			MockTemplate: "newapi_list_grpc_mock",
			ProtobufImport: "cloud.google.com/go/appengine/apiv1/appenginepb",
			MockImports: []string{"cloud.google.com/go/appengine/apiv1"},
		},
	}

	Resources = append(Resources, resources...)
}

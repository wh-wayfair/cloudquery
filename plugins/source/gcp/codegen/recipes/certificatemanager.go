package recipes

import (
	certificatemanager "cloud.google.com/go/certificatemanager/apiv1"
	certificatemanagerpb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
)


func init() {
	resources := []*Resource{
		{
			Service: "certificatemanager",
			SubService: "certificates",
			Struct: &certificatemanagerpb.Certificate{},
			NewFunction: certificatemanager.NewClient,
			ListFunction: (&certificatemanager.Client{}).ListCertificates,
			RequestStruct: &certificatemanagerpb.ListCertificatesRequest{},
			ResponseStruct: &certificatemanagerpb.ListCertificatesResponse{},
			UnimplementedServer: &certificatemanagerpb.UnimplementedCertificateManagerServer{},
			RegisterServer: certificatemanagerpb.RegisterCertificateManagerServer,
			Template: "newapi_list",
			MockTemplate: "newapi_list_grpc_mock",
			ProtobufImport: "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb",
			MockImports: []string{"cloud.google.com/go/certificatemanager/apiv1"},
		},
		{
			Service: "certificatemanager",
			SubService: "dns_authorizations",
			Struct: &certificatemanagerpb.DnsAuthorization{},
			NewFunction: certificatemanager.NewClient,
			ListFunction: (&certificatemanager.Client{}).ListDnsAuthorizations,
			RequestStruct: &certificatemanagerpb.ListDnsAuthorizationsRequest{},
			ResponseStruct: &certificatemanagerpb.ListDnsAuthorizationsResponse{},
			UnimplementedServer: &certificatemanagerpb.UnimplementedCertificateManagerServer{},
			RegisterServer: certificatemanagerpb.RegisterCertificateManagerServer,
			Template: "newapi_list",
			MockTemplate: "newapi_list_grpc_mock",
			ProtobufImport: "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb",
			MockImports: []string{"cloud.google.com/go/certificatemanager/apiv1"},
		},
	}

	Resources = append(Resources, resources...)
}

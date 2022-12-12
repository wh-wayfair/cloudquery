package recipes

import (
	accessapproval "cloud.google.com/go/accessapproval/apiv1"
	accessapprovalpb "cloud.google.com/go/accessapproval/apiv1/accessapprovalpb"
)


func init() {
	resources := []*Resource{
		{
			Service: "accessapproval",
			SubService: "requests",
			Struct: &accessapprovalpb.ApprovalRequest{},
			NewFunction: accessapproval.NewClient,
			ListFunction: (&accessapproval.Client{}).ListApprovalRequests,
			RequestStruct: &accessapprovalpb.ListApprovalRequestsMessage{},
			ResponseStruct: &accessapprovalpb.ListApprovalRequestsResponse{},
			UnimplementedServer: &accessapprovalpb.UnimplementedAccessApprovalServer{},
			RegisterServer: accessapprovalpb.RegisterAccessApprovalServer,
			Template: "newapi_list",
			MockTemplate: "newapi_list_grpc_mock",
			ProtobufImport: "cloud.google.com/go/accessapproval/apiv1/accessapprovalpb",
			MockImports: []string{"cloud.google.com/go/accessapproval/apiv1"},
		},
	}

	Resources = append(Resources, resources...)
}

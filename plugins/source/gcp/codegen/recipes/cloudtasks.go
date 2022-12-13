package recipes

import (
	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	cloudtaskspb "cloud.google.com/go/cloudtasks/apiv2/cloudtaskspb"
)


func init() {
	resources := []*Resource{
		{
			Service: "cloudtasks",
			SubService: "queues",
			Struct: &cloudtaskspb.Queue{},
			NewFunction: cloudtasks.NewClient,
			ListFunction: (&cloudtasks.Client{}).ListQueues,
			RequestStruct: &cloudtaskspb.ListQueuesRequest{},
			ResponseStruct: &cloudtaskspb.ListQueuesResponse{},
			UnimplementedServer: &cloudtaskspb.UnimplementedCloudTasksServer{},
			RegisterServer: cloudtaskspb.RegisterCloudTasksServer,
			Template: "newapi_list",
			MockTemplate: "newapi_list_grpc_mock",
			ProtobufImport: "cloud.google.com/go/cloudtasks/apiv2/cloudtaskspb",
			MockImports: []string{"cloud.google.com/go/cloudtasks/apiv2"},
		},
	}

	Resources = append(Resources, resources...)
}

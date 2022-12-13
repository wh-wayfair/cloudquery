package recipes

import (
	batch "cloud.google.com/go/batch/apiv1"
	batchpb "cloud.google.com/go/batch/apiv1/batchpb"
)


func init() {
	resources := []*Resource{
		{
			Service: "batch",
			SubService: "jobs",
			Struct: &batchpb.Job{},
			NewFunction: batch.NewClient,
			ListFunction: (&batch.Client{}).ListJobs,
			RequestStruct: &batchpb.ListJobsRequest{},
			ResponseStruct: &batchpb.ListJobsResponse{},
			UnimplementedServer: &batchpb.UnimplementedBatchServiceServer{},
			RegisterServer: batchpb.RegisterBatchServiceServer,
			Template: "newapi_list",
			MockTemplate: "newapi_list_grpc_mock",
			ProtobufImport: "cloud.google.com/go/batch/apiv1/batchpb",
			MockImports: []string{"cloud.google.com/go/batch/apiv1"},
		},
	}

	Resources = append(Resources, resources...)
}

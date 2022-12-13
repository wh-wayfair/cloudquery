package recipes

import (
	cloudbuild "cloud.google.com/go/cloudbuild/apiv1"
	cloudbuildpb "google.golang.org/genproto/googleapis/devtools/cloudbuild/v1"
)


func init() {
	resources := []*Resource{
		{
			Service: "cloudbuild",
			SubService: "worker_pools",
			Struct: &cloudbuildpb.WorkerPool{},
			NewFunction: cloudbuild.NewClient,
			ListFunction: (&cloudbuild.Client{}).ListWorkerPools,
			RequestStruct: &cloudbuildpb.ListWorkerPoolsRequest{},
			ResponseStruct: &cloudbuildpb.ListWorkerPoolsResponse{},
			UnimplementedServer: &cloudbuildpb.UnimplementedCloudBuildServer{},
			RegisterServer: cloudbuildpb.RegisterCloudBuildServer,
			Template: "newapi_list",
			MockTemplate: "newapi_list_grpc_mock",
			ProtobufImport: "google.golang.org/genproto/googleapis/devtools/cloudbuild/v1",
			MockImports: []string{"cloud.google.com/go/cloudbuild/apiv1"},
			SkipFetch: true,
		},
	}

	Resources = append(Resources, resources...)
}

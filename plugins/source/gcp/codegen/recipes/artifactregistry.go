package recipes

import (
	artifactregistry "cloud.google.com/go/artifactregistry/apiv1"
	artifactregistrypb "cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb"
)


func init() {
	resources := []*Resource{
		{
			Service: "artifactregistry",
			SubService: "repostiories",
			Struct: &artifactregistrypb.Repository{},
			NewFunction: artifactregistry.NewClient,
			ListFunction: (&artifactregistry.Client{}).ListRepositories,
			RequestStruct: &artifactregistrypb.ListRepositoriesRequest{},
			ResponseStruct: &artifactregistrypb.ListRepositoriesResponse{},
			UnimplementedServer: &artifactregistrypb.UnimplementedArtifactRegistryServer{},
			RegisterServer: artifactregistrypb.RegisterArtifactRegistryServer,
			Template: "newapi_list",
			MockTemplate: "newapi_list_grpc_mock",
			ProtobufImport: "cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb",
			MockImports: []string{"cloud.google.com/go/artifactregistry/apiv1"},
		},
		{
			Service: "artifactregistry",
			SubService: "packages",
			Struct: &artifactregistrypb.Package{},
			NewFunction: artifactregistry.NewClient,
			ListFunction: (&artifactregistry.Client{}).ListPackages,
			RequestStruct: &artifactregistrypb.ListPackagesRequest{},
			ResponseStruct: &artifactregistrypb.ListPackagesResponse{},
			UnimplementedServer: &artifactregistrypb.UnimplementedArtifactRegistryServer{},
			RegisterServer: artifactregistrypb.RegisterArtifactRegistryServer,
			Template: "newapi_list",
			MockTemplate: "newapi_list_grpc_mock",
			ProtobufImport: "cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb",
			MockImports: []string{"cloud.google.com/go/artifactregistry/apiv1"},
		},
		{
			Service: "artifactregistry",
			SubService: "dockerimages",
			Struct: &artifactregistrypb.DockerImage{},
			NewFunction: artifactregistry.NewClient,
			ListFunction: (&artifactregistry.Client{}).ListDockerImages,
			RequestStruct: &artifactregistrypb.ListDockerImagesRequest{},
			ResponseStruct: &artifactregistrypb.ListDockerImagesResponse{},
			UnimplementedServer: &artifactregistrypb.UnimplementedArtifactRegistryServer{},
			RegisterServer: artifactregistrypb.RegisterArtifactRegistryServer,
			Template: "newapi_list",
			MockTemplate: "newapi_list_grpc_mock",
			ProtobufImport: "cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb",
			MockImports: []string{"cloud.google.com/go/artifactregistry/apiv1"},
		},
	}

	Resources = append(Resources, resources...)
}

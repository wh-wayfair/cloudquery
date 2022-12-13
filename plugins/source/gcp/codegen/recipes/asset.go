package recipes

import (
	asset "cloud.google.com/go/asset/apiv1"
	assetpb "cloud.google.com/go/asset/apiv1/assetpb"
)


func init() {
	resources := []*Resource{
		{
			Service: "asset",
			SubService: "assets",
			Struct: &assetpb.Asset{},
			NewFunction: asset.NewClient,
			ListFunction: (&asset.Client{}).ListAssets,
			RequestStruct: &assetpb.ListAssetsRequest{},
			ResponseStruct: &assetpb.ListAssetsResponse{},
			UnimplementedServer: &assetpb.UnimplementedAssetServiceServer{},
			RegisterServer: assetpb.RegisterAssetServiceServer,
			Template: "newapi_list",
			MockTemplate: "newapi_list_grpc_mock",
			ProtobufImport: "cloud.google.com/go/asset/apiv1/assetpb",
			MockImports: []string{"cloud.google.com/go/asset/apiv1"},
		},
		{
			Service: "asset",
			SubService: "feeds",
			Struct: &assetpb.Feed{},
			NewFunction: asset.NewClient,
			ListFunction: (&asset.Client{}).ListFeeds,
			RequestStruct: &assetpb.ListFeedsRequest{},
			ResponseStruct: &assetpb.ListFeedsResponse{},
			UnimplementedServer: &assetpb.UnimplementedAssetServiceServer{},
			RegisterServer: assetpb.RegisterAssetServiceServer,
			Template: "newapi_list",
			MockTemplate: "newapi_list_grpc_mock",
			// ProtobufImport: "cloud.google.com/go/asset/apiv1/assetpb",
			// MockImports: []string{"cloud.google.com/go/asset/apiv1"},
			SkipFetch: true,
			SkipMock: true,
		},
	}

	Resources = append(Resources, resources...)
}

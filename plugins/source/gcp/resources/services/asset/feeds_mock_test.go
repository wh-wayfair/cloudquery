package asset

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"

	pb "cloud.google.com/go/asset/apiv1/assetpb"
)

func createFeeds(gsrv *grpc.Server) error {
	fakeServer := &fakeFeedsServer{}
	pb.RegisterAssetServiceServer(gsrv, fakeServer)
	return nil
}

type fakeFeedsServer struct {
	pb.UnimplementedAssetServiceServer
}

func (f *fakeFeedsServer) ListFeeds(context.Context, *pb.ListFeedsRequest) (*pb.ListFeedsResponse, error) {
	resp := pb.ListFeedsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	return &resp, nil
}

func TestFeeds(t *testing.T) {
	client.MockTestGrpcHelper(t, Feeds(), createFeeds, client.TestOptions{})
}

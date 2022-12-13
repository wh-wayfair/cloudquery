package asset

import (
	"context"

	asset "cloud.google.com/go/asset/apiv1"
	pb "cloud.google.com/go/asset/apiv1/assetpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)
func fetchFeeds(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.ListFeedsRequest{}
	gcpClient, err := asset.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it, err := gcpClient.ListFeeds(ctx, req, c.CallOptions...)
	if err != nil {
		return err
	}
	res <- it.Feeds

	return nil
}
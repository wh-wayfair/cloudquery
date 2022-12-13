package cloudbuild

import (
	"context"

	cloudbuild "cloud.google.com/go/cloudbuild/apiv1"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	cloudbuildpb "google.golang.org/genproto/googleapis/devtools/cloudbuild/v1"
)

func fetchWorkerPools(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &cloudbuildpb.ListWorkerPoolsRequest{}
	gcpClient, err := cloudbuild.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	for {
		it, err := gcpClient.ListWorkerPools(ctx, req, c.CallOptions...)
		if err != nil {
			return err
		}
		res <- it.WorkerPools
		if it.NextPageToken == "" {
			break
		}
		req.PageToken = it.NextPageToken
	}
	return nil
}

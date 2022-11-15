package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchCloudfrontCachePolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config cloudfront.ListCachePoliciesInput
	c := meta.(*client.Client)
	s := c.Services()
	svc := s.Cloudfront
	for {
		response, err := svc.ListCachePolicies(ctx, nil)
		if err != nil {
			return err
		}

		if response.CachePolicyList != nil {
			res <- response.CachePolicyList.Items
		}

		if aws.ToString(response.CachePolicyList.NextMarker) == "" {
			break
		}
		config.Marker = response.CachePolicyList.NextMarker
	}
	return nil
}

func resolveCachePolicyARN(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.CachePolicySummary)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.CloudfrontService),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "cache-policy/" + aws.ToString(item.CachePolicy.Id),
	}
	return resource.Set(c.Name, a.String())
}

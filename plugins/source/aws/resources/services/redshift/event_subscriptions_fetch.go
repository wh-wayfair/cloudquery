package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchRedshiftEventSubscriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Redshift
	var params redshift.DescribeEventSubscriptionsInput
	params.MaxRecords = aws.Int32(100)
	for {
		result, err := svc.DescribeEventSubscriptions(ctx, &params)
		if err != nil {
			return err
		}
		res <- result.EventSubscriptionsList
		if aws.ToString(result.Marker) == "" {
			break
		}
		params.Marker = result.Marker
	}
	return nil
}

func resolveEventSubscriptionARN(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	sub := resource.Item.(types.EventSubscription)
	arn := arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.RedshiftService),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource: "eventsubscription:" + aws.ToString(sub.CustSubscriptionId),
	}
	return resource.Set(c.Name, arn.String())
}

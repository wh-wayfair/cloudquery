package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchDirectconnectGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config directconnect.DescribeDirectConnectGatewaysInput
	c := meta.(*client.Client)
	svc := c.Services().Directconnect
	for {
		output, err := svc.DescribeDirectConnectGateways(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.DirectConnectGateways
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func fetchDirectconnectGatewayAssociations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	gateway := parent.Item.(types.DirectConnectGateway)
	c := meta.(*client.Client)
	svc := c.Services().Directconnect
	config := directconnect.DescribeDirectConnectGatewayAssociationsInput{DirectConnectGatewayId: gateway.DirectConnectGatewayId}
	for {
		output, err := svc.DescribeDirectConnectGatewayAssociations(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.DirectConnectGatewayAssociations
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func fetchDirectconnectGatewayAttachments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	gateway := parent.Item.(types.DirectConnectGateway)
	c := meta.(*client.Client)
	svc := c.Services().Directconnect
	config := directconnect.DescribeDirectConnectGatewayAttachmentsInput{DirectConnectGatewayId: gateway.DirectConnectGatewayId}
	for {
		output, err := svc.DescribeDirectConnectGatewayAttachments(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.DirectConnectGatewayAttachments
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func resolveGatewayARN(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.DirectConnectGateway)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.DirectConnectService),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "dx-gateway/" + aws.ToString(item.DirectConnectGatewayId),
	}
	return resource.Set(c.Name, a.String())
}

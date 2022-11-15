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

func fetchRedshiftSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Redshift
	cluster := parent.Item.(types.Cluster)
	params := redshift.DescribeClusterSnapshotsInput{
		ClusterExists:     aws.Bool(true),
		ClusterIdentifier: cluster.ClusterIdentifier,
		MaxRecords:        aws.Int32(100),
	}
	for {
		result, err := svc.DescribeClusterSnapshots(ctx, &params)
		if err != nil {
			return err
		}
		res <- result.Snapshots
		if aws.ToString(result.Marker) == "" {
			break
		}
		params.Marker = result.Marker
	}
	return nil
}

func resolveSnapshotARN(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	snapshot := resource.Item.(types.Snapshot)
	arn := arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.RedshiftService),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource: "snapshot:" + aws.ToString(snapshot.ClusterIdentifier) + "/" + aws.ToString(snapshot.SnapshotIdentifier),
	}
	return resource.Set(c.Name, arn.String())
}


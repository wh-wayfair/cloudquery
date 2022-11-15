package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchWafregionalRateBasedRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Wafregional
	var params wafregional.ListRateBasedRulesInput
	for {
		result, err := svc.ListRateBasedRules(ctx, &params, func(o *wafregional.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		for _, item := range result.Rules {
			detail, err := svc.GetRateBasedRule(
				ctx,
				&wafregional.GetRateBasedRuleInput{RuleId: item.RuleId},
				func(o *wafregional.Options) {
					o.Region = cl.Region
				},
			)
			if err != nil {
				return err
			}
			if detail.Rule == nil {
				continue
			}
			res <- *detail.Rule
		}
		if aws.ToString(result.NextMarker) == "" {
			break
		}
		params.NextMarker = result.NextMarker
	}
	return nil
}
func resolveWafregionalRateBasedRuleArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.RateBasedRule)
	arn := arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.WAFRegional),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "ratebasedrule/" + aws.ToString(item.RuleId),
	}
	return resource.Set(c.Name, arn.String())
}
func resolveWafregionalRateBasedRuleTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Wafregional
	item := resource.Item.(types.RateBasedRule)
	arn := arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.WAFRegional),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "ratebasedrule/" + aws.ToString(item.RuleId),
	}
	params := wafregional.ListTagsForResourceInput{ResourceARN: aws.String(arn.String())}
	tags := make(map[string]string)
	for {
		result, err := svc.ListTagsForResource(ctx, &params)
		if err != nil {
			return err
		}
		if result.TagInfoForResource != nil {
			client.TagsIntoMap(result.TagInfoForResource.TagList, tags)
		}
		if aws.ToString(result.NextMarker) == "" {
			break
		}
		params.NextMarker = result.NextMarker
	}
	return resource.Set(c.Name, tags)
}

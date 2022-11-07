package client

import (
	"github.com/cloudquery/plugin-sdk/v2/plugins"
)

func (c *Client) Metrics() plugins.DestinationMetrics {
	return c.metrics
}

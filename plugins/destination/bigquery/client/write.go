package client

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/cloudquery/plugin-sdk/schema"
)

const (
	batchSize    = 1000
	writeTimeout = 5 * time.Minute
)

type worker struct {
	writeChan chan []interface{}
}

type item struct {
	cols map[string]bigquery.Value
}

func (i *item) Save() (map[string]bigquery.Value, string, error) {
	// we're not doing de-dup at the moment
	return i.cols, bigquery.NoDedupeID, nil
}


func (c *Client) WriteTableBatch(ctx context.Context, writeClient interface{}, table *schema.Table, resources [][]interface{}) error {
	bqClient := writeClient.(*bigquery.Client)
	inserter := bqClient.Dataset(c.pluginSpec.DatasetID).Table(table.Name).Inserter()
	inserter.IgnoreUnknownValues = true
	inserter.SkipInvalidRows = false
	batch := make([]*item, 0)
	for _, resource := range resources {
		c.logger.Debug().Msg("Got resource")
		saver := &item{
			cols: make(map[string]bigquery.Value, len(table.Columns)),
		}
		for i := range resource {
			if resource[i] == nil {
				// save some bandwidth by not sending nil values
				continue
			}
			saver.cols[table.Columns[i].Name] = resource[i]
		}
		batch = append(batch, saver)
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, writeTimeout)
	err := inserter.Put(timeoutCtx, batch)
	if err != nil {
		cancel()
		return fmt.Errorf("failed to put item into BigQuery table %s: %w", table.Name, err)
	}
	// release resources from timeout context if it finished early
	batch = nil
	cancel()

	return nil
}

func (c *Client) PreWrite(ctx context.Context, tables schema.Tables, sourceName string, syncTime time.Time) (interface{}, error) {
	bqClient, err := c.bqClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}
	return bqClient, nil
}

func (c *Client) PostWrite(ctx context.Context, writeClient interface{}, tables schema.Tables, sourceName string, syncTime time.Time) error {
	bqClient := writeClient.(*bigquery.Client)
	return bqClient.Close()
}

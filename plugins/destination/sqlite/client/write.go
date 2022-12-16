package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
)

func (c *Client) PreWrite(ctx context.Context, tables schema.Tables, sourceName string, syncTime time.Time) error {
	return nil
}

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, resources [][]interface{}) error {
	var sql string
	if c.spec.WriteMode == specs.WriteModeAppend {
		sql = c.insert(table)
	} else {
		sql = c.upsert(table)
	}
	for _, r := range resources {
		if _, err := c.db.Exec(sql, r...); err != nil {
			return fmt.Errorf("failed to execute '%s': %w", sql, err)
		}
	}
	return nil
}


func (*Client) insert(table *schema.Table) string {
	var sb strings.Builder
	sb.WriteString("insert into ")
	sb.WriteString(`"` + table.Name + `"`)
	sb.WriteString(" (")
	columns := table.Columns
	columnsLen := len(columns)
	for i, c := range columns {
		sb.WriteString(`"` + c.Name + `"`)
		if i < columnsLen-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString(") values (")
		}
	}
	for i := range columns {
		sb.WriteString(fmt.Sprintf("$%d", i+1))
		if i < columnsLen-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString(")")
		}
	}
	return sb.String()
}

func (*Client) upsert(table *schema.Table) string {
	var sb strings.Builder
	sb.WriteString("insert or replace into ")
	sb.WriteString(`"` + table.Name + `"`)
	sb.WriteString(" (")
	columns := table.Columns
	columnsLen := len(columns)
	for i, c := range columns {
		sb.WriteString(`"` + c.Name + `"`)
		if i < columnsLen-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString(") values (")
		}
	}
	for i := range columns {
		sb.WriteString(fmt.Sprintf("$%d", i+1))
		if i < columnsLen-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString(")")
		}
	}
	return sb.String()
}

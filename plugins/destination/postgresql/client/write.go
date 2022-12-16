package client

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

var cqStatusToPgStatus = map[schema.Status]pgtype.Status{
	schema.Null:      pgtype.Null,
	schema.Undefined: pgtype.Null,
	schema.Present:   pgtype.Present,
}

func (c *Client) PreWrite(ctx context.Context, tables schema.Tables, sourceName string, syncTime time.Time) (interface{}, error) {
	return nil, nil
}

func (c *Client) PostWrite(ctx context.Context, writeClient interface{}, tables schema.Tables, sourceName string, syncTime time.Time) error {
	return nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRandomString() string {
	b := make([]byte, 40)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (c *Client) WriteTableBatch(ctx context.Context, writeClient interface{}, table *schema.Table, resources [][]interface{}) error {
	// var sql string
	// batch := &pgx.Batch{}
	// if c.spec.WriteMode == specs.WriteModeAppend {
	// 	sql = c.insert(table)
	// } else {
	// 	sql = c.upsert(table)
	// }
	// for _, r := range resources {
	// 	batch.Queue(sql, r...)
	// }
	tx, err := c.conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	tempTable := generateRandomString()
	_, err = tx.Exec(ctx, fmt.Sprintf("create temp table \"%s\" (like %s including all) on commit drop", tempTable, table.Name))
	if err != nil {
		tx.Rollback(ctx)
		return fmt.Errorf("failed to create temporary table: %w", err)
	}

	_, err = tx.CopyFrom(ctx, pgx.Identifier{tempTable}, table.Columns.Names(), pgx.CopyFromRows(resources))
	if err != nil {
		tx.Rollback(ctx)
		return fmt.Errorf("failed to copy data to temporary table: %w", err)
	}
	sql := c.upsertX(table, tempTable)
	_, err = tx.Exec(ctx, sql)
	if err != nil {
		tx.Rollback(ctx)
		return fmt.Errorf("failed to upsert data: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (*Client) insert(table *schema.Table) string {
	var sb strings.Builder
	sb.WriteString("insert into ")
	sb.WriteString(pgx.Identifier{table.Name}.Sanitize())
	sb.WriteString(" (")
	columns := table.Columns
	columnsLen := len(columns)
	for i, c := range columns {
		sb.WriteString(pgx.Identifier{c.Name}.Sanitize())
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


func (c *Client) upsertX(table *schema.Table, tmpTable string) string {
	var sb strings.Builder

	sb.WriteString("insert into ")
	sb.WriteString(pgx.Identifier{table.Name}.Sanitize())
	sb.WriteString(" (")
	columns := table.Columns
	columnsLen := len(columns)
	for i, c := range columns {
		sb.WriteString(pgx.Identifier{c.Name}.Sanitize())
		if i < columnsLen-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString(") select * from ")
			sb.WriteString(pgx.Identifier{tmpTable}.Sanitize())
		}
	}

	constraintName := fmt.Sprintf("%s_cqpk", table.Name)
	sb.WriteString(" on conflict on constraint ")
	sb.WriteString(constraintName)
	sb.WriteString(" do update set ")
	for i, column := range columns {
		sb.WriteString(pgx.Identifier{column.Name}.Sanitize())
		sb.WriteString("=excluded.") // excluded references the new values
		sb.WriteString(pgx.Identifier{column.Name}.Sanitize())
		if i < columnsLen-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString("")
		}
	}

	return sb.String()
}

func (c *Client) upsert(table *schema.Table) string {
	var sb strings.Builder

	sb.WriteString(c.insert(table))
	columns := table.Columns
	columnsLen := len(columns)

	constraintName := fmt.Sprintf("%s_cqpk", table.Name)
	sb.WriteString(" on conflict on constraint ")
	sb.WriteString(constraintName)
	sb.WriteString(" do update set ")
	for i, column := range columns {
		sb.WriteString(pgx.Identifier{column.Name}.Sanitize())
		sb.WriteString("=excluded.") // excluded references the new values
		sb.WriteString(pgx.Identifier{column.Name}.Sanitize())
		if i < columnsLen-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString("")
		}
	}

	return sb.String()
}


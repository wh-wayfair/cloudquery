package queries

import (
	"encoding/json"
	"testing"

	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/stretchr/testify/require"
)

func TestCreateTable(t *testing.T) {
	query, err := CreateTable(schema.CQSchemaToArrow(&schema.Table{
		Name: "table_name",
		Columns: schema.ColumnList{
			schema.CqIDColumn,
			schema.CqParentIDColumn,
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
			schema.Column{
				Name:            "extra_col",
				Type:            schema.TypeFloat,
				CreationOptions: schema.ColumnCreationOptions{NotNull: true},
			},
			schema.Column{Name: "extra_inet_col", Type: schema.TypeInet},
			schema.Column{Name: "extra_inet_arr_col", Type: schema.TypeInetArray},
		},
	}), "", DefaultEngine())
	require.NoError(t, err)
	ensureContents(t, query, "create_table.sql")
}

func TestCreateTableNoOrderBy(t *testing.T) {
	query, err := CreateTable(schema.CQSchemaToArrow(&schema.Table{
		Name: "table_name",
		Columns: schema.ColumnList{
			schema.Column{Name: "extra_col", Type: schema.TypeFloat},
			schema.Column{Name: "extra_inet_col", Type: schema.TypeInet},
			schema.Column{Name: "extra_inet_arr_col", Type: schema.TypeInetArray},
		},
	}), "", DefaultEngine())
	require.NoError(t, err)
	ensureContents(t, query, "create_table_no_order_by.sql")
}

func TestCreateTableOnCluster(t *testing.T) {
	query, err := CreateTable(schema.CQSchemaToArrow(&schema.Table{
		Name: "table_name",
		Columns: schema.ColumnList{
			schema.CqIDColumn,
			schema.CqParentIDColumn,
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
			schema.Column{
				Name:            "extra_col",
				Type:            schema.TypeFloat,
				CreationOptions: schema.ColumnCreationOptions{NotNull: true},
			},
			schema.Column{Name: "extra_inet_col", Type: schema.TypeInet},
			schema.Column{Name: "extra_inet_arr_col", Type: schema.TypeInetArray},
		},
	}), "my_cluster", DefaultEngine())
	require.NoError(t, err)
	ensureContents(t, query, "create_table_cluster.sql")
}

func TestCreateTableWithEngine(t *testing.T) {
	query, err := CreateTable(schema.CQSchemaToArrow(&schema.Table{
		Name: "table_name",
		Columns: schema.ColumnList{
			schema.CqIDColumn,
			schema.CqParentIDColumn,
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
			schema.Column{
				Name:            "extra_col",
				Type:            schema.TypeFloat,
				CreationOptions: schema.ColumnCreationOptions{NotNull: true},
			},
			schema.Column{Name: "extra_inet_col", Type: schema.TypeInet},
			schema.Column{Name: "extra_inet_arr_col", Type: schema.TypeInetArray},
		},
	}), "", &Engine{
		Name:       "ReplicatedMergeTree",
		Parameters: []any{"a", "b", 1, int32(2), int64(3), float32(1.2), float64(3.4), json.Number("327"), false, true},
	})
	require.NoError(t, err)
	ensureContents(t, query, "create_table_engine.sql")
}

func TestDropTable(t *testing.T) {
	query := DropTable(schema.CQSchemaToArrow(&schema.Table{Name: "table_name"}), "")

	ensureContents(t, query, "drop_table.sql")
}

func TestDropTableOnCLuster(t *testing.T) {
	query := DropTable(schema.CQSchemaToArrow(&schema.Table{Name: "table_name"}), "my_cluster")

	ensureContents(t, query, "drop_table_cluster.sql")
}

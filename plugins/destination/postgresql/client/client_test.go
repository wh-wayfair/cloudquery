package client

import (
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/v2/plugins"
	"github.com/cloudquery/plugin-sdk/v2/specs"
)

func getTestConnection() string {
	testConn := os.Getenv("CQ_DEST_PG_TEST_CONN")
	if testConn == "" {
		return "postgresql://postgres:pass@localhost:5432/postgres?sslmode=disable"
	}
	return testConn
}


func TestPgPlugin(t *testing.T) {
	p := plugins.NewDestinationPlugin("postgresql", "development", New)
	plugins.DestinationPluginTestSuite(t, p, specs.Destination{
		WriteMode: specs.WriteModeAppend,
		Spec: Spec{
			ConnectionString: getTestConnection(),
			PgxLogLevel:      LogLevelTrace,
		},
	})
}

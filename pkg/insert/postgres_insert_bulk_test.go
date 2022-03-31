package insert

import (
	"github.com/SbstnErhrdt/env"
	"github.com/max-planck-innovation-competition/go-patstat/connections"
	"testing"
)

func TestBulkReadFile(t *testing.T) {
	env.LoadEnvFiles("../../.env")
	connections.ConnectToSQL()
	BulkReadFile("/ingest/tls206_part01.csv")
	BulkReadFile("/ingest/tls206_part02.csv")
}

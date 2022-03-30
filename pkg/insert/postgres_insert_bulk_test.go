package insert

import (
	"github.com/SbstnErhrdt/env"
	"github.com/max-planck-innovation-competition/go-patstat/connections"
	"testing"
)

func TestBulkReadFile(t *testing.T) {
	env.LoadEnvFiles("../../.env")
	connections.ConnectToSQL()
	BulkReadFile("/ingest/tls201_part03.csv")
}

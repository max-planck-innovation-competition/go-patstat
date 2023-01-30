package insert

import (
	"github.com/max-planck-innovation-competition/go-patstat/connections"
	"github.com/max-planck-innovation-competition/go-patstat/pkg/models"
	"testing"
)

func TestBulkReadFile(t *testing.T) {
	connections.ConnectToSQL()
	connections.SQLClient.Config.DisableForeignKeyConstraintWhenMigrating = true
	err := connections.SQLClient.AutoMigrate(&models.Tls224ApplnCpc{})
	if err != nil {
		t.Fatal(err)
	}
	BulkReadFile("/var/lib/postgresql/data/ingest/tls224_part01.csv")
	BulkReadFile("/var/lib/postgresql/data/ingest/tls224_part02.csv")
	BulkReadFile("/var/lib/postgresql/data/ingest/tls224_part03.csv")
}

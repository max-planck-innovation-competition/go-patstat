package insert

import (
	"github.com/max-planck-innovation-competition/go-patstat/connections"
	"github.com/max-planck-innovation-competition/go-patstat/pkg/models"
	"testing"
)

func TestBulkReadFile(t *testing.T) {
	connections.ConnectToSQL()
	connections.SQLClient.Config.DisableForeignKeyConstraintWhenMigrating = true
	err := connections.SQLClient.AutoMigrate(&models.Tls202ApplnTitle{})
	if err != nil {
		t.Fatal(err)
	}
	err = connections.SQLClient.AutoMigrate(&models.Tls203ApplnAbstr{})
	if err != nil {
		t.Fatal(err)
	}
	BulkInsertFile("/var/lib/postgresql/data/ingest/tls224_part01.csv")
	// BulkInsertFile("/var/lib/postgresql/data/ingest/tls224_part02.csv")
	// BulkInsertFile("/var/lib/postgresql/data/ingest/tls224_part03.csv")
}

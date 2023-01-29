package insert

import (
	"github.com/max-planck-innovation-competition/go-patstat/connections"
	"github.com/max-planck-innovation-competition/go-patstat/pkg/models"
	"testing"
)

func TestBulkReadFile(t *testing.T) {
	connections.ConnectToSQL()
	connections.SQLClient.Config.DisableForeignKeyConstraintWhenMigrating = true
	err := connections.SQLClient.AutoMigrate(&models.Tls201Appln{})
	if err != nil {
		t.Fatal(err)
	}
	BulkReadFile("ingest/tls201_part01.csv")
	BulkReadFile("ingest/tls201_part02.csv")
	BulkReadFile("ingest/tls201_part03.csv")
}

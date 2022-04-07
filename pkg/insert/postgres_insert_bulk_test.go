package insert

import (
	"github.com/SbstnErhrdt/env"
	"github.com/max-planck-innovation-competition/go-patstat/connections"
	"testing"
)

func TestBulkReadFile(t *testing.T) {
	env.LoadEnvFiles("../../.env")
	connections.ConnectToSQL()

	/*
		err := connections.SQLClient.Migrator().DropTable(&models.Tls227PersPubln{})
		if err != nil {
			panic(err)
		}
		err = connections.SQLClient.Migrator().CreateTable(&models.Tls227PersPubln{})
		if err != nil {
			panic(err)
		}

	*/
	BulkReadFile("/ingest/tls211_part01.csv")
	BulkReadFile("/ingest/tls211_part02.csv")
}

/*

 */

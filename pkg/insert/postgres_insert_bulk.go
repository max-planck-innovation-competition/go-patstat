package insert

import (
	"github.com/max-planck-innovation-competition/go-patstat/connections"
	log "github.com/sirupsen/logrus"
)

// BulkReadFile reads a file and creates the db entries in batches
func BulkReadFile(filePath string) {
	logger := log.WithFields(log.Fields{"file": filePath})
	logger.Info("reading file")
	// get the table name of the file
	tableName, err := File2Table(filePath)
	if err != nil {
		logger.Fatal(err)
		return
	}
	logger.WithField("table_name", tableName).Info("sql table name")
	// copy csv
	err = BulkInsert(tableName, filePath)
	if err != nil {
		logger.Fatal(err)
		return
	}
	logger.Info("finished copying csv")
	return
}

func BulkInsert(tableName string, filePath string) (err error) {
	logger := log.WithFields(log.Fields{"file": filePath, "table": tableName})
	logger.Info("start copying csv")
	err = connections.SQLClient.Exec(`COPY ` + tableName + ` FROM '` + filePath + `' WITH CSV HEADER;`).Error
	if err != nil {
		logger.Fatal(err)
		return
	}
	logger.Info("successfully finished")
	return
}

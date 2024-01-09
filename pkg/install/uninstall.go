package install

import (
	"fmt"
	"github.com/max-planck-innovation-competition/go-patstat/connections"
	log "github.com/sirupsen/logrus"
)

// Uninstall deletes all migrateTables and data from the database
func Uninstall() (err error) {
	log.Println("Start De installation")
	connections.ConnectToSQL()

	// check if db exists
	exists, err := checkIfDBExists()
	if err != nil {
		log.Panic(err)
		return
	}

	if exists {

		stmt := fmt.Sprintf("DROP DATABASE '%s' WITH (FORCE);", PatstatDatabaseName)
		if rs := connections.SQLClient.Exec(stmt); rs.Error != nil {
			log.Panic(rs.Error)
			err = rs.Error
			return
		}
		return
	}
	return nil
}

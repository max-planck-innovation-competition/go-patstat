package install

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"patstat/connections"
)

// Uninstall deletes all tables and data from the database
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
		stmt := fmt.Sprintf("DROP DATABASE %s WITH (FORCE);", dbName)
		if rs := connections.SQLClient.Exec(stmt); rs.Error != nil {
			log.Panic(rs.Error)
			err = rs.Error
			return
		}
		return
	}
	return nil
}

package install

import (
	"fmt"
	"github.com/max-planck-innovation-competition/go-patstat/connections"
	"github.com/rs/zerolog/log"
)

// Uninstall deletes all migrateTables and data from the database
func Uninstall(dbName string) {
	log.Info().Msg("Start De installation")
	connections.ConnectToSQL()

	// check if db exists
	exists, err := checkIfDBExists(dbName)
	if err != nil {
		log.Panic().Msg(err.Error())
		return
	}

	if exists {
		stmt := fmt.Sprintf("DROP DATABASE %s WITH (FORCE);", dbName)
		if rs := connections.SQLClient.Exec(stmt); rs.Error != nil {
			log.Panic().Msg(rs.Error.Error())
			err = rs.Error
			return
		}
		return
	}
}

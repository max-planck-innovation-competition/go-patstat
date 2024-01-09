package install

import (
	"fmt"
	"github.com/max-planck-innovation-competition/go-patstat/connections"
	"github.com/max-planck-innovation-competition/go-patstat/pkg/models"
	"github.com/rs/zerolog/log"
)

func checkIfDBExists(dbName string) (exists bool, err error) {
	// check if db exists
	stmt := fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s';", dbName)
	rs := connections.SQLClient.Raw(stmt)
	if rs.Error != nil {
		log.Err(rs.Error).Msg("error while checking if db exists")
		return
	}

	// if not create it
	var rec = make(map[string]interface{})
	if rs.Error != nil {
		err = rs.Error
		return
	}

	// if exists return true
	if rs.Find(rec); len(rec) == 0 {
		return false, nil
	}

	// if not return false
	return true, nil
}

// CreateDatabase the database
func CreateDatabase(dbName string) {
	log.Info().Msg("Start Installation")
	connections.ConnectToSQL()

	// check if db exists
	log.Info().Msg("check if database exists before installation")
	exists, err := checkIfDBExists(dbName)
	if err != nil {
		log.Err(err).Msg("error while checking if db exists")
		return
	}
	log.Info().Msg("database exists: " + fmt.Sprintf("%t", exists))

	// if not create it
	if !exists {
		log.Info().Str("dbName", dbName).Msg("creating new db")
		stmt := fmt.Sprintf("CREATE DATABASE %s;", dbName)
		if rs := connections.SQLClient.Exec(stmt); rs.Error != nil {
			log.Err(err).Str("dbName", dbName).Msg("error while creating new db")
			return
		}
		log.Info().Str("dbName", dbName).Msg("successfully created new db")
	}
}

func CreateTableConstraints() {
	log.Info().Msg("adding table constraints")
	// add constraints to the migrateTables
	connections.SQLClient.Config.DisableForeignKeyConstraintWhenMigrating = false
	err := migrateTables()
	if err != nil {
		log.Panic().AnErr("err", err).Msg("error while adding the table constraints migrateTables")
		return
	}
	return
}

func CreateTables() {
	// create the migrateTables
	connections.SQLClient.Config.DisableForeignKeyConstraintWhenMigrating = true
	err := migrateTables()
	if err != nil {
		log.Panic().AnErr("err", err).Msg("error while creating migrateTables")
		return
	}
	return
}

var Models = []interface{}{
	&models.Tls904Nuts{},
	&models.Tls902IpcNace2{},
	&models.Tls901TechnFieldIpc{},
	&models.Tls803LegalEventCode{},
	&models.Tls801Country{},
	&models.Tls231InpadocLegalEvent{},
	&models.Tls230ApplnTechnField{},
	&models.Tls229ApplnNace2{},
	&models.Tls228DocdbFamCitn{},
	&models.Tls227PersPubln{},
	&models.Tls226PersonOrig{},
	&models.Tls225DocdbFamCpc{},
	&models.Tls224ApplnCpc{},
	&models.Tls222ApplnJpClass{},
	&models.Tls216ApplnContn{},
	&models.Tls215CitnCateg{},
	&models.Tls214NplPubln{},
	&models.Tls212Citation{},
	&models.Tls211PatPubln{},
	&models.Tls210ApplnNCls{},
	&models.Tls209ApplnIpc{},
	&models.Tls207PersAppln{},
	&models.Tls206Person{},
	&models.Tls205TechRel{},
	&models.Tls204ApplnPrior{},
	&models.Tls203ApplnAbstr{},
	&models.Tls202ApplnTitle{},
	&models.Tls201Appln{},
}

func migrateTables() (err error) {
	// all migrateTables
	err = connections.SQLClient.AutoMigrate(Models...)
	return
}

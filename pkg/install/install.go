package install

import (
	"fmt"
	"github.com/max-planck-innovation-competition/go-patstat/connections"
	"github.com/max-planck-innovation-competition/go-patstat/pkg/models"
	log "github.com/sirupsen/logrus"
	"log/slog"
)

func checkIfDBExists() (exists bool, err error) {
	// check if db exists
	stmt := fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s';", PatstatDatabaseName)
	rs := connections.SQLClient.Raw(stmt)
	if rs.Error != nil {
		log.Panic(rs.Error)
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

// Install creates all migrateTables in the database
func Install() (err error) {
	log.Println("Start Installation")
	connections.ConnectToSQL()

	// check if db exists
	exists, err := checkIfDBExists()
	if err != nil {
		log.Panic(err)
		return
	}

	// if not create it
	if !exists {
		stmt := fmt.Sprintf("CREATE DATABASE %s;", PatstatDatabaseName)
		if rs := connections.SQLClient.Exec(stmt); rs.Error != nil {
			log.Panic(rs.Error)
			return
		}
	}

	// add constraints to the migrateTables
	connections.SQLClient.Config.DisableForeignKeyConstraintWhenMigrating = false
	err = migrateTables()
	if err != nil {
		log.Panic(err)
		return
	}

	// close db connection
	sql, err := connections.SQLClient.DB()
	defer func() {
		_ = sql.Close()
	}()
	if err != nil {
		log.Panic(err)
		return
	}
	return
}

func CreateTables() (err error) {
	// create the migrateTables
	connections.SQLClient.Config.DisableForeignKeyConstraintWhenMigrating = true
	err = migrateTables()
	if err != nil {
		slog.Error("error creating tables", err)
		panic(err)
		return
	}
	slog.Info("tables created")
	return
}

func migrateTables() (err error) {
	// all migrateTables
	err = connections.SQLClient.AutoMigrate(
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
	)
	return
}

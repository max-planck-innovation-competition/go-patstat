package install

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"patstat/connections"
	"patstat/pkg/models"
)

func checkIfDBExists() (exists bool, err error) {
	// check if db exists
	stmt := fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s';", dbName)
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

// Install creates all tables in the database
func Install() {
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
		stmt := fmt.Sprintf("CREATE DATABASE %s;", dbName)
		if rs := connections.SQLClient.Exec(stmt); rs.Error != nil {
			log.Panic(rs.Error)
			return
		}
	}

	// create all the tables
	err = connections.SQLClient.AutoMigrate(
		&models.Tls201Appln{},
		&models.Tls202ApplnTitle{},
		&models.Tls203ApplnAbstr{},
		&models.Tls204ApplnPrior{},
		&models.Tls205TechRel{},
		&models.Tls206Person{},
		&models.Tls207PersAppln{},
		&models.Tls209ApplnIpc{},
		&models.Tls210ApplnNCls{},
		&models.Tls211PatPubln{},
		&models.Tls212Citation{},
		&models.Tls214NplPubln{},
		&models.Tls215CitnCateg{},
		&models.Tls216ApplnContn{},
		&models.Tls222ApplnJpClass{},
		&models.Tls224ApplnCpc{},
		&models.Tls225DocdbFamCpc{},
		&models.Tls226PersonOrig{},
		&models.Tls227PersPubln{},
		&models.Tls228DocdbFamCitn{},
		&models.Tls229ApplnNace2{},
		&models.Tls230ApplnTechnField{},
		&models.Tls231InpadocLegalEvent{},
		&models.Tls801Country{},
		&models.Tls803LegalEventCode{},
		&models.Tls901TechnFieldIpc{},
		&models.Tls902IpcNace2{},
		&models.Tls904Nuts{},
	)
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
}

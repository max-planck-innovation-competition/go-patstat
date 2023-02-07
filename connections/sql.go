package connections

import (
	"github.com/SbstnErhrdt/go-gorm-all-sql/pkg/sql"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strings"
	"time"
)

// SQLClient is a gorm sql client
// must be initialized with ConnectToSQL()
// can be closed with CloseSqlConnection()
var SQLClient *gorm.DB

// ConnectToSQL inits a sql client
func ConnectToSQL() {
	log.Println("Try to connect to sql database")
	client, err := sql.ConnectToDatabase()
	if err != nil {
		log.Error("Failed to connected to sql database")
		log.Panic(err)
		return
	}
	SQLClient = client
	db, err := SQLClient.DB()
	if err != nil {
		log.Error("Failed to get to sql db client")
		log.Panic(err)
		return
	}
	db.SetConnMaxLifetime(time.Minute * 20) // needs to be long enough for the whole request
	db.SetConnMaxIdleTime(time.Minute * 20) // needs to be long enough for the whole request
	log.Info("Successfully connected to sql database")
	if strings.ToUpper(os.Getenv("ENVIRONMENT")) != "PRODUCTION" {
		SQLClient.Config.Logger = logger.Default.LogMode(logger.Info)
	}
	return
}

// CloseSqlConnection closes a sql client
func CloseSqlConnection() {
	log.Println("Try to close connection to sql database")
	db, err := SQLClient.DB()
	if err != nil {
		log.Error("Failed to get sql db")
		log.Panic(err)
		return
	}
	err = db.Close()
	if err != nil {
		log.Error("Failed to close connection")
		log.Panic(err)
		return
	}
	log.Info("Successfully closed connection")
}

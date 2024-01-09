package connections

import (
	"github.com/SbstnErhrdt/go-gorm-all-sql/pkg/sql"
	"github.com/rs/zerolog/log"
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
	log.Info().Msg("Try to connect to sql database")
	client, err := sql.ConnectToDatabase()
	if err != nil {
		log.Err(err).Msg("Failed to connected to sql database")
		return
	}
	SQLClient = client
	db, err := SQLClient.DB()
	if err != nil {
		log.Err(err).Msg("Failed to get to sql db client")
		return
	}
	db.SetConnMaxLifetime(time.Minute * 40) // needs to be long enough for the whole request
	db.SetConnMaxIdleTime(time.Minute * 40) // needs to be long enough for the whole request
	log.Info().Msg("Successfully connected to sql database")
	if strings.ToUpper(os.Getenv("ENVIRONMENT")) != "PRODUCTION" {
		SQLClient.Config.Logger = logger.Default.LogMode(logger.Info)
	}
	return
}

// CloseSqlConnection closes a sql client
func CloseSqlConnection() {
	log.Info().Msg("Try to close connection to sql database")
	db, err := SQLClient.DB()
	if err != nil {
		log.Err(err).Msg("Failed to get sql db")
		return
	}
	err = db.Close()
	if err != nil {
		log.Err(err).Msg("Failed to close connection")
		return
	}
	log.Info().Msg("Successfully closed connection")
}

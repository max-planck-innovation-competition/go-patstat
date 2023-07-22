package main

import (
	"flag"
	"github.com/SbstnErhrdt/env"
	"github.com/max-planck-innovation-competition/go-patstat/connections"
	"github.com/max-planck-innovation-competition/go-patstat/pkg/insert"
	"github.com/max-planck-innovation-competition/go-patstat/pkg/install"
	"github.com/rs/zerolog/log"
	"os"
)

func init() {
	env.LoadEnvFiles()
}

// DefaultDbName is the default database name
var DefaultDbName = "patstat"

// DefaultDirectoryPath is the default directory path of the data
var DefaultDirectoryPath = "."

func main() {
	log.Debug().
		Msg("Start Go Patstat script")

	// check db connection
	connections.ConnectToSQL()

	// get first arg
	args := os.Args
	if len(args) < 2 {
		return
	}
	// get the first arg
	arg := args[1]
	// switch on the arg
	switch arg {
	case "install":
		installMode()
		break
	case "uninstall":
		uninstallMode()
		break
	case "insert":
		insertMode()
	default:
		log.Info().Msg("No mode selected")
	}

	// close db connection
	connections.CloseSqlConnection()
	log.Info().Msg("End Go Patstat script")
}

func installMode() {
	log.Info().Msg("Start Installation")
	// get flag db name
	flag.StringVar(&DefaultDbName, "db", "patstat", "database name")
	flag.StringVar(&DefaultDirectoryPath, "directory", ".", "the directory path to the data")
	// install
	install.CreateDatabase(DefaultDbName)
	install.CreateTables()
	// insert data
	insert.ProcessDirectory(DefaultDirectoryPath)
	// create indexes
	install.CreateTableConstraints()
	log.Info().Msg("End Installation")
}

func insertMode() {
	log.Info().Msg("Start Insertion")
	// get flag db name
	flag.StringVar(&DefaultDbName, "db", "patstat", "database name")
	flag.StringVar(&DefaultDirectoryPath, "directory", ".", "the directory path to the data")
	// insert data
	insert.ProcessDirectory(DefaultDirectoryPath)
	log.Info().Msg("End Insertion")
}

func uninstallMode() {
	log.Info().Msg("Start uninstall")
	flag.StringVar(&DefaultDbName, "db", "patstat", "database name")
	install.Uninstall(DefaultDbName)
	log.Info().Msg("End uninstall")
}

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
var DefaultDbName = "patstat_2023_spring"

// DefaultDirectoryPath is the default directory path of the data
var DefaultDirectoryPath = "."

var DefaultPostgresDirectoryPath = "/var/lib/postgresql/data/ingest/"

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
	// get flag db name
	flag.StringVar(&DefaultDbName, "db", "patstat", "database name")
	flag.StringVar(&DefaultDirectoryPath, "directory", ".", "the directory path to the data")
	flag.StringVar(&DefaultPostgresDirectoryPath, "postgres-directory", "/var/lib/postgresql/data/ingest/", "the directory path to the data")
	log.Info().
		Str("db", DefaultDbName).
		Str("directory", DefaultDirectoryPath).
		Str("postgres-directory", DefaultPostgresDirectoryPath).
		Msg("Start Installation")

	// install
	install.CreateDatabase(DefaultDbName)
	install.CreateTables()
	// insert data
	insert.ProcessDirectory(DefaultDirectoryPath, DefaultPostgresDirectoryPath)
	// create indexes
	install.CreateTableConstraints()
	log.Info().Msg("End Installation")
}

func insertMode() {
	// get flag db name
	flag.StringVar(&DefaultDbName, "db", "patstat_2023_spring", "database name")
	flag.StringVar(&DefaultDirectoryPath, "directory", ".", "the directory path to the data")
	flag.StringVar(&DefaultPostgresDirectoryPath, "postgres-directory", "/var/lib/postgresql/data/ingest/", "the directory path to the data")
	flag.Parse()
	log.Info().
		Str("db", DefaultDbName).
		Str("directory", DefaultDirectoryPath).
		Str("postgres-directory", DefaultPostgresDirectoryPath).
		Msg("Start Insertion")
	// insert data
	insert.ProcessDirectory(DefaultDirectoryPath, DefaultPostgresDirectoryPath)
	log.Info().Msg("End Insertion")
}

func uninstallMode() {
	log.Info().
		Str("db", DefaultDbName).
		Msg("Start uninstall")
	flag.StringVar(&DefaultDbName, "db", "patstat", "database name")
	flag.Parse()
	install.Uninstall(DefaultDbName)
	log.Info().Msg("End uninstall")
}

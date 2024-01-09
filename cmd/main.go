package main

import (
	"flag"
	"github.com/SbstnErhrdt/env"
	"github.com/max-planck-innovation-competition/go-patstat/connections"
	"github.com/max-planck-innovation-competition/go-patstat/pkg/insert"
	"github.com/max-planck-innovation-competition/go-patstat/pkg/install"
	"github.com/rs/zerolog/log"
)

func init() {
	env.LoadEnvFiles()
}

func main() {
	log.Debug().
		Msg("Start Go Patstat script")

	// load env files
	env.LoadEnvFiles()

	// check db connection
	connections.ConnectToSQL()

	mode := flag.String("mode", "install", "the mode to run the script in (install, uninstall, insert)")
	DefaultDbName := flag.String("db", "patstat", "database name")
	DefaultDirectoryPath := flag.String("directory", ".", "the directory path to the data")
	DefaultPostgresDirectoryPath := flag.String("postgres-directory", "/ingest", "the directory path to the data")
	flag.Parse()

	// switch on the arg
	switch *mode {
	case "install":
		installMode(*DefaultDbName, *DefaultDirectoryPath, *DefaultPostgresDirectoryPath)
		break
	case "uninstall":
		uninstallMode(*DefaultDbName, *DefaultDirectoryPath, *DefaultPostgresDirectoryPath)
		break
	case "insert":
		insertMode(*DefaultDbName, *DefaultDirectoryPath, *DefaultPostgresDirectoryPath)
	default:
		log.Info().Msg("No mode selected")
	}

	// close db connection
	connections.CloseSqlConnection()
	log.Info().Msg("End Go Patstat script")
}

func installMode(DefaultDbName, DefaultDirectoryPath, DefaultPostgresDirectoryPath string) {
	// get flag db name
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

func insertMode(DefaultDbName, DefaultDirectoryPath, DefaultPostgresDirectoryPath string) {
	// get flag db name
	log.Info().
		Str("db", DefaultDbName).
		Str("directory", DefaultDirectoryPath).
		Str("postgres-directory", DefaultPostgresDirectoryPath).
		Msg("Start Insertion")
	// insert data
	insert.ProcessDirectory(DefaultDirectoryPath, DefaultPostgresDirectoryPath)
	log.Info().Msg("End Insertion")
}

func uninstallMode(DefaultDbName, DefaultDirectoryPath, DefaultPostgresDirectoryPath string) {
	log.Info().
		Str("db", DefaultDbName).
		Msg("Start uninstall")
	flag.StringVar(&DefaultDbName, "db", "patstat", "database name")
	flag.Parse()
	err := install.Uninstall(DefaultDbName)
	if err != nil {
		log.Err(err).Msg("error while uninstalling")
		return
	}
	log.Info().Msg("End uninstall")
}

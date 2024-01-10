package insert

import (
	"github.com/max-planck-innovation-competition/go-patstat/connections"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
)

var ErrNoFilesFound = errors.New("no files found")

func BulkReadDirectory(dirPath, postGresPath string) (files []string, err error) {
	files = []string{}
	// check if last character is a slash
	if dirPath[len(dirPath)-1] != '/' {
		dirPath = dirPath + "/"
	}
	// walk over the directory
	err = filepath.Walk(dirPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Err(err).Msg("error while walking over directory")
				return err
			}
			// check if the path is the directory itself
			if path == dirPath {
				return nil
			}
			// skip child directories
			if path != "." && info.IsDir() {
				log.Info().Str("dir", path).Msg("skipping directory")
				return filepath.SkipDir
			}
			// check if the file is a csv
			if filepath.Ext(path) == ".csv" {
				log.Info().Str("file", path).Msg("found file")
				// remove the directory path
				path = path[len(dirPath):]
				// add path prefix
				path = filepath.Join(postGresPath, path)
				// append to files
				files = append(files, path)
			} else {
				log.Info().Str("file", path).Msg("skipping file")
			}
			return nil
		})
	if err != nil {
		log.Err(err).Msg("error while walking over directory")
	}
	if len(files) == 0 {
		err = ErrNoFilesFound
		log.Err(err).Str("path", dirPath).Msg("no files found")
	}
	return
}

// BulkInsertFile reads a file and creates the db entries in batches
func BulkInsertFile(filePath string) {
	logger := log.With().Str("file", filePath).Logger()
	logger.Info().Msg("reading file")
	// get the table name of the file
	tableName, err := File2Table(filePath)
	if err != nil {
		logger.Fatal().Msg(err.Error())
		return
	}
	logger.Info().Str("table_name", tableName).Msg("sql table name")
	// copy csv
	err = BulkInsert(tableName, filePath)
	if err != nil {
		logger.Fatal().Msg(err.Error())
		return
	}
	logger.Info().Msg("finished copying csv")
	return
}

func BulkInsert(tableName string, filePath string) (err error) {
	logger := log.With().Str("file", filePath).Str("table", tableName).Logger()
	logger.Info().Msg("start copying csv")
	// use the COPY command of Postgres to copy the csv into the table
	err = connections.SQLClient.Exec(`COPY ` + tableName + ` FROM '` + filePath + `' WITH CSV HEADER;`).Error
	if err != nil {
		logger.Fatal().Msg(err.Error())
		return
	}
	logger.Info().Msg("successfully finished")
	return
}

func ProcessDirectory(fileDirectoryPath, postgresDirectoryPath string) {
	// read all files in the directory
	filePaths, err := BulkReadDirectory(fileDirectoryPath, postgresDirectoryPath)
	if err != nil {
		log.Fatal().Msg(err.Error())
		return
	}
	// read each file
	for _, filePath := range filePaths {
		BulkInsertFile(filePath)
	}
}

package insert

import (
	"encoding/csv"
	"github.com/jszwec/csvutil"
	"github.com/max-planck-innovation-competition/go-patstat/connections"
	"github.com/max-planck-innovation-competition/go-patstat/pkg/models"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm/clause"
	"io"
	"os"
	"time"
)

// ReadFile reads a file and creates the db entries in batches
func ReadFile(filePath string) {
	logger := log.With().
		Str("file", filePath).
		Logger()
	logger.Info().Msg("reading file")
	// open file
	file, err := os.Open(filePath)
	if err != nil {
		logger.Err(err).Msg("failed to open file")
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logger.Err(err).Msg("failed to close file")
		}
	}(file)
	// init reader
	csvReader := csv.NewReader(file)
	// init csv decoder
	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		logger.Err(err).Msg("failed to decode file")
	}

	// init slice
	var data []*models.Tls206Person
	var counter = 0
	// iterate over lines
	for {
		var obj models.Tls206Person
		if err := dec.Decode(&obj); err == io.EOF {
			break
		} else if err != nil {
			logger.Err(err).Msg("failed to open file")
			return
		}
		// skip if not a company
		// if obj.PsnSector != "COMPANY" {
		// 	continue
		// }

		// append decoded object to slice
		data = append(data, &obj)
		// every 1000 lines create a batch
		if counter > 0 && counter%1000 == 0 {
			create(&logger, 0, data)
			// empty slice
			data = []*models.Tls206Person{}
			log.Info().Int("counter", counter).Msg("inserted")
		}
		counter++
	}
	// final batch
	create(&logger, 0, data)
	logger.Info().Msg("successfully finished")
}

func create(logger *zerolog.Logger, attempt int, data []*models.Tls206Person) {
	if len(data) == 0 {
		return
	}
	// if there are too many attempts
	if attempt > 10 {
		logger.Fatal().Msg("could not create data in db")
		return
	}
	errCreate := connections.SQLClient.Omit(clause.Associations).Create(&data).Error
	if errCreate != nil {
		logger.Warn().Msg(errCreate.Error())
		time.Sleep(time.Second * time.Duration(attempt*5))
		create(logger, attempt+1, data)
		return
	}
}

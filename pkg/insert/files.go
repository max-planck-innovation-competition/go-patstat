package insert

import (
	"encoding/csv"
	"github.com/jszwec/csvutil"
	"github.com/max-planck-innovation-competition/go-patstat/connections"
	"github.com/max-planck-innovation-competition/go-patstat/pkg/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
	"io"
	"os"
	"time"
)

// ProcessDirectory processes all files in a directory
func ProcessDirectory() {
	/*
		tls201 appln.go
		tls202 appln_title.go
		tls203 appln_abstract.go
		tls204 appln_prior.go
		tls205 tech_rel.go
		tls206 person.go
		tls207 pers_appln.go
		tls209 appln_ipc.go
		tls210 appln_n_cls.go
		tls211 pat_publn.go
		tls212 citation.go
		tls214 npl_publn.go
		tls215 citn_categ.go
		tls216 appln_contn.go
		tls222 appln_jp_class.go
		tls223 appln_docus.go
		tls224 appln_cpc.go
		tls225 docdb_fam_cpc.go
		tls226 person_orig.go
		tls227 pers_publn.go
		tls228 docdb_fam_citn.go
		tls229 appln_nace2.go
		tls230 appln_techn_field.go
		tls231 inpadoc_legal_event.go
		tls801 country.go
		tls803 legal_event_code.go
		tls901 techn_field_ipc.go
		tls902 ipc_nace2.go
		tls904 nuts.go
	*/

}

// ReadFile reads a file and creates the db entries in batches
func ReadFile(filePath string) {
	// open file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	// init reader
	csvReader := csv.NewReader(file)
	// init csv decoder
	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		log.Fatal(err)
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
			log.Fatal(err)
			return
		}
		// append decoded object to slice
		data = append(data, &obj)
		// every 1000 lines create a batch
		if counter > 0 && counter%1000 == 0 {
			create(0, data)
			// empty slice
			data = []*models.Tls206Person{}
		}
		counter++
	}
}

func create(attempt int, data []*models.Tls206Person) {
	if len(data) == 0 {
		return
	}
	// if there are too many attempts
	if attempt > 10 {
		log.Fatal("could not create data in db")
		return
	}
	errCreate := connections.SQLClient.Omit(clause.Associations).Create(&data).Error
	if errCreate != nil {
		time.Sleep(time.Second * time.Duration(attempt*5))
		create(attempt+1, data)
		return
	}
}

package insert

import (
	"github.com/max-planck-innovation-competition/go-patstat/pkg/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFile2Table(t *testing.T) {
	ass := assert.New(t)
	// s.th else
	_, err := File2Table("../../../_fixtures/insert/file_2_table.go")
	ass.Error(err)
	// empty
	_, err = File2Table("")
	ass.Error(err)
	// real file
	tls201TableName, err := File2Table("../../../tls201_part06.csv")
	tls201Obj := models.Tls201Appln{}
	tableName := tls201Obj.TableName()
	ass.Equal(tableName, tls201TableName)
	// zip
	tls201TableName, err = File2Table("../../../tls201_part06.zip")
	ass.Equal(tableName, tls201TableName)

}

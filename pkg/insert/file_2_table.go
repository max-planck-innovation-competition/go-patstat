package insert

import (
	"github.com/max-planck-innovation-competition/go-patstat/pkg/models"
	"github.com/pkg/errors"
	"log/slog"
	"path/filepath"
	"regexp"
)

// RegexTls is the regexp for matching the prefix of a file name
var RegexTls = regexp.MustCompile(`^(tls[0-9]{3})`)

// ErrorNoTableName is the error returned when no table name is found based on the filename
var ErrorNoTableName = errors.New("can not find a corresponding table name")

// ErrorNoPrefix is the error returned when no prefix can be extracted
var ErrorNoPrefix = errors.New("can not find a corresponding prefix")

// File2Table is a map of filename to table name
func File2Table(filePath string) (tableName string, err error) {
	// extract the filename from the filepath
	fileName := filepath.Base(filePath)
	// extract the prefix of the filename using the regexp
	prefix := RegexTls.FindStringSubmatch(fileName)
	// if no prefix is found, return an error
	if len(prefix) == 0 {
		slog.With("fileName", fileName).Error("no prefix found")
		return "", ErrorNoPrefix
	}
	// based on the prefix, find the table name
	switch prefix[0] {
	case "tls201":
		obj := models.Tls201Appln{}
		return obj.TableName(), nil
	case "tls202":
		obj := models.Tls202ApplnTitle{}
		return obj.TableName(), nil
	case "tls203":
		obj := models.Tls203ApplnAbstr{}
		return obj.TableName(), nil
	case "tls204":
		obj := models.Tls204ApplnPrior{}
		return obj.TableName(), nil
	case "tls205":
		obj := models.Tls205TechRel{}
		return obj.TableName(), nil
	case "tls206":
		obj := models.Tls206Person{}
		return obj.TableName(), nil
	case "tls207":
		obj := models.Tls207PersAppln{}
		return obj.TableName(), nil
	case "tls209":
		obj := models.Tls209ApplnIpc{}
		return obj.TableName(), nil
	case "tls210":
		obj := models.Tls210ApplnNCls{}
		return obj.TableName(), nil
	case "tls211":
		obj := models.Tls211PatPubln{}
		return obj.TableName(), nil
	case "tls212":
		obj := models.Tls212Citation{}
		return obj.TableName(), nil
	case "tls214":
		obj := models.Tls214NplPubln{}
		return obj.TableName(), nil
	case "tls215":
		obj := models.Tls215CitnCateg{}
		return obj.TableName(), nil
	case "tls216":
		obj := models.Tls216ApplnContn{}
		return obj.TableName(), nil
	case "tls222":
		obj := models.Tls222ApplnJpClass{}
		return obj.TableName(), nil
		// Deprecated
	// case "tls223":
	//	obj := models.Tls223ApplnDocus{}
	//	return obj.TableName(), nil
	case "tls224":
		obj := models.Tls224ApplnCpc{}
		return obj.TableName(), nil
	case "tls225":
		obj := models.Tls225DocdbFamCpc{}
		return obj.TableName(), nil
	case "tls226":
		obj := models.Tls226PersonOrig{}
		return obj.TableName(), nil
	case "tls227":
		obj := models.Tls227PersPubln{}
		return obj.TableName(), nil
	case "tls228":
		obj := models.Tls228DocdbFamCitn{}
		return obj.TableName(), nil
	case "tls229":
		obj := models.Tls229ApplnNace2{}
		return obj.TableName(), nil
	case "tls230":
		obj := models.Tls230ApplnTechnField{}
		return obj.TableName(), nil
	case "tls231":
		obj := models.Tls231InpadocLegalEvent{}
		return obj.TableName(), nil
	case "tls801":
		obj := models.Tls801Country{}
		return obj.TableName(), nil
	case "tls803":
		obj := models.Tls803LegalEventCode{}
		return obj.TableName(), nil
	case "tls901":
		obj := models.Tls901TechnFieldIpc{}
		return obj.TableName(), nil
	case "tls902":
		obj := models.Tls902IpcNace2{}
		return obj.TableName(), nil
	case "tls904":
		obj := models.Tls904Nuts{}
		return obj.TableName(), nil
	}

	return "", ErrorNoTableName
}

package read_csv

import (
	"github.com/SbstnErhrdt/env"
	"github.com/max-planck-innovation-competition/go-patstat/connections"
	"testing"
)

func TestReadFile(t *testing.T) {
	env.LoadEnvFiles()
	connections.ConnectToSQL()
	ReadFile("../../data/tls206_part01.csv")
}

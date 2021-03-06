package insert

import (
	"github.com/SbstnErhrdt/env"
	"github.com/max-planck-innovation-competition/go-patstat/connections"
	"testing"
)

func TestReadFile(t *testing.T) {
	env.LoadEnvFiles("../../.env")
	connections.ConnectToSQL()
	ReadFile("../../data/tls206_part01.csv")
}

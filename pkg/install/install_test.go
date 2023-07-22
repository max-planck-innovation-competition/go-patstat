package install

import "testing"
import "github.com/SbstnErhrdt/env"

func TestInstall(t *testing.T) {
	t.Log("Test Install")
	env.LoadEnvFiles()
	CreateDatabase("patstat_test")
	CreateTables()
	CreateTableConstraints()
}

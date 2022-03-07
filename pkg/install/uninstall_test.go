package install

import (
	"github.com/SbstnErhrdt/env"
	"testing"
)

func TestUninstall(t *testing.T) {
	t.Log("Test Install")
	env.LoadEnvFiles()
	Uninstall()
}

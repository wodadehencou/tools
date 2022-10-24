package version_test

import (
	"testing"

	"github.com/wodadehencou/tools/version"
)

func TestGlobal(t *testing.T) {
	version.Name = "TestVersionProject"
	version.Major = 1
	version.Minor = 1
	version.Patch = 3

	t.Log(version.Full())
	t.Log(version.Short())

	version.Hash = "a0371d"
	t.Log(version.Full())
	t.Log(version.Short())
}

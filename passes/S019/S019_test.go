package S019_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/Codelax/tfproviderlint/passes/S019"
)

func TestS019(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S019.Analyzer, "a")
}

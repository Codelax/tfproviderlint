package S022_test

import (
	"testing"

	"github.com/Codelax/tfproviderlint/passes/S022"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestS022(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S022.Analyzer, "a")
}

package R002_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/Codelax/tfproviderlint/passes/R002"
)

func TestR002(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, R002.Analyzer, "a")
}

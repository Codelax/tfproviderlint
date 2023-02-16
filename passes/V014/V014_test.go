package V014_test

import (
	"testing"

	"github.com/Codelax/tfproviderlint/passes/V014"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestV014(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, V014.Analyzer, "a")
}

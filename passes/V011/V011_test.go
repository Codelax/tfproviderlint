package V011_test

import (
	"testing"

	"github.com/Codelax/tfproviderlint/passes/V011"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestV011(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, V011.Analyzer, "a")
}

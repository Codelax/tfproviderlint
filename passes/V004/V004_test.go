package V004_test

import (
	"testing"

	"github.com/Codelax/tfproviderlint/passes/V004"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestV004(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, V004.Analyzer, "a")
}

func TestAnalyzerFixes(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, V004.Analyzer, "a")
}

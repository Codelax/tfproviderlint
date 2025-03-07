package V008_test

import (
	"testing"

	"github.com/Codelax/tfproviderlint/passes/V008"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestV008(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, V008.Analyzer, "a")
}

func TestAnalyzerFixes(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, V008.Analyzer, "a")
}

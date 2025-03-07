package V005_test

import (
	"testing"

	"github.com/Codelax/tfproviderlint/passes/V005"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestV005(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, V005.Analyzer, "a")
}

func TestAnalyzerFixes(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, V005.Analyzer, "a")
}

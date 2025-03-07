package V002_test

import (
	"testing"

	"github.com/Codelax/tfproviderlint/passes/V002"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestV002(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, V002.Analyzer, "a")
}

func TestAnalyzerFixes(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, V002.Analyzer, "a")
}

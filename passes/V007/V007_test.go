package V007_test

import (
	"testing"

	"github.com/Codelax/tfproviderlint/passes/V007"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestV007(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, V007.Analyzer, "a")
}

func TestAnalyzerFixes(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, V007.Analyzer, "a")
}

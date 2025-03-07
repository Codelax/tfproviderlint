package V006_test

import (
	"testing"

	"github.com/Codelax/tfproviderlint/passes/V006"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestV006(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, V006.Analyzer, "a")
}

func TestAnalyzerFixes(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, V006.Analyzer, "a")
}

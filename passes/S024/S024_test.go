package S024_test

import (
	"testing"

	"github.com/Codelax/tfproviderlint/passes/S024"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestS024(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S024.Analyzer, "a")
}

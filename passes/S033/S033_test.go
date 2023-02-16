package S033_test

import (
	"testing"

	"github.com/Codelax/tfproviderlint/passes/S033"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestS033(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S033.Analyzer, "a")
}

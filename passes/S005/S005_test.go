package S005_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/Codelax/tfproviderlint/passes/S005"
)

func TestS005(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S005.Analyzer, "a")
}

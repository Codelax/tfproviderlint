package XR001_test

import (
	"testing"

	"github.com/Codelax/tfproviderlint/xpasses/XR001"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestXR001(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, XR001.Analyzer, "a")
}

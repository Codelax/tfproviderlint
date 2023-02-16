package fmterrorfcallexpr

import (
	"github.com/Codelax/tfproviderlint/helper/analysisutils"
)

var Analyzer = analysisutils.StdlibFunctionCallExprAnalyzer(
	"fmterrorfcallexpr",
	"fmt",
	"Errorf",
)

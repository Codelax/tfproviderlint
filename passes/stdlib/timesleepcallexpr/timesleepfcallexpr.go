package timesleepcallexpr

import (
	"github.com/Codelax/tfproviderlint/helper/analysisutils"
)

var Analyzer = analysisutils.StdlibFunctionCallExprAnalyzer(
	"timesleepcallexpr",
	"time",
	"Sleep",
)

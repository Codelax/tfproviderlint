package XR007

import (
	"github.com/Codelax/tfproviderlint/helper/analysisutils"
	"github.com/Codelax/tfproviderlint/passes/stdlib/osexeccommandcallexpr"
	"github.com/Codelax/tfproviderlint/passes/stdlib/osexeccommandselectorexpr"
)

var Analyzer = analysisutils.AvoidSelectorExprAnalyzer(
	"XR007",
	osexeccommandcallexpr.Analyzer,
	osexeccommandselectorexpr.Analyzer,
	"os/exec",
	"Command",
)

package XR008

import (
	"github.com/Codelax/tfproviderlint/helper/analysisutils"
	"github.com/Codelax/tfproviderlint/passes/stdlib/osexeccommandcontextcallexpr"
	"github.com/Codelax/tfproviderlint/passes/stdlib/osexeccommandcontextselectorexpr"
)

var Analyzer = analysisutils.AvoidSelectorExprAnalyzer(
	"XR008",
	osexeccommandcontextcallexpr.Analyzer,
	osexeccommandcontextselectorexpr.Analyzer,
	"os/exec",
	"CommandContext",
)

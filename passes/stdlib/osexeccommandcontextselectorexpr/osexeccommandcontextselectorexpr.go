package osexeccommandcontextselectorexpr

import (
	"github.com/Codelax/tfproviderlint/helper/analysisutils"
)

var Analyzer = analysisutils.StdlibFunctionSelectorExprAnalyzer(
	"osexeccommandselectorexpr",
	"os/exec",
	"CommandContext",
)

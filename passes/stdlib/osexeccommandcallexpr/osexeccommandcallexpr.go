package osexeccommandcallexpr

import (
	"github.com/Codelax/tfproviderlint/helper/analysisutils"
)

var Analyzer = analysisutils.StdlibFunctionCallExprAnalyzer(
	"osexeccommandcallexpr",
	"os/exec",
	"Command",
)

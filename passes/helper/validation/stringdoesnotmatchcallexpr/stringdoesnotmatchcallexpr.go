package stringdoesnotmatchcallexpr

import (
	"github.com/Codelax/tfproviderlint/helper/analysisutils"
	"github.com/Codelax/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.FunctionCallExprAnalyzer(
	"stringdoesnotmatchcallexpr",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameStringDoesNotMatch,
)

package stringinslicecallexpr

import (
	"github.com/Codelax/tfproviderlint/helper/analysisutils"
	"github.com/Codelax/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.FunctionCallExprAnalyzer(
	"stringinslicecallexpr",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameStringInSlice,
)

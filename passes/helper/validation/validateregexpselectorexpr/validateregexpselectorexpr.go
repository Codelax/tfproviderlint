package validateregexpselectorexpr

import (
	"github.com/Codelax/tfproviderlint/helper/analysisutils"
	"github.com/Codelax/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"validateregexpselectorexpr",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameValidateRegexp,
)

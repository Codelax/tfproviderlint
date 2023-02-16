package validatejsonstringselectorexpr

import (
	"github.com/Codelax/tfproviderlint/helper/analysisutils"
	"github.com/Codelax/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"validatejsonstringselectorexpr",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameValidateJsonString,
)

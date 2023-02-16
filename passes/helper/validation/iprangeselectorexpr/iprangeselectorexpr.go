package iprangeselectorexpr

import (
	"github.com/Codelax/tfproviderlint/helper/analysisutils"
	"github.com/Codelax/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"iprangeselectorexpr",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameIPRange,
)

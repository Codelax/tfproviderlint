package V005

import (
	"github.com/Codelax/tfproviderlint/helper/analysisutils"
	"github.com/Codelax/tfproviderlint/helper/terraformtype/helper/validation"
	"github.com/Codelax/tfproviderlint/passes/helper/validation/validatejsonstringselectorexpr"
)

var Analyzer = analysisutils.DeprecatedWithReplacementSelectorExprAnalyzer(
	"V005",
	validatejsonstringselectorexpr.Analyzer,
	validation.PackagePath,
	validation.FuncNameValidateJsonString,
	validation.PackagePath,
	validation.FuncNameStringIsJSON,
)

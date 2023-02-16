package V007

import (
	"github.com/Codelax/tfproviderlint/helper/analysisutils"
	"github.com/Codelax/tfproviderlint/helper/terraformtype/helper/validation"
	"github.com/Codelax/tfproviderlint/passes/helper/validation/validateregexpselectorexpr"
)

var Analyzer = analysisutils.DeprecatedWithReplacementSelectorExprAnalyzer(
	"V007",
	validateregexpselectorexpr.Analyzer,
	validation.PackagePath,
	validation.FuncNameValidateRegexp,
	validation.PackagePath,
	validation.FuncNameStringIsValidRegExp,
)

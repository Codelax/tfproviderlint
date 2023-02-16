package V004

import (
	"github.com/Codelax/tfproviderlint/helper/analysisutils"
	"github.com/Codelax/tfproviderlint/helper/terraformtype/helper/validation"
	"github.com/Codelax/tfproviderlint/passes/helper/validation/singleipcallexpr"
	"github.com/Codelax/tfproviderlint/passes/helper/validation/singleipselectorexpr"
)

var Analyzer = analysisutils.DeprecatedEmptyCallExprWithReplacementSelectorExprAnalyzer(
	"V004",
	singleipcallexpr.Analyzer,
	singleipselectorexpr.Analyzer,
	validation.PackagePath,
	validation.FuncNameSingleIP,
	validation.PackagePath,
	validation.FuncNameIsIPAddress,
)

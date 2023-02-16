package V003

import (
	"github.com/Codelax/tfproviderlint/helper/analysisutils"
	"github.com/Codelax/tfproviderlint/helper/terraformtype/helper/validation"
	"github.com/Codelax/tfproviderlint/passes/helper/validation/iprangecallexpr"
	"github.com/Codelax/tfproviderlint/passes/helper/validation/iprangeselectorexpr"
)

var Analyzer = analysisutils.DeprecatedEmptyCallExprWithReplacementSelectorExprAnalyzer(
	"V003",
	iprangecallexpr.Analyzer,
	iprangeselectorexpr.Analyzer,
	validation.PackagePath,
	validation.FuncNameIPRange,
	validation.PackagePath,
	validation.FuncNameIsIPv4Range,
)

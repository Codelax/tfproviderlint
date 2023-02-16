package V002

import (
	"github.com/Codelax/tfproviderlint/helper/analysisutils"
	"github.com/Codelax/tfproviderlint/helper/terraformtype/helper/validation"
	"github.com/Codelax/tfproviderlint/passes/helper/validation/cidrnetworkselectorexpr"
)

var Analyzer = analysisutils.DeprecatedWithReplacementSelectorExprAnalyzer(
	"V002",
	cidrnetworkselectorexpr.Analyzer,
	validation.PackagePath,
	validation.FuncNameCIDRNetwork,
	validation.PackagePath,
	validation.FuncNameIsCIDRNetwork,
)

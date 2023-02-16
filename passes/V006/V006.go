package V006

import (
	"github.com/Codelax/tfproviderlint/helper/analysisutils"
	"github.com/Codelax/tfproviderlint/helper/terraformtype/helper/validation"
	"github.com/Codelax/tfproviderlint/passes/helper/validation/validatelistuniquestringsselectorexpr"
)

var Analyzer = analysisutils.DeprecatedWithReplacementSelectorExprAnalyzer(
	"V006",
	validatelistuniquestringsselectorexpr.Analyzer,
	validation.PackagePath,
	validation.FuncNameValidateListUniqueStrings,
	validation.PackagePath,
	validation.FuncNameListOfUniqueStrings,
)

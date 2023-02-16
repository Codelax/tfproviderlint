package validatelistuniquestringsselectorexpr

import (
	"github.com/Codelax/tfproviderlint/helper/analysisutils"
	"github.com/Codelax/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"validatelistuniquestringsselectorexpr",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameValidateListUniqueStrings,
)

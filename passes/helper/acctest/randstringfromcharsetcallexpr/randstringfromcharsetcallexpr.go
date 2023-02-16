package randstringfromcharsetcallexpr

import (
	"github.com/Codelax/tfproviderlint/helper/analysisutils"
	"github.com/Codelax/tfproviderlint/helper/terraformtype/helper/acctest"
)

var Analyzer = analysisutils.FunctionCallExprAnalyzer(
	"randstringfromcharsetcallexpr",
	acctest.IsFunc,
	acctest.PackagePath,
	acctest.FuncNameRandStringFromCharSet,
)

package providertypeassertexpr

import (
	"github.com/Codelax/tfproviderlint/helper/analysisutils"
	"github.com/Codelax/tfproviderlint/helper/terraformtype/helper/schema"
)

var Analyzer = analysisutils.TypeAssertExprAnalyzer(
	"providertypeassertexpr",
	schema.IsFunc,
	schema.PackagePath,
	schema.TypeNameProvider,
)

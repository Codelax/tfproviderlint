package resourcedatasetpartialselectorexpr

import (
	"github.com/Codelax/tfproviderlint/helper/analysisutils"
	"github.com/Codelax/tfproviderlint/helper/terraformtype/helper/schema"
)

var Analyzer = analysisutils.ReceiverMethodSelectorExprAnalyzer(
	"resourcedatasetpartialselectorexpr",
	schema.IsReceiverMethod,
	schema.PackagePath,
	schema.TypeNameResourceData,
	"SetPartial",
)

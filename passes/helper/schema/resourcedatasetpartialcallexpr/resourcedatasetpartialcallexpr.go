package resourcedatasetpartialcallexpr

import (
	"github.com/Codelax/tfproviderlint/helper/analysisutils"
	"github.com/Codelax/tfproviderlint/helper/terraformtype/helper/schema"
)

var Analyzer = analysisutils.ReceiverMethodCallExprAnalyzer(
	"resourcedatasetpartialcallexpr",
	schema.IsReceiverMethod,
	schema.PackagePath,
	schema.TypeNameResourceData,
	"SetPartial",
)

package R008

import (
	"github.com/Codelax/tfproviderlint/helper/analysisutils"
	"github.com/Codelax/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/Codelax/tfproviderlint/passes/helper/schema/resourcedatasetpartialcallexpr"
	"github.com/Codelax/tfproviderlint/passes/helper/schema/resourcedatasetpartialselectorexpr"
)

var Analyzer = analysisutils.DeprecatedReceiverMethodSelectorExprAnalyzer(
	"R008",
	resourcedatasetpartialcallexpr.Analyzer,
	resourcedatasetpartialselectorexpr.Analyzer,
	schema.PackagePath,
	schema.TypeNameResourceData,
	"SetPartial",
)

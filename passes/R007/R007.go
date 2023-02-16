package R007

import (
	"github.com/Codelax/tfproviderlint/helper/analysisutils"
	"github.com/Codelax/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/Codelax/tfproviderlint/passes/helper/schema/resourcedatapartialcallexpr"
	"github.com/Codelax/tfproviderlint/passes/helper/schema/resourcedatapartialselectorexpr"
)

var Analyzer = analysisutils.DeprecatedReceiverMethodSelectorExprAnalyzer(
	"R007",
	resourcedatapartialcallexpr.Analyzer,
	resourcedatapartialselectorexpr.Analyzer,
	schema.PackagePath,
	schema.TypeNameResourceData,
	"Partial",
)

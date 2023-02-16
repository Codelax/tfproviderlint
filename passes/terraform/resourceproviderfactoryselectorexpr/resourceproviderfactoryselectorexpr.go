package resourceproviderfactoryselectorexpr

import (
	"github.com/Codelax/tfproviderlint/helper/analysisutils"
	"github.com/Codelax/tfproviderlint/helper/terraformtype/terraform"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"resourceproviderfactoryselectorexpr",
	terraform.IsFunc,
	terraform.PackagePath,
	terraform.TypeNameResourceProviderFactory,
)

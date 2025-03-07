package analysisutils

import (
	"go/ast"

	"github.com/Codelax/tfproviderlint/helper/astutils"
	"github.com/Codelax/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/Codelax/tfproviderlint/passes/commentignore"
	"github.com/Codelax/tfproviderlint/passes/helper/schema/schemainfo"
	"golang.org/x/tools/go/analysis"
)

// SchemaAttributeReferencesRunner returns an Analyzer runner for fields that use schema attribute references
func SchemaAttributeReferencesRunner(analyzerName string, fieldName string) func(*analysis.Pass) (interface{}, error) {
	return func(pass *analysis.Pass) (interface{}, error) {
		ignorer := pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
		schemaInfos := pass.ResultOf[schemainfo.Analyzer].([]*schema.SchemaInfo)

		for _, schemaInfo := range schemaInfos {
			if ignorer.ShouldIgnore(analyzerName, schemaInfo.AstCompositeLit) {
				continue
			}

			if !schemaInfo.DeclaresField(fieldName) {
				continue
			}

			switch value := schemaInfo.Fields[fieldName].Value.(type) {
			case *ast.CompositeLit:
				if !astutils.IsExprTypeArrayString(value.Type) {
					continue
				}

				for _, elt := range value.Elts {
					attributeReference := astutils.ExprStringValue(elt)

					if attributeReference == nil {
						continue
					}

					if _, err := schema.ParseAttributeReference(*attributeReference); err != nil {
						pass.Reportf(elt.Pos(), "%s: invalid %s attribute reference: %s", analyzerName, fieldName, err)
					}
				}
			}
		}

		return nil, nil
	}
}

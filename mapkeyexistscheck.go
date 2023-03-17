package mapkeyexistscheck

import (
	"go/ast"

	"github.com/k0kubun/pp/v3"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "mapkeyexistscheck",
	Doc:  "checking an existence check is being performed when accessing a map",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {

		ast.Inspect(file, func(n ast.Node) bool {
			if stmt, ok := n.(*ast.AssignStmt); ok {
				if indexExpr, ok := stmt.Rhs[0].(*ast.IndexExpr); ok {
					if indexExprXIdent, ok := indexExpr.X.(*ast.Ident); ok {
						if indexExprXIdentObjDecl, ok := indexExprXIdent.Obj.Decl.(*ast.AssignStmt); ok {
							if indexExprXIdentObjDeclRhsCompositeLit, ok := indexExprXIdentObjDecl.Rhs[0].(*ast.CompositeLit); ok {
								if _, ok := indexExprXIdentObjDeclRhsCompositeLit.Type.(*ast.MapType); ok {
									if len(stmt.Lhs) == 1 {
										pp.Println(stmt)
										pass.Report(
											analysis.Diagnostic{
												Pos:     stmt.Pos(),
												End:     stmt.End(),
												Message: "map key exists check is not performed",
											})
									}
								}
							}
						}
					}
				}
			}
			return true
		})
	}
	return nil, nil
}

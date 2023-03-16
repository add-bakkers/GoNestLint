package mylinter

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "mylinter is a static analysis tool that detects unnecessarily deep nested levels of control flow."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "GoNestLint",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	inspect.Preorder(nil, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.IfStmt:
			pass.Reportf(n.Pos(), "unnecessarily nested")
		default:
			fmt.Print(n)
		}
	})

	return nil, nil
}

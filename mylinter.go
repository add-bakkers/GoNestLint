package mylinter

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"

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

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	inspect.Preorder(nil, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.IfStmt:
			if detectUnnecessaryNesting(n, pass) {
				pass.Reportf(n.Pos(), "unnecessarily nested")
			}
		}
	})

	return nil, nil
}

func detectUnnecessaryNesting(ifStmt *ast.IfStmt, pass *analysis.Pass) bool {
	if ifStmt.Else == nil || ifStmt.Body == nil || isIfStmt(ifStmt.Else) {
		return false
	}

	ifBlock, ok1 := ifStmt.Body.List[len(ifStmt.Body.List)-1].(*ast.ExprStmt)
	elseBlock, ok2 := ifStmt.Else.(*ast.BlockStmt)
	if !ok1 || !ok2 || len(elseBlock.List) == 0 {
		return false
	}

	elseLastStmt, ok3 := elseBlock.List[len(elseBlock.List)-1].(*ast.ExprStmt)
	if !ok3 {
		return false
	}

	return nodesCompare(ifBlock, elseLastStmt, pass.Fset)
}

func isIfStmt(stmt ast.Stmt) bool {
	_, ok := stmt.(*ast.IfStmt)
	return ok
}

func nodesCompare(a, b ast.Node, fset *token.FileSet) bool {
	bufA := new(bytes.Buffer)
	bufB := new(bytes.Buffer)

	errA := printer.Fprint(bufA, fset, a)
	errB := printer.Fprint(bufB, fset, b)

	if errA != nil || errB != nil {
		return false
	}

	return bufA.String() == bufB.String()
}

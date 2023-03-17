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
	if ifStmt.Body == nil {
		return false
	}

	if len(ifStmt.Body.List) == 1 {
		if _, ok := ifStmt.Body.List[0].(*ast.IfStmt); ok {
			return true
		}
	}

	if ifStmt.Else != nil {
		if elseBlock, ok := ifStmt.Else.(*ast.BlockStmt); ok {
			lastStmtInIf := getLastStmt(ifStmt.Body)
			lastStmtInElse := getLastStmt(elseBlock)
			if lastStmtInIf == nil || lastStmtInElse == nil {
				return false
			}

			if nodesEqual(lastStmtInIf, lastStmtInElse, pass.Fset) {
				return true
			}
		} else if nestedIf, ok := ifStmt.Else.(*ast.IfStmt); ok {
			return detectUnnecessaryNesting(nestedIf, pass)
		}
	}

	return false
}

func getLastStmt(block *ast.BlockStmt) ast.Stmt {
	if len(block.List) == 0 {
		return nil
	}
	return block.List[len(block.List)-1]
}

func nodesEqual(a, b ast.Node, fset *token.FileSet) bool {
	bufA := new(bytes.Buffer)
	bufB := new(bytes.Buffer)

	errA := printer.Fprint(bufA, fset, a)
	errB := printer.Fprint(bufB, fset, b)

	if errA != nil || errB != nil {
		return false
	}

	return bufA.String() == bufB.String()
}

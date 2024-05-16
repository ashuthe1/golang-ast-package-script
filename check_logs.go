package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func checkLogStatement(funcName string, block *ast.BlockStmt) bool {
	for _, stmt := range block.List {
		if exprStmt, ok := stmt.(*ast.ExprStmt); ok {
			if callExpr, ok := exprStmt.X.(*ast.CallExpr); ok {
				// Check if the called function is a selector (e.g., log.Info)
				if selExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
					if strings.HasPrefix(selExpr.Sel.Name, "Print") {
						if len(callExpr.Args) > 0 {
							if basicLit, ok := callExpr.Args[0].(*ast.BasicLit); ok {
								if basicLit.Kind == token.STRING {
									logMessage := strings.Trim(basicLit.Value, `"`)
									if strings.HasPrefix(logMessage, funcName) {
										fmt.Println("Function: ", funcName, " is following convention :)")
										return true
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return false
}

func main() {

	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, "./test/main.go", nil, parser.AllErrors)
	if err != nil {
		fmt.Printf("Failed to parse file: %v\n", err)
		return
	}

	hasMissingLog := false
	ast.Inspect(node, func(n ast.Node) bool {
		if fn, ok := n.(*ast.FuncDecl); ok {
			if fn.Body != nil {
				funcName := fn.Name.Name
				if !checkLogStatement(funcName, fn.Body) {
					fmt.Printf("Function:  %s is missing log statement starting with its name :(\n", funcName)
					hasMissingLog = true
				}

			}
		}
		return true
	})

	if !hasMissingLog {
		fmt.Println("All functions have log statements starting with their name.")
	}
}

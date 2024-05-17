package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// regex pattern to match log functions (e.g., log.Error(), log.Info() etc.)
var logFuncPattern = regexp.MustCompile(`^log\.[A-Za-z]+$`)
// regex pattern to match log functions (e.g., fmt.Print(), fmt.Println() etc.)
var fmtFuncPattern = regexp.MustCompile(`^fmt\.[A-Za-z]+$`)

// checkLogStatements checks if all log statements in the function body follow the convention.
func checkLogStatements(funcName string, block *ast.BlockStmt) bool {
	hasLogStatement := false
	allFollowConvention := true

	for _, stmt := range block.List {
		if containsLogCall(stmt) {
			hasLogStatement = true
			if !followingConvention(funcName, stmt) {
				allFollowConvention = false
			}
		}
	}

	if hasLogStatement && !allFollowConvention {
		fmt.Printf("		\033[33m Function: %s has some log statements not following the convention \033[0m\n", funcName)
		return false
	}
	if hasLogStatement {
		fmt.Printf("		\033[32m Function: %s is following convention \033[0m\n", funcName)
	}
	return true
}

func followingConvention(funcName string, stmt ast.Stmt) bool {
	exprStmt, ok := stmt.(*ast.ExprStmt)
	if !ok {
		return false
	}

	callExpr, ok := exprStmt.X.(*ast.CallExpr)
	if !ok {
		return false
	}

	if len(callExpr.Args) == 0 {
		return false
	}

	basicLit, ok := callExpr.Args[0].(*ast.BasicLit)
	if !ok || basicLit.Kind != token.STRING {
		return false
	}

	logMessage := strings.Trim(basicLit.Value, `"`)
	// fmt.Println("logMessage:", logMessage)
	return strings.HasPrefix(logMessage, funcName)
}

func containsLogCall(stmt ast.Stmt) bool {
	exprStmt, ok := stmt.(*ast.ExprStmt)
	if !ok {
		return false
	}

	callExpr, ok := exprStmt.X.(*ast.CallExpr)
	if !ok {
		return false
	}

	selExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	// Combine the package name and function name to match with the regex
	fullFuncName := fmt.Sprintf("%s.%s", selExpr.X, selExpr.Sel.Name)
	// fmt.Println("fullFuncName", fullFuncName)
	return (logFuncPattern.MatchString(fullFuncName) || fmtFuncPattern.MatchString(fullFuncName))
}

func inspectNode(n ast.Node, hasMissingLog *bool) bool {
	fn, ok := n.(*ast.FuncDecl)
	if !ok || fn.Body == nil {
		return true
	}

	funcName := fn.Name.Name
	fmt.Println("	Checking Function: ",funcName,"()")
	if !checkLogStatements(funcName, fn.Body) {
		*hasMissingLog = true
	}
	return true
}

// processFile processes a single Go file to check for log statement conventions.
func processFile(filePath string, hasMissingLog *bool) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
	if err != nil {
		fmt.Printf("Failed to parse file %s: %v\n", filePath, err)
		return
	}

	// Inspect the AST and check each function declaration.
	ast.Inspect(node, func(n ast.Node) bool {
		return inspectNode(n, hasMissingLog)
	})
}

// processDirectory recursively processes all .go files in a directory.
func processDirectory(dirPath string, hasMissingLog *bool) {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			fmt.Println("Processing file: ", path)
			processFile(path, hasMissingLog)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Failed to process directory %s: %v\n", dirPath, err)
	}
}

func main() {
	dirPath := "./src"
	hasMissingLog := false

	processDirectory(dirPath, &hasMissingLog)

	if !hasMissingLog {
		fmt.Println("All functions have log statements following the convention.")
	}
}

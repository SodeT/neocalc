// TODO:
// implement a runtime logging system to log warnings, errors and general information to the user
// implement integrals and derivative solvers
// switch to a arbitrary precision math lbrary (math/big)

package main

import (
	"fmt"
	"neocalc/src/ast"
	"neocalc/src/runtime"
	"neocalc/src/tokenizer"
	"neocalc/src/utils"
)

var (
	pl = fmt.Println
)

func main() {
	runInput("f(A)")
}

func runInput(input string) float64 {
	utils.OriginalInput = input
	toks := tokenizer.Tokenize(input)
	ast := ast.Parse(toks)
	ans := runtime.Execute(ast)
	pl(">>>", ans)
	return ans
}


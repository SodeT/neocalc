package main

import (
	"fmt"
	"neocalc/src/ast"
	"neocalc/src/runtime"
	"neocalc/src/tokenizer"
)

var (
	pl = fmt.Println
)

func main() {
	runInput("PI = 3.1415")
	runInput("f(X,Y)=3X+Y*PI")
	runInput("f(3, 1)")
}

func runInput(input string) float64 {
	toks := tokenizer.Tokenize(input)
	ast := ast.Parse(toks)
	ans := runtime.Execute(ast)
	pl(input)
	pl(">>>", ans)
	return ans
}


// TODO:
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
	runInput("5+++2")
}

func runInput(input string) float64 {
	utils.Input = input
	toks := tokenizer.Tokenize(input)
	ast := ast.Parse(toks)
	ans, msg := runtime.Execute(ast)
	pl(">>>", ans)
	displayMsg(msg)
	return ans
}

func displayMsg(msg utils.Message) {
	switch msg.Level {
	case utils.NIL_LOG:
		return
	case utils.INFO_LOG:
		pl("INFO: ", msg.Message)
	case utils.WARN_LOG:
		pl("WARNING: ", msg.Message)
	case utils.ERR_LOG:
		pl("ERROR: ", msg.Message)
	}
}


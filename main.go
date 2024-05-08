package main

import (
	"fmt"
	"math"
	"strconv"
	"neocalc/src/ast"
	"neocalc/src/runtime"
	"neocalc/src/tokenizer"
	"neocalc/src/utils"
)

var (
	pl = fmt.Println
	nodes = 0
)

func main() {
	toks := tokenizer.Tokenize("100-10*2+5")
	ast := ast.Parse(toks)
	ans := runtime.Execute(ast)
	pl(ans)
	pl(rebuildInput(ast))
	//testing()
}

// Rebuilds the input string from the ast, this assumes that the tokenizer works, otherwise these tests will fail aswell
func rebuildInput(node *utils.ASTNode) string {
	res := ""
	switch node.Class {
	case utils.NUMBER_LIT:
		res = strconv.FormatFloat(node.Value, 'f', -1, 64)
	case utils.SUBTRACTION_LIT:
		res = rebuildInput(node.Left) + "-" + rebuildInput(node.Right)
	case utils.ADDITION_LIT:
		res = rebuildInput(node.Left) + "+" + rebuildInput(node.Right)
	case utils.MULTIPLICATION_LIT:
		res = rebuildInput(node.Left) + "*" + rebuildInput(node.Right)
	case utils.DIVISION_LIT:
		res = rebuildInput(node.Left) + "/" + rebuildInput(node.Right)
	default:
		res = "invalid"
	}
	return res
}



func testing() {
	wrap("5 + 3", 8.0)
	wrap("10 - 4", 6.0)
	wrap("2 * 6", 12.0)
	wrap("15 / 3", 5.0)
	wrap("7 + 2 * 4", 15.0)
	wrap("10 - 2 / 2", 9.0)
	wrap("8 * 2 - 5", 11.0)
	wrap("20 / 4 - 1", 4.0)
	wrap("6 + 2 * 3", 12.0)
	wrap("9 - 3 / 3", 8.0)
	wrap("1 + 2 + 3 + 4", 10.0) // Sum of consecutive numbers
	wrap("100 - 10 * 2 + 5", 85.0) // Mix of addition, subtraction, and multiplication
	wrap("0 * 100 / 10", 0.0) // Multiplication with zero
	wrap("1 / 0", math.Inf(1)) // Division by zero (expects positive infinity)
	wrap("10 / 3", 3.3333333333333335) // Division resulting in a repeating decimal
	wrap("1000000 * 1000000", 1e12) // Large multiplication
	wrap("9999999999999999 + 1", 1e16) // Large number addition
}

func wrap(input string, expected float64) {
	toks := tokenizer.Tokenize(input)
	ast := ast.Parse(toks)
	ans := runtime.Execute(ast)
	pl(input, "=", ans)
	if ans != expected {
		pl("///////////////////////////// wrong /////////////////////////////")
	}
	pl("---------------")
}

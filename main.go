package main

import (
	"fmt"
	"neocalc/src/ast"
	"neocalc/src/runtime"
	"neocalc/src/tokenizer"
	"neocalc/src/utils"
	"strconv"
)

var (
	pl = fmt.Println
)

func main() {
	a := runInput("2^10000/2^9998")
	pl("ans: ", a)
	
}

func runInput(input string) float64 {
	toks := tokenizer.Tokenize(input)
	ast := ast.Parse(toks)
	ans := runtime.Execute(ast)
	return ans
}

func rebuildInput(node *utils.ASTNode) string {
	res := ""
	switch node.Class {
	case utils.NUMBER_LIT:
		res = strconv.FormatFloat(node.Value, 'f', -1, 64)
	case utils.VARIABLE_LIT:
		res = node.Identifier
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

func displayTree(node *utils.ASTNode, depth int) {
	pl("class: ", node.Class, depth)

	if node.Left != nil {
		displayTree(node.Left, depth+1)
	}
	if node.Right != nil {
		displayTree(node.Right, depth+1)
	}
}

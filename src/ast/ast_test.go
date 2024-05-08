package ast

import (
	"neocalc/src/utils"
	"neocalc/src/tokenizer"
	"testing"
	"strconv"
)


// Rebuilds the input string from the ast, this assumes that the tokenizer works, otherwise these tests will fail aswell
// TODO: make it check the actuall tree becauce this test does not find all errors in the tree
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

func TestCase1(t *testing.T) {
	const input = "1+2*3+4+5"
	toks := tokenizer.Tokenize(input)
	ast := Parse(toks)
	rebuilt := rebuildInput(ast)
	if input != rebuilt {
		t.Fatal("Rebuilt input did not match original input.", input, rebuilt)
	}
}

func TestCase2(t *testing.T) {
	const input = "1-2+12-3+2"
	toks := tokenizer.Tokenize(input)
	ast := Parse(toks)
	rebuilt := rebuildInput(ast)
	if input != rebuilt {
		t.Fatal("Rebuilt input did not match original input. ", input, rebuilt)
	}
}

func TestCase3(t *testing.T) {
	const input = "100-10*2+5"
	toks := tokenizer.Tokenize(input)
	ast := Parse(toks)
	rebuilt := rebuildInput(ast)
	if input != rebuilt {
		t.Fatal("Rebuilt input did not match original input. ", input, rebuilt)
	}
}

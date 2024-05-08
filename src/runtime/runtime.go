package runtime

import (
	"fmt"
	"neocalc/src/utils"
	"math"
)

var (
	pl = fmt.Println
	variables = make(map[string]float64)
)

func Execute(head *utils.ASTNode) float64 {
	pl("vars: ", variables)
	return evaluateNode(head)
}


func evaluateNode(node *utils.ASTNode) float64 {
	switch node.Class {
	case utils.NUMBER_LIT:
		return node.Value
	case utils.VARIABLE_LIT:
		val := variables[node.Identifier]
		return val
	case utils.SUBTRACTION_LIT:
		return evaluateNode(node.Left) - evaluateNode(node.Right)
	case utils.ADDITION_LIT:
		return evaluateNode(node.Left) + evaluateNode(node.Right)
	case utils.MULTIPLICATION_LIT:
		return evaluateNode(node.Left) * evaluateNode(node.Right)
	case utils.DIVISION_LIT:
		return evaluateNode(node.Left) / evaluateNode(node.Right)
	case utils.POWER_LIT:
		return math.Pow(evaluateNode(node.Left), evaluateNode(node.Right))
	case utils.EQUALITY_LIT:
		variables[node.Left.Identifier] = evaluateNode(node.Right)
		return 0
	default:
		return 0
	}
}



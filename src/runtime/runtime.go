package runtime

import (
	"neocalc/src/utils"
)

func Execute(head *utils.ASTNode) float64 {
	return evaluateNode(head)
}


func evaluateNode(node *utils.ASTNode) float64 {
	switch node.Class {
	case utils.NUMBER_LIT:
		return node.Value
	case utils.SUBTRACTION_LIT:
		return evaluateNode(node.Left) - evaluateNode(node.Right)
	case utils.ADDITION_LIT:
		return evaluateNode(node.Left) + evaluateNode(node.Right)
	case utils.MULTIPLICATION_LIT:
		return evaluateNode(node.Left) * evaluateNode(node.Right)
	case utils.DIVISION_LIT:
		return evaluateNode(node.Left) / evaluateNode(node.Right)
	default:
		return 0
	}
}



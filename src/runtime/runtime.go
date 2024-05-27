package runtime

import (
	"fmt"
	"neocalc/src/utils"
	"math"
)

var (
	pl = fmt.Println
	variables = make(map[string]float64)
	functions = []utils.ASTNode{}
)

func Execute(head *utils.ASTNode) float64 {
	return evaluateNode(head)
}

func evaluateNode(node *utils.ASTNode) float64 {
	switch node.Class {
	case utils.NUMBER_LIT:
		return node.Value
	case utils.VARIABLE_LIT:
		return variables[node.Identifier]
	case utils.FUNCTION_LIT:
		for _, fun := range functions {
			if fun.Identifier == node.Identifier {
				// set all the variables
				for i, param := range fun.Parameters {
					variables[param.Identifier] = evaluateNode(&node.Parameters[i])
				}
				return evaluateNode(fun.Right)
			}
		}
		return 0
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
		if node.Left.Class == utils.VARIABLE_LIT {
			// TODO: check that the variable is allowed to be assigned, pi & e should not be allowed to change
			variables[node.Left.Identifier] = evaluateNode(node.Right)
		} else if node.Left.Class == utils.FUNCTION_LIT {
			fun := node.Left
			fun.Right = node.Right
			functions = append(functions, *fun)
		}
		return 0
	default:
		return 0
	}
}



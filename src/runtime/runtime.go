package runtime

import (
	"fmt"
	"neocalc/src/utils"
	"math"
)

var (
	pl = fmt.Println
	variables = make(map[string]float64)
	functions = make(map[string]utils.ASTNode)
	predefinedVar = []string{"PI", "E"}
	predefinedFunc = []string{"sin", "asin", "cos", "acos", "tan", "atan", "sec", "csc", "cot", "sinh", "asinh", "cosh", "acosh", "tanh", "atanh", "log", "ln"}
)

func Execute(head *utils.ASTNode) float64 {
	variables["PI"] = 3.141592653589793
	variables["E"] = 2.718281828459045

	return evaluateNode(head)
}

func evaluateNode(node *utils.ASTNode) float64 {
	switch node.Class {
	case utils.NUMBER_LIT:
		return node.Value
	case utils.VARIABLE_LIT:
		return variables[node.Identifier]
	case utils.FUNCTION_LIT:
		if inSlice(node.Identifier, predefinedFunc) {
			return callFunc(node)
		}
		fun := functions[node.Identifier]
		if len(node.Parameters) != len(fun.Parameters) {
			// TODO: Throw warning or error here
		}
		// set all the variables
		for i, param := range fun.Parameters {
			variables[param.Identifier] = evaluateNode(&node.Parameters[i])
		}
		return evaluateNode(fun.Right)
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
		defineVarFunc(node)
		return 0
	default:
		return 0
	}
}

// TODO: make storing variables work
func defineVarFunc(node *utils.ASTNode) {
	// dissallow overriding of default predefined variables and functions (pi, sine...)
	if inSlice(node.Left.Identifier, append(predefinedVar, predefinedFunc...)) {
		return
	}

	if node.Left.Class == utils.VARIABLE_LIT {
		val := evaluateNode(node.Right)
		variables[node.Left.Identifier] = val
		/*
		err := storage.SaveVariable(node.Left.Identifier, val)
		if err != nil {
			panic(err)
		}
		*/
	} else if node.Left.Class == utils.FUNCTION_LIT {
		fun := node.Left
		fun.Right = node.Right
		functions[fun.Identifier] = *fun
		/*
		err := storage.SaveFunction(fun.Identifier, utils.OriginalInput)
		if err != nil {
			panic(err)
		}
		*/
	}
}

func callFunc(node *utils.ASTNode) float64 {
	if len(node.Parameters) != 1 {
		return 0
	}

	x := evaluateNode(&node.Parameters[0])
	// TODO: make a config option that automaticually makes trig functions use degrees or radians
	switch node.Identifier {
	case "sin":
		return math.Sin(x)
	case "asin":
		return math.Asin(x)
	case "cos":
		return math.Cos(x)
	case "acos":
		return math.Acos(x)
	case "tan":
		return math.Tan(x)
	case "atan":
		return math.Atan(x)
	case "sec":
		return 1 / math.Cos(x)
	case "csc":
		return 1 / math.Sin(x)
	case "cot":
		return 1 / math.Tan(x)
	case "sinh":
		return math.Sinh(x)
	case "asinh":
		return math.Asinh(x)
	case "cosh":
		return math.Cosh(x)
	case "acosh":
		return math.Acosh(x)
	case "tanh":
		return math.Tanh(x)
	case "atanh":
		return math.Atanh(x)
	case "log":
		return math.Log(x)
	case "ln":
		return math.Log(x) / math.Log(2.718281828459045)
	default:
		return 0
	}
}

func inSlice(identifier string, identifiers []string) bool {
	for _, str := range identifiers {
		if str == identifier {
			return true
		}
	}
	return false
}

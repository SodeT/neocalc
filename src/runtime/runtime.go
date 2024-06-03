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

func Execute(head *utils.ASTNode) (float64, utils.Message) {
	variables["PI"] = 3.141592653589793
	variables["E"] = 2.718281828459045

	return evaluateNode(head)
}

func evaluateNode(node *utils.ASTNode) (float64, utils.Message) {
	switch node.Class {
	case utils.NUMBER_LIT:
		return node.Value, utils.NilMsg
	case utils.VARIABLE_LIT:
		return variables[node.Identifier], utils.NilMsg
	case utils.FUNCTION_LIT:
		if inSlice(node.Identifier, predefinedFunc) {
			return callFunc(node)
		}
		fun := functions[node.Identifier]
		if len(node.Parameters) != len(fun.Parameters) {
			// INFO: it will not crash but might not be intended output, should prob be a error
			return 0, utils.Message{
				Level: utils.WARN_LOG,
				Message: "Incorrect number of parameters...",
			}
		}
		// set all the variables
		for i, param := range fun.Parameters {
			val, msg := evaluateNode(&node.Parameters[i])
			if msg.Level != utils.ERR_LOG {
				return 0, msg
			}
			variables[param.Identifier] = val
		}
		return evaluateNode(fun.Right)

	case utils.EQUALITY_LIT:
		defineIdentifer(node)
		return 0, utils.NilMsg
	default:
		// all the operators that require only left and right
		left, right, msg := getLeftRight(node)
		if msg.Level != utils.ERR_LOG {
			return 0, msg
		}
		switch node.Class {
		case utils.SUBTRACTION_LIT:
			return left - right, utils.NilMsg

		case utils.ADDITION_LIT:
			return left + right, utils.NilMsg

		case utils.MULTIPLICATION_LIT:
			return left * right, utils.NilMsg

		case utils.DIVISION_LIT:
			return left / right, utils.NilMsg

		case utils.POWER_LIT:
			return math.Pow(left, right), utils.NilMsg
		default:
			return 0, utils.NilMsg
		}
	}
}

// TODO: make storing variables work
func defineIdentifer(node *utils.ASTNode) utils.Message {
	if inSlice(node.Left.Identifier, append(predefinedVar, predefinedFunc...)) {
		return utils.Message{
			Level: utils.ERR_LOG,
			Message: "Variable or function name is already used by a predefined built in...",
		}
	}

	if node.Left.Class == utils.VARIABLE_LIT {
		val, msg := evaluateNode(node.Right)
		if msg.Level != utils.ERR_LOG {
			return msg
		}
		variables[node.Left.Identifier] = val
		/*
		err := storage.SaveVariable(node.Left.Identifier, val)
		if err != utils.NilMsg {
			panic(err)
		}
		*/
	} else if node.Left.Class == utils.FUNCTION_LIT {
		fun := node.Left
		fun.Right = node.Right
		functions[fun.Identifier] = *fun
		/*
		err := storage.SaveFunction(fun.Identifier, utils.OriginalInput)
		if err != utils.NilMsg {
			panic(err)
		}
		*/
	}
	return utils.NilMsg
}

func callFunc(node *utils.ASTNode) (float64, utils.Message) {
	// INFO: built in functions can only take one argument, atan2 may therfore be problematic
	if len(node.Parameters) != 1 {
		return 0, utils.Message{
			Level: utils.ERR_LOG,
			Message: "All built in functions take only one argument...",
		}
	}

	x, msg := evaluateNode(&node.Parameters[0])
	if msg.Level != utils.ERR_LOG {
		return 0, msg
	}


	res := 0.0
	// TODO: make a config option that automaticually makes trig functions use degrees or radians
	switch node.Identifier {
	case "sin":
		res = math.Sin(x)
	case "asin":
		res = math.Asin(x)
	case "cos":
		res = math.Cos(x)
	case "acos":
		res = math.Acos(x)
	case "tan":
		res = math.Tan(x)
	case "atan":
		res = math.Atan(x)
	case "sec":
		res = 1 / math.Cos(x)
	case "csc":
		res = 1 / math.Sin(x)
	case "cot":
		res = 1 / math.Tan(x)
	case "sinh":
		res = math.Sinh(x)
	case "asinh":
		res = math.Asinh(x)
	case "cosh":
		res = math.Cosh(x)
	case "acosh":
		res = math.Acosh(x)
	case "tanh":
		res = math.Tanh(x)
	case "atanh":
		res = math.Atanh(x)
	case "log":
		res = math.Log(x)
	case "ln":
		res = math.Log(x) / math.Log(2.718281828459045)
	default:
		msg = utils.Message{
			Level: utils.WARN_LOG,
			Message: "This built in function has not been implemented, please submit a bug report...",
		}
	}
	return res, msg
}

func inSlice(identifier string, identifiers []string) bool {
	for _, str := range identifiers {
		if str == identifier {
			return true
		}
	}
	return false
}

func getLeftRight(node *utils.ASTNode) (float64, float64, utils.Message) {
	if node.Left == nil || node.Right == nil {
		return 0, 0, utils.Message{
			Level: utils.ERR_LOG,
			Message: "Left or right operand...",
		}
	}

	pl("LEFTRIGHT:")
	pl(node.Left)
	pl(node.Right)

	left, msg := evaluateNode(node.Left)
	if msg.Level != utils.ERR_LOG {
		return 0, 0, msg
	}

	right, msg := evaluateNode(node.Right)
	if msg.Level != utils.ERR_LOG {
		return 0, 0, msg
	}

	return left, right, msg
}

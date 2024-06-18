package runtime

import (
	"neocalc/src/utils"
	"math"
)

var (
	Variables = make(map[string]float64)
	Functions = make(map[string]utils.ASTNode)
	predefinedVar = []string{"PI", "E"}
	predefinedFunc = []string{"sin", "asin", "cos", "acos", "tan", "atan", "sec", "csc", "cot", "sinh", "asinh", "cosh", "acosh", "tanh", "atanh", "log", "ln"}
)

func Execute(head *utils.ASTNode) (float64, utils.Message) {
	Variables["PI"] = 3.141592653589793
	Variables["E"] = 2.718281828459045

	return evaluateNode(head)
}

func evaluateNode(node *utils.ASTNode) (float64, utils.Message) {
	switch node.Class {
	case utils.NUMBER_LIT:
		return node.Value, utils.NilMsg
	case utils.VARIABLE_LIT:
		val := Variables[node.Identifier]
		if val == 0 && len(node.Identifier) > 1 {
			return val, utils.Message{
				Level: utils.INFO_LOG,
				Message: "Possible misinterpretation of variable, \"" + node.Identifier + "\" treated as one variable... (AB != A*B)",
			}
		}
		return val, utils.NilMsg
	case utils.FUNCTION_LIT:
		if inSlice(node.Identifier, predefinedFunc) {
			return callFunc(node)
		}
		fun := Functions[node.Identifier]
		if len(node.Parameters) != len(fun.Parameters) {
			return 0, utils.Message{
				Level: utils.WARN_LOG,
				Message: "Incorrect number of parameters...",
			}
			}
			// set all the variables
		for i, param := range fun.Parameters {
			val, msg := evaluateNode(&node.Parameters[i])
			if msg.Level == utils.ERR_LOG {
				return 0, msg
			}
			Variables[param.Identifier] = val
		}
		return evaluateNode(fun.Right)

	case utils.EQUALITY_LIT:
		defineIdentifer(node)
		return 0, utils.NilMsg
	default:
		// all the operators that require only left and right
		left, right, msg := getLeftRight(node)
		if msg.Level == utils.ERR_LOG {
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

func defineIdentifer(node *utils.ASTNode) utils.Message {
	if inSlice(node.Left.Identifier, append(predefinedVar, predefinedFunc...)) {
		return utils.Message{
			Level: utils.ERR_LOG,
			Message: "Variable or function name is already used by a predefined built in...",
		}
	}

	if node.Left.Class == utils.VARIABLE_LIT {
		val, msg := evaluateNode(node.Right)
		if msg.Level == utils.ERR_LOG {
			return msg
		}
		Variables[node.Left.Identifier] = val

	} else if node.Left.Class == utils.FUNCTION_LIT {
		fun := node.Left
		fun.Right = node.Right
		Functions[fun.Identifier] = *fun
	}
	return utils.NilMsg
}

func callFunc(node *utils.ASTNode) (float64, utils.Message) {
	// INFO: built in functions can only take one argument, note to self when implementing other functions
	if len(node.Parameters) != 1 {
		return 0, utils.Message{
			Level: utils.ERR_LOG,
			Message: "All built in functions take only one argument...",
		}
	}

	x, msg := evaluateNode(&node.Parameters[0])
	if msg.Level == utils.ERR_LOG {
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
			Message: "The specified function was not recognized...",
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
			Message: "Missing left or right operand...",
		}
	}

	left, msg := evaluateNode(node.Left)
	if msg.Level == utils.ERR_LOG {
		return 0, 0, msg
	}

	right, msg := evaluateNode(node.Right)
	if msg.Level == utils.ERR_LOG {
		return 0, 0, msg
	}

	return left, right, msg
}

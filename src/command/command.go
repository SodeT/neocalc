package command

import (
	"fmt"
	"neocalc/src/ast"
	"neocalc/src/runtime"
	"neocalc/src/tokenizer"
	"neocalc/src/utils"
	"neocalc/src/validate"
	"os"
	"strconv"
	"strings"
)

func Run(input string) (float64, utils.Message) {
	args := strings.Fields(input)
	switch args[0] {
	case "#deriv":
		// #deriv variable value expression...
		// #deriv X 5 5X/X^2
		varID := args[1]
		x, err := strconv.ParseFloat(args[2], 64)
		if err != nil {
			return 0, utils.Message{
				Level: utils.ERR_LOG,
				Message: "Failed to parse variable...",
			}
		}
		
		toks := tokenizer.Tokenize(strings.Join(args[2:], " "))
		msg := validate.Tokens(toks)
		if msg.Level == utils.ERR_LOG {
			return 0, msg
		}

		ast := ast.Parse(toks)
		return fnDeriv(varID, x, 0.000001, ast)

	case "#integ":
		// #integ variable lower upper expression...
		// #integ X 5 10 5X/X^2
		varID := args[1]
		a, err := strconv.ParseFloat(args[2], 64)
		b, err := strconv.ParseFloat(args[3], 64)
		if err != nil {
			return 0, utils.Message{
				Level: utils.ERR_LOG,
				Message: "Failed to parse variable...",
			}
		}

		toks := tokenizer.Tokenize(strings.Join(args[4:], " "))
		msg := validate.Tokens(toks)
		if msg.Level == utils.ERR_LOG {
			return 0, msg
		}

		ast := ast.Parse(toks)
		return fnInteg(varID, a, b, 16000, ast)

	case "#solve":
		// #solve variable expression...
		// #integ X 5X+7=-23
		varID := args[1]

		expression := strings.Join(args[2:], " ")
		toks := tokenizer.Tokenize(expression)
		msg := validate.Tokens(toks)
		if msg.Level == utils.ERR_LOG {
			return 0, msg
		}

		ast := ast.Parse(toks)
		expr := utils.ASTNode{
			Class: utils.SUBTRACTION_LIT,
			Left: ast.Left,
			Right: ast.Right,
		}

		return fnSolve(varID, 0, &expr)

	case "#help":
		fmt.Println(utils.HelpText)
		return 0, utils.NilMsg
	case "#quit":
		os.Exit(0)
		return 0, utils.NilMsg
	default:
		return 0, utils.Message{
			Level: utils.ERR_LOG,
			Message: "Command not found...",
		}
	}
}

func fnDeriv(varID string, x float64, h float64, expr *utils.ASTNode) (float64, utils.Message) {
	runtime.Variables[varID] = x + h
	y1, msg := runtime.Execute(expr)
	if msg.Level == utils.ERR_LOG {
		return 0, msg
	}

	runtime.Variables[varID] = x
	y2, msg := runtime.Execute(expr)
	if msg.Level == utils.ERR_LOG {
		return 0, msg
	}

	res := (y1-y2) / h
	return res, msg
}

func fnInteg(varID string, a float64, b float64, stepsPerUnit int, expr *utils.ASTNode) (float64, utils.Message) {
	steps := int(b-a)*stepsPerUnit
	stepSize := 1/float64(stepsPerUnit)
	runtime.Variables[varID] = a
	sum := 0.0

	for i := 0; i < steps; i++ {
		y, msg := runtime.Execute(expr)
		if msg.Level == utils.ERR_LOG {
			return 0, msg
		}
		a += stepSize
		runtime.Variables[varID] = a
		sum += y * stepSize
	}
	return sum, utils.NilMsg
}

func fnSolve(varID string, guess float64, expr *utils.ASTNode) (float64, utils.Message) {
	for {
		runtime.Variables[varID] = guess

		y, msg := runtime.Execute(expr)
		if msg.Level == utils.ERR_LOG {
			return 0, msg
		}

		yPrim, msg := fnDeriv(varID, guess, 0.00001, expr)
		if msg.Level == utils.ERR_LOG {
			return 0, msg
		}

		if y == 0 {
			return guess, msg
		}

		guess -= y/yPrim

	}
}


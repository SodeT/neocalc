// TODO:
// search for buggs and edge cases (write tests)
// more debugging and error handeling
// switch to a arbitrary precision math lbrary (math/big)

package main

import (
	"fmt"
	"neocalc/src/ast"
	"neocalc/src/runtime"
	"neocalc/src/tokenizer"
	"neocalc/src/utils"
	"strconv"
	"strings"

)

var (
	pl = fmt.Println
)

func main() {
	runInput("#solve X sin(X)=0.5")
}

func runInput(input string) float64 {

	ans := 0.0
	msg := utils.NilMsg

	if []rune(input)[0] == '#' {
		ans, msg = runCommand(input)
	} else {
		utils.Input = input
		toks := tokenizer.Tokenize(input)
		ast := ast.Parse(toks)
		ans, msg = runtime.Execute(ast)
	}

	pl(">>>", ans)
	displayMsg(msg)
	return ans
}

func displayMsg(msg utils.Message) {
	switch msg.Level {
	case utils.NIL_LOG:
		return
	case utils.INFO_LOG:
		pl("INFO: ", msg.Message)
	case utils.WARN_LOG:
		pl("WARNING: ", msg.Message)
	case utils.ERR_LOG:
		pl("ERROR: ", msg.Message)
	}
}

func runCommand(input string) (float64, utils.Message) {
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
		ast := ast.Parse(toks)
		return fnInteg(varID, a, b, 16000, ast)

	case "#solve":
		// #solve variable expression...
		// #integ X 5X+7=-23
		varID := args[1]

		expression := strings.Join(args[2:], " ")
		toks := tokenizer.Tokenize(expression)
		ast := ast.Parse(toks)
		expr := utils.ASTNode{
			Class: utils.SUBTRACTION_LIT,
			Left: ast.Left,
			Right: ast.Right,
		}

		return fnSolve(varID, 0, &expr)

	case "#help":
	default:
		return 0, utils.Message{
			Level: utils.ERR_LOG,
			Message: "Command not found...",
		}
	}
	return 0, utils.NilMsg
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

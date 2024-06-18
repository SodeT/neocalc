package validate

import (
	"neocalc/src/utils"
	"strings"
	"unicode"
)

var (
	allowedChars = "0123456789.+-*/^()=, #"
)

func Input(input string) utils.Message {
	chars := []rune(input)
	for _, ch := range chars {
		if !(unicode.IsLower(ch) || unicode.IsUpper(ch) || strings.Contains(allowedChars, string(ch))) {
			return utils.Message{
				Level: utils.ERR_LOG,
				Message: "Illegal token detected: " + string(ch),
			}
		}
	}
	return utils.NilMsg
}

func Tokens(toks []utils.Token) utils.Message {
	lparen := 0
	rparen := 0
	eqChar := 0
	lastTok := utils.UNUSED_LIT

	for _, tok := range toks {
		switch tok.Class {
		case utils.LPAREN_LIT:
			lparen++
		case utils.RPAREN_LIT:
			rparen++
		case utils.EQUALITY_LIT:
			eqChar++
		}
		if lastTok == utils.FUNCTION_LIT && tok.Class != utils.LPAREN_LIT {
			return utils.Message{
				Level: utils.ERR_LOG,
				Message: "Function has to be followed by opening parenthesis...",
			}
		}
		lastTok = tok.Class
	}

	if eqChar > 1 {
		return utils.Message{
			Level: utils.ERR_LOG,
			Message: "To many equal signs...",
		}
	}

	if rparen < lparen {
		return utils.Message{
			Level: utils.ERR_LOG,
			Message: "Missing closing parenthesis...",
		}
	} else if rparen > lparen {
		return utils.Message{
			Level: utils.ERR_LOG,
			Message: "Missing opening parenthesis...",
		}
	}

	return utils.NilMsg
}

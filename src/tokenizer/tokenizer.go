package tokenizer

import (
	"fmt"
	"strings"
	"neocalc/src/utils"
)

const (
	NUMBERS = "0123456789."
)

var (
	pl = fmt.Println
)

func Tokenize(input string) []utils.Token {
	if len(input) == 0 {
		return []utils.Token{}
	}

	output := []utils.Token{}
	chars := []rune(input)
	tok := ""
	litClass := getLitClass(chars[0])

	for _, ch := range chars {
		if ch == ' ' {
			continue
		}

		if litClass != getLitClass(ch) || litClass == -1 {
			if len(tok) != 0 {
				output = append(output, utils.NewToken(tok, litClass))
			}
			litClass = getLitClass(ch)
			tok = string(ch)
		} else {
			tok += string(ch)
		}
	}
	output = append(output, utils.NewToken(tok, litClass))
	return output
}

func getLitClass(ch rune) int {
	if strings.Contains(NUMBERS, string(ch)) {
		return utils.NUMBER_LIT
	}

	switch ch {
	case '+':
		return utils.ADDITION_LIT
	case '-':
		return utils.SUBTRACTION_LIT
	case '*':
		return utils.MULTIPLICATION_LIT
	case '/':
		return utils.DIVISION_LIT
	case '(':
		return utils.LPAREN_LIT
	case ')':
		return utils.RPAREN_LIT
	case '=':
		return utils.EQUALITY_LIT
	default:
		return -1
	}
}


package tokenizer

import (
	"fmt"
	"neocalc/src/utils"
	"strings"
	"unicode"
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
				output = append(output, utils.Token{
					Token: tok,
					Class: litClass,
				})
			}
			litClass = getLitClass(ch)
			tok = string(ch)
		} else {
			tok += string(ch)
		}
	}
	output = append(output, utils.Token{
		Token: tok,
		Class: litClass,
	})
	return output
}

func getLitClass(ch rune) int {
	if strings.Contains(NUMBERS, string(ch)) {
		return utils.NUMBER_LIT
	} else if unicode.IsLower(ch) {
		return utils.FUNCTION_LIT
	} else if unicode.IsUpper(ch) {
		return utils.VARIABLE_LIT
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
	case '^':
		return utils.POWER_LIT
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


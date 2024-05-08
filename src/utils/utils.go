package utils

const (
	UNUSED_LIT = iota
	NUMBER_LIT
	FUNCTION_LIT
	VARIABLE_LIT
	ADDITION_LIT
	SUBTRACTION_LIT
	MULTIPLICATION_LIT
	DIVISION_LIT
	EQUALITY_LIT
	LPAREN_LIT
	RPAREN_LIT
)

type Token struct {
	Token string
	Class int
	InTree bool
}

func NewToken(token string, class int) Token {
	return Token{Token: token, Class: class, InTree: false}
}

type ASTNode struct {
	Class int
	Value float64
	Left *ASTNode
	Right *ASTNode
	Parameters []ASTNode
}

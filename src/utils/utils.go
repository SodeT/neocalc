package utils


// the order of the literals is important, higher values corespond to higher precedene
const (
	UNUSED_LIT = iota

	// binary literals
	NONBINARY_LIT
	POWER_LIT
	DIVISION_LIT
	MULTIPLICATION_LIT
	SUBTRACTION_LIT
	ADDITION_LIT
	EQUALITY_LIT

	// nonbinary literals
	NUMBER_LIT
	VARIABLE_LIT

	// prebuilt ast tree when encountering parenthesis
	PREBUILT_LIT
	LPAREN_LIT
	RPAREN_LIT

	FUNCTION_LIT
	SEPARATOR_LIT

	MODULO_LIT
)

type Token struct {
	Token string
	Class int
	Subtree *ASTNode
}

type ASTNode struct {
	Class int
	Value float64
	Identifier string
	Left *ASTNode
	Right *ASTNode
	Parameters []ASTNode
}

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


	// TODO: implement these
	PREBUILT_LIT
	FUNCTION_LIT

	MODULO_LIT
	LPAREN_LIT
	RPAREN_LIT
)

type Token struct {
	Token string
	Class int
	Prebuild *ASTNode
}

type ASTNode struct {
	Class int
	Value float64
	Identifier string
	Left *ASTNode
	Right *ASTNode
	Parameters []ASTNode
}

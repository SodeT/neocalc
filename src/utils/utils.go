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
	FUNCTION_LIT

	// prebuilt ast tree when encountering parenthesis
	PREBUILT_LIT
	LPAREN_LIT
	RPAREN_LIT

	SEPARATOR_LIT

	MODULO_LIT // TODO: implement this one
)

var Input string

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

const (
	NIL_LOG = iota
	INFO_LOG
	WARN_LOG
	ERR_LOG
)

type Message struct {
	Level int
	Message string
}

var NilMsg = Message{
	Level: NIL_LOG,
	Message: "",
}

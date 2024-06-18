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

var HelpText string = `
Usage: neocalc [options]

neocalc - A modern more feature rich replacement of calc

Options:
  -h, --help                 Show this help message and exit
  -v, --version              Show the program version and exit

  -i, --input FILE           Specify the input file


Examples:
  neocalc -i math.txt -o results.txt
  ; f(X) = 5(X+2) - 3X
  ; #solve X f(X) = 0
  >>> -5
  ; area(R) = PI*R^2
  ; A = area(5)
  ; #deriv R 5 area(R)
  >>> 157.0796483747472


Description:
  Neocalc is a CLI calculator that interperets mathematical expressions and evaluates them.

Notes:
  - Variables are defined by uppercase letters
    * A = 5
    * SIZE = 128
  - Functions are defined by lowercase letters followed by their parameters
    * f(X) = 5X
	* sigmoid(X) = E^X / (1+E^X)
  - Every trigonometric funtion is already defined and the constants PI and E are also defined (note that 'E' is uppercase)
`

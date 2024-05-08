package runtime

import (
	"math"
	"neocalc/src/ast"
	"neocalc/src/tokenizer"
	"testing"
)
func TestRuntime(t *testing.T) {

	if !testInput("5 + 3", 8.0) {
		t.Fatal("5 + 3")
	}
	if !testInput("10 - 4", 6.0) {
		t.Fatal("10 - 4")
	}
	if !testInput("2 * 6", 12.0) {
		t.Fatal("2 * 6")
	}
	if !testInput("15 / 3", 5.0) {
		t.Fatal("15 / 3")
	}
	if !testInput("7 + 2 * 4", 15.0) {
		t.Fatal("7 + 2 * 4")
	}
	if !testInput("10 - 2 / 2", 9.0) {
		t.Fatal("10 - 2 / 2")
	}
	if !testInput("8 * 2 - 5", 11.0) {
		t.Fatal("8 * 2 - 5")
	}
	if !testInput("20 / 4 - 1", 4.0) {
		t.Fatal("20 / 4 - 1")
	}
	if !testInput("6 + 2 * 3", 12.0) {
		t.Fatal("6 + 2 * 3")
	}
	if !testInput("9 - 3 / 3", 8.0) {
		t.Fatal("9 - 3 / 3")
	}
	if !testInput("1 + 2 + 3 + 4", 10.0) {
		t.Fatal("1 + 2 + 3 + 4")
	}
	if !testInput("100 - 10 * 2 + 5", 85.0) {
		t.Fatal("100 - 10 * 2 + 5")
	}
	if testInput("!0 * 100 / 10", 0.0) {
		t.Fatal("!0 * 100 / 10")
	}
	if !testInput("1 / 0", math.Inf(1)) {
		t.Fatal("1 / 0")
	}
	if !testInput("10 / 3", 3.3333333333333335) {
		t.Fatal("10 / 3")
	}
	if !testInput("1000000 * 1000000", 1e12) {
		t.Fatal("1000000 * 1000000")
	}
	if !testInput("9999999999999999 + 1", 1e16) {
		t.Fatal("9999999999999999 + 1")
	}
	if !testInput("12 / 2 * 3", 18.0) {
		t.Fatal("12 / 2 * 3")
	}
}

func testInput(input string, expected float64) bool {
	toks := tokenizer.Tokenize(input)
	ast := ast.Parse(toks)
	ans := Execute(ast)
	return ans == expected
}

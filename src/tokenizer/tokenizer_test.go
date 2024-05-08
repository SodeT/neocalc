package tokenizer

import (
	"testing"
	"neocalc/src/utils"
)

func compareTokens(literals []utils.Token, tokens []string) bool {
	if len(literals) != len(tokens) {
		return false
	}

	for i, lit := range literals {
		if lit.Token != tokens[i] {
			return false
		}
	}
	return true
}

func TestCompact(t *testing.T) {
	input := string("2+43*65.2")
	toks := Tokenize(input)
	want := []string{"2", "+", "43", "*", "65.2"}
	if !compareTokens(toks, want) {
		t.Fatalf("input: %s, output: %v, want: %v", input, toks, want)
	}
}

func TestSpaces(t *testing.T) {
	input := string("   2 + 43 * 65.2  ")
	toks := Tokenize(input)
	want := []string{"2", "+", "43", "*", "65.2"}
	if !compareTokens(toks, want) {
		t.Fatalf("input: %s, output: %v, want: %v", input, toks, want)
	}
}

func TestMixed(t *testing.T) {
	input := string("2+ 43 *65.2")
	toks := Tokenize(input)
	want := []string{"2", "+", "43", "*", "65.2"}
	if !compareTokens(toks, want) {
		t.Fatalf("input: %s, output: %v, want: %v", input, toks, want)
	}
}

func TestEmpty(t *testing.T) {
	input := string("")
	toks := Tokenize(input)
	want := []string{}
	if !compareTokens(toks, want) {
		t.Fatalf("input: %s, output: %v, want: %v", input, toks, want)
	}
}

func TestSpacedNumbers(t *testing.T) {
	input := string("100 000 / 23.22")
	toks := Tokenize(input)
	want := []string{"100000", "/", "23.22"}
	if !compareTokens(toks, want) {
		t.Fatalf("input: %s, output: %v, want: %v", input, toks, want)
	}
}

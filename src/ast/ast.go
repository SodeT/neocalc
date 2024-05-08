package ast

import (
	"log"
	"neocalc/src/utils"
	"strconv"
)

var (
	toks = []utils.Token{}
	current = 0
)

func Parse(toks []utils.Token) *utils.ASTNode {
	return getTerms(toks)
}

// BUG: 100-10*2+5, 5 attaches to the 10*2 making the tree evaluate to 100-25
// make the tree take into consideration the order in which the additions and subtracions need to be assembled
func getTerms(toks []utils.Token) *utils.ASTNode {
	first := true
	head := &utils.ASTNode{}
	current := head
	left := []utils.Token{}
	for _, tok := range toks {
		if tok.Class != utils.ADDITION_LIT && tok.Class != utils.SUBTRACTION_LIT {
			left = append(left, tok)
			continue
		}
		leftTree := getFactor(left)
		node := utils.ASTNode{
			Class: tok.Class,
			Left: leftTree,
		}

		if first {
			head = &node
			first = false
		} else {
			current.Right = &node
		}
		current = &node
		left = []utils.Token{}
	}
	leftTree := getFactor(left)
	if first {
		return leftTree
	}
	current.Right = leftTree
	return head
}

func getFactor(toks []utils.Token) *utils.ASTNode {
	first := true
	head := &utils.ASTNode{}
	current := head
	left := []utils.Token{}
	for _, tok := range toks {
		if tok.Class != utils.MULTIPLICATION_LIT && tok.Class != utils.DIVISION_LIT {
			left = append(left, tok)
			continue
		}

		// TODO: clean up this if statement when everything is properly tested
		if len(left) != 1 {
			log.Fatal("left array in getFactor i above 1", left)
		}

		leftTree := getLiteral(left)
		node := utils.ASTNode{
			Class: tok.Class,
			Left: leftTree,
		}

		if first {
			head = &node
			first = false
		} else {
			current.Right = &node
		}
		current = &node
		left = []utils.Token{}
	}
	leftTree := getLiteral(left)
	if first {
		return leftTree
	}
	current.Right = leftTree
	return head
}

func getLiteral(tok []utils.Token) *utils.ASTNode {
	val, err := strconv.ParseFloat(tok[0].Token, 64)
	if err != nil {
		log.Fatal(err)
	}
	return &utils.ASTNode{
		Class: utils.NUMBER_LIT,
		Value: val,
	}
}

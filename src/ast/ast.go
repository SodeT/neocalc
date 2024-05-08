package ast

import (
	"log"
	"neocalc/src/utils"
	"strconv"
)

func Parse(toks []utils.Token) *utils.ASTNode {
	return getBinary(toks, utils.EQUALITY_LIT)
}

func getBinary(toks []utils.Token, litClass int) *utils.ASTNode {
	if litClass == utils.NONBINARY_LIT {
		return getNonBinary(toks)
	}

	first := true
	head := &utils.ASTNode{}
	current := head
	left := []utils.Token{}
	for _, tok := range toks {
		if tok.Class != litClass {
			left = append(left, tok)
			continue
		}

		leftTree := getBinary(left, litClass-1)
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
	leftTree := getBinary(left, litClass-1)
	if first {
		return leftTree
	}
	current.Right = leftTree
	return head
}

func getNonBinary(tokens []utils.Token) *utils.ASTNode {
	if len(tokens) == 1 {
		return tokenToNode(tokens[0])
	}

	head := &utils.ASTNode{
		Class: utils.MULTIPLICATION_LIT,
		Left: tokenToNode(tokens[0]),
	}
	current := head
	for i := 1; i < len(tokens) -1; i++ {
		if i == 1 {
			head.Right = &utils.ASTNode{
				Class: utils.MULTIPLICATION_LIT,
				Left: tokenToNode(tokens[i]),
			}
			current = head.Right
			continue
		}
		
		current.Right = &utils.ASTNode{
			Class: utils.MULTIPLICATION_LIT,
			Left: tokenToNode(tokens[i]),
		}
		current = current.Right
	}
	current.Right = tokenToNode(tokens[len(tokens)-1])

	return head
}

func tokenToNode(token utils.Token) *utils.ASTNode {
	switch token.Class {
	case utils.NUMBER_LIT:
		val, err := strconv.ParseFloat(token.Token, 64)
		if err != nil {
			log.Fatal(err)
		}
		return &utils.ASTNode{
			Class: utils.NUMBER_LIT,
			Value: val,
		}
	case utils.VARIABLE_LIT:
		return &utils.ASTNode{
			Class: utils.VARIABLE_LIT,
			Identifier: token.Token,
		}
	default:
		return &utils.ASTNode{}
	}
}


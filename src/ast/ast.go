package ast

import (
	"fmt"
	"log"
	"neocalc/src/utils"
	"strconv"
)

var (
	pl = fmt.Println
)

func Parse(tokens []utils.Token) *utils.ASTNode {
	tokens = preprocess(tokens)
	return getBinary(tokens, utils.EQUALITY_LIT)
}

func preprocess(tokens []utils.Token) []utils.Token {
	for i := 0; i < len(tokens); i++ {
		tok := tokens[i]
		if tok.Class == utils.FUNCTION_LIT {
			funcToken, end := buildFunc(tokens, i)
			tokens[i] = funcToken
			tokens = append(tokens[:i+1], tokens[end+1:]...)
			continue
		} else if tok.Class == utils.LPAREN_LIT {
			subtreeTok, end := prebuild(tokens, i)
			tokens[i] = subtreeTok
			tokens = append(tokens[:i+1], tokens[end+1:]...)
			continue
		}

	}
	return tokens
}

func prebuild(tokens []utils.Token, start int) (utils.Token, int) {
	end := start
	parenDepth := 0
	for i := start; i < len(tokens); i++ {
		if tokens[i].Class == utils.LPAREN_LIT {
			parenDepth++
		} else if tokens[i].Class == utils.RPAREN_LIT {
			parenDepth--
		}

		if parenDepth == 0 {
			end = i
			break
		}
	}
	subtree := Parse(tokens[start+1:end])
	return utils.Token{
		Class: utils.PREBUILT_LIT,
		Subtree: subtree,
	}, end
}

func buildFunc(tokens []utils.Token, start int) (utils.Token, int) {
	node := &utils.ASTNode{
		Class: utils.FUNCTION_LIT,
		Identifier: tokens[start].Token,
	}

	start += 2 // skipping the function identifier and first parenthesis
	end := start
	parenDepth := 1
	for i := start; i < len(tokens); i++ {
		if tokens[i].Class == utils.LPAREN_LIT {
			parenDepth++
		} else if tokens[i].Class == utils.RPAREN_LIT {
			parenDepth--
		}

		if parenDepth == 0 {
			end = i
			break
		}

		if tokens[i].Class == utils.SEPARATOR_LIT && parenDepth == 1 {
			param := tokens[start:i]
			start = i + 1
			node.Parameters = append(node.Parameters, *Parse(param))
			
		}
	}
	param := tokens[start:end]
	node.Parameters = append(node.Parameters, *Parse(param))
	
	return utils.Token{
		Class: utils.PREBUILT_LIT,
		Subtree: node,
	}, end
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
	case utils.PREBUILT_LIT:
		return token.Subtree
	default:
		return &utils.ASTNode{}
	}
}


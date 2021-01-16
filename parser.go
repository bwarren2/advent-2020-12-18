package advent20201218

import (
	"fmt"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

const (
	MULTIPLICATION = iota
	ADDITION
)

var operators [2]int = [2]int{MULTIPLICATION, ADDITION}

type Evaluator interface {
	Evaluate() int
}

type ExprNode struct {
	Children []Evaluator
}

func (e ExprNode) Evaluate() int {
	start := e.Children[0].Evaluate()
	for i := 1; i < len(e.Children); i += 2 {
		operator := e.Children[i]
		next := e.Children[i+1]
		switch operator.Evaluate() {
		case 1:
			start += next.Evaluate()
		case 0:
			start *= next.Evaluate()
		}
	}

	return start
}

type ValueNode struct {
	Value int
}

func (v ValueNode) Evaluate() int {
	return v.Value
}

type OperatorNode struct {
	Value string
}

func (o OperatorNode) Evaluate() int {
	if o.Value == "*" {
		return 0
	}
	return 1
}

type Parser struct {
	root   ExprNode
	tokens *[]Token
}

func NewParser(tokens *[]Token) Parser {
	return Parser{
		tokens: tokens,
	}
}

func (p *Parser) ParseTerm() ExprNode {
	root := ExprNode{}
	needsMatchingParen := false
	if p.FirstToken().TokenType == OPAREN {
		p.ConsumeTokens(1)
		needsMatchingParen = true
	}
	for len(*p.tokens) > 0 && p.FirstToken().TokenType != CPAREN {
		fmt.Println(p.tokens)
		switch (*p.tokens)[0].TokenType {
		case OPAREN:
			root.Children = append(root.Children, p.ParseTerm())
		case ADD, MUL:
			root.Children = append(root.Children, p.ParseOperator())
		case INTEGER:
			root.Children = append(root.Children, p.ParseInteger())
		case WHITESPACE:
			p.ParseWhitespace()
		case CPAREN:
			break
		default:
			panic("Didn't know what to do!")
		}
	}

	if needsMatchingParen {
		if p.FirstToken().TokenType == CPAREN {
			p.ConsumeTokens(1)
		} else {
			spew.Dump(root)
			panic("Did not have a matching paren!")
		}
	}
	return root
}

func (p *Parser) ParseWhitespace() {
	p.ConsumeTokens(1)
}

func (p Parser) FirstToken() Token {
	return (*p.tokens)[0]
}

func (p *Parser) ParseInteger() ValueNode {
	tokenValue, err := strconv.ParseInt(p.FirstToken().Content, 10, 64)
	if err != nil {
		panic(err)
	}
	p.ConsumeTokens(1)
	return ValueNode{Value: int(tokenValue)}
}
func (p *Parser) ParseOperator() OperatorNode {
	tokenContent := p.FirstToken().Content
	p.ConsumeTokens(1)
	switch tokenContent {
	case "+":
		return OperatorNode{Value: "+"}
	case "*":
		return OperatorNode{Value: "*"}
	}
	return OperatorNode{Value: ""}
}

func (p *Parser) ConsumeTokens(count int) {
	sublist := (*p.tokens)[count:]
	p.tokens = &sublist
}

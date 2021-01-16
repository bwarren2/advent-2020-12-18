package advent20201218

import (
	"strconv"
)

const (
	MULTIPLICATION = iota
	ADDITION
)

var operators [2]int = [2]int{MULTIPLICATION, ADDITION}

type Evaluator interface {
	FlatPrecedenceEvaluate() int
}

type ExprNode struct {
	Children []Evaluator
}

func (e ExprNode) FlatPrecedenceEvaluate() int {
	start := e.Children[0].FlatPrecedenceEvaluate()
	for i := 1; i < len(e.Children); i += 2 {
		operator := e.Children[i]
		next := e.Children[i+1]
		switch operator.FlatPrecedenceEvaluate() {
		case 1:
			start += next.FlatPrecedenceEvaluate()
		case 0:
			start *= next.FlatPrecedenceEvaluate()
		}
	}

	return start
}

type ValueNode struct {
	Value int
}

func (v ValueNode) FlatPrecedenceEvaluate() int {
	return v.Value
}

type OperatorNode struct {
	Value string
}

func (o OperatorNode) FlatPrecedenceEvaluate() int {
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

	for len(*p.tokens) > 0 && p.FirstToken().TokenType != CPAREN {
		switch (*p.tokens)[0].TokenType {
		case OPAREN:
			p.ConsumeTokens(1)
			root.Children = append(root.Children, p.ParseTerm())
			if p.FirstToken().TokenType != CPAREN {
				panic("Need a closing brace after an opening brace!")
			}
			p.ConsumeTokens(1)
		case ADD, MUL:
			root.Children = append(root.Children, p.ParseOperator())
		case INTEGER:
			root.Children = append(root.Children, p.ParseInteger())
		case WHITESPACE:
			p.ConsumeTokens(1)
		default:
			panic("Didn't know what to do!")
		}
	}

	return root
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

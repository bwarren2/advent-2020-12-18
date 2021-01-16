package advent20201218

import (
	"fmt"
	"regexp"
)

// Operators
const (
	OPAREN = iota
	CPAREN
	INTEGER
	ADD
	MUL
	WHITESPACE
)

var tokenTypes [6]int = [6]int{OPAREN, CPAREN, INTEGER, ADD, MUL, WHITESPACE}

var oparenRegex = regexp.MustCompile(`\A\(`)
var cparenRegex = regexp.MustCompile(`\A\)`)
var integerRegex = regexp.MustCompile(`\A\d+`)
var addRegex = regexp.MustCompile(`\A\+`)
var mulRegex = regexp.MustCompile(`\A\*`)
var whitespaceRegex = regexp.MustCompile(`\A\s`)

// Token represents a lexeme of an equation
type Token struct {
	TokenType int
	Content   string
}

// Lexer breaks a string into Tokens
type Lexer struct {
	input    string
	tokens   []Token
	position int
	patterns map[int]*regexp.Regexp
}

// Tokenize converts a string into a stream of tokens
func (l *Lexer) Tokenize() (tokens []Token) {
	for l.position < len(l.input) {
		if !l.ParseToken() {
			panic(fmt.Sprintf("Couldn't find a match in `%v`!", l.input[l.position:]))
		}
	}
	return l.tokens
}

// ParseToken consumes a token from the input, advances the position, and returns true,
// or returns false if there is no match.
func (l *Lexer) ParseToken() bool {
	for _, tokenType := range tokenTypes {
		substring := l.input[l.position:]
		match := l.patterns[tokenType].FindString(substring)
		if match != "" {
			l.tokens = append(l.tokens, Token{tokenType, match})
			l.position += len(match)
			return true
		}
	}
	return false
}

// NewLexer initializes a Lexer
func NewLexer(input string) Lexer {
	return Lexer{
		input:    input,
		tokens:   []Token{},
		position: 0,
		patterns: map[int]*regexp.Regexp{
			OPAREN:     oparenRegex,
			CPAREN:     cparenRegex,
			INTEGER:    integerRegex,
			ADD:        addRegex,
			MUL:        mulRegex,
			WHITESPACE: whitespaceRegex,
		},
	}
}

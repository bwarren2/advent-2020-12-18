package advent20201218_test

import (
	advent "advent20201218"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRecodsFromFile(t *testing.T) {
	fmt.Println(advent.RecordsFromFile("sample.txt"))
	fmt.Println(advent.RecordsFromFile("input.txt"))
	// t.Fail()
}

func TestLexing(t *testing.T) {
	tcs := []struct {
		input  string
		output []advent.Token
	}{
		{
			"1 + 2", []advent.Token{
				advent.Token{2, "1"},
				advent.Token{5, " "},
				advent.Token{3, "+"},
				advent.Token{5, " "},
				advent.Token{2, "2"},
			},
		},
		{
			"(1 + 2) * 3", []advent.Token{
				advent.Token{0, "("},
				advent.Token{2, "1"},
				advent.Token{5, " "},
				advent.Token{3, "+"},
				advent.Token{5, " "},
				advent.Token{2, "2"},
				advent.Token{1, ")"},
				advent.Token{5, " "},
				advent.Token{4, "*"},
				advent.Token{5, " "},
				advent.Token{2, "3"},
			},
		},
	}
	for _, tc := range tcs {
		lexer := advent.NewLexer(tc.input)
		got := lexer.Tokenize()
		want := tc.output
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("Token map() mismatch for %s (-want +got):\n%s", tc.input, diff)
		}
	}
}

func TestParsing(t *testing.T) {
	tcs := []struct {
		input  string
		output advent.ExprNode
	}{
		{" 1", advent.ExprNode{Children: []advent.Evaluator{advent.ValueNode{Value: 1}}}},
		{"2 + 1", advent.ExprNode{Children: []advent.Evaluator{
			advent.ValueNode{Value: 2},
			advent.OperatorNode{Value: "+"},
			advent.ValueNode{Value: 1},
		}}},
		// {"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", advent.ExprNode{Children: []advent.Evaluator{advent.ValueNode{Value: 1}}}},
	}
	for _, tc := range tcs {
		lexer := advent.NewLexer(tc.input)
		tokens := lexer.Tokenize()
		parser := advent.NewParser(&tokens)
		got := parser.ParseTerm()
		want := tc.output
		if diff := cmp.Diff(want, got); diff != "" {
			fmt.Println(got)
			t.Errorf("Token map() mismatch for %s (-want +got):\n%s", tc.input, diff)
		}
	}
}

func TestFlatPrecedenceEvaluate(t *testing.T) {
	tcs := []struct {
		input  string
		output int
	}{
		{" 1", 1},
		{"2 + 1", 3},
	}
	for _, tc := range tcs {
		lexer := advent.NewLexer(tc.input)
		tokens := lexer.Tokenize()
		parser := advent.NewParser(&tokens)
		got := parser.ParseTerm().FlatPrecedenceEvaluate()
		want := tc.output
		if want != got {
			t.Errorf("Wanted %v but got %v for %v", want, got, tc.input)
		}
	}
}

func TestFlatE2E(t *testing.T) {
	tcs := []struct {
		input  string
		output int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 71},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 26},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
	}
	for _, tc := range tcs {
		got := advent.FlatPrecedenceMath(tc.input)
		if got != tc.output {
			t.Errorf("Wanted %v but got %v for %v", tc.output, got, tc.input)
		}
	}
}

func TestAdditionFirstE2E(t *testing.T) {
	tcs := []struct {
		input  string
		output int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 231},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 46},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 1445},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 669060},
		{"(2 + 4 * 9)", 6 * 9},
		{"6 + 9 * 8 + 6", 15 * 14},
		{"(6 + 9 * 8 + 6)", 15 * 14},
		{"(6 + 9 * 8 + 6) + 6", 15*14 + 6},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 23340},
	}
	for _, tc := range tcs {
		got := advent.AdditionPrecedenceMath(tc.input)
		if got != tc.output {
			t.Errorf("Wanted %v but got %v for %v", tc.output, got, tc.input)
		}
	}
}

func TestPart1(t *testing.T) {
	fmt.Println(advent.Part1("sample.txt"))
	fmt.Println(advent.Part1("input.txt"))
	// t.Fail()
}

func TestPart2(t *testing.T) {
	fmt.Println(advent.Part2("sample.txt"))
	fmt.Println(advent.Part2("input.txt"))
	// t.Fail()
}

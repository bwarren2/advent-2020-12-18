package advent20201218

func NewMath(input string) int {
	lexer := NewLexer(input)
	tokens := lexer.Tokenize()
	parser := NewParser(&tokens)
	return parser.ParseTerm().FlatPrecedenceEvaluate()
}

// Part1 solves part1
func Part1(filename string) (total int) {
	for _, tc := range RecordsFromFile(filename) {
		total += NewMath(tc)
	}
	return
}

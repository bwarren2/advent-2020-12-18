package advent20201218

func FlatPrecedenceMath(input string) int {
	lexer := NewLexer(input)
	tokens := lexer.Tokenize()
	parser := NewParser(&tokens)
	return parser.ParseTerm().FlatPrecedenceEvaluate()
}

func AdditionPrecedenceMath(input string) int {
	lexer := NewLexer(input)
	tokens := lexer.Tokenize()
	parser := NewParser(&tokens)
	return parser.ParseTerm().AdditionFirstEvaluate()
}

// Part1 solves part1
func Part1(filename string) (total int) {
	for _, tc := range RecordsFromFile(filename) {
		total += FlatPrecedenceMath(tc)
	}
	return
}

// Part2 solves part2
func Part2(filename string) (total int) {
	for _, tc := range RecordsFromFile(filename) {
		total += AdditionPrecedenceMath(tc)
	}
	return
}

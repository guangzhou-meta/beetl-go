package core

type GrammarToken struct {
	Line int
	Text string
	Col  int
}

func NewGrammarToken(text string, line int, col int) *GrammarToken {
	return &GrammarToken{
		Text: text,
		Line: line,
		Col:  col,
	}
}

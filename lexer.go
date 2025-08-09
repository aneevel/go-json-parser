package main

type TokenType string

const (
	BeginObject TokenType = "{"
	EndObject   TokenType = "}"
)

type Lexer struct {
}

func NewLexer() *Lexer {
	l := &Lexer{}
	return l
}

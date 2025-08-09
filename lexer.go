package main

type TokenType string

const (
	BeginObject TokenType = "{"
	EndObject   TokenType = "}"
)

type Lexer struct {
	tokens []byte
}

func NewLexer(tokens []byte) *Lexer {
	l := &Lexer{
		tokens: tokens,
	}
	return l
}

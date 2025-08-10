package main

type Parser struct {
	lexer *Lexer
}

func NewParser(tokens []byte) *Parser {
	return &Parser{
		lexer: NewLexer(tokens),
	}
}

func (p *Parser) Parse() (status int) {

	return 1
}

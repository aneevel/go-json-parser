package main

import (
	"fmt"
)

type Parser struct {
	lexer *Lexer
}

func NewParser(tokens []byte) *Parser {
	return &Parser{
		lexer: NewLexer(tokens),
	}
}

func (p *Parser) Parse() (status int) {

	for {
		token := p.lexer.GetToken()
		fmt.Printf("Found token: %v\n", token.Type)

		if token.Type == EndOfFile {
			return 0
		}

		p.lexer.NextToken()
	}

}

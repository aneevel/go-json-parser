package main

type Parser struct {
	tokens []byte
}

func NewParser(tokens []byte) *Parser {
	return &Parser{
		tokens: tokens,
	}
}

func (p *Parser) Parse() {

}

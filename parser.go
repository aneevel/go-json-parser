package main

type Parser struct {
	lexer *Lexer
	token Token
}

func NewParser(input string) *Parser {

	p := &Parser{
		lexer: NewLexer(input),
	}
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.token = p.lexer.NextToken()
}

func (p *Parser) Parse() (status int) {

	if p.token.Type != BEGIN_OBJECT {
		return 1
	}

	p.nextToken()

	if p.token.Type != END_OBJECT {
		return 1
	}

	p.nextToken()

	if p.token.Type != END_OF_FILE {
		return 1
	}

	return 0
}

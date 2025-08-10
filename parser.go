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

func (p *Parser) Parse() (status int, data []string) {

	var output []string

	if p.token.Type != BEGIN_OBJECT {
		return 1, output
	}

	for p.token.Type != INVALID {
		p.nextToken()

		output = append(output, p.token.Literal)
	}

	if p.token.Type != END_OBJECT {
		return 1, output
	}

	p.nextToken()

	if p.token.Type != END_OF_FILE {
		return 1, output
	}

	return 0, output
}

package main

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	INVALID      TokenType = "INVALID"
	BEGIN_OBJECT TokenType = "{"
	END_OBJECT   TokenType = "}"
	END_OF_FILE  TokenType = "EOF"
)

/**
* By utilizing "readPosition", we have a lookahead
* pointer that helps us understand what is coming next and
* react to it
 */
type Lexer struct {
	input        string
	position     int
	readPosition int
	current      byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}

/**
* Read the next character into current, progressing
* the readPosition pointer
 */
func (l *Lexer) readChar() {

	// We have progressed past the extent of the input and
	// should signal EOF
	if l.readPosition >= len(l.input) {
		l.current = 0
	} else {
		l.current = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

/**
* Get the next token in the string, advancing the
* pointer in the process
 */
func (l *Lexer) NextToken() Token {

	var token Token

	switch l.current {
	case '{':
		token = Token{BEGIN_OBJECT, string(l.current)}
	case '}':
		token = Token{END_OBJECT, string(l.current)}
	case 0:
		token = Token{END_OF_FILE, string(l.current)}
	default:
		token = Token{INVALID, string(l.current)}
	}

	l.readChar()
	return token
}

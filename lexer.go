package main

import (
	"fmt"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	Unknown     TokenType = "Unknown"
	BeginObject TokenType = "{"
	EndObject   TokenType = "}"
	EndOfFile   TokenType = "EOF"
)

type Lexer struct {
	tokens   []byte
	position int
	current  byte
}

func NewLexer(tokens []byte) *Lexer {
	l := &Lexer{
		tokens:   tokens,
		position: 0,
	}
	return l
}

func (l *Lexer) NextToken() {
	l.position++

	if l.position < len(l.tokens) {
		l.current = l.tokens[l.position]
	}
}

func (l *Lexer) GetToken() Token {

	switch l.current {
	case '{':
		return Token{BeginObject, string(l.current)}
	case '}':
		return Token{EndObject, string(l.current)}
	case 0:
		return Token{EndOfFile, string(l.current)}
	default:
		return Token{Unknown, string(l.current)}
	}
}

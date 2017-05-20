package lexer

import "gitlab.com/fharding/brainf/token"

// Lexer stores info for lexing input
type Lexer struct {
	input        string
	ch           byte
	position     int
	readPosition int
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition + 1
	l.readPosition++
}

// NextToken gets the next token from input
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.eatWhitespace()

	switch l.ch {
	case token.GT:
		tok = newToken(token.GT, l.ch)
	case token.LT:
		tok = newToken(token.LT, l.ch)
	case token.ADD:
		tok = newToken(token.ADD, l.ch)
	case token.SUB:
		tok = newToken(token.SUB, l.ch)
	case token.LBRAC:
		tok = newToken(token.LBRAC, l.ch)
	case token.RBRAC:
		tok = newToken(token.RBRAC, l.ch)
	case token.DOT:
		tok = newToken(token.DOT, l.ch)
	case token.COMMA:
		tok = newToken(token.COMMA, l.ch)
	case token.EOF:
		tok = token.Token{Type: token.EOF, Literal: ""}
	}

	l.readChar()

	return tok
}

func (l *Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(t token.TokenType, ch byte) token.Token {
	return token.Token{Type: t, Literal: string(ch)}
}

// New creates a new lexer from an input
func New(in string) *Lexer {
	l := &Lexer{input: in}
	l.readChar()
	return l
}

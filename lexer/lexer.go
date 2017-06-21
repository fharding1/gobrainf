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
		l.ch = token.EOF
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

	if l.ch == token.EOF {
		tok = token.Token{Type: token.EOF}
	} else {
		tok = token.Token{Type: token.TokenType(l.ch)}
	}

	l.readChar()

	return tok
}

func contains(arr []byte, b byte) bool {
	for _, v := range arr {
		if b == v {
			return true
		}
	}
	return false
}

func (l *Lexer) eatWhitespace() {
	for !contains(append([]byte("<>[].,+-"), 0), l.ch) {
		l.readChar()
	}
}

// New creates a new lexer from an input
func New(in string) *Lexer {
	l := &Lexer{input: in}
	l.readChar()
	return l
}

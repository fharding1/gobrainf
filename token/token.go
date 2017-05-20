package token

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
}

const (
	LT    = '<'
	GT    = '>'
	ADD   = '+'
	SUB   = '-'
	LBRAC = '['
	RBRAC = ']'
	DOT   = '.'
	COMMA = ','
	EOF   = 0
)

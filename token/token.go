package token

// TokenType an int that represents a char like '<'
type TokenType int

// Token a token type and literal value
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

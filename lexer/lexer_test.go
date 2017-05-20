package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/fharding/brainf/token"
)

func TestLex(t *testing.T) {
	input := "<   >[  ]+ \n-"
	res := []token.Token{
		{token.LT, "<"},
		{token.GT, ">"},
		{token.LBRAC, "["},
		{token.RBRAC, "]"},
		{token.ADD, "+"},
		{token.SUB, "-"},
		{token.EOF, ""},
	}

	l := New(input)

	for _, v := range res {
		tok := l.NextToken()
		assert.Equal(t, v.Type, tok.Type)
		assert.Equal(t, v.Literal, tok.Literal)
	}
}

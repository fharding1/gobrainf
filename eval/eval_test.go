package eval

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/fharding/brainf/lexer"
)

func TestEval(t *testing.T) {
	cases := []struct {
		in  string
		err error
	}{
		{"+-", nil},
		{">>", nil},
		{"<<", ErrDataPtrOutOfRange},
		{">><<<", ErrDataPtrOutOfRange},
		{">+", nil},
		{"+++>++++<[->+<]", nil},
		{"<-++>-<++->-", ErrDataPtrOutOfRange},
		{"<<>>", ErrDataPtrOutOfRange},
		{"-+<-+>+<>+", ErrDataPtrOutOfRange},
		{",", ErrNoReadFromNil},
		{".", ErrNoWriteToNil},
	}

	for _, v := range cases {
		l := lexer.New(v.in)

		env := New(l, nil, nil)

		_, err := env.Eval()
		assert.Equal(t, v.err, err, v.in)
	}
}

package eval

import (
	"fmt"
	"io"

	"gitlab.com/fharding/brainf/lexer"
	"gitlab.com/fharding/brainf/token"
)

// Stack is an array of 30,000 integers
type Stack [30000]int

// Environment is a environment for running bf in with a Stack and instruction array, and indexes
type Environment struct {
	instructions     []token.Token
	stack            Stack
	error            error
	out              io.Writer
	in               io.Reader
	dataIndex        int
	instructionIndex int
}

var (
	// ErrDataPtrOutOfRange : the pointer (or index) for data has gone out of range of the Stack
	ErrDataPtrOutOfRange = fmt.Errorf("data pointer out of range")
	// ErrInstructionPtrOutOfRange : the pointer (or index) for instructions has gone out of the range of instructions
	ErrInstructionPtrOutOfRange = fmt.Errorf("instruction pointer out of range")
	// ErrNoWriteToNil cannot write to nil
	ErrNoWriteToNil = fmt.Errorf("cannot write to nil")
	// ErrNoReadFromNil cannot read from nil
	ErrNoReadFromNil = fmt.Errorf("cannot read from nil")
	// ErrNoMatchingBrac there were unmatched brackets
	ErrNoMatchingBrac = fmt.Errorf("unmatched brackets")
)

// New creates a new environment from a lexer
func New(l *lexer.Lexer, in io.Reader, out io.Writer) *Environment {
	var instructions []token.Token

	var tok token.Token
	for tok = l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		instructions = append(instructions, tok)
	}

	instructions = append(instructions, tok)
	return &Environment{instructions: instructions, in: in, out: out}
}

// Eval evaluates the program and returns the Stack as well as any error
func (env *Environment) Eval() (Stack, error) {
	for env.instructionIndex != -1 {
		env.eval()
		if env.error != nil {
			return Stack{}, env.error
		}
	}
	return env.stack, env.error
}

func outOfBounds(arrLen, index int) bool {
	if index < 0 {
		return true
	} else if index > arrLen-1 {
		return true
	}
	return false
}

func (env *Environment) indexOfMatchingRBrac() int {
	n := 0
	for ind := env.instructionIndex + 1; ind < len(env.instructions); ind++ {
		tok := env.instructions[ind]
		if tok.Type == token.RBRAC && n == 0 {
			return ind
		} else if tok.Type == token.LBRAC {
			n++
		} else if tok.Type == token.RBRAC {
			n--
		}
	}
	return -1
}

func (env *Environment) indexOfMatchingLBrac() int {
	n := 0
	for ind := env.instructionIndex - 1; ind >= 0; ind-- {
		tok := env.instructions[ind]
		if tok.Type == token.LBRAC && n == 0 {
			return ind
		} else if tok.Type == token.RBRAC {
			n++
		} else if tok.Type == token.LBRAC {
			n--
		}
	}
	return -1
}

func (env *Environment) eval() {
	tok := env.instructions[env.instructionIndex]

	switch tok.Type {
	case token.LT:
		if outOfBounds(len(env.stack), env.dataIndex-1) {
			env.error = ErrDataPtrOutOfRange
			return
		}
		env.dataIndex--

	case token.GT:
		if outOfBounds(len(env.stack), env.dataIndex+1) {
			env.error = ErrDataPtrOutOfRange
			return
		}
		env.dataIndex++

	case token.ADD:
		env.stack[env.dataIndex]++

	case token.SUB:
		env.stack[env.dataIndex]--

	case token.LBRAC:
		env.lLoop()

	case token.RBRAC:
		env.rLoop()

	case token.DOT:
		if env.out == nil {
			env.error = ErrNoWriteToNil
			return
		}
		env.out.Write([]byte{byte(env.stack[env.dataIndex])})

	case token.COMMA:
		if env.in == nil {
			env.error = ErrNoReadFromNil
			return
		}
		var dat [1]byte
		_, err := env.in.Read(dat[:])
		if err != nil {
			env.error = err
			return
		}
		env.stack[env.dataIndex] = int(dat[0])

	case token.EOF:
		env.instructionIndex = -1
		return

	}

	env.instructionIndex++
}

func (env *Environment) lLoop() {
	if env.stack[env.dataIndex] == 0 {
		if env.instructionIndex = env.indexOfMatchingRBrac(); env.instructionIndex == -1 {
			env.error = ErrNoMatchingBrac
			return
		}
	}
}

func (env *Environment) rLoop() {
	if env.stack[env.dataIndex] != 0 {
		if env.instructionIndex = env.indexOfMatchingLBrac(); env.instructionIndex == -1 {
			env.error = ErrNoMatchingBrac
			return
		}
	}
}

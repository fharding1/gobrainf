[![Go Report Card](https://goreportcard.com/badge/gitlab.com/fharding/gobrainf)](https://goreportcard.com/report/gitlab.com/fharding/gobrainf)

# brainf is a library for brainf**k lexing and evaluating.
## usage:
```go
code := "++>+++"
l := lexer.New(code)
env := eval.New(l, os.Stdin, os.Stdout)
stack, err := env.Eval()
fmt.Println(stack, err)
```
## brainf rules:
* brackets do not wrap, brackets must be matched
* stack is array of 30,000 ints
* start index or 'pointer' is at 0
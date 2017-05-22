package main

import (
	"flag"
	"io/ioutil"
	"os"

	"fmt"

	"gitlab.com/fharding/brainf/eval"
	"gitlab.com/fharding/brainf/lexer"
)

func main() {
	stackFlag := flag.Int("stack", 0, "elements of stack [0-30000] to print")
	flag.Parse()

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	stack, err := eval.New(lexer.New(string(bytes)), os.Stdin, os.Stdout).Eval()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	stackN := *stackFlag
	if stackN != 0 {
		fmt.Println(stack[:stackN])
	}
}

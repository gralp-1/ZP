package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "./example.zp"
	// check if file exists
	// check if file is .pb

	_, err := os.Stat(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File %s doesn't exist\n", filename)
		os.Exit(1)
	}
	if string(filename[len(filename)-3:]) != ".zp" {
		fmt.Fprintf(os.Stderr, "File %s is not a .zp file\n", filename)
		os.Exit(1)
	}
	f_contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File %s could not be read\n", filename)
		os.Exit(1)
	}
	Interpreter := NewInterpreter(string(f_contents)).InitStdLib()
	Interpreter.Interpret()
	// check if file is .zp
}

package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "./example.zp"
	// check if file exists
	// check if file is .zp

	_, err := os.Stat(filename)
	if err != nil {
		Fatal(fmt.Sprintf("File %s doesn't exist\n", filename))
	}
	if string(filename[len(filename)-3:]) != ".zp" {
		Fatal(fmt.Sprintf("File %s is not a .zp file\n", filename))
	}
	f_contents, err := os.ReadFile(filename)
	if err != nil {
		Fatal(fmt.Sprintf("File %s could not be read\n", filename))
	}
	Preprocessor := NewPreprocessor(string(f_contents))
	processed := Preprocessor.Preprocess()
	Interpreter := NewInterpreter(processed).AddLibrary(StdLib())
	Interpreter.Interpret()
}

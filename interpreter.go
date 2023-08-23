package main

import (
	"fmt"
	"os"
	"strings"
)

type Func struct {
	funct func(*Interpreter, []string)
	argc  int
}

type Interpreter struct {
	funcs       map[string]Func
	vars        map[string]string
	content     string
	line_number int
}

func NewInterpreter(content string) *Interpreter {
	return &Interpreter{
		funcs:       map[string]Func{},
		vars:        map[string]string{},
		content:     content,
		line_number: 0,
	}
}

func (self *Interpreter) GetVariable(identifier string) string {
	val, ok := self.vars[identifier]
	if !ok {
		fmt.Fprintf(os.Stderr,"Tried to access non-existant variable %s on line %d\n", identifier, self.line_number)
		os.Exit(1)
	}
	return val
}
func (self *Interpreter) InitStdLib() *Interpreter {
	self.funcs = std_lib()
	return self
}
func (self *Interpreter) Interpret() {
	// start processing the file
	lines := strings.Split(self.content, ";")
	for idx, line := range lines {
		line = strings.Trim(line, "\n")
		// check for comments
		if strings.HasPrefix(line, "//") || len(line) == 0 {
			continue
		}
		self.line_number = idx + 1
		line = strings.TrimSpace(line)
		line = strings.ReplaceAll(line, "\n", "")
		parts := strings.Split(line, " ")
		function := self.funcs[parts[0]]
		args := parts[1:]
		fmt.Printf("Running %s\n", line)
		if len(args) != function.argc {
			fmt.Fprintf(os.Stderr, "Error: expected %d arguments on line %d\n", function.argc, self.line_number)
			os.Exit(1)
		}
		function.funct(self, parts[1:])
	}
}

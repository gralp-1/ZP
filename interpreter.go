package main

import (
	"fmt"
	"maps"
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
	content     []string
	line_number int
}

func NewInterpreter(content []string) *Interpreter {
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
		Fatal(fmt.Sprintf("Tried to access non-existant variable %s on line %d\n", identifier, self.line_number))
	}
	return val
}
func (self *Interpreter) SetVariable(identifier string, val string) {
	// check that the identifier is acceptable
	self.vars[identifier] = val
}
func (self *Interpreter) AddLibrary(library map[string]Func) *Interpreter {
	maps.Copy(self.funcs, library)
	return self
}
func (self *Interpreter) Interpret() {
	for idx, line := range self.content {
		// check for comments
		self.line_number = idx + 1
		parts := strings.Split(line, " ")
		function := self.funcs[parts[0]]
		args := parts[1:]
		if len(args) != function.argc {
			Fatal(fmt.Sprintf("(%s interpreter %s) expected %d arguments on line %d", "", GetFnName(function.funct), function.argc, self.line_number))
			os.Exit(1)
		}
		function.funct(self, parts[1:])
	}
}

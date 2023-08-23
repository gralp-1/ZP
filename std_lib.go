package main

import (
	"fmt"
	"os"
	"strconv"
)

func add(self *Interpreter, inps []string) {
	int_a, err_a := strconv.Atoi(self.GetVariable(inps[1]))
	int_b, err_b := strconv.Atoi(self.GetVariable(inps[2]))
	if err_a != nil || err_b != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err_a)
		os.Exit(1)
	}
	if err_b != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err_b)
	}
	self.vars[inps[0]] = fmt.Sprint(int_a + int_b)
}
func exists(self *Interpreter, inps []string) {
	_, ok := self.vars[inps[0]]
	if ok {
		fmt.Printf("key %s exists on line %d\n", inps[0], self.line_number)
	} else {
		fmt.Printf("key %s doesn't exist on line %d\n", inps[0], self.line_number)
	}
}
func store(self *Interpreter, inps []string) {
	self.vars[inps[0]] = inps[1]
}
func printVal(self *Interpreter, inps []string) {
	fmt.Printf("%s\n", inps[0])
}
func printVar(self *Interpreter, inps []string) {

	fmt.Printf("%s\n", self.GetVariable(inps[0]))
}
func del(self *Interpreter, inps []string) {
	delete(self.vars, inps[0])
}
func stdLib() map[string]Func {
	lib := map[string]Func{}
	Register(lib, store, 2)
	Register(lib, add, 3)
	Register(lib, del, 1)
	Register(lib, exists, 1)
	Register(lib, printVar, 1)
	Register(lib, printVal, 1)

	return lib
}

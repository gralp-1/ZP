package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func add(self *Interpreter, inps []string) {
	int_a, err_a := strconv.ParseFloat(self.GetVariable(inps[1]), 64)
	int_b, err_b := strconv.ParseFloat(self.GetVariable(inps[2]), 64)
	if err_a != nil {
		Fatal(fmt.Sprintf("(std_lib) Error: %s\n", err_a))
		os.Exit(1)
	}
	if err_b != nil {
		Fatal(fmt.Sprintf("(std_lib) Error: %s\n", err_b))
	}
	self.SetVariable(inps[0], fmt.Sprint(int_a+int_b))
}
func sub(self *Interpreter, inps []string) {
	int_a, err_a := strconv.ParseFloat(self.GetVariable(inps[1]), 64)
	int_b, err_b := strconv.ParseFloat(self.GetVariable(inps[2]), 64)
	if err_a != nil || err_b != nil {
		Fatal(fmt.Sprintf("Error: %s\n", err_a))
	}
	if err_b != nil {
		log.Fatal(fmt.Sprintf("Error: %s\n", err_b))
	}
	self.SetVariable(inps[0], fmt.Sprint(int_a-int_b))
}
func mul(self *Interpreter, inps []string) {
	int_a, err_a := strconv.ParseFloat(self.GetVariable(inps[1]), 64)
	int_b, err_b := strconv.ParseFloat(self.GetVariable(inps[2]), 64)
	if err_a != nil || err_b != nil {
		log.Fatal(fmt.Sprintf("Error: %s\n", err_a))
	}
	if err_b != nil {
		log.Fatal(fmt.Sprintf("Error: %s\n", err_b))
	}
	self.SetVariable(inps[0], fmt.Sprint(int_a*int_b))
}

func div(self *Interpreter, inps []string) {
	int_a, err_a := strconv.ParseFloat(self.GetVariable(inps[1]), 64)
	int_b, err_b := strconv.ParseFloat(self.GetVariable(inps[2]), 64)
	if err_a != nil || err_b != nil {
		log.Fatal(fmt.Sprintf("Error: %s\n", err_a))
	}
	if err_b != nil {
		log.Fatal(fmt.Sprintf("Error: %s\n", err_b))
	}
	self.SetVariable(inps[0], fmt.Sprint(int_a/int_b))
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
	self.SetVariable(inps[0], inps[1])
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
func dbg(self *Interpreter, inps []string) {
	val := self.GetVariable(inps[0])
	fmt.Printf("[%s DEBUG line %d]: %s has value %s\n", time.Now().Format("03:04:05"), self.line_number, inps[0], val)
}

func read(self *Interpreter, inps []string) {
	// read <variable> <path>
	fPath := inps[1]
	Info(fPath)
	val, err := os.ReadFile(fPath)
	if err != nil {
		//Fatal(fmt.Sprintf("read instruction failed on line %d", self.line_number))
		fmt.Printf("%s\n", err)
	}
	self.SetVariable(inps[0], string(val))
}

//	func write(self *Interpreter, inps []string) {
//		// write path content
//		fPath := self.GetVariable(inps[0])
//		content := []byte(self.GetVariable(inps[1]))
//		err := os.WriteFile(fPath, content, 0666)
//		if err != nil {
//			fmt.Printf("%s\n", err)
//		}
//	}
func StdLib() map[string]Func {
	lib := map[string]Func{}
	Register(lib, store, 2)
	Register(lib, del, 1)
	Register(lib, exists, 1)
	Register(lib, dbg, 1)

	Register(lib, add, 3)
	Register(lib, sub, 3)
	Register(lib, mul, 3)
	Register(lib, div, 3)

	Register(lib, printVar, 1)
	Register(lib, printVal, 1)

	// Register(lib, read, 2)
	// Register(lib, write, 2)
	return lib
}

package main

import (
	"reflect"
	"runtime"
	"strings"
)

func GetFnName(i interface{}) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	fnName := strings.Split(fullName, ".")[1]
	return fnName
}
func Register(lib map[string]Func, function func(*Interpreter, []string), argc int) {
	lib[GetFnName(function)] = Func{funct: function, argc: argc}
}

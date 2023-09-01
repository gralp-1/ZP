package main

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"time"

	colour "github.com/fatih/color"
)

func GetFnName(i interface{}) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	fnName := strings.Split(fullName, ".")[1]
	return fnName
}
func Register(lib map[string]Func, function func(*Interpreter, []string), argc int) {
	lib[GetFnName(function)] = Func{funct: function, argc: argc}
}
func RemoveBetween(str, start, end string) string {

	anyIncludingEndLine := fmt.Sprintf(`%s[\r\n\s\w]*%s`, start, end)

	return regexp.MustCompile(anyIncludingEndLine).ReplaceAllString(str, "")
}

func DeleteEmptyFromSlice(slice []string) []string {
	var r []string
	for _, str := range slice {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

type LogLevel int

const (
	LevelInfo LogLevel = iota
	LevelWarn
	LevelError
	LevelFatal
)

func logPrefix(level LogLevel) string {
	levelPrefix := ""
	switch level {
	case LevelInfo:
		levelPrefix = "Info"
	case LevelWarn:
		levelPrefix = "Warn"
	case LevelError:
		levelPrefix = "Error"
	case LevelFatal:
		levelPrefix = "Fatal"
	}
	return fmt.Sprintf("[%s %s]: ", time.Now().Format("03:04:05"), levelPrefix)
}

type Log struct {
	level int
}

func Info(info string) {
	colour.Green(logPrefix(LevelInfo) + info)
}
func Warn(info string) {
	colour.Yellow(logPrefix(LevelWarn) + info)
}
func Error(info string) {
	colour.Red(logPrefix(LevelError) + info)
}

func Fatal(info string) {
	colour.Red(logPrefix(LevelFatal) + info)
	os.Exit(1)
}

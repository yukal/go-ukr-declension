package main

import (
	"reflect"
	"runtime"
	"strings"
	// . "github.com/franela/goblin"
)

// go test -v -count=1 ./...						- to run all tests
//
// go test -v -count=1 -run Basic ./...				- to run all tests within Basic group
// go test -v -count=1 -run General ./...			- to run all tests within General group
//
// go test -v -count=1 -run BasicFeminine ./...		- to run all Feminine tests within Basic group
// go test -v -count=1 -run GeneralFeminine ./...	- to run all Feminine tests within General group

// Get function name at runtime
func GetFunctionName(fn interface{}) string {
	fnFullName := strings.Split(
		runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(), ".")

	return fnFullName[len(fnFullName)-1]
	// return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func getLastLetters(word string, count int) string {
	letters := []rune(word)
	return string(letters[len(letters)-count:])
}

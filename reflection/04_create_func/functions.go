package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func main() {
	ourTitle := "the go standard library"
	timed := MakeTimedFunction(properTitle).(func(string) string)
	fmt.Println(timed(ourTitle))

	timedToo := MakeTimedFunction(doubleOurNumber).(func(int) int)
	fmt.Println(timedToo(2))
}

// MakeTimedFunction takes a function interface and times the execution
func MakeTimedFunction(f interface{}) interface{} {
	tf := reflect.TypeOf(f)
	if tf.Kind() != reflect.Func {
		panic("expects a function")
	}

	vf := reflect.ValueOf(f)
	// Create wrapper function via reflection to time the function execution
	wrapperF := reflect.MakeFunc(tf, func(in []reflect.Value) []reflect.Value {
		start := time.Now()
		out := vf.Call(in)
		elapsed := time.Since(start)
		fmt.Printf("Calling %s took %v\n", runtime.FuncForPC(vf.Pointer()).Name(), elapsed)
		return out
	})
	return wrapperF.Interface()
}

func properTitle(input string) string {
	words := strings.Fields(input)
	smallwords := " a an on the to "

	for index, word := range words {
		if !strings.Contains(smallwords, " "+word+" ") {
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}

func doubleOurNumber(a int) int {
	time.Sleep(time.Second)
	return a * 2
}

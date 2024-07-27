package main

import (
	"strconv"
	"syscall/js"
)

// This function is imported from JavaScript, as it doesn't define a body.
// You should define a function named 'add' in the WebAssembly 'env'
// module from JavaScript.
//
//export add
func add(x, y int) int

// This function is exported to JavaScript, so can be called using
// exports.multiply() in JavaScript.
//
//export multiply
func multiply(x, y int) int {
	return x * y
}

//export update
func update() bool {
	document := js.Global().Get("document")
	aStr := document.Call("getElementById", "a").Get("value").String()
	bStr := document.Call("getElementById", "b").Get("value").String()
	a, _ := strconv.Atoi(aStr)
	b, _ := strconv.Atoi(bStr)
	result := add(a, b)
	document.Call("getElementById", "result").Set("value", strconv.Itoa(result))
	return true
}

// This calls a JS function from Go.
func main() {
	println("adding two numbers:", add(2, 3)) // expecting 5
}

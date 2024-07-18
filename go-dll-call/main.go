package main

// #cgo LDFLAGS: -L. ./hello.dll
// #include "hello.h"
import "C"

func main() {
	C.SayHello("Bro")
}

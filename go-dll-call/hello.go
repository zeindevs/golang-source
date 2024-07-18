package main

import "C"
import "fmt"

//export SayHello
func SayHello(name string) {
	fmt.Println("Hello", name)
}

func main() {}

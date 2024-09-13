package main

import "fmt"

func main() {
	msgch := make(chan int)

	go func() {
		msg := <-msgch // eat the cookie
		fmt.Println(msg)
	}()

	msgch <- 10 // blocking
}

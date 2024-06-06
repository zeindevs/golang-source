package main

import (
	"fmt"
	"testing"
)

func TestAddUser(t *testing.T) {
	server := NewServer()
	server.Start()

	//  you are processing work here
	for i := 0; i < 10; i++ {
		go func(i int) {
			server.userch <- (fmt.Sprintf("user_%d", i))
			// server.addUser(
		}(i)
	}
	// you want the work to be done

	fmt.Println("the loop is done!!")
}

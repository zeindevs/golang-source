package main

/*
#include <stdint.h>
*/
import "C"
import (
	"time"
)

var ch = make(chan int)

//export StartChannel
func StartChannel() {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(1 * time.Second)

		}
		close(ch)
	}()
}

//export SubscribeChannel
func SubscribeChannel() C.int {
	val, ok := <-ch
	if !ok {
		return -1
	}
	return C.int(val)
}

func main() {
}

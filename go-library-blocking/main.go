package main

/*
#include <stdint.h>

extern void CallGoFunction();
*/

import "C"
import (
	"fmt"
	"time"
)

//export CallGoFunction
func CallGoFunction() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Go function running: %d\n", i)
		time.Sleep(1 * time.Second) // Blocking
	}
}

func main() {
	select {}
}

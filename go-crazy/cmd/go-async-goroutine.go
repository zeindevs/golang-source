package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func main() {
	go printNumbers()
	printNumbers()
}

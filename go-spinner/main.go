package main

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
	for {
		for _, x := range `-\|/` {
			fmt.Printf("\r%c", x)
			time.Sleep(delay)
		}
	}
}

func waitAndPrint(delay time.Duration) {
	fmt.Println("Working")
	time.Sleep(delay)
	fmt.Println("\nFunction Function Function")
}

func main() {
	go spinner(100 * time.Millisecond)
	waitAndPrint(10 * time.Second)
}

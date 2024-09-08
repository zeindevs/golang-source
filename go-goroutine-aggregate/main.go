package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)
	return "BOB"
}

func fetchUserLikes(userName string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150)
	respch <- 11
	wg.Done()
}

func fetchUserMatch(userName string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	respch <- "ANNA"
	wg.Done()
}

func main() {
	start := time.Now()
	userName := fetchUser()
	respch := make(chan any, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go fetchUserLikes(userName, respch, wg)
	go fetchUserMatch(userName, respch, wg)

	wg.Wait()
	close(respch)

	for msg := range respch {
		fmt.Println("resp:", msg)
	}

	fmt.Println("took:", time.Since(start))
}

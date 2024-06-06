package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	userID := 10
	respch := make(chan string, 100)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go fetchUserData(userID, respch, wg)
	wg.Add(1)
	go fetchUserRecommendations(userID, respch, wg)
	wg.Add(1)
	go fetchUserLikes(userID, respch, wg)
	wg.Wait()

	close(respch)

	for resp := range respch {
		fmt.Println(resp)
	}

	fmt.Println(time.Since(now))
}

func fetchUserData(userID int, respch chan string, wg *sync.WaitGroup) {
	time.Sleep(80 * time.Millisecond)

	respch <- "user data"
	wg.Done()
}

func fetchUserRecommendations(userID int, respch chan string, wg *sync.WaitGroup) {
	time.Sleep(120 * time.Millisecond)

	respch <- "user recommendations"
	wg.Done()
}

func fetchUserLikes(userID int, respch chan string, wg *sync.WaitGroup) {
	time.Sleep(150 * time.Millisecond)

	respch <- "user likes"
	wg.Done()
}

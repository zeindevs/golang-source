package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func doWorkWithValue(ctx context.Context) {
	out := fmt.Sprintf("Doing name work with value: %s", ctx.Value("myKey"))
	fmt.Println(out)
}

func doWorkWithTimeout(ctx context.Context, workCh <-chan int, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("Something went wrong: %s\n", err.Error())
				wg.Done()
			}
			fmt.Printf("worker: finished\n")
		case i := <-workCh:
			fmt.Printf("squaring int: %d -> %d\n", i, i+i)
			wg.Done()
		}
	}
}

func main() {
	ctx := context.WithValue(context.TODO(), "myKey", "Hello Bro")

	doWorkWithValue(ctx)

	workCh := make(chan int, 1)

	ctxTwo, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	wg := &sync.WaitGroup{}

	fmt.Println("starting worker...")

	go doWorkWithTimeout(ctxTwo, workCh, wg)

	ints := []int{2, 4, 5, 6, 11}
	wg.Add(1)

	go func() {
		for i := range ints {
			wg.Add(1)
			workCh <- i
			time.Sleep(time.Second * 1)
		}
	}()

	wg.Wait()
}

package main

import (
	"fmt"
	"go-concurrent-currency/internal/currency"
	"sync"
	"time"
)

func runCurrencyWorker(
	workerId int,
	currencyCh <-chan currency.Currency,
	resultCh chan<- currency.Currency) {
	fmt.Printf("Worker %d started\n", workerId)
	for c := range currencyCh {
		rates, err := currency.FetchCurrencyRates(c.Code)
		if err != nil {
			panic(err)
		}
		c.Rates = rates
		resultCh <- c
	}
	fmt.Printf("Worker %d stopped\n", workerId)
}

func main() {
	cre := currency.MyCurrencyExchange{
		Currencies: make(map[string]currency.Currency),
	}
	err := cre.FetchAllCurrencies()
	if err != nil {
		panic(err)
	}

	currencyCh := make(chan currency.Currency, len(cre.Currencies))
	resultCh := make(chan currency.Currency, len(cre.Currencies))

	for i := 0; i < 5; i++ {
		go runCurrencyWorker(i, currencyCh, resultCh)
	}

	startTime := time.Now()

	resultCount := 0

	for _, curr := range cre.Currencies {
		currencyCh <- curr
	}

	for {
		if resultCount == len(cre.Currencies) {
			fmt.Println("Closing resultCh")
			close(currencyCh)
			break
		}
		select {
		case c := <-resultCh:
			cre.Currencies[c.Code] = c
			resultCount++
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout")
			return
		}
	}

	endTime := time.Now()

	fmt.Println("===== Results =====")
	for _, curr := range cre.Currencies {
		fmt.Printf("%s (%s): %d rates\n", curr.Name, curr.Code, len(curr.Rates))
	}
	fmt.Println("===================")
	fmt.Println("Time taken:", endTime.Sub(startTime))
}

func withMutex() {
	cre := currency.MyCurrencyExchange{
		Currencies: make(map[string]currency.Currency),
	}
	err := cre.FetchAllCurrencies()
	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}

	startTime := time.Now()

	go func() {
		for {
			cre.Lock()
			usd, ok := cre.Currencies["usd"]
			cre.Unlock()
			if ok {
				fmt.Println("USD:", usd.Rates)
			}
		}
	}()

	for code := range cre.Currencies {
		wg.Add(1)
		go func(code string) {
			rates, err := currency.FetchCurrencyRates(code)
			if err != nil {
				panic(err)
			}
			cre.Lock()
			cre.Currencies[code] = currency.Currency{
				Code:  code,
				Name:  cre.Currencies[code].Name,
				Rates: rates,
			}
			cre.Unlock()
			wg.Done()
		}(code)
	}

	wg.Wait()

	endTime := time.Now()

	fmt.Println("===== Results =====")
	for _, curr := range cre.Currencies {
		fmt.Printf("%s (%s): %d rates\n", curr.Name, curr.Code, len(curr.Rates))
	}
	fmt.Println("===================")
	fmt.Println("Time taken:", endTime.Sub(startTime))
}

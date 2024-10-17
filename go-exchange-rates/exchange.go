package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

var conversionError = errors.New("Error calculation exchange rate.")

type GetRatesFunc func(date string) map[string]float64

func CalcExchangeDate(value float64, base string, conv string, date string, getter GetRatesFunc) (float64, error) {
	var rates map[string]float64 = getter(date)
	if rates == nil {
		return -1, conversionError
	}

	baseRate, ok1 := rates[base]
	convRate, ok2 := rates[conv]
	if !ok1 || !ok2 {
		return -1, conversionError
	}

	net := value * (convRate / baseRate)

	return math.Round(net*100) / 100, nil
}

func main() {
	rate, err := CalcExchangeDate(100, "USD", "IDR", "2024-10-06", getLocalExchangeRatesUSD)
	// rate, err := CalcExchangeDate(10000000, "IDR", "USD", "2024-10-06", getLocalExchangeRatesUSD)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Rate USD to IDR:", int64(rate))
}

package main

import "fmt"

type Candle struct {
	open  int
	close int
	high  int
	low   int
}

func (c *Candle) reset() {
	c.open = 0
	c.close = 0
	c.high = 0
	c.low = 0
}

func main() {
	candles := map[int]*Candle{}
	count := 5
	for i := 0; i < count; i++ {
		candles[i] = &Candle{open: i}
	}
	for _, candle := range candles {
		candle.reset()
		// candles[i] = &Candle{
		// 	open: 69,
		// }
		// candle = &Candle{
		// 	open: 69,
		// }
		// more logic to do here
		fmt.Println(candle)
	}
	for _, candle := range candles {
		fmt.Println(candle)
	}
}

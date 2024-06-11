package orderbook

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewOrderbook(t *testing.T) {
	asksFile := "../data/asks.gob"
	bidsFile := "../data/bids.gob"
	ob, err := NewOrderbookFromFile("BTCUSDT", asksFile, bidsFile)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("asks volume:", ob.totalAskVolume())
	fmt.Println("bids volume:", ob.totalBidsVolume())

	for price, limit := range ob.asks.limits {
		fmt.Printf("%.2f -> %+v\n", price, limit)
	}
}

func TestLimitFillMultiOrder(t *testing.T) {
	l := NewLimit(10000)
	askOrderA := NewAskOrder(10)
	askOrderB := NewAskOrder(5)
	l.addOrder(askOrderA)
	l.addOrder(askOrderB)

	marketOrder := NewAskOrder(12)
	l.fillOrder(marketOrder)
	assert.Equal(t, 3.0, l.totalVolume)

	assert.Equal(t, 1, len(l.orders))
	assert.Equal(t, 3.0, l.orders[0].size)
	assert.True(t, marketOrder.isFilled())

	fmt.Printf("%+v\n", l)
}

func TestLimitFillSingleOrder(t *testing.T) {

}

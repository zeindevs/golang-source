package orderbook

import (
	"encoding/gob"
	"os"
)

type ByBestAsk struct{ LimitMap }

func (ba ByBestAsk) Len() int { return len(ba.LimitMap.limits) }

type LimitMap struct {
	isBids      bool
	limits      map[float64]*Limit
	totalVolume float64
}

func NewLimitMap(isBids bool) *LimitMap {
	return &LimitMap{
		isBids: isBids,
		limits: make(map[float64]*Limit),
	}
}

type AskMap struct {
	limits      map[float64]*Limit
	totalVolume float64
}

func NewAskMap() *AskMap {
	return &AskMap{
		limits: make(map[float64]*Limit),
	}
}

func (m *LimitMap) loadFromFile(src string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}

	var asks map[float64]float64
	if err := gob.NewDecoder(f).Decode(&asks); err != nil {
		return err
	}

	for price, size := range asks {
		l := NewLimit(price)
		l.totalVolume += size
		m.limits[price] = l
		m.totalVolume += size

	}

	return nil
}

type Orderbook struct {
	ticker string
	asks   *LimitMap
	bids   *LimitMap
	size   int
}

func NewOrderbookFromFile(ticker, askSrc, bidSrc string) (*Orderbook, error) {
	asks := NewLimitMap(false)
	if err := asks.loadFromFile(askSrc); err != nil {
		return nil, err
	}

	bids := NewLimitMap(true)
	if err := bids.loadFromFile(askSrc); err != nil {
		return nil, err
	}

	return &Orderbook{
		ticker: ticker,
		asks:   asks,
		bids:   bids,
	}, nil
}

func NewOrderbook(ticker string) *Orderbook {
	return &Orderbook{
		ticker: ticker,
	}
}

func (ob *Orderbook) totalAskVolume() float64 {
	return ob.asks.totalVolume
}

func (ob *Orderbook) totalBidsVolume() float64 {
	return ob.bids.totalVolume
}

type Limit struct {
	price       float64
	totalVolume float64
	orders      []*Orderbook
}

func NewLimit(price float64) *Limit {
	return &Limit{
		price:  price,
		orders: []*Orderbook{},
	}
}

func NewAskOrder(price int) *Orderbook {
	return &Orderbook{}
}

func (l *Limit) addOrder(order *Orderbook) {

}

func (l *Limit) fillOrder(order *Orderbook) {

}

func (ob *Orderbook) isFilled() bool {
	return false
}

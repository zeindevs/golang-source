package main

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

var prices = map[string]float64{
	"BTC": 29_000.0,
	"ETH": 200.0,
	"GG":  100_000.0,
}

// PriceService is an interface that can fetch the price any given ticker
type PriceService interface {
	FetchPrice(ctx context.Context, ticker string) (float64, error)
}

type priceService struct{}

// is the business logic
func (s *priceService) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	price, ok := prices[ticker]
	if !ok {
		return 0.0, fmt.Errorf("the given ticker (%s) is not supported", ticker)
	}
	return price, nil
}

type loggingService struct {
	next PriceService
}

func NewLoggingService(next PriceService) PriceService {
	return &loggingService{
		next: next,
	}
}

func (s *loggingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"requestID": ctx.Value("requestID"),
			"took":      time.Since(begin),
			"err":       err,
			"price":     price,
		}).Info("fetchPrice")
	}(time.Now())

	return s.next.FetchPrice(ctx, ticker)
}

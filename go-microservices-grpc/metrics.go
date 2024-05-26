package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type metricsService struct {
	next PriceService
}

func NewMetricsService(next PriceService) PriceService {
	return &metricsService{
		next: next,
	}
}

func (s *metricsService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"requestID": ctx.Value("requestID"),
			"took":      time.Since(begin),
			"err":       err,
			"price":     price,
		}).Info("pushing metrics to prometheus")
	}(time.Now())

	// your metrics storage, Push to prometheus (gauage, counters)
	return s.next.FetchPrice(ctx, ticker)
}

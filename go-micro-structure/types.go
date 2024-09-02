package main

import (
	"context"
	"fmt"
	"time"
)

type CatFact struct {
	Fact string `json:"fact"`
}

func NewLoggingService(next Service) Service {
	return &LoggingService{
		next: next,
	}
}

func (s *LoggingService) GetCatFact(ctx context.Context) (fact *CatFact, err error) {
	defer func(start time.Time) {
		fmt.Printf("fact=%s err=%s took=%v\n", fact.Fact, err, time.Since(start))
	}(time.Now())
	return s.next.GetCatFact(ctx)
}

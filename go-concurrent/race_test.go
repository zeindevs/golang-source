package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestDataRaceCondition(t *testing.T) {
	var state int32

	for i := 0; i < 10; i++ {
		go func(i int) {
			state += int32(i)
		}(i)
	}
}

func TestDataRaceConditionMutex(t *testing.T) {
	var state int32
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		go func(i int) {
			mu.Lock()
			state += int32(i)
			mu.Unlock()
		}(i)
	}
}

func TestDataRaceConditionAtomic(t *testing.T) {
	var state int32

	for i := 0; i < 10; i++ {
		go func(i int) {
			atomic.AddInt32(&state, int32(i))
		}(i)
	}
}

type server struct {
	gamerount atomic.Value
}

func TestDataRaceConditionServer(t *testing.T) {
	s := server{
		gamerount: atomic.Value{},
	}

	for i := 0; i < 10; i++ {
		go func(i int) {
			s.gamerount.Store(i)
		}(i)
	}
}

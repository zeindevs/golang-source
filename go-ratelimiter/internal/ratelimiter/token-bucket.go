package ratelimiter

import (
	"log"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type TokenBucketLimiter struct {
	sync.RWMutex
	clients map[string]*Client
	config  Config
}

func NewTokenBucketLimiter(cfg Config) *TokenBucketLimiter {
	rl := &TokenBucketLimiter{
		clients: make(map[string]*Client),
		config:  cfg,
	}

	go rl.cleanupClients()

	return rl
}

type Client struct {
	Limiter  *rate.Limiter
	LastSeen time.Time
}

func (rl *TokenBucketLimiter) getClient(ip string) *Client {
	rl.RLock()
	_, exists := rl.clients[ip]
	rl.Unlock()

	if !exists {
		limiter := rate.NewLimiter(rate.Every(rl.config.TimeFrame), rl.config.RequestPerTimeFrame)
		rl.Lock()

		rl.clients[ip] = &Client{
			Limiter:  limiter,
			LastSeen: time.Now(),
		}

		rl.Unlock()
	}

	v := rl.clients[ip]
	v.LastSeen = time.Now()

	return v
}

func (rl *TokenBucketLimiter) Allow(ip string) (bool, float64) {
	client := rl.getClient(ip)

	allowed := client.Limiter.Allow()
	tokens := client.Limiter.Tokens()

	return allowed, tokens
}

func (rl *TokenBucketLimiter) cleanupClients() {
	for {
		time.Sleep(time.Minute)
		log.Println("Cleaning up clients", rl.clients)

		rl.Lock()
		for ip, v := range rl.clients {
			if time.Since(v.LastSeen) > 3*time.Minute {
				delete(rl.clients, ip)
			}

			rl.Unlock()
		}
	}
}

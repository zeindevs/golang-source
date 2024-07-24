package config

import (
	_ "github.com/joho/godotenv/autoload"
)

import "os"

type Config struct {
	CLIENT_ID     string
	CLIENT_SECRET string
}

func NewConfig() *Config {
	return &Config{
		CLIENT_ID:     os.Getenv("CLIENT_ID"),
		CLIENT_SECRET: os.Getenv("CLIENT_SECRET"),
	}
}

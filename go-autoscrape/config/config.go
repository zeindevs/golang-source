package config

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"github.com/zeindevs/goautoscrape/util"
)

type Config struct {
	Interval int
}

func NewConfig() *Config {
	interval, err := strconv.Atoi(os.Getenv("INTERVAL"))
	if err != nil {
		util.ErrorPanic(err)
	}
	return &Config{
		Interval: interval,
	}
}

func (Config) Get(key string) string {
	return os.Getenv(key)
}

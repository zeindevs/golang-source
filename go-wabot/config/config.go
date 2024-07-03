package config

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"github.com/zeindevs/gowabot/util"
)

type Config struct {
	Public  bool
	BotName string
	Owner   string
	Phone   string
}

func NewConfig() *Config {
	public, err := strconv.ParseBool(os.Getenv("BOTMODE"))
	if err != nil {
		util.ErrorPanic(err)
	}

	return &Config{
		Public:  public,
		BotName: os.Getenv("BOTNAME"),
		Owner:   os.Getenv("WAOWNER"),
		Phone:   os.Getenv("WAPHONE"),
	}
}

func (Config) Get(key string) string {
	return os.Getenv(key)
}

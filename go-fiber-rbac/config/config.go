package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func New(log *logrus.Logger) *viper.Viper {
	cfg := viper.New()
	cfg.SetConfigName(".env")
	cfg.SetConfigType("env")
	cfg.AddConfigPath(".")
	if err := cfg.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	return cfg
}

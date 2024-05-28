package config

import "github.com/spf13/viper"

type Config struct {
	// MySQL Setup
	DBHost     string `mapstructure:"SQL_HOST"`
	DBUsername string `mapstructure:"SQL_USER"`
	DBPassword string `mapstructure:"SQL_PASSWORD"`
	DBName     string `mapstructure:"SQL_DB"`
	DBPort     int    `mapstructure:"SQL_PORT"`

	// Redis Setup
	RedisUrl string `mapstructure:"REDIS_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	// Handle null
	err = viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)

	// after this lets make database connection
	return config, nil
}

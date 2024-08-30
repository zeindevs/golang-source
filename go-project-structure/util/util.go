package util

import "os"

type Map map[string]any

func IsEnvProd() bool {
	return os.Getenv("EN") == "production"
}

func AppEnv() string {
	return os.Getenv("ENV")
}

package main

import (
	"fmt"

	"github.com/zeindevs/gospotify/config"
	"github.com/zeindevs/gospotify/internal"
)

func main() {
	cfg := config.NewConfig()
	auth := internal.NewAuthService(cfg)

	res, err := auth.ClientLogin()
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

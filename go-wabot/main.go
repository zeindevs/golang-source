package main

import (
	"github.com/zeindevs/gowabot/config"
	"github.com/zeindevs/gowabot/internal"
	"github.com/zeindevs/gowabot/pkg"
	"github.com/zeindevs/gowabot/util"
)

func main() {
	cfg := config.NewConfig()
	client := internal.NewWhatsapp(cfg)

	err := client.QRConnect()
	util.ErrorPanic(err)

	client.RegisterCommand(pkg.CommandMenu())
	client.RegisterCommand(pkg.CommandPing())
	client.RegisterCommand(pkg.CommandSource())
	client.RegisterCommand(pkg.CommandMode(cfg))

	client.Wait()
	client.Disconnect()
}

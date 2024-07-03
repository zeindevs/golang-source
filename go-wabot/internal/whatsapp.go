package internal

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal"
	"github.com/zeindevs/gowabot/config"
	"github.com/zeindevs/gowabot/util"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waCompanionReg"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"
)

type Whatsapp struct {
	Config  *config.Config
	Client  *whatsmeow.Client
	Handler *Handler
}

func NewWhatsapp(cfg *config.Config) Whatsapp {
	store.DeviceProps.PlatformType = waCompanionReg.DeviceProps_SAFARI.Enum()
	store.DeviceProps.Os = proto.String(cfg.BotName)

	dbLog := waLog.Stdout("Database", cfg.Get("LOGGING"), true)
	container, err := sqlstore.New("sqlite3", "file:store.db?_foreign_keys=on", dbLog)
	util.ErrorPanic(err)

	handler := NewHandler(container)
	handler.Config = cfg

	client := handler.NewClient()
	client.PrePairCallback = func(jid types.JID, platform, businessName string) bool {
		log.Println("Connected")
		return true
	}

	return Whatsapp{
		Config:  cfg,
		Client:  client,
		Handler: handler,
	}
}

func (wc Whatsapp) QRConnect() error {
	if wc.Client.Store.ID == nil {
		qrChan, err := wc.Client.GetQRChannel(context.Background())
		if err != nil {
			return err
		}

		if err := wc.Client.Connect(); err != nil {
			return err
		}

		for evt := range qrChan {
			if evt.Event == "code" {
				log.Println("QT code:")
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
			} else {
				log.Println("Login event:", evt.Event)
			}
		}

		return nil
	}

	return wc.Client.Connect()
}

func (wc Whatsapp) PhoneConnect() error {
	if err := wc.Client.Connect(); err != nil {
		return err
	}

	code, err := wc.Client.PairPhone(wc.Config.Phone, true, whatsmeow.PairClientChrome, "Chrome (Linux)")
	if err != nil {
		return err
	}

	log.Println("Your code:", code)
	return nil
}

func (wc Whatsapp) Wait() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	log.Println("SIGTERM received")
}

func (wc Whatsapp) Disconnect() {
	log.Println("Disconnect...")
	wc.Client.Disconnect()
}

func (wc Whatsapp) RegisterCommand(cmd *Command) {
	wc.Handler.AddCommand(cmd)
}

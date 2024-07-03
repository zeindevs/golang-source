package internal

import (
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
)

type WAClient struct {
	WA       *whatsmeow.Client
	Commands []*Command
}

type Message struct {
	Info            types.MessageInfo
	From            types.JID
	IsBot           bool
	Sender          types.JID
	OwnerNumber     []string
	PushName        string
	IsOwner         bool
	IsGroup         bool
	Query           string
	Body            string
	Command         string
	IsImage         bool
	IsVideo         bool
	IsQuotedImage   bool
	IsQuotedVideo   bool
	IsQuotedSticker bool
	IsAdmin         bool
	IsBotAdmin      bool
	Media           whatsmeow.DownloadableMessage
	ID              *waE2E.ContextInfo
	QuotedMsg       *waE2E.ContextInfo
	Reply           func(text string, opts ...whatsmeow.SendRequestExtra) (whatsmeow.SendResponse, error)
	React           func(text string) (whatsmeow.SendResponse, error)
}

type Command struct {
	Name        string
	As          []string
	Description string
	Tags        string
	IsPrefix    bool
	IsOwner     bool
	IsMedia     bool
	IsQuery     bool
	IsGroup     bool
	IsAdmin     bool
	IsBotAdmin  bool
	IsWaitt     bool
	IsPrivate   bool
	After       func(client *WAClient, m *Message)
	Exec        func(client *WAClient, m *Message)
}

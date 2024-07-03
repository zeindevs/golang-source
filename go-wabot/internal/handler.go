package internal

import (
	"context"
	"regexp"
	"strings"

	"github.com/zeindevs/gowabot/config"
	"github.com/zeindevs/gowabot/util"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"
)

type Handler struct {
	Config    *config.Config
	Container *store.Device
	Client    *WAClient
}

func NewHandler(container *sqlstore.Container) *Handler {
	deviceStore, err := container.GetFirstDevice()
	util.ErrorPanic(err)

	return &Handler{
		Container: deviceStore,
		Client:    &WAClient{},
	}
}

func (h *Handler) AddCommand(cmd *Command) {
	h.Client.Commands = append(h.Client.Commands, cmd)
}

func (h *Handler) NewClient(makebot ...bool) *whatsmeow.Client {
	clientLog := waLog.Stdout("Client", h.Config.Get("LOGGING"), true)
	client := whatsmeow.NewClient(h.Container, clientLog)
	h.Client = NewClient(client)
	client.AddEventHandler(h.RegisterHandler(makebot...))
	return client
}

func (h *Handler) RegisterHandler(makebot ...bool) func(evt interface{}) {
	return func(evt interface{}) {
		switch v := evt.(type) {
		case *events.Message:
			m := h.MessageParser(v, makebot...)
			if !h.Config.Public && !m.IsOwner {
				return
			}
			go h.Handler(h.Client, m)
			break
		}
	}
}

func (h *Handler) MessageParser(mess *events.Message, makebot ...bool) *Message {
	var command string
	var media whatsmeow.DownloadableMessage
	var isOwner = false
	var owner []string

	quotedMsg := mess.Message.GetExtendedTextMessage().GetContextInfo().GetQuotedMessage()
	owner = append(owner, h.Config.Owner)
	if makebot == nil {
		owner = append(owner, h.Client.WA.Store.ID.ToNonAD().String())
	}

	for _, own := range owner {
		if own == mess.Info.Sender.ToNonAD().String() {
			isOwner = true
		}
	}

	if pe := mess.Message.GetExtendedTextMessage().GetText(); pe != "" {
		command = pe
	} else if pe := mess.Message.GetImageMessage().GetCaption(); pe != "" {
		command = pe
	} else if pe := mess.Message.GetVideoMessage().GetCaption(); pe != "" {
		command = pe
	} else if pe := mess.Message.GetConversation(); pe != "" {
		command = pe
	}

	if quotedMsg != nil && (quotedMsg.ImageMessage != nil || quotedMsg.VideoMessage != nil || quotedMsg.StickerMessage != nil) {
		if msg := quotedMsg.GetImageMessage(); msg != nil {
			media = msg
		} else if msg := quotedMsg.GetVideoMessage(); msg != nil {
			media = msg
		} else if msg := quotedMsg.GetStickerMessage(); msg != nil {
			media = msg
		}
	} else if mess.Message != nil && (mess.Message.ImageMessage != nil || mess.Message.VideoMessage != nil) {
		if msg := mess.Message.GetImageMessage(); msg != nil {
			media = msg
		} else if msg := mess.Message.GetVideoMessage(); msg != nil {
			media = msg
		}
	} else {
		media = nil
	}

	if strings.HasPrefix(command, "@"+h.Client.WA.Store.ID.ToNonAD().User) {
		command = strings.Trim(strings.Replace(command, "@"+h.Client.WA.Store.ID.ToNonAD().User, "", 1), " ")
	}

	return &Message{
		Info:        mess.Info,
		From:        mess.Info.Chat,
		Sender:      mess.Info.Sender,
		PushName:    mess.Info.PushName,
		OwnerNumber: owner,
		IsOwner:     isOwner,
		IsBot:       mess.Info.IsFromMe,
		IsGroup:     mess.Info.IsGroup,
		Query:       strings.Join(strings.Split(command, " ")[1:], ` `),
		Body:        command,
		Command:     strings.ToLower(strings.Split(command, " ")[0]),
		Media:       media,
		IsImage: func() bool {
			return mess.Message.GetImageMessage() != nil
		}(),
		IsAdmin: func() bool {
			if !mess.Info.IsGroup {
				return false
			}
			admin, err := h.Client.FetchGroupAdmin(mess.Info.Chat)
			if err != nil {
				return false
			}
			for _, v := range admin {
				if v == mess.Info.Sender.String() {
					return true
				}
			}
			return false
		}(),
		IsBotAdmin: func() bool {
			if !mess.Info.IsGroup {
				return false
			}
			admin, err := h.Client.FetchGroupAdmin(mess.Info.Chat)
			if err != nil {
				return false
			}
			for _, v := range admin {
				if v == h.Client.WA.Store.ID.ToNonAD().String() {
					return true
				}
			}
			return false
		}(),
		QuotedMsg: mess.Message.GetExtendedTextMessage().GetContextInfo(),
		ID: &waE2E.ContextInfo{
			StanzaID:      &mess.Info.ID,
			Participant:   proto.String(mess.Info.Sender.String()),
			QuotedMessage: mess.Message,
		},
		IsQuotedImage: func() bool {
			return quotedMsg.GetImageMessage() != nil
		}(),
		IsQuotedSticker: func() bool {
			return quotedMsg.GetStickerMessage() != nil
		}(),
		Reply: func(text string, opts ...whatsmeow.SendRequestExtra) (whatsmeow.SendResponse, error) {
			ok, err := h.Client.SendText(mess.Info.Chat, text, &waE2E.ContextInfo{
				StanzaID:      &mess.Info.ID,
				Participant:   proto.String(mess.Info.Sender.String()),
				QuotedMessage: mess.Message,
			}, opts...)

			return *ok, err
		},
		React: func(react string) (whatsmeow.SendResponse, error) {
			return h.Client.WA.SendMessage(context.Background(), mess.Info.Chat, h.Client.WA.BuildReaction(mess.Info.Chat, mess.Info.Sender, mess.Info.ID, react))
		},
	}
}

func (h *Handler) Handler(c *WAClient, m *Message) {
	var prefix string

	pattern := regexp.MustCompile(`[?!.#]`)
	for _, f := range pattern.FindAllString(m.Command, -1) {
		prefix = f
	}

	for _, cmd := range h.Client.Commands {
		if cmd.After != nil {
			cmd.After(c, m)
		}

		re := regexp.MustCompile(`^` + cmd.Name + `$`)
		if valid := len(re.FindAllString(strings.ReplaceAll(m.Command, prefix, ""), -1)) > 0; valid {
			var cmdWithPref bool
			var cmdWithoutPref bool

			if cmd.IsPrefix && (prefix != "" && strings.HasPrefix(m.Command, prefix)) {
				cmdWithPref = true
			} else {
				cmdWithPref = false
			}

			if !cmd.IsPrefix {
				cmdWithoutPref = true
			} else {
				cmdWithoutPref = false
			}

			if !cmdWithPref && !cmdWithoutPref {
				continue
			}

			if cmd.IsOwner && !m.IsOwner {
				m.Reply("This command is only for owner!")
				continue
			}

			if cmd.IsMedia && m.Media == nil {
				m.Reply("Please use command with added media!")
				continue
			}

			if cmd.IsQuery && m.Query == "" {
				m.Reply("Please use command with added query!")
				continue
			}

			if cmd.IsGroup && !m.IsGroup {
				m.Reply("This command is only for groups!")
				continue
			}

			if cmd.IsPrivate && m.IsGroup {
				m.Reply("This command is only for private!")
				continue
			}

			if (m.IsGroup && cmd.IsAdmin) && !m.IsAdmin {
				m.Reply("This command is for group admins only!")
				continue
			}

			if (m.IsGroup && cmd.IsBotAdmin) && !m.IsBotAdmin {
				m.Reply("Before using this command, please make the bot an admin!")
				continue
			}

			if cmd.IsWaitt {
				m.Reply("The request is being processed!")
			}

			cmd.Exec(c, m)
		}
	}
}

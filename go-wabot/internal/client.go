package internal

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"go.mau.fi/util/random"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waCommon"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

func NewClient(client *whatsmeow.Client) *WAClient {
	return &WAClient{
		WA:       client,
		Commands: []*Command{},
	}
}

func (client *WAClient) SendText(from types.JID, txt string, opts *waE2E.ContextInfo, optn ...whatsmeow.SendRequestExtra) (*whatsmeow.SendResponse, error) {
	ok, err := client.WA.SendMessage(context.Background(), from, &waE2E.Message{
		ExtendedTextMessage: &waE2E.ExtendedTextMessage{
			Text:        proto.String(txt),
			ContextInfo: opts,
		},
	}, optn...)

	return &ok, err
}

func (client *WAClient) SendWithNewsLetter(from types.JID, text string, newjid string, newserver int32, name string, opts *waE2E.ContextInfo) (*whatsmeow.SendResponse, error) {
	ok, err := client.SendText(from, text, &waE2E.ContextInfo{
		ForwardedNewsletterMessageInfo: &waE2E.ContextInfo_ForwardedNewsletterMessageInfo{
			NewsletterJID:     proto.String(newjid),
			NewsletterName:    proto.String(name),
			ServerMessageID:   proto.Int32(newserver),
			ContentType:       waE2E.ContextInfo_ForwardedNewsletterMessageInfo_UPDATE.Enum(),
			AccessibilityText: proto.String(""),
		},
		IsForwarded:   proto.Bool(true),
		StanzaID:      opts.StanzaID,
		Participant:   opts.Participant,
		QuotedMessage: opts.QuotedMessage,
	})

	return ok, err
}

func (client *WAClient) SendImage(from types.JID, data []byte, caption string, opts *waE2E.ContextInfo) (*whatsmeow.SendResponse, error) {
	uploaded, err := client.WA.Upload(context.Background(), data, whatsmeow.MediaImage)
	if err != nil {
		log.Printf("Failed to upload file: %v\n", err)
		return nil, err
	}

	resultImg := &waE2E.Message{
		ImageMessage: &waE2E.ImageMessage{
			URL:           proto.String(uploaded.URL),
			DirectPath:    proto.String(uploaded.DirectPath),
			MediaKey:      uploaded.MediaKey,
			Caption:       proto.String(caption),
			Mimetype:      proto.String(http.DetectContentType(data)),
			FileEncSHA256: uploaded.FileEncSHA256,
			FileSHA256:    uploaded.FileSHA256,
			FileLength:    proto.Uint64(uint64(len(data))),
			ContextInfo:   opts,
		},
	}
	ok, _ := client.WA.SendMessage(context.Background(), from, resultImg)

	return &ok, nil
}

func (client *WAClient) SendVideo(from types.JID, data []byte, caption string, opts *waE2E.ContextInfo) (*whatsmeow.SendResponse, error) {
	uploaded, err := client.WA.Upload(context.Background(), data, whatsmeow.MediaVideo)
	if err != nil {
		log.Printf("Failed to upload file: %v\n", err)
		return nil, err
	}

	resultVideo := &waE2E.Message{
		VideoMessage: &waE2E.VideoMessage{
			URL:           proto.String(uploaded.URL),
			DirectPath:    proto.String(uploaded.DirectPath),
			MediaKey:      uploaded.MediaKey,
			Caption:       proto.String(caption),
			Mimetype:      proto.String(http.DetectContentType(data)),
			FileEncSHA256: uploaded.FileEncSHA256,
			FileSHA256:    uploaded.FileSHA256,
			FileLength:    proto.Uint64(uint64(len(data))),
			ContextInfo:   opts,
		},
	}
	ok, err := client.WA.SendMessage(context.Background(), from, resultVideo)

	return &ok, err
}

func (client *WAClient) SendDocument(from types.JID, data []byte, fileName string, caption string, opts *waE2E.ContextInfo) (*whatsmeow.SendResponse, error) {
	uploaded, err := client.WA.Upload(context.Background(), data, whatsmeow.MediaDocument)
	if err != nil {
		log.Printf("Failed to upload file: %v\n", err)
		return nil, err
	}

	resultDoc := &waE2E.Message{
		DocumentMessage: &waE2E.DocumentMessage{
			URL:           proto.String(uploaded.URL),
			DirectPath:    proto.String(uploaded.DirectPath),
			MediaKey:      uploaded.MediaKey,
			FileName:      proto.String(fileName),
			Caption:       proto.String(caption),
			Mimetype:      proto.String(http.DetectContentType(data)),
			FileEncSHA256: uploaded.FileEncSHA256,
			FileSHA256:    uploaded.FileSHA256,
			FileLength:    proto.Uint64(uint64(len(data))),
			ContextInfo:   opts,
		},
	}
	ok, err := client.WA.SendMessage(context.Background(), from, resultDoc)

	return &ok, err
}

func (client *WAClient) SendSticker(jid types.JID, data []byte, opts *waE2E.ContextInfo) (*whatsmeow.SendResponse, error) {
	uploaded, err := client.WA.Upload(context.Background(), data, whatsmeow.MediaImage)
	if err != nil {
		return nil, fmt.Errorf("Failed to upload file: %v\n", err)
	}

	ok, err := client.WA.SendMessage(context.Background(), jid, &waE2E.Message{
		StickerMessage: &waE2E.StickerMessage{
			URL:           proto.String(uploaded.URL),
			DirectPath:    proto.String(uploaded.DirectPath),
			MediaKey:      uploaded.MediaKey,
			Mimetype:      proto.String(http.DetectContentType(data)),
			FileEncSHA256: uploaded.FileEncSHA256,
			FileSHA256:    uploaded.FileSHA256,
			FileLength:    proto.Uint64(uint64(len(data))),
			ContextInfo:   opts,
		},
	})

	return &ok, err
}

func (client *WAClient) UploadImage(data []byte) (string, error) {
	bodyy := &bytes.Buffer{}
	writer := multipart.NewWriter(bodyy)
	part, _ := writer.CreateFormFile("file", "file")
	_, err := io.Copy(part, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	writer.Close()

	// Create request
	req, err := http.NewRequest("POST", "https://telegra.ph/upload", bodyy)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send request and handle response
	htt := &http.Client{}
	resp, err := htt.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("HTTP Error: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var uploads []struct {
		Path string `json:"src"`
	}
	if err := json.Unmarshal(body, &uploads); err != nil {
		m := map[string]string{}
		if err := json.Unmarshal(data, &m); err != nil {
			return "", err
		}

		return "", fmt.Errorf("telegraph: %s", m["error"])
	}

	return "https://telegra.ph/" + uploads[0].Path, nil
}

func (client *WAClient) DeleteMsg(from types.JID, id string, me bool) (*whatsmeow.SendResponse, error) {
	ok, err := client.WA.SendMessage(context.Background(), from, &waE2E.Message{
		ProtocolMessage: &waE2E.ProtocolMessage{
			Type: waE2E.ProtocolMessage_REVOKE.Enum(),
			Key: &waCommon.MessageKey{
				FromMe: proto.Bool(me),
				ID:     proto.String(id),
			},
		},
	})

	return &ok, err
}

func (client *WAClient) FetchGroupAdmin(Jid types.JID) ([]string, error) {
	var Admin []string
	resp, err := client.WA.GetGroupInfo(Jid)
	if err != nil {
		return Admin, err
	}

	for _, group := range resp.Participants {
		if group.IsAdmin || group.IsSuperAdmin {
			Admin = append(Admin, group.JID.String())
		}
	}

	return Admin, err
}

func (client *WAClient) ParseJID(arg string) (*types.JID, bool) {
	if arg[0] == '+' {
		arg = arg[1:]
	}

	if !strings.ContainsRune(arg, '@') {
		jid := types.NewJID(arg, types.DefaultUserServer)
		return &jid, true
	}

	recipient, err := types.ParseJID(arg)
	if err != nil {
		return &recipient, false
	} else if recipient.User == "" {
		return &recipient, false
	}

	return &recipient, true
}

func (client *WAClient) GenerateMessageID(cust string) types.MessageID {
	data := make([]byte, 8, 8+20+16)
	binary.BigEndian.PutUint64(data, uint64(time.Now().Unix()))
	data = append(data, random.Bytes(16)...)
	hash := sha256.Sum256(data)

	return cust + strings.ToUpper(hex.EncodeToString(hash[:12])) + "NM4O"
}

func (client *WAClient) GetBytes(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return bytes, nil
}

func (client *WAClient) AddCommand(cmd *Command) {
	client.Commands = append(client.Commands, cmd)
}

func (client *WAClient) GetCommandList() []*Command {
	return client.Commands
}

package pkg

import (
	"fmt"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/zeindevs/gowabot/config"
	"github.com/zeindevs/gowabot/internal"
	"go.mau.fi/whatsmeow/proto/waE2E"
)

type item struct {
	Name     []string
	IsPrefix bool
}

type tagSlice []string

func (t tagSlice) Len() int {
	return len(t)
}

func (t tagSlice) Less(i int, j int) bool {
	return t[i] < t[j]
}

func (t tagSlice) Swap(i int, j int) {
	t[i], t[j] = t[j], t[i]
}

func CommandMenu() *internal.Command {
	return &internal.Command{
		Name:     "menu",
		As:       []string{"menu"},
		Tags:     "main",
		IsPrefix: true,
		Exec: func(client *internal.WAClient, m *internal.Message) {
			var str string
			str += fmt.Sprintf("Hello %s\n\n⇒ Library: whatsmeow\n⇒ Language: Golang\n\n", m.PushName)
			var tags map[string][]item

			for _, list := range client.GetCommandList() {
				if tags == nil {
					tags = make(map[string][]item)
				}
				if _, ok := tags[list.Tags]; !ok {
					tags[list.Tags] = []item{}
				}
				tags[list.Tags] = append(tags[list.Tags], item{Name: list.As, IsPrefix: list.IsPrefix})
			}

			var keys tagSlice
			for key := range tags {
				keys = append(keys, key)
			}

			sort.Sort(keys)

			for _, key := range keys {
				str += fmt.Sprintf("「 %s MENU 」\n", strings.ToUpper(key))
				for _, e := range tags[key] {
					var prefix string
					if e.IsPrefix {
						prefix = m.Command[:1]
					} else {
						prefix = ""
					}
					for _, nm := range e.Name {
						str += fmt.Sprintf("ゝ *%s%s*\n", prefix, nm)
					}
				}
				str += "\n"
			}

			client.SendText(m.From, strings.TrimSpace(str), &waE2E.ContextInfo{})
		},
	}
}

func CommandPing() *internal.Command {
	return &internal.Command{
		Name:     "ping",
		As:       []string{"ping"},
		Tags:     "info",
		IsPrefix: true,
		Exec: func(client *internal.WAClient, m *internal.Message) {
			now := time.Now()
			mdate := time.Unix(m.Info.Timestamp.Unix(), 0)
			mtime := now.Sub(mdate)
			m.Reply(fmt.Sprintf("%.3f seconds", mtime.Seconds()))
		},
	}
}

func CommandSource() *internal.Command {
	return &internal.Command{
		Name:     "(sc|source)",
		As:       []string{"sc"},
		Tags:     "info",
		IsPrefix: true,
		Exec: func(client *internal.WAClient, m *internal.Message) {
			m.Reply("https://github.com/zeindevs/gowabot")
		},
	}
}

func CommandMode(cfg *config.Config) *internal.Command {
	return &internal.Command{
		Name:     "(setmode|mode)",
		As:       []string{"setmode"},
		Tags:     "owner",
		IsPrefix: true,
		IsOwner:  true,
		Exec: func(client *internal.WAClient, m *internal.Message) {
			if m.Query == "public" {
				cfg.Public = true
				m.Reply("Public Mode: " + strconv.FormatBool(cfg.Public))
			} else if m.Query == "private" {
				cfg.Public = false
				m.Reply("Public Mode: " + strconv.FormatBool(cfg.Public))
			} else {
				m.Reply("Mode " + m.Query + " Not Found")
			}
		},
	}
}

func CommandExec() *internal.Command {
	return &internal.Command{
		Name:     `\$`,
		As:       []string{"$"},
		Tags:     "owner",
		IsPrefix: false,
		IsOwner:  true,
		Exec: func(client *internal.WAClient, m *internal.Message) {
			out, err := exec.Command("bash", "-c", m.Query).Output()
			if err != nil {
				m.Reply(fmt.Sprintf("%v", err))
				return
			}
			m.Reply(string(out))
		},
	}
}

package sshserver

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/TwiN/go-color"
	"github.com/gliderlabs/ssh"
	"github.com/spaolacci/murmur3"
	"github.com/zeindevs/sendit/data"
	"github.com/zeindevs/sendit/logger"
	"github.com/zeindevs/sendit/util"
	"github.com/zeindevs/sendit/version"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/ratelimit"
	gossh "golang.org/x/crypto/ssh"
)

var (
	maxMsgLen      = 10
	maxFilenameLen = 10
)

type Pipe struct {
	Donech chan struct{}
}

type UserOptions struct {
	Email    string
	Filename string
	Form     string
	Msg      string
	Hook     string
	Airdrop  bool
}

type Handler struct {
	rl           ratelimit.Limiter
	generateLink func() string
	pipes        map[string]Pipe
}

const (
	optionEmail    = "mailto"
	optionHook     = "hook"
	optionFilename = "filename"
	optionMsg      = "msg"
	optionFrom     = "from"
	optionAirdrop  = "airdrop"
)

func ListenAndServe(signer gossh.Signer, pipes map[string]Pipe) error {
	handler := &Handler{
		generateLink: generateLink,
		pipes:        pipes,
	}
	sshPort := util.Getenv("SENDIT_SSH_PORT", ":2222")
	server := ssh.Server{
		Addr:    sshPort,
		Handler: handler.handleSSHSession,
		ServerConfigCallback: func(ctx ssh.Context) *gossh.ServerConfig {
			cfg := &gossh.ServerConfig{
				ServerVersion: "SSH-2.0-sendit",
			}
			cfg.Ciphers = []string{"chacha20-poly1305@openssh.com"}
			return cfg
		},
		PublicKeyHandler: func(ctx ssh.Context, key ssh.PublicKey) bool {
			return true
		},
	}
	server.AddHostKey(signer)
	return server.ListenAndServe()
}

func parseUserOptions(options []string, user *data.User) (*UserOptions, error) {
	userOpts := &UserOptions{}
	options = preParseOptions(options)
	for _, o := range options {
		parts := strings.Split(o, "=")
		if len(parts) < 2 {
			return nil, fmt.Errorf("invalid options")
		}
		key := parts[0]
		value := parts[1]
		switch key {
		case optionAirdrop:
		case optionMsg:
			if len(value) > maxMsgLen {
				errMsg := fmt.Sprintf("message longer than %d characters", maxMsgLen)
				return nil, fmt.Errorf(errMsg)
			}
			if len(value) < 2 {
				return nil, fmt.Errorf("message too short")
			}
			userOpts.Msg = value
		case optionEmail:
			if !util.ValidateEmail(value) {
				return nil, fmt.Errorf("invalid email address")
			}
			userOpts.Email = value
		case optionHook:
			userOpts.Hook = value
		case optionFilename:
			if len(value) > maxFilenameLen {
				return nil, fmt.Errorf("")
			}
		}
	}
	_ = user
	return userOpts, nil
}

func preParseOptions(opts []string) []string {
	return opts
}

type PeekSession struct {
	err    error
	reader io.Reader
}

func peekSession(s ssh.Session) chan PeekSession {
	var (
		peekch = make(chan PeekSession)
		pr     = bufio.NewReader(s)
	)
	go func(r *bufio.Reader) {
		b, err := r.Peek(1)
		if err != nil {
			peekch <- PeekSession{err: err}
			return
		}
		if len(b) == 0 {
			peekch <- PeekSession{err: fmt.Errorf("empty bytes")}
		} else {
			peekch <- PeekSession{reader: r}
		}
	}(pr)
	return peekch
}

type limitReader struct {
	r           io.Reader
	total, size int
}

func LimitReader(r io.Reader, size int64) io.Reader {
	return &limitReader{
		r:     io.LimitReader(r, size),
		size:  int(size),
		total: int(size),
	}
}

func (l *limitReader) Read(p []byte) (n int, err error) {
	return len(p), nil
}

type Context struct {
	userOpts *UserOptions
	user     *data.User
	link     string
}

func (sh *Handler) handleSSHSession(s ssh.Session) {
	sh.rl.Take()

	var (
		start   = time.Now()
		ctx     = new(Context)
		pubkey  = s.PublicKey()
		keyHash = murmur3.Sum32(pubkey.Marshal())
	)
	where := bson.M{
		"settings.sshKeys": bson.M{
			"$elemMatch": bson.M{
				"hash": keyHash,
			},
		},
	}
	user, _ := data.FindUser(where)
	if user != nil {
		ctx.user = user
	}

	logger.Log(
		"msg", "new SSH Session started",
		"ssh username", s.User(),
		"verified", user != nil,
		"serverVersion", s.Context().ServerVersion(),
		"clientVersion", s.Context().ClientVersion(),
	)
	opts, err := parseUserOptions(s.Command(), user)
	if err != nil {
		writeError(s, err.Error())
		return
	}

	if user == nil {
		if opts.Airdrop {
			writeVerifiedOptionsError(s, optionAirdrop)
			return
		}
		if len(opts.Email) > 0 {
			writeVerifiedOptionsError(s, optionEmail)
			return
		}
		if len(opts.Hook) > 0 {
			writeSenditPlusOptionsError(s, optionHook)
			return
		}
	}
	if user != nil {
		if opts.Airdrop {
			transfer, err := sh.generateTransfer(s, ctx)
			if err != nil {
				logger.Log("err", err.Error())
			}
			fmt.Printf("%+v\n", transfer)
			os.Exit(1)
		}
		if len(opts.Hook) > 0 {
			writeSenditPlusOptionsError(s, optionHook)
			return
		}
	}
	if len(opts.Filename) == 0 {
		opts.Filename = "sendit"
	}
	_ = start
}
func (sh *Handler) finalizeAndCleanupTransfer(t data.Transfer) error {
	close(sh.pipes[t.Link].Donech)
	delete(sh.pipes, t.Link)
	return data.FinalizeTransfer(t)
}

func (sh *Handler) generateTransfer(s ssh.Session, ctx *Context) (*data.Transfer, error) {
	_ = s
	_ = ctx
	return &data.Transfer{}, nil
}

func writeLink(s ssh.Session, ctx *Context) {
	builder := strings.Builder{}
	builder.WriteString("\n")
	builder.WriteString(color.InPurple("sendit.io version " + version.Latest))
	builder.WriteString("\n")
	builder.WriteString("\n")

	if ctx.user != nil {
		msg := fmt.Sprintf(" Detected verified user domain https://%s.sendit.sh", ctx.user.Settings.Subdomain)
		builder.WriteString(color.InGreen(msg))
		builder.WriteString("\n")
	} else {
		builder.WriteString(" Want your own personal verified link ")
		builder.WriteString(color.InGreen("username.sendit.io"))
		builder.WriteString("?\n")
		builder.WriteString("\n")
		builder.WriteString("\t-> Visit https://sendit.io to signup, it's FREE")
		builder.WriteString("\n")
	}

	builder.WriteString("Your connection says open untuk someone downloads your file.\n")
	s.Write([]byte(builder.String()))
}

func writeError(s ssh.Session, err string) {

}

func writeVerifiedOptionsError(s ssh.Session, op string) {
	builder := strings.Builder{}
	builder.WriteString("\n")
	msg := fmt.Sprintf(" %s is only available for verified accounts.", op)
	builder.WriteString(color.InYellow(msg))
	builder.WriteString("\n")
	builder.WriteString("\n")
	builder.WriteString(" Sign up for a FREE account and get your personal verified Sendit domain")
	builder.WriteString("\n")
	builder.WriteString("\n")
	builder.WriteString(color.InBlue(" -> https://sendit.io/signup"))
	builder.WriteString("\n")
	builder.WriteString("\n")
	s.Write([]byte(builder.String()))
}

func writeSenditPlusOptionsError(s ssh.Session, op string) {
	builder := strings.Builder{}
	builder.WriteString("\n")
	msg := fmt.Sprintf(" %s is only available for Sentit+ accounts.", op)
	builder.WriteString(color.InYellow(msg))
	builder.WriteString("\n")
	builder.WriteString("\n")
	builder.WriteString(" Upgrade your account and enjoy the FULL Sendit experiance")
	builder.WriteString("\n")
	builder.WriteString("\n")
	builder.WriteString(color.InBlue(" -> https://sendit.io/upgrade"))
	builder.WriteString("\n")
	builder.WriteString("\n")
	s.Write([]byte(builder.String()))
}

func writeLinkExpired(s ssh.Session) {
	builder := strings.Builder{}
	builder.WriteString("\n")
	builder.WriteString(color.InYellow(" Link expired "))
	builder.WriteString("\n")
	builder.WriteString("\n")
	s.Write([]byte(builder.String()))
}

func writeTransferComplete(s ssh.Session, took time.Duration, total int64, speed float64) {
	builder := strings.Builder{}
	builder.WriteString("\n")
	msg := fmt.Sprintf(" Transfer completed in %.2f sec - %d MB - %.2f MB/s", took.Seconds(), total/(1024*1024), speed)
	builder.WriteString(color.InGreen(msg))
	builder.WriteString("\n")
	builder.WriteString("\n")
	s.Write([]byte(builder.String()))
}

type Usage struct {
	command string
	role    string
	desc    string
	example string
}

func generateLink() string {
	return util.GenerateRandomString(12)
}

func makeFullLink(link string, domain string) string {
	if util.IsEnvDev() {
		port := util.Getenv("SENDIT_HTTP_PORT", ":3000")
		if len(domain) > 0 {
			return fmt.Sprintf("http://localhost%s/u/%s/%s", port, domain, link)
		}
		return fmt.Sprintf("http://localhost%s/%s", port, link)
	}
	if len(domain) > 0 {
		return fmt.Sprintf("https://%s.sendit.io/%s", domain, link)
	}
	return fmt.Sprintf("https://sendit.io/%s", link)
}

func makeDirectLink(link string, domain string) string {
	if util.IsEnvDev() {
		port := util.Getenv("SENDIT_DOWNLOAD_PORT", ":4000")
		return fmt.Sprintf("http://localhost%s/%s", port, link)
	}
	_ = domain
	return fmt.Sprintf("https://d.sendit.io/%s", link)
}

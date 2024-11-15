package mailer

import "embed"

const (
	FromName             = "RepoPattern"
	maxRetries           = 3
	UserWelcomeTemplates = "user_invitation.html"
)

//go:embed "templates"
var FS embed.FS

type Client interface {
	Send(templateFile, username, email string, data any, isSandBox bool) (int, error)
}

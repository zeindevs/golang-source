package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/zeindevs/sendit/data"
	"github.com/zeindevs/sendit/logger"
	"github.com/zeindevs/sendit/sshserver"
	"github.com/zeindevs/sendit/util"
)

type TransderHandler struct {
	pipes map[string]sshserver.Pipe
}

func NewTransferHandler(pipes map[string]sshserver.Pipe) *TransderHandler {
	return &TransderHandler{
		pipes: pipes,
	}
}

func (h *TransderHandler) HandleDirectDownload(w http.ResponseWriter, r *http.Request) {}

func (h *TransderHandler) HandleGetLink(c *fiber.Ctx) error {
	link := c.Params("link")
	if util.IsValidLink(link) {
		return c.Redirect("/", 301)
	}
	transfer, err := data.FindTransferByLink(link)
	if err != nil {
		logger.Log("err", err, "link", link)
		return c.Render("link_not_found", locals(c), "layouts/simple_main")
	}
	if _, ok := h.pipes[link]; !ok {
		return c.Render("link_not_found", locals(c), "layouts/simple_main")
	}
	data := struct {
		From             string
		Link             string
		Expires          string
		Filename         string
		Message          string
		IsVerified       bool
		VerifiedUsername string
		Created          string
	}{
		From:             transfer.Initiator,
		Link:             transfer.Link,
		Filename:         transfer.Filename,
		Message:          transfer.Message,
		IsVerified:       transfer.IsVerified,
		VerifiedUsername: c.Params("subdomain"),
	}
	local := locals(c)
	local["transfer"] = data
	return c.Render("download", local, "layouts/simple_main")
}

func (h *TransderHandler) HandleDownload(w http.ResponseWriter, r *http.Request) {}

func HandleGetLink(c *fiber.Ctx) error {
	return nil
}

func HandleDownload(c *fiber.Ctx) error {
	return nil
}

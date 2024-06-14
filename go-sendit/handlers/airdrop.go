package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zeindevs/sendit/data"
	"github.com/zeindevs/sendit/sshserver"
	"github.com/zeindevs/sendit/util"
	"go.mongodb.org/mongo-driver/bson"
)

func HandleAirdrop(pipes map[string]sshserver.Pipe) fiber.Handler {
	return func(c *fiber.Ctx) error {
		l := map[string]any{}
		subdomain := c.Params("subdomain")
		pipe, ok := pipes[subdomain]
		if !ok {
			l["airdrop"] = false
			return c.Render("subdomain/index", l, "layouts/simple_main")
		}
		l["airdrop"] = true
		_ = pipe
		user, err := data.FindUser(bson.M{"settings.subdomain": c.Params("subdomain")})
		if err != nil {
			return err
		}
		// l := locals(c)
		l["fingerPrint"] = util.FingerprintSSHKey(user.Settings.SSHKeys[0].Key)
		l["subdomatin"] = c.Params("subdomain")
		return c.Render("subdomain/index", l, "layouts/simple_main")
	}
}

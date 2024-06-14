package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zeindevs/sendit/data"
	"github.com/zeindevs/sendit/util"
	"go.mongodb.org/mongo-driver/bson"
)

func HandleUserSubdomain(c *fiber.Ctx) error {
	user, err := data.FindUser(bson.M{"settings.subdomain": c.Params("subdomain")})
	if err != nil {
		return err
	}
	// l := locals(c)
	l := map[string]string{}
	l["fingerPrint"] = util.FingerprintSSHKey(user.Settings.SSHKeys[0].Key)
	l["subdomatin"] = c.Params("subdomain")
	return c.Render("subdomain/index", l, "layouts/simple_main")
}

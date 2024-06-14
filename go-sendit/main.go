package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/mustache/v2"
	"github.com/joho/godotenv"
	"github.com/zeindevs/sendit/db"
	"github.com/zeindevs/sendit/handlers"
	"github.com/zeindevs/sendit/logger"
	"github.com/zeindevs/sendit/sshserver"
	"github.com/zeindevs/sendit/util"

	gossh "golang.org/x/crypto/ssh"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	logger.Init()
	logEnv()

	s := &http.Server{
		ReadTimeout:  15 * time.Minute,
		WriteTimeout: 15 * time.Minute,
		Addr:         ":4000",
	}
	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	engine := mustache.New("www", ".html")
	engine.Reload(true)

	pipes := make(map[string]sshserver.Pipe)
	transferHandler := handlers.NewTransferHandler(pipes)
	go func() {
		http.HandleFunc("/", transferHandler.HandleDirectDownload)
		logger.Log("msg", "download server running", "port")
		log.Fatal(s.ListenAndServe())
	}()

	app := fiber.New(fiber.Config{
		Views:        engine,
		WriteTimeout: time.Minute * 10,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			logger.Log("error", err)
			return c.Render("errors/500", nil, "layouts/simple_main")
		},
	})
	app.Static("/assets", "./www/assets")
	app.Use(recover.New(recover.Config{EnableStackTrace: true}))
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: os.Getenv("SENDIT_SECRET"),
	}))
	app.Use(favicon.New(favicon.Config{
		File: "./www/favicon.io",
		URL:  "/favicon.ico",
	}))
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
		return c.Next()
	})
	app.Use(handlers.AuthMiddleware)
	app.Get("/", handlers.HandleLanding)
	app.Get("/terms", handlers.HandleTerms)
	app.Get("/upgrade", handlers.HandleGetUpgrade)
	app.Get("/signup", handlers.HandleSignup)
	app.Get("/login", handlers.HandleLoginPage)
	app.Get("/login/github", handlers.HandleGithubAuth)
	app.Get("/login/github/callback", handlers.HandleGithubAuthCallback)
	app.Get("/logout", handlers.HandleLogout)
	app.Get("/:link", handlers.HandleGetLink)
	app.Get("/d/:link", handlers.HandleDownload)

	app.Get("/a/b/c/metrics", basicauth.New(basicauth.Config{
		Users: map[string]string{
			"therealdeal": "1458@superline",
		},
	}), handlers.HandleMetrics)

	// Must be authenticated
	mustAuth := app.Group("/m", handlers.MustAuthMiddleware)
	mustAuth.Get("/settings/account", handlers.HandleUserSettings)
	mustAuth.Post("/settings/account", handlers.HandlePostSubdomain)
	mustAuth.Get("/settings/keys", handlers.HandleUserSettingsKeys)
	mustAuth.Post("/settings/keys", handlers.HandleUserSettingsDeleteKey)
	mustAuth.Get("/settings/keys/add", handlers.HandleGetUserSettingsAddKey)
	mustAuth.Post("/settings/keys/add", handlers.HandlePostUserSettingsAddKey)

	app.Get("/u/:subdomain", handlers.SubdomainCheckerMiddleware, handlers.HandleAirdrop(pipes))
	app.Get("/u/:subdomain/:link", handlers.SubdomainCheckerMiddleware, transferHandler.HandleGetLink)

	go func() {
		app.Listen(util.Getenv("SENDIT_HTTP_PORT", ":3000"))
	}()

	b, err := os.ReadFile("privateKey")
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := gossh.ParsePrivateKey(b)
	if err != nil {
		log.Fatal("Faied to parse private key: ", err)
	}

	log.Fatal(sshserver.ListenAndServe(privateKey, pipes))
}

func logEnv() {

}

package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/ashrhmn/go-logger/config"
	"github.com/ashrhmn/go-logger/constants"
	"github.com/ashrhmn/go-logger/guards"
	"github.com/ashrhmn/go-logger/middlewares"
	"github.com/ashrhmn/go-logger/modules/storage"
	"github.com/ashrhmn/go-logger/types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"go.uber.org/fx"
)

type Server struct {
	app *fiber.App
}

func proxyClient(c *fiber.Ctx) error {
	path := c.Path()
	if os.Getenv(config.ENV_KEY) == "production" {
		parts := strings.Split(strings.Split(path, "?")[0], "/")
		if strings.Contains(parts[len(parts)-1], ".") {
			return c.SendFile(fmt.Sprintf("./views%s", path))
		}
		return c.SendFile("./views/index.html")
	}
	proxyBaseUrl := os.Getenv("PROXY_BASE_URL")
	if proxyBaseUrl == "" {
		proxyBaseUrl = "http://localhost:5173"
	}
	proxyUrl := fmt.Sprintf("%s%s", proxyBaseUrl, path)
	if err := proxy.Do(c, proxyUrl); err != nil {
		log.Error(err, proxyUrl)
		if strings.Contains(err.Error(), "connection refused") {
			c.Set("Content-Type", "text/html")
			return c.Status(500).SendString(constants.Html500)
		}
		return err
	}
	return nil
}

func newServer(controllers []types.Controller, mongoCollection storage.MongoCollection) *Server {
	app := fiber.New()
	app.Use(middlewares.AuthCookieMiddleware(mongoCollection.AuthSessionCollection))
	api := app.Group("/api")
	for _, controller := range controllers {
		controller.RegisterRoutes(api)
	}

	if os.Getenv(config.ENV_KEY) == "production" {
		app.Use("/___assets___", filesystem.New(filesystem.Config{
			Root:   http.Dir("./views/___assets___"),
			Browse: false,
			MaxAge: 3600,
		}))
	}

	app.Get("/login", guards.NoneLoggedIn("/dashboard"), proxyClient)
	app.Get("/dashboard", guards.AnyLoggedIn("/login"), proxyClient)
	app.Get("/dashboard/*", guards.AnyLoggedIn("/login"), proxyClient)
	app.Get("*", proxyClient)
	return &Server{
		app: app,
	}
}

func startServer(s *Server, lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				println("Server started.")
				if err := s.app.Listen(":4000"); err != nil {
					panic(err)
				}
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			s.app.Shutdown()
			println("Server stopped.")
			return nil
		},
	})
}

var Module = fx.Module("server",
	fx.Provide(fx.Annotate(
		newServer,
		fx.ParamTags(`group:"controllers"`)),
	),
	fx.Invoke(startServer),
)

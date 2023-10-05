package logging

import (
	"github.com/ashrhmn/go-logger/utils"
	"github.com/gofiber/fiber/v2"
)

type LoggingController struct {
	loggingService LoggingService
}

func (lc LoggingController) RegisterRoutes(app fiber.Router) {
	app.Get("/logging", func(c *fiber.Ctx) error {
		lc.loggingService.Log()
		return c.SendString("Hello, World 👋!")
	})
}

var loggingControllerProvider = utils.ProvideController(func(loggingService LoggingService) LoggingController {
	return LoggingController{
		loggingService: loggingService,
	}
})

package logging

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ashrhmn/go-logger/constants"
	"github.com/ashrhmn/go-logger/guards"
	"github.com/ashrhmn/go-logger/middlewares"
	"github.com/ashrhmn/go-logger/utils"
	"github.com/gofiber/fiber/v2"
)

type LoggingController struct {
	loggingService LoggingService
}

func (lc LoggingController) RegisterRoutes(app fiber.Router) {

	logging := app.Group("/logging")

	logging.Get(
		"/all-log-levels",
		guards.WithPermission(constants.PermissionViewLogs, ""),
		func(c *fiber.Ctx) error {
			levels, err := lc.loggingService.GetLogLevels()
			if err != nil {
				return err
			}
			return c.JSON(levels)
		},
	)

	logging.Get(
		"/selected-log-levels",
		guards.WithPermission(constants.PermissionViewLogs, ""),
		func(c *fiber.Ctx) error {
			token := middlewares.GetAuthTokenFromRequest(c)
			levels, err := lc.loggingService.GetSelectedLogLevel(token)
			if err != nil {
				return err
			}
			return c.JSON(levels)
		},
	)

	logging.Get(
		"/logs",
		guards.WithPermission(constants.PermissionViewLogs, ""),
		func(c *fiber.Ctx) error {
			limit, offset := utils.ExtractLimitOffset(c)
			levelsJoined := c.Query("levels")
			levels := []string{}
			if levelsJoined != "" {
				levels = strings.Split(levelsJoined, ",")
			}
			now := time.Now().Unix()

			fromTimestampStr := c.Query("from", fmt.Sprintf("%d", now-60*60*6))
			fromTimestamp, err := strconv.ParseInt(fromTimestampStr, 10, 64)
			if err != nil {
				return fiber.NewError(fiber.StatusBadRequest, "Bad Request, Invalid From Timestamp")
			}

			toTimestampStr := c.Query("to", fmt.Sprintf("%d", now))
			toTimestamp, err := strconv.ParseInt(toTimestampStr, 10, 64)
			if err != nil {
				return fiber.NewError(fiber.StatusBadRequest, "Bad Request, Invalid to Timestamp")
			}

			logs, err := lc.loggingService.GetLogs(levels, fromTimestamp, toTimestamp, limit, offset)
			if err != nil {
				return err
			}
			return c.JSON(logs)
		},
	)

	logging.Patch(
		"/selected-log-levels",
		guards.WithPermission(constants.PermissionViewLogs, ""),
		func(c *fiber.Ctx) error {
			token := middlewares.GetAuthTokenFromRequest(c)
			var levels []string
			err := c.BodyParser(&levels)
			if err != nil {
				return fiber.NewError(fiber.StatusBadRequest, "Bad Request, Invalid Body")
			}
			err = lc.loggingService.UpdateSelectedLogLevel(token, levels)
			if err != nil {
				return err
			}
			return c.SendStatus(201)
		},
	)
}

var loggingControllerProvider = utils.ProvideController(func(loggingService LoggingService) LoggingController {
	return LoggingController{
		loggingService: loggingService,
	}
})

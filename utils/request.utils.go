package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ExtractLimitOffset(c *fiber.Ctx) (int64, int64) {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		offset = 0
	}

	return int64(limit), int64(offset)
}

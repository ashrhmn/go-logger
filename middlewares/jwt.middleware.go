package middlewares

import (
	"github.com/ashrhmn/go-logger/constants"
	"github.com/ashrhmn/go-logger/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func AuthCookieMiddleware(authSessionCollection *mongo.Collection) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies(constants.AUTH_TOKEN_ONE_KEY) + "." + c.Cookies(constants.AUTH_TOKEN_TWO_KEY) + "." + c.Cookies(constants.AUTH_TOKEN_THREE_KEY)
		user, err := types.AuthPayloadFromToken(token, authSessionCollection)
		if err != nil {
			return c.Next()
		}
		c.Locals(constants.AUTH_USER_CONTEXT_KEY, user)
		return c.Next()
	}
}

func GetAuthUserFromRequest(c *fiber.Ctx) (types.AuthPayload, error) {
	user := types.AuthPayload{}
	localsUser := c.Locals(constants.AUTH_USER_CONTEXT_KEY)
	if localsUser == nil {
		return user, fiber.ErrUnauthorized
	}
	user, ok := localsUser.(types.AuthPayload)
	if !ok {
		return user, fiber.ErrUnauthorized
	}
	return user, nil
}

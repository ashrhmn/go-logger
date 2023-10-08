package auth

import (
	"strings"

	"github.com/ashrhmn/go-logger/constants"
	"github.com/ashrhmn/go-logger/guards"
	"github.com/ashrhmn/go-logger/middlewares"
	"github.com/ashrhmn/go-logger/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService AuthService
}

func newAuthController(authService AuthService) AuthController {
	return AuthController{
		authService: authService,
	}
}

func (ac AuthController) RegisterRoutes(app fiber.Router) {
	auth := app.Group("/auth")

	auth.Post("/login", func(c *fiber.Ctx) error {
		body := LoginInput{}

		err := c.BodyParser(&body)

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Bad Request, Invalid Body")
		}
		if err := utils.ValidateStruct(body); err != "" {
			return c.Status(400).SendString(err)
		}
		token, err := ac.authService.Login(body)
		if err != nil {
			return err
		}
		tokens := strings.Split(token, ".")
		if len(tokens) != 3 {
			return fiber.ErrInternalServerError
		}

		c.Cookie(&fiber.Cookie{
			Name:     constants.AUTH_TOKEN_ONE_KEY,
			Value:    tokens[0],
			Secure:   true,
			HTTPOnly: true,
		})

		c.Cookie(&fiber.Cookie{
			Name:     constants.AUTH_TOKEN_TWO_KEY,
			Value:    tokens[1],
			Secure:   true,
			HTTPOnly: true,
		})

		c.Cookie(&fiber.Cookie{
			Name:     constants.AUTH_TOKEN_THREE_KEY,
			Value:    tokens[2],
			Secure:   true,
			HTTPOnly: true,
		})

		return c.SendStatus(201)
	})

	auth.Get(
		"/whoami",
		guards.AnyLoggedIn(""),
		func(c *fiber.Ctx) error {
			user, err := middlewares.GetAuthUserFromRequest(c)
			if err != nil {
				return err
			}
			return c.JSON(user)
		},
	)

	auth.Delete(
		"/",
		guards.AnyLoggedIn(""),
		func(c *fiber.Ctx) error {
			token := middlewares.GetAuthTokenFromRequest(c)
			err := ac.authService.Logout(token)
			if err != nil {
				return err
			}
			return c.SendStatus(201)
		},
	)

	auth.Get(
		"/permissions",
		func(c *fiber.Ctx) error {
			return c.JSON(ac.authService.GetAllPermissions())
		},
	)
}

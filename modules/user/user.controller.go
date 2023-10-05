package user

import (
	"github.com/ashrhmn/go-logger/utils"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService UserService
}

func newUserController(userService UserService) UserController {
	return UserController{
		userService: userService,
	}
}

func (uc UserController) RegisterRoutes(app fiber.Router) {
	usersRoute := app.Group("/users")

	usersRoute.Post("/", func(c *fiber.Ctx) error {
		body := AddUserDto{}

		err := c.BodyParser(&body)

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Bad Request, Invalid Body")
		}

		if err := utils.ValidateStruct(body); err != "" {
			return c.Status(400).SendString(err)
		}
		err = uc.userService.AddUser(body)
		if err != nil {
			return err
		}
		return c.SendStatus(201)
	})
}

package user

import (
	"github.com/ashrhmn/go-logger/constants"
	"github.com/ashrhmn/go-logger/guards"
	"github.com/ashrhmn/go-logger/middlewares"
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

	usersRoute.Post(
		"/",
		guards.WithPermission(constants.PermissionAddUser, ""),
		func(c *fiber.Ctx) error {
			body := AddUserDto{}

			err := c.BodyParser(&body)

			if err != nil {
				return fiber.NewError(fiber.StatusBadRequest, "Bad Request, Invalid Body")
			}

			if err := utils.ValidateStruct(body); err != "" {
				return c.Status(400).SendString(err)
			}
			loggedInUser, err := middlewares.GetAuthUserFromRequest(c)
			if err != nil {
				return err
			}
			err = uc.userService.AddUser(body, loggedInUser)
			if err != nil {
				return err
			}
			return c.SendStatus(201)
		},
	)
}

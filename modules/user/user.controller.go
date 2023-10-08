package user

import (
	"github.com/ashrhmn/go-logger/constants"
	"github.com/ashrhmn/go-logger/guards"
	"github.com/ashrhmn/go-logger/middlewares"
	"github.com/ashrhmn/go-logger/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	usersRoute.Get(
		"/",
		guards.WithPermission(constants.PermissionViewUsers, ""),
		func(c *fiber.Ctx) error {
			limit, offset := utils.ExtractLimitOffset(c)
			filter := c.Query("filter")
			users, err := uc.userService.GetAllUsers(limit, offset, filter)
			if err != nil {
				return err
			}
			return c.JSON(users)
		},
	)

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

	usersRoute.Get(
		"/:id",
		guards.WithPermission(constants.PermissionViewUsers, ""),
		func(c *fiber.Ctx) error {
			hexId := c.Params("id")
			id, err := primitive.ObjectIDFromHex(hexId)
			if err != nil {
				return fiber.NewError(fiber.StatusBadRequest, "Invalid ID")
			}
			user, err := uc.userService.GetUserById(id)
			if err != nil {
				return err
			}
			return c.JSON(user)
		},
	)

	usersRoute.Patch(
		"/:id",
		guards.WithPermission(constants.PermissionModifyUser, ""),
		func(c *fiber.Ctx) error {
			hexId := c.Params("id")
			id, err := primitive.ObjectIDFromHex(hexId)
			if err != nil {
				return fiber.NewError(fiber.StatusBadRequest, "Invalid ID")
			}
			loggedInUser, err := middlewares.GetAuthUserFromRequest(c)
			if err != nil {
				return err
			}
			body := UpdateUserDto{}
			err = c.BodyParser(&body)
			if err != nil {
				return fiber.NewError(fiber.StatusBadRequest, "Bad Request, Invalid Body")
			}
			if err := utils.ValidateStruct(body); err != "" {
				return c.Status(400).SendString(err)
			}
			err = uc.userService.UpdateUser(id, body, loggedInUser)
			if err != nil {
				return err
			}

			return c.SendStatus(204)
		},
	)

	usersRoute.Delete(
		"/:id",
		guards.WithPermission(constants.PermissionModifyUser, ""),
		func(c *fiber.Ctx) error {
			hexId := c.Params("id")
			id, err := primitive.ObjectIDFromHex(hexId)
			if err != nil {
				return fiber.NewError(fiber.StatusBadRequest, "Invalid ID")
			}
			loggedInUser, err := middlewares.GetAuthUserFromRequest(c)
			if err != nil {
				return err
			}
			err = uc.userService.DeleteUser(id, loggedInUser)
			if err != nil {
				return err
			}

			return c.SendStatus(204)
		},
	)
}

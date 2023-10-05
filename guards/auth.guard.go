package guards

import (
	"slices"

	"github.com/ashrhmn/go-logger/middlewares"
	"github.com/gofiber/fiber/v2"
)

var getUser = middlewares.GetAuthUserFromRequest

func AnyLoggedIn(redirect string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, err := getUser(c)
		if err != nil {
			if redirect != "" {
				return c.Redirect(redirect)
			}
			return err
		}
		return c.Next()
	}
}

func NoneLoggedIn(redirect string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, err := getUser(c)
		if err != nil {
			return c.Next()
		}
		if redirect != "" {
			return c.Redirect(redirect)
		}
		return fiber.ErrForbidden
	}
}

func WithPermission(permission string, redirect string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, err := getUser(c)
		if err != nil {
			if redirect != "" {
				return c.Redirect(redirect)
			}
			return err
		}
		if !slices.Contains(user.Permissions, permission) {
			if redirect != "" {
				return c.Redirect(redirect)
			}
			return fiber.ErrForbidden
		}
		return c.Next()
	}
}

func WithAnyPermission(permissions []string, redirect string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, err := getUser(c)
		if err != nil {
			if redirect != "" {
				return c.Redirect(redirect)
			}
			return err
		}
		for _, permission := range permissions {
			if slices.Contains(user.Permissions, permission) {
				return c.Next()
			}
		}
		if redirect != "" {
			return c.Redirect(redirect)
		}
		return fiber.ErrForbidden
	}
}

func WithAllPermissions(permissions []string, redirect string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, err := getUser(c)
		if err != nil {
			if redirect != "" {
				return c.Redirect(redirect)
			}
			return err
		}
		for _, permission := range permissions {
			if !slices.Contains(user.Permissions, permission) {
				if redirect != "" {
					return c.Redirect(redirect)
				}
				return fiber.ErrForbidden
			}
		}
		return c.Next()
	}
}

package auth

import (
	"github.com/ashrhmn/go-logger/modules/storage"
	"github.com/ashrhmn/go-logger/modules/user"
	"github.com/ashrhmn/go-logger/types"
	"github.com/ashrhmn/go-logger/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type AuthService struct {
	userService     user.UserService
	mongoCollection storage.MongoCollection
}

func newAuthService(userService user.UserService, mongoCollection storage.MongoCollection) AuthService {
	return AuthService{
		userService:     userService,
		mongoCollection: mongoCollection,
	}
}

func (as AuthService) Login(loginInput LoginInput) (token string, err error) {
	user, err := as.userService.GetUserByEmailOrUsername(loginInput.UsernameOrEmail)
	if err != nil {
		log.Error(err)
		return "", fiber.ErrInternalServerError
	}

	match, err := utils.ComparePasswordAndHash(loginInput.Password, user.Password)
	if err != nil {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Invalid username or password")
	}
	if !match {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Invalid username or password")
	}
	payload := types.AuthPayload{
		Username:    user.Username,
		Email:       user.Email,
		Permissions: user.Permissions,
	}
	token, err = payload.GenerateToken(as.mongoCollection.AuthSessionCollection)
	if err != nil {
		log.Error(err)
		return "", fiber.ErrInternalServerError
	}
	return token, nil
}

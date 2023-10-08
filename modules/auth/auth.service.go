package auth

import (
	"context"

	"github.com/ashrhmn/go-logger/constants"
	"github.com/ashrhmn/go-logger/modules/logging"
	"github.com/ashrhmn/go-logger/modules/storage"
	"github.com/ashrhmn/go-logger/modules/user"
	"github.com/ashrhmn/go-logger/types"
	"github.com/ashrhmn/go-logger/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson"
)

type AuthService struct {
	userService     user.UserService
	loggingService  logging.LoggingService
	mongoCollection storage.MongoCollection
}

func newAuthService(
	userService user.UserService,
	mongoCollection storage.MongoCollection,
	loggingService logging.LoggingService,
) AuthService {
	return AuthService{
		userService:     userService,
		mongoCollection: mongoCollection,
		loggingService:  loggingService,
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
		ID:          user.ID,
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

func (as AuthService) GetAllPermissions() []string {
	return constants.PermissionsAll
}

func (as AuthService) Logout(token string) error {
	_, err := as.mongoCollection.AuthSessionCollection.DeleteOne(
		context.Background(),
		bson.D{{Key: "token", Value: token}},
	)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

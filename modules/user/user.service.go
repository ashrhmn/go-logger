package user

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/ashrhmn/go-logger/constants"
	"github.com/ashrhmn/go-logger/modules/storage"
	"github.com/ashrhmn/go-logger/types"
	"github.com/ashrhmn/go-logger/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserService struct {
	mongoCollection storage.MongoCollection
}

func newUserService(mongoCollection storage.MongoCollection) UserService {
	go ensureAdminUser(mongoCollection)
	return UserService{
		mongoCollection: mongoCollection,
	}
}

func (us UserService) GetUserByEmailOrUsername(emailOrUsername string) (types.User, error) {
	lower := strings.ToLower(emailOrUsername)
	user := types.User{}
	err := us.mongoCollection.UserCollection.FindOne(
		context.Background(),
		bson.D{{Key: "$or", Value: bson.A{
			bson.D{{Key: "email", Value: lower}},
			bson.D{{Key: "username", Value: lower}},
		}}},
		&options.FindOneOptions{},
	).Decode(&user)

	if err != nil {
		log.Error(err)
		return user, err
	}
	return user, nil
}

func (us UserService) AddUser(addUserDto AddUserDto, loggedInUser types.AuthPayload) error {

	_, err := us.GetUserByEmailOrUsername(addUserDto.Username)
	if err == nil {
		return fiber.NewError(400, "Username already in use")
	}
	_, err = us.GetUserByEmailOrUsername(addUserDto.Email)
	if err == nil {
		return fiber.NewError(400, "Email already in use")
	}

	password, err := utils.HashPassword(addUserDto.Password)
	if err != nil {
		return err
	}

	user := types.User{
		Username:    addUserDto.Username,
		Email:       addUserDto.Email,
		Password:    password,
		FirstName:   addUserDto.FirstName,
		LastName:    addUserDto.LastName,
		Permissions: []string{},
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
		CreatedBy:   loggedInUser.Username,
		UpdatedBy:   loggedInUser.Username,
	}
	_, err = us.mongoCollection.UserCollection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

func ensureAdminUser(mongoCollection storage.MongoCollection) {
	cursor, err := mongoCollection.UserCollection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Error("Error finding users", err)
		return
	}
	defer cursor.Close(context.Background())
	if cursor.RemainingBatchLength() > 0 {
		return
	}

	plainPassword := os.Getenv("ADMIN_PASSWORD")
	if plainPassword == "" {
		plainPassword = "admin"
	}

	password, err := utils.HashPassword(plainPassword)
	if err != nil {
		panic(err)
	}

	adminUser := types.User{
		Username:    "admin",
		Password:    password,
		Permissions: []string{constants.PermissionAdmin},
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
		CreatedBy:   "system",
		UpdatedBy:   "system",
	}

	_, err = mongoCollection.UserCollection.InsertOne(context.Background(), adminUser)
	if err != nil {
		log.Error("Error creating init admin", err)
	}

	log.Info("Init admin created")

}

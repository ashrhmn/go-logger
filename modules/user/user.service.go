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
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		ID:          primitive.NewObjectID(),
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

func (us UserService) GetUserById(id primitive.ObjectID) (types.User, error) {
	user := types.User{}
	err := us.mongoCollection.UserCollection.FindOne(
		context.Background(),
		bson.M{"_id": id},
		&options.FindOneOptions{},
	).Decode(&user)
	if err != nil {
		log.Error(err)
		return user, err
	}
	return user, nil
}

func (us UserService) UpdateUser(
	id primitive.ObjectID,
	updateUserDto UpdateUserDto,
	loggedInUser types.AuthPayload,
) error {
	user, err := us.GetUserById(id)
	if err != nil {
		log.Error(err)
		return err
	}

	if updateUserDto.Password != "" {
		password, err := utils.HashPassword(updateUserDto.Password)
		if err != nil {
			return err
		}
		user.Password = password
	}

	if updateUserDto.Username != "" {
		if user.Username != updateUserDto.Username {
			_, err := us.GetUserByEmailOrUsername(updateUserDto.Username)
			if err == nil {
				return fiber.NewError(400, "Username already in use")
			}
		}
		user.Username = updateUserDto.Username
	}

	if updateUserDto.Email != "" {
		if user.Email != updateUserDto.Email {
			_, err := us.GetUserByEmailOrUsername(updateUserDto.Email)
			if err == nil {
				return fiber.NewError(400, "Email already in use")
			}
		}
		user.Email = updateUserDto.Email
	}

	if updateUserDto.FirstName != "" {
		user.FirstName = updateUserDto.FirstName
	}

	if updateUserDto.LastName != "" {
		user.LastName = updateUserDto.LastName
	}

	if updateUserDto.Permissions != nil {
		user.Permissions = updateUserDto.Permissions
	}

	user.UpdatedBy = loggedInUser.Username
	user.UpdatedAt = time.Now().Unix()

	_, err = us.mongoCollection.UserCollection.UpdateOne(
		context.Background(),
		bson.D{{Key: "_id", Value: id}},
		bson.D{{Key: "$set", Value: user}},
	)

	return err
}

func (us UserService) GetAllUsers(limit int64, offset int64, filter string) ([]types.User, error) {

	var filterBson interface{}

	if filter != "" {
		err := bson.UnmarshalExtJSON([]byte(filter), true, &filterBson)
		if err != nil {
			log.Error("Error parsing filter", err)
			return nil, fiber.NewError(400, "Invalid filter")
		}
	} else {
		filterBson = bson.D{{}}
	}

	cursor, err := us.mongoCollection.UserCollection.Find(
		context.Background(),
		filterBson,
		&options.FindOptions{
			Limit: &limit,
			Skip:  &offset,
		},
	)
	if err != nil {
		log.Error("Error finding users", err)
		return nil, err
	}
	defer cursor.Close(context.Background())
	users := []types.User{}
	for cursor.Next(context.Background()) {
		var user types.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Error("Error decoding user", err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (us UserService) DeleteUser(id primitive.ObjectID, loggedInUser types.AuthPayload) error {
	user, err := us.GetUserById(id)
	if err != nil {
		return err
	}
	if user.DeletedAt != 0 {
		_, err = us.mongoCollection.UserCollection.DeleteOne(
			context.Background(),
			bson.D{{Key: "_id", Value: id}},
		)
		if err != nil {
			log.Error("Error deleting user", err)
			return err
		}
	} else {
		_, err = us.mongoCollection.UserCollection.UpdateOne(
			context.Background(),
			bson.D{{Key: "_id", Value: id}},
			bson.D{{Key: "$set", Value: bson.D{
				{Key: "deletedAt", Value: time.Now().Unix()},
				{Key: "deletedBy", Value: loggedInUser.ID},
			}}},
		)
		if err != nil {
			log.Error("Error soft deleting user", err)
			return err
		}
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
		ID:          primitive.NewObjectID(),
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

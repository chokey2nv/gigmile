package services

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/chokey2nv/gigmile/config"
	"github.com/chokey2nv/gigmile/dtos"
	"github.com/chokey2nv/gigmile/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

const User_collection = "users"

var (
	ErrPasswordFormat = errors.New("unable to has password")
)

type UserService struct {
	*config.AppConfig
}

var Errorf = config.Errorf

func (service *UserService) CreateUser(ctx context.Context, userDto *dtos.UserDto) (*models.User, error) {
	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	col := service.DBClient.Database(service.Database).Collection(User_collection)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, ErrPasswordFormat
	}
	user := &models.User{
		LastName:  userDto.LastName,
		FirstName: userDto.FirstName,
		Email:     userDto.Email,
		UserRole:  userDto.UserRole,
		Hash:      string(hashedPassword),
		Id:        uuid.NewString(),
	}
	_, err = col.InsertOne(c, user)
	if err != nil {
		return nil, Errorf(err)
	}
	return user, nil
}
func (service *UserService) GetUsers(ctx context.Context, pageOption dtos.PageOption) ([]*models.User, error) {
	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	collection := config.Config.DBClient.Database(config.Config.Database).Collection(User_collection)
	cursor, err := collection.Find(
		c, bson.M{},
		&options.FindOptions{Limit: &pageOption.Limit, Skip: &pageOption.Skip},
	)
	if err != nil {
		return nil, Errorf(err)
	}
	defer cursor.Close(c)
	// Iterate through the cursor and collect users
	var users []*models.User
	for cursor.Next(context.Background()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			log.Println(err)
			continue
		}
		users = append(users, &user)
	}
	return users, nil
}
func (service *UserService) GetUser(ctx context.Context, getUserDto dtos.GetUserDto) (*models.User, error) {
	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	collection := config.Config.DBClient.Database(config.Config.Database).Collection(User_collection)
	filter := bson.M{}
	err := config.StructToBSON(getUserDto, &filter)
	if err != nil {
		return nil, Errorf(err)
	}
	user := models.User{}
	err = collection.FindOne(
		c,
		filter,
	).Decode(&user)
	if err != nil {
		return nil, Errorf(err)
	}
	return &user, nil
}

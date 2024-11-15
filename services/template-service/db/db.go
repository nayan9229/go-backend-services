package db

import (
	"github.com/nayan9229/go-backend-services/services/template-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DB interface {
	GetUsers() ([]*model.User, error)
	GetUserByID(UserID int) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(UserID int, user *model.User) error
	DeleteUser(UserID int) error
	Close() error
}

type JsonDB interface {
	InsertUser(user model.UserBson) (*mongo.InsertOneResult, error)
	FindUserByID(id primitive.ObjectID) (*model.UserBson, error)
	UpdateUser(id primitive.ObjectID, update bson.M) (*mongo.UpdateResult, error)
	DeleteUser(id primitive.ObjectID) (*mongo.DeleteResult, error)
	Close() error
}

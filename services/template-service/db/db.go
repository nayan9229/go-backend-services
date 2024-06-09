package db

import "github.com/nayan9229/go-backend-services/services/template-service/model"

type DB interface {
	GetUsers() ([]*model.User, error)
	GetUserByID(UserID int) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(UserID int, user *model.User) error
	DeleteUser(UserID int) error
}

type JsonDB interface {
}

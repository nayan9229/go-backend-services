package db

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/nayan9229/go-backend-services/services/template-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const qGetUsers = `SELECT user_id, first_name, last_name, email FROM USERS OFFSET 0 LIMIT 1000;`

const qGetUserByID = `SELECT user_id, first_name, last_name, email FROM USERS WHERE user_id = $1;`

const qCreateUser = `INSERT INTO USERS (user_id, first_name, last_name, email, password) 
						VALUES (:user_id, :first_name, :last_name, :email, :password);`

const qUpdateUser = `UPDATE USERS SET first_name = :first_name, last_name = :last_name, 
						email = :email, password = :password WHERE user_id = :user_id;`

const qDeleteUser = `DELETE FROM USERS WHERE user_id = $1;`

func (pg *PGClient) GetUsers() ([]*model.User, error) {
	users := []*model.User{}
	if err := pg.DB.Select(&users, qGetUsers); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("users not found")
		}
		return nil, err
	}
	return users, nil
}

func (pg *PGClient) GetUserByID(UserID int) (*model.User, error) {
	user := model.User{}
	if err := pg.DB.Get(&user, qGetUserByID, UserID); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (pg *PGClient) CreateUser(user *model.User) error {
	_, err := pg.DB.NamedExec(qCreateUser, user)
	if err != nil {
		return err
	}
	return nil
}

func (pg *PGClient) UpdateUser(UserID int, user *model.User) error {
	user.UserID = UserID
	_, err := pg.DB.NamedExec(qUpdateUser, user)
	if err != nil {
		return err
	}
	return nil
}
func (pg *PGClient) DeleteUser(UserID int) error {
	_, err := pg.DB.Exec(qDeleteUser, UserID)
	if err != nil {
		return err
	}
	return nil
}

func (lh *LHClient) GetUsers() ([]*model.User, error) {
	users := []*model.User{}
	if err := lh.DB.Select(&users, qGetUsers); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("users not found")
		}
		return nil, err
	}
	return users, nil
}

func (lh *LHClient) GetUserByID(UserID int) (*model.User, error) {
	user := model.User{}
	if err := lh.DB.Get(&user, qGetUserByID, UserID); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (lh *LHClient) CreateUser(user *model.User) error {
	_, err := lh.DB.NamedExec(qCreateUser, user)
	if err != nil {
		return err
	}
	return nil
}

func (lh *LHClient) UpdateUser(UserID int, user *model.User) error {
	user.UserID = UserID
	_, err := lh.DB.NamedExec(qUpdateUser, user)
	if err != nil {
		return err
	}
	return nil
}
func (lh *LHClient) DeleteUser(UserID int) error {
	_, err := lh.DB.Exec(qDeleteUser, UserID)
	if err != nil {
		return err
	}
	return nil
}

// InsertUser inserts a new user into the users collection
func (h *MongoClient) InsertUser(user model.UserBson) (*mongo.InsertOneResult, error) {
	return h.DB.Collection("users").InsertOne(context.TODO(), user)
}

// FindUserByID finds a user by their ID in the users collection
func (h *MongoClient) FindUserByID(id primitive.ObjectID) (*model.UserBson, error) {
	var user model.UserBson
	collection := h.DB.Collection("users")
	timeout, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	err := collection.FindOne(timeout, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates an existing user in the users collection
func (h *MongoClient) UpdateUser(id primitive.ObjectID, update bson.M) (*mongo.UpdateResult, error) {
	timeout, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	return h.DB.Collection("users").UpdateOne(timeout, bson.M{"_id": id}, bson.M{"$set": update})
}

// DeleteUser deletes a user from the users collection
func (h *MongoClient) DeleteUser(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	timeout, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	return h.DB.Collection("users").DeleteOne(timeout, bson.M{"_id": id})
}

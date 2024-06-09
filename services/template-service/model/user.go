package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	UserID    int    `db:"user_id" json:"user_id"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"password"`
}

type UserBson struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"user_id"`
	FirstName string             `bson:"first_name" json:"first_name"`
	LastName  string             `bson:"last_name" json:"last_name"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
}

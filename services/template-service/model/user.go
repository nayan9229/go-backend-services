package model

type User struct {
	UserID    int    `db:"user_id" json:"user_id"`
	FirstName string `db:"first_name"  json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"password"`
}

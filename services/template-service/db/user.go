package db

import (
	"database/sql"
	"errors"

	"github.com/nayan9229/go-backend-services/services/template-service/model"
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

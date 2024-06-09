package server

import (
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/nayan9229/go-backend-services/chassis"
	"github.com/nayan9229/go-backend-services/services/template-service/model"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Server) GetUsers(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	users, err := s.sqlDb.GetUsers()
	if err != nil {
		log.Err(err).Msg("error while fetching user list")
		return chassis.BadRequest(w, r, err)
	}
	return users, nil
}

func (s *Server) GetUserByID(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	user_id := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(user_id)
	if err != nil {
		log.Err(err).Msg("error while fetching user by id")
		return chassis.BadRequest(w, r, err)
	}
	if userID <= 0 {
		err = errors.New("invalid user id")
		log.Err(err).Msg("error while fetching user by id")
		return chassis.BadRequest(w, r, err)
	}
	user, err := s.sqlDb.GetUserByID(userID)
	if err != nil {
		log.Err(err).Msg("error while fetching user by id")
		return chassis.BadRequest(w, r, err)
	}
	return user, nil
}

func (s *Server) CreateUsers(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var user model.User
	body, err := chassis.ReadBody(r, 0)
	if err != nil {
		log.Error().Err(err).Msg("chassis reading req body")
		return chassis.BadRequest(w, r, err)
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Error().Err(err).Msg("chassis reading req body")
		return chassis.BadRequest(w, r, err)
	}
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	if user.UserID <= 0 {
		user.UserID = int(rand.Uint32())
	}
	err = s.sqlDb.CreateUser(&user)
	if err != nil {
		log.Err(err).Msg("error while creating user")
		return chassis.BadRequest(w, r, err)
	}
	return chassis.NoContent(w, r)
}

func (s *Server) UpdateUserByID(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	user_id := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(user_id)
	if err != nil {
		log.Err(err).Msg("error while fetching user by id")
		return chassis.BadRequest(w, r, err)
	}
	if userID <= 0 {
		err = errors.New("invalid user id")
		log.Err(err).Msg("error while fetching user by id")
		return chassis.BadRequest(w, r, err)
	}

	var user model.User
	body, err := chassis.ReadBody(r, 0)
	if err != nil {
		log.Error().Err(err).Msg("chassis reading req body")
		return chassis.BadRequest(w, r, err)
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Error().Err(err).Msg("chassis reading req body")
		return chassis.BadRequest(w, r, err)
	}
	err = s.sqlDb.UpdateUser(userID, &user)
	if err != nil {
		log.Err(err).Msg("error while updating user")
		return chassis.BadRequest(w, r, err)
	}
	return user, nil
}

func (s *Server) DeleteUserByID(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	user_id := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(user_id)
	if err != nil {
		log.Err(err).Msg("error while fetching user by id")
		return chassis.BadRequest(w, r, err)
	}
	if userID <= 0 {
		err = errors.New("invalid user id")
		log.Err(err).Msg("error while fetching user by id")
		return chassis.BadRequest(w, r, err)
	}
	err = s.sqlDb.DeleteUser(userID)
	if err != nil {
		log.Err(err).Msg("error while deleting user")
		return chassis.BadRequest(w, r, err)
	}
	return chassis.NoContent(w, r)
}

func (s *Server) InsertUser(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var user model.UserBson
	body, err := chassis.ReadBody(r, 0)
	if err != nil {
		log.Error().Err(err).Msg("chassis reading req body")
		return chassis.BadRequest(w, r, err)
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Error().Err(err).Msg("chassis reading req body")
		return chassis.BadRequest(w, r, err)
	}

	_, err = s.jsonDb.InsertUser(user)
	if err != nil {
		log.Err(err).Msg("error while creating user")
		return chassis.BadRequest(w, r, err)
	}
	return chassis.NoContent(w, r)
}

func (s *Server) FindUserByID(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	UserID := chi.URLParam(r, "userID")
	objID, err := primitive.ObjectIDFromHex(UserID)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return chassis.BadRequest(w, r, err)
	}
	user, err := s.jsonDb.FindUserByID(objID)
	if err != nil {
		log.Err(err).Msg("error while fetching user by id")
		return chassis.BadRequest(w, r, err)
	}

	return user, nil
}

func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	UserID := chi.URLParam(r, "userID")
	objID, err := primitive.ObjectIDFromHex(UserID)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return chassis.BadRequest(w, r, err)
	}
	var user model.UserBson
	body, err := chassis.ReadBody(r, 0)
	if err != nil {
		log.Error().Err(err).Msg("chassis reading req body")
		return chassis.BadRequest(w, r, err)
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Error().Err(err).Msg("chassis reading req body")
		return chassis.BadRequest(w, r, err)
	}
	s.jsonDb.UpdateUser(objID, bson.M{
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
		"password":   user.Password,
	})
	return user, nil
}

func (s *Server) DeleteUser(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	UserID := chi.URLParam(r, "userID")
	objID, err := primitive.ObjectIDFromHex(UserID)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return chassis.BadRequest(w, r, err)
	}
	_, err = s.jsonDb.DeleteUser(objID)
	if err != nil {
		log.Err(err).Msg("error while deleting user")
		return chassis.BadRequest(w, r, err)
	}
	return chassis.NoContent(w, r)
}

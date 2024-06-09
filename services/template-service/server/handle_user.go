package server

import "net/http"

func (s *Server) GetUsers(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return "GET USERS", nil
}

func (s *Server) GetUserByID(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return "GET USER BY ID", nil
}

func (s *Server) CreateUsers(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return "CREAT USERS", nil
}

func (s *Server) UpdateUserByID(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return "UPDATE USER BY ID", nil
}

func (s *Server) DeleteUserByID(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return "DELETE USER BY ID", nil
}

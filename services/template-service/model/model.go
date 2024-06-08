package model

type Response struct {
	Message string `json:"message"`
}

type Request struct {
	Name string `json:"name"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
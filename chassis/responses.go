package chassis

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

type ErrResp struct {
	Message string `json:"message"`
}

func BadRequest(w http.ResponseWriter, r *http.Request, err error) (interface{}, error) {
	rsp := ErrResp{err.Error()}
	body, _ := json.Marshal(rsp)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(body)
	return nil, nil
}

func BadRequestWithMessage(w http.ResponseWriter, r *http.Request, msg string) (interface{}, error) {
	rsp := ErrResp{msg}
	body, _ := json.Marshal(rsp)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(body)
	return nil, nil
}

func NotFound(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	http.NotFound(w, nil)
	return nil, nil
}

// NotFoundWithMessage sets up an HTTP 404 Not Found with a given error
// message and returns the (nil, nil) pair used by SimpleHandler to
// signal that the response has been dealt with.
func NotFoundWithMessage(w http.ResponseWriter, r *http.Request, msg string) (interface{}, error) {
	rsp := ErrResp{msg}
	body, _ := json.Marshal(rsp)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(body)
	return nil, nil
}

// Forbidden sets up an HTTP 403 Forbidden and returns the (nil, nil)
// pair used by SimpleHandler to signal that the response has been
// dealt with.
func Forbidden(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	w.WriteHeader(http.StatusForbidden)
	return nil, nil
}

// Unauthorized sets up an HTTP 401 StatusUnauthorized and returns the (nil, nil)
// pair used by SimpleHandler to signal that the response has been dealt with.
func Unauthorized(w http.ResponseWriter, r *http.Request, msg string) (interface{}, error) {
	rsp := ErrResp{msg}
	body, _ := json.Marshal(rsp)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	w.Write(body)
	return nil, nil
}

// NoContent sets up an HTTP 204 No Content and returns the (nil, nil)
// pair used by SimpleHandler to signal that the response has been
// dealt with.
func NoContent(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	w.WriteHeader(http.StatusNoContent)
	return nil, nil
}

// TooManyRequests sets up an HTTP 429 Too Many Requests and returns the (nil, nil)
// pair used by SimpleHandler to signal that the response has been dealt with.
func TooManyRequests(w http.ResponseWriter, r *http.Request, retryAfter time.Duration) (interface{}, error) {
	if seconds := int(retryAfter.Seconds() + 0.5); seconds > 0 {
		w.Header().Set("Retry-After", strconv.Itoa(seconds))
	}
	w.WriteHeader(http.StatusTooManyRequests)
	return nil, nil
}

// BadRequest sets up an HTTP 400 Bad Request with a given error
// message and returns the (nil, nil) pair used by SimpleHandler to
// signal that the response has been dealt with.
func InternalServerError(w http.ResponseWriter, r *http.Request, err error) (interface{}, error) {
	log.Warn().Msgf("internal error: %s", err)
	w.WriteHeader(http.StatusInternalServerError)
	return nil, nil
}

// NotAcceptableError This response is sent when the web server,
// after performing server-driven content negotiation,
// doesn't find any content that conforms to the criteria given by the user agent.
func NotAcceptableError(w http.ResponseWriter, r *http.Request, msg string) (interface{}, error) {
	rsp := ErrResp{msg}
	body, _ := json.Marshal(rsp)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotAcceptable)
	w.Write(body)
	return nil, nil
}
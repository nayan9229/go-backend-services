package chassis

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/c2h5oh/datasize"
	"github.com/pkg/errors"
)

const defaultMaxBodySize = 10 * datasize.MB

// ReadBody reads a request body, limiting to a maximum size.
func ReadBody(r *http.Request, limit64 uint64) ([]byte, error) {
	limit := datasize.ByteSize(limit64)
	allowEmpty := false
	if limit == 1 {
		allowEmpty = true
	}
	if limit == 0 || limit == 1 {
		limit = defaultMaxBodySize
	}

	limiter := &io.LimitedReader{
		R: r.Body,
		N: int64(limit),
	}
	data, err := io.ReadAll(limiter)
	defer r.Body.Close()
	if err != nil || (!allowEmpty && len(data) == 0) {
		return nil, errors.New("Invalid request body")
	}
	if limiter.N <= 0 {
		return nil, errors.New(fmt.Sprintf("Request body too large (limit is %s)",
			limit.String()))
	}
	return data, nil
}

// Unmarshal does JSON unmarshalling disallowing unknown fields and
// limiting the permitted body size.
func Unmarshal(r io.ReadCloser, v interface{}) error {
	limiter := &io.LimitedReader{
		R: r,
		N: int64(defaultMaxBodySize),
	}
	dec := json.NewDecoder(limiter)
	defer r.Close()
	dec.DisallowUnknownFields()
	err := dec.Decode(&v)
	if err == nil {
		return nil
	}
	if limiter.N <= 0 {
		return errors.New(fmt.Sprintf("Request body too large (limit is %s)",
			defaultMaxBodySize.String()))
	}
	return errors.New("Invalid request body")
}

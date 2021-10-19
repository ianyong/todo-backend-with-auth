package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ianyong/todo-backend/internal/services"
)

type Response struct {
	Payload     json.RawMessage `json:"payload"`
	Messages    StatusMessages  `json:"messages"`
	Code        int             `json:"-"`
	AccessToken string          `json:"-"`
}

type Handler = func(*http.Request, *services.Services) (*Response, error)

// WrapHandler converts the internal Handler type into a standard http.HandlerFunc.
func WrapHandler(s *services.Services, handler Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		w.Header().Set("Content-Type", "application/json")

		res, err := handler(r, s)

		// Handle errors
		if err != nil {
			serveHTTPError(w, err)
			return
		}

		// Handle response
		serveHTTPResponse(w, res)
	}
}

// serveHTTPResponse takes in a http.ResponseWriter and a *Response and writes
// the appropriate response to the response body.
func serveHTTPResponse(w http.ResponseWriter, response *Response) {
	if response == nil {
		response = &Response{}
	}

	if response.Messages == nil {
		response.Messages = []StatusMessage{}
	}

	if response.AccessToken != "" {
		w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", response.AccessToken))
	}

	if response.Code > 0 {
		w.WriteHeader(response.Code)
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err)
	}
}

// serveHTTPError takes in a http.ResponseWriter and an error and writes
// the appropriate error response to the response body.
func serveHTTPError(w http.ResponseWriter, err error) {
	userErr, ok := asExternalError(err)

	// Since `err` is not a user-facing error, it is an internal error.
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(userErr.StatusCode())
	err = json.NewEncoder(w).Encode(&Response{
		Messages: StatusMessages{
			ErrorMessage(userErr.Error()),
		},
	})
	if err != nil {
		panic(err)
	}
}

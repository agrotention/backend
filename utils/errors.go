package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var (
	ErrInternal           = NewError(500, "internal server error")
	ErrInvalidPayload     = NewError(400, "invalid payload")
	ErrUnauthorized       = NewError(401, "unauthorized")
	ErrForbidden          = NewError(403, "permission denied")
	ErrNotFound           = NewError(404, "not found")
	ErrConflict           = NewError(409, "conflict")
	ErrTooManyRequests    = NewError(429, "too many requests")
	ErrServiceUnavailable = NewError(503, "service unavailable")
	ErrGatewayTimeout     = NewError(504, "gateway timeout")
	ErrNotImplemented     = NewError(501, "not implemented")
)

type HTTPError struct {
	code    int
	message string
}

func NewError(code int, message string) *HTTPError {
	return &HTTPError{code, message}
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP ERROR %d: %s", e.code, e.message)
}

func (e *HTTPError) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.code)
	err := json.NewEncoder(w).Encode(&Response{
		Success: false,
		Message: e.message,
	})
	// Fallback
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, `{ "success": false, "message": %q }`, e.message)
	}
}

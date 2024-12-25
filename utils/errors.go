package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Predefined common HTTP errors
var (
	ErrInternal           = NewErrWithMessage(http.StatusInternalServerError, "internal server error")
	ErrInvalidPayload     = NewErrWithMessage(http.StatusBadRequest, "invalid payload")
	ErrUnauthorized       = NewErrWithMessage(http.StatusUnauthorized, "unauthorized")
	ErrForbidden          = NewErrWithMessage(http.StatusForbidden, "permission denied")
	ErrNotFound           = NewErrWithMessage(http.StatusNotFound, "not found")
	ErrConflict           = NewErrWithMessage(http.StatusConflict, "conflict")
	ErrTooManyRequests    = NewErrWithMessage(http.StatusTooManyRequests, "too many requests")
	ErrServiceUnavailable = NewErrWithMessage(http.StatusServiceUnavailable, "service unavailable")
	ErrGatewayTimeout     = NewErrWithMessage(http.StatusGatewayTimeout, "gateway timeout")
	ErrNotImplemented     = NewErrWithMessage(http.StatusNotImplemented, "not implemented")
)

// ErrWithMessage represents an HTTP error with a status code and message
type ErrWithMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewErrWithMessage creates a new ErrWithMessage instance
func NewErrWithMessage(code int, message string) *ErrWithMessage {
	return &ErrWithMessage{Code: code, Message: message}
}

// Error implements the error interface for ErrWithMessage
func (e *ErrWithMessage) Error() string {
	return fmt.Sprintf("HTTP ERROR %d: %s", e.Code, e.Message)
}

// Send sends the error response as JSON to the client
func (e *ErrWithMessage) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Code)
	if err := json.NewEncoder(w).Encode(&Response[any]{
		Success: false,
		Message: e.Message,
	}); err != nil {
		log.Println("Failed to encode error response:", err)
		fallbackErrorResponse(w, e.Code, e.Message)
	}
}

// ErrWithData represents an HTTP error with additional data
type ErrWithData[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

// NewErrWithData creates a new ErrWithData instance
func NewErrWithData[T any](code int, message string, data T) *ErrWithData[T] {
	return &ErrWithData[T]{Code: code, Message: message, Data: data}
}

// Error implements the error interface for ErrWithData
func (e *ErrWithData[T]) Error() string {
	return fmt.Sprintf("HTTP ERROR %d: %s", e.Code, e.Message)
}

// Send sends the error response with data as JSON to the client
func (e *ErrWithData[T]) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Code)
	if err := json.NewEncoder(w).Encode(&Response[T]{
		Success: false,
		Message: e.Message,
		Data:    e.Data,
	}); err != nil {
		log.Println("Failed to encode error response with data:", err)
		fallbackErrorResponse(w, e.Code, e.Message)
	}
}

// fallbackErrorResponse writes a plain JSON error response as a fallback
func fallbackErrorResponse(w http.ResponseWriter, _ int, message string) {
	fallback := fmt.Sprintf(`{"success": false, "message": %q}`, message)
	_, err := w.Write([]byte(fallback))
	if err != nil {
		log.Println("Failed to write fallback error response:", err)
	}
}

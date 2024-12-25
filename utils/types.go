package utils

import "net/http"

type HTTPError interface {
	Send(http.ResponseWriter)
	Error() string
}

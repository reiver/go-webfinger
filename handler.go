package webfinger

import (
	"net/http"
)

// A Handler responds to a WebFinger request.
type Handler interface {
	ServeWebFinger(responseWriter http.ResponseWriter, resource string)
}

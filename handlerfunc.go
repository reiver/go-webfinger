package webfinger

import (
	"net/http"
)

// HandlerFunc can be used to turn a func(http.ResponseWriter,string) into a Handler.
type HandlerFunc func(responseWriter http.ResponseWriter, resource string)

// ServeWebFinger calls fn(responseWriter, resource)
func (fn HandlerFunc) ServeWebFinger(responseWriter http.ResponseWriter, resource string) {
	fn(responseWriter, resource)
}

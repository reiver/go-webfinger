package webfinger

import (
	"net/http"
)

// HandlerFunc can be used to turn a func(http.ResponseWriter,string) into a Handler.
type HandlerFunc func(responseWriter http.ResponseWriter, resource string, rels ...string)

var _ Handler = HandlerFunc(nil)

// ServeWebFinger calls fn(responseWriter, resource)
func (fn HandlerFunc) ServeWebFinger(responseWriter http.ResponseWriter, resource string, rels ...string) {
	fn(responseWriter, resource, rels...)
}

package webfinger

import (
	"net/http"
	"net/url"
)

// HTTPHandler returns an http.Handler based on the passed 'webFingerHandler'.
//
// You can think of it as a way of turning a webfinger.Handler into an http.Handler.
func HTTPHandler(webFingerHandler Handler) http.Handler {
	return internalHTTPHandler{
		handler: webFingerHandler,
	}
}

type internalHTTPHandler struct {
	handler Handler
}

func (receiver internalHTTPHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {

	if nil == responseWriter {
		return
	}
	if nil == request {
		httpError(responseWriter, http.StatusInternalServerError)
		return
	}

	var handler Handler = receiver.handler
	if nil == handler {
		httpError(responseWriter, http.StatusInternalServerError)
		return
	}

	var queryValues url.Values
	{
		var httpRequestURL *url.URL = request.URL
		if nil == httpRequestURL {
			httpError(responseWriter, http.StatusInternalServerError)
			return
		}

		queryValues = httpRequestURL.Query()
		if nil == queryValues {
			httpError(responseWriter, http.StatusBadRequest)
			return
		}
	}

	var resource string
	{
		var resources []string
		var found bool
		resources, found = queryValues["resource"]
		if !found {
			httpError(responseWriter, http.StatusBadRequest)
			return
		}
		if 1 != len(resources) {
			httpError(responseWriter, http.StatusBadRequest)
			return
		}

		resource = resources[0]
	}

	var rels []string = queryValues["rel"]

	var bytes []byte
	{
		var err error

		bytes, err = handler.ServeWebFinger(resource, rels...)
		if nil != err {
			switch casted := err.(type) {
			case ErrHTTP:
				httpError(responseWriter, casted.ErrHTTP())
				return
			default:
				httpError(responseWriter, http.StatusInternalServerError)
				return
			}
		}
	}

	ServeJRDBytes(responseWriter, request, bytes)
}

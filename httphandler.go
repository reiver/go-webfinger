package webfinger

import (
	"net/http"
	"net/url"
)

HTTPHandler returns an http.Handler based on the passed 'webFingerHandler'.
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

	var resource string
	{
		var httpRequestURL *url.URL = request.URL
		if nil == httpRequestURL {
			httpError(responseWriter, http.StatusInternalServerError)
			return
		}

		var query url.Values = httpRequestURL.Query()
		if nil == query {
			httpError(responseWriter, http.StatusBadRequest)
			return
		}

		var resources []string
		var found bool
		resources, found = query["resource"]
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

	{
		const contentType string = "application/jrd+json"

		var header http.Header = responseWriter.Header()
		if nil == header {
			httpError(responseWriter, http.StatusInternalServerError)
			return
		}

		header.Add("Content-Type", contentType)
	}

	handler.ServeWebFinger(responseWriter, resource)
}

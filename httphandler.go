package webfinger

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
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
			httpError(responseWriter, http.StatusInternalServerError)
			return
		}
	}

	var digest [sha256.Size]byte
	{
		digest = sha256.Sum256(bytes)
	}

	var cacheDigest string
	{
		cacheDigest = fmt.Sprintf("sha-256=:%s:", base64.StdEncoding.EncodeToString(digest[:]))
	}

	var eTag string
	{
		var format string = fmt.Sprintf("sha256=0x%%0%dX", sha256.Size*2)
		eTag = fmt.Sprintf(format, digest[:])
	}

	{
		const contentType string = "application/jrd+json"

		var header http.Header = responseWriter.Header()
		if nil == header {
			httpError(responseWriter, http.StatusInternalServerError)
			return
		}

		header.Add("Access-Control-Allow-Origin", "*")
		header.Add("Cache-Control", "max-age=907")
		header.Add("Content-Digest", cacheDigest)
		header.Add("Content-Type", contentType)
		header.Add("ETag", `"`+eTag+`"`)
	}

	{
		responseWriter.WriteHeader(http.StatusOK)
		responseWriter.Write(bytes)
	}
}

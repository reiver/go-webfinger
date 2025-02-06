package webfinger

import (
	"encoding/json"
	"io"
	"net/http"
	liburl "net/url"

	"github.com/reiver/go-erorr"
	"github.com/reiver/go-errhttp"
)

const (
	errBadHTTPURL          = erorr.Error("webfinger: bad http-url")
	errEmptyHost           = erorr.Error("webfinger: empty host")
	errEmptyResource       = erorr.Error("webfinger: empty resource")
	errNilDestination      = erorr.Error("webfinger: nil destination")
	errNilHTTPResponse     = erorr.Error("webfinger: nil http-response")
	errNilHTTPResponseBody = erorr.Error("webfinger: nil http-response-body")
	errNilParsedURL        = erorr.Error("webdinger: nil parsed-url")
)

// GetRaw makes an HTTP or HTTPS request (sending the HTTP request header "Accept: application/jrd+json")
// to any HTTP or HTTPS URL provided, and expects and interprets the HTTP response as JRD.
//
// For example:
//
//	import "github.com/reiver/go-webfinger"
//	
//	// ...
//	
//	var wfResponse webfinger.Response
//	
//	err := webfinger.GetRaw(&wfResponse, "https://example.com/.well-known/webfinger?resource=acct:joeblow@example.com")
//
// See also: [Get]
func GetRaw(response *Response, httpurl string) error {
	if nil == response {
		return errNilDestination
	}

	var httprequest http.Request
	{
		var urloc *liburl.URL
		{
			var err error
			urloc, err = liburl.Parse(httpurl)
			if nil != err {
				return erorr.Errorf("webfinger: problem parsing HTTP(S) URL %q: %w", httpurl, err)
			}
			if nil == urloc {
				return errNilParsedURL
			}
		}

		switch urloc.Scheme {
		case "http","https":
			// nothing here
		default:
			return erorr.Errorf("webfinger: not an HTTP(S) URL — %q", httpurl)
		}

		var header http.Header = http.Header{}
		header.Add("Accept", ContentTypeJRD)

		httprequest = http.Request{
			Method: http.MethodGet,
			URL: urloc,
			Header: header,
		}
	}

	var httpresponse *http.Response
	{
		var err error
		httpresponse, err = http.DefaultClient.Do(&httprequest)
		if nil != err {
			return erorr.Errorf("webfinger: problem making HTTP(S) request to %q: %w", httpurl, err)
		}
		if nil == httpresponse {
			return errNilHTTPResponse
		}
	}

	{
		if 400 <= httpresponse.StatusCode && httpresponse.StatusCode <= 599 {
			return errhttp.Return(httpresponse.StatusCode)
		}
		if 200 != httpresponse.StatusCode {
			return erorr.Errorf("webfinger: not HTTP 200 response from %q — %d", httpurl, httpresponse.StatusCode)
		}
	}

	var body []byte
	{
		if nil == httpresponse.Body {
			return errNilHTTPResponseBody
		}

		var err error
		body, err = io.ReadAll(httpresponse.Body)
		if nil != err {
			return erorr.Errorf("webfinger: problem reading-all from HTTP-response-body from %q: %w", httpurl, err)
		}
		httpresponse.Body.Close()
	}

	{
		err := json.Unmarshal(body, response)
		if nil != err {
			return erorr.Errorf("webfinger: problem unmarshal JRD JSON response from %q: %w", httpurl, err)
		}
	}

	return nil
}

// Get makes an HTTPS request (sending the HTTP request header "Accept: application/jrd+json")
// to the provided host to the default WebFinger Path with the provided resource as its query,
// and expects and interprets the HTTP response as JRD.
//
// For example:
//
//	import "github.com/reiver/go-webfinger"
//	
//	// ...
//	
//	var wfResponse webfinger.Response
//	
//	err := webfinger.Get(&wfResponse, "example.com", "acct:reiver@mastodon.social")
//
// See also: [GetRaw]
func Get(response *Response, host string, resource string) error {
	if "" == host {
		return errEmptyHost
	}
	if "" == resource {
		return errEmptyResource
	}

	var httpurl string = DefaultHTTPSURL(host, resource)
	if "" == httpurl {
		return errBadHTTPURL
	}

	return GetRaw(response, httpurl)
}

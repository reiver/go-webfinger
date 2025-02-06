# go-webfinger

Package **webfinger** provides tools for working with the **WebFinger** protocol, for the Go programming language.

WebFinger is a protocol, building on top of HTTP/HTTPS, that allows for discovery of things (such as people, users, objects, etc) that can be identified by a URI.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-webfinger

[![GoDoc](https://godoc.org/github.com/reiver/go-webfinger?status.svg)](https://godoc.org/github.com/reiver/go-webfinger)

## Client Examples

Here is an example on how to make a WebFinger request:

```golang
host := "example.com"
resource := "acct:reiver@mastodon.social"


var response webfinger.Response
err := webfinger.Get(&response, host, resource)
```

Alternatively, if you want to specify the example HTTPS URL, you can instead do something similar to:

```golang
url := "https://example.com/.well-known/webfinger?resource=acct:reiver@mastodon.social"

var response webfinger.Response
err := webfinger.GetRaw(&response, url)
```

## Server Examples

If you want to create a WebFinger server, you can either do something similar to:

```golang
func ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {

	// ...

	var jrdBytes []byte = ???


	// ...

	webfinger.ServeJRDBytes(responseWriter, request, jrdBytes)
}
```

Or, alternatively, you can do something similar to:

```golang
func serveWebFinger(responseWriter http.ResponseWriter, resource string, rels ...string) {
	//@TODO

	var response webfinger.Response

	// ...

	err := json.NewEncoder(responseWriter).Encode(response)

	//@TODO
}

var webFingerHandler webfinger.Handler = webfinger.HandlerFunc(serveWebFinger)

var httpHandler http.Handler = webfinger.HTTPHandler(webFingerHandler)

// ...

var path string = webfinger.webFingerHandler // == "/.well-known/webfinger"

// Replace this line of code with however you register handlers with your favorite HTTP mux.
http.Handle(path, httpHandler)
```

## Import

To import package **webfinger** use `import` code like the follownig:
```
import "github.com/reiver/go-webfinger"
```

## Installation

To install package **webfinger** do the following:
```
GOPROXY=direct go get github.com/reiver/go-webfinger
```

## Author

Package **webfinger** was written by [Charles Iliya Krempeaux](http://reiver.link)

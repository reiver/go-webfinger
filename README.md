# go-webfinger

Package **webfinger** provides tools for working with the **WebFinger** protocol, for the Go programming language.

WebFinger is a protocol, building on top of HTTP/HTTPS, that allows for discovery of things (such as people, users, objects, etc) that can be identified by a URI.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-webfinger

[![GoDoc](https://godoc.org/github.com/reiver/go-webfinger?status.svg)](https://godoc.org/github.com/reiver/go-webfinger)

## Example

Here is an example:

```golang
func serveWebFinger(responseWriter http.ResponseWriter, resource string, rels ...string) {
	//@TODO

	var response webfinger.Response

	// ...

	err := json.NewEncoder(responseWriter).Encode(response)

	//@TODO
}

var webFingerHandler webfinger.Handler = webfinger.HandlerFunc(serveWebFinger)

var httpHandler http.Handler = HTTPHandler(webFingerHandler)

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
GOPROXY=direct go get https://github.com/reiver/go-webfinger
```

## Author

Package **webfinger** was written by [Charles Iliya Krempeaux](http://reiver.link)

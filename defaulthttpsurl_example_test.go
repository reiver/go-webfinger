package webfinger_test

import (
	"fmt"

	"github.com/reiver/go-webfinger"
)

func ExampleDefaultURL() {

	var host string     = "example.com"
	var resource string = "acct:reiver@mastodon.social"

	var url string = webfinger.DefaultHTTPSURL(host, resource) // <---------

	fmt.Printf("URL: %s", url)

	// Output:
	// URL: https://example.com/.well-known/webfinger?resource=acct%3Areiver%40mastodon.social
}

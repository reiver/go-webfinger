package webfinger

import (
	liburl "net/url"
)

// DefaultHTTPSURL returns the default WebFinger URL for host and resource.
//
// (Note this does NOT account for "/.well-known/host-meta" and "/.well-known/host-meta.js".)
func DefaultHTTPSURL(host string, resource string) string {
	if "" == host {
		return ""
	}
	if "" == resource {
		return ""
	}

	var url liburl.URL
	url.Scheme = "https"
	url.Host   = host
	url.Path   = DefaultPath
	url.RawQuery = "resource=" + liburl.QueryEscape(resource)

	return url.String()
}

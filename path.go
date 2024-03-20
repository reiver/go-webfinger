package webfinger

// DefaultPath is the default well-known HTTP request path to WebFinger.
//
// Not that WebFinger COULD be located at a different HTTP request path.
// This would be specified using /.well-known/host-meta
const DefaultPath string = "/.well-known/webfinger"

package webfinger

// A Handler responds to a WebFinger request.
type Handler interface {
	ServeWebFinger(resource string, rels ...string) ([]byte, error)
}

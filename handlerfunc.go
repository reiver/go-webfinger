package webfinger

// HandlerFunc can be used to turn a func(string,...string)([]byte,error) into a Handler.
type HandlerFunc func(resource string, rels ...string) ([]byte, error)

var _ Handler = HandlerFunc(nil)

// ServeWebFinger calls fn(rresource, rels...)
func (fn HandlerFunc) ServeWebFinger(resource string, rels ...string) ([]byte, error) {
	return fn(resource, rels...)
}

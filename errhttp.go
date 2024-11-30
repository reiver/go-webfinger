package webfinger

// If webfinger.Handler's ServeWebFinger() method returns an error that is an webfinger.ErrHTTP,
// then webfinger.HTTPHandler() calls webfinger.ErrHTTP ErrHTTP() method for the HTTP code.
//
// For example, if ErrHTTP() returns 404, then the HTTP error returned it 404 (Not Found).
type ErrHTTP interface {
	ErrHTTP() int
}

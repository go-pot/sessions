package cookiestore

import (
	nSessions "gopkg.in/go-pot/sessions.v1"
	gSessions "github.com/gorilla/sessions"
)

// New returns a new CookieStore.
func New(keyPairs ...[]byte) nSessions.Store {
	return &cookieStore{gSessions.NewCookieStore(keyPairs...)}
}

type cookieStore struct {
	*gSessions.CookieStore
}

func (c *cookieStore) Options(options nSessions.Options) {
	c.CookieStore.Options = &gSessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HTTPOnly,
	}
}

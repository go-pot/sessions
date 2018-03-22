package filesystemstore

import (
	gSessions "github.com/gorilla/sessions"
	nSessions "gopkg.in/go-pot/sessions.v1"
)

// New returns a new FilesystemStore.
func New(path string, keyPairs ...[]byte) nSessions.Store {
	return &filesystemStore{gSessions.NewFilesystemStore(path, keyPairs...)}
}

type filesystemStore struct {
	*gSessions.FilesystemStore
}

func (c *filesystemStore) Options(options nSessions.Options) {
	c.FilesystemStore.Options = &gSessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HTTPOnly,
	}
}

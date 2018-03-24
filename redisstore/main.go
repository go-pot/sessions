package redisstore

import (
	"github.com/boj/redistore"
	gSessions "github.com/gorilla/sessions"
	nSessions "gopkg.in/go-pot/sessions.v1"
)

//New returns a new Redis store
func New(size int, network, address, password string, keyPairs ...[]byte) (nSessions.Store, error) {
	store, err := redistore.NewRediStore(size, network, address, password, keyPairs...)
	if err != nil {
		return nil, err
	}
	return &rediStore{store}, nil
}

//New returns a new Redis store
func NewWithDB(size int, network, address, password, db, keyPrefix string, keyPairs ...[]byte) (nSessions.Store, error) {
	store, err := redistore.NewRediStoreWithDB(size, network, address, password, db, keyPairs...)
	if err != nil {
		return nil, err
	}
	store.SetKeyPrefix(keyPrefix)
	return &rediStore{store}, nil
}

type rediStore struct {
	*redistore.RediStore
}

func (c *rediStore) Options(options nSessions.Options) {
	c.RediStore.Options = &gSessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HTTPOnly,
	}
	c.RediStore.SetSerializer(redistore.JSONSerializer{})
}

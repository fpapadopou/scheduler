package redis

import (
	"github.com/fpapadopou/scheduler/internal/storage"
	"github.com/go-redis/redis/v7"
)

// Store satisfies the Store interface. Jobs are stored in Redis.
type Store struct {
	c redis.Client
}

// Get returns a Job object by its ID.
func (s *Store) Get(ID int) (*storage.Job, error) {

	return nil, nil
}

// Put updates a Job in the database.
func (s *Store) Put(job storage.Job) error {

	return nil
}

// Delete removes a Job from the store.
func (s *Store) Delete(ID int) error {

	return nil
}

// New returns a new Redis store. Redis configuration is required.
func New(o redis.Options) *Store {
	return &Store{}
}

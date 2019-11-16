package redis

import (
	"github.com/fpapadopou/scheduler/internal/storage"
	"github.com/go-redis/redis/v7"
)

type Store struct {
	c redis.Client
}

func (s *Store) Get(ID int) (*storage.Job, error) {

	return nil, nil
}

func (s *Store) Put(job storage.Job) error {

	return nil
}

func (s *Store) Delete(ID int) error {

	return nil
}

func New() *Store {
	return &Store{}
}

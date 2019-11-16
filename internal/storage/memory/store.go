package memory

import "github.com/fpapadopou/scheduler/internal/storage"

type Store struct {
	jobs map[string]*storage.Job
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

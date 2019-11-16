package memory

import "github.com/fpapadopou/scheduler/internal/storage"

// Store satisfies the Store interface. Jobs are stored in map via pointers.
type Store struct {
	jobs map[string]*storage.Job
}

// Get returns a Job object by its ID.
func (s *Store) Get(ID int) (*storage.Job, error) {

	return nil, nil
}

// Put updates a Job in the jobs map.
func (s *Store) Put(job storage.Job) error {

	return nil
}

// Delete removes a Job from the store.
func (s *Store) Delete(ID int) error {

	return nil
}

// New returns a new memory Store object. No configuration is needed.
func New() *Store {
	return &Store{}
}

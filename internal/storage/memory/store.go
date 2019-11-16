package memory

import (
	"errors"
	"fmt"

	"github.com/fpapadopou/scheduler/internal/storage"
)

// Store satisfies the Store interface. Jobs are stored in map via pointers.
type Store struct {
	jobs map[int64]*storage.Job
}

// Get returns a Job object by its ID.
func (s *Store) Get(ID int64) (*storage.Job, error) {

	job, ok := s.jobs[ID]
	if !ok {
		return nil, errors.New(fmt.Sprintf("no job with ID %d exists", ID))
	}
	return job, nil
}

// Put updates a Job in the jobs map.
func (s *Store) Put(job storage.Job) error {

	// TODO: Add validation, make thread safe(?), get next ID if job does not exist.
	s.jobs[job.ID] = &job
	return nil
}

// Delete removes a Job from the store.
func (s *Store) Delete(ID int64) error {

	_, ok := s.jobs[ID]
	if !ok {
		return errors.New(fmt.Sprintf("no job with ID %d exists", ID))
	}
	delete(s.jobs, ID)
	return nil
}

// New returns a new memory Store object. No configuration is needed.
func New() *Store {
	return &Store{
		jobs: map[int64]*storage.Job{},
	}
}

package storage

// Job represents a job that has been scheduled for execution.
type Job struct {
	ID     int64
	Status string
}

// Store provides methods for retrieving and updating a Job in database.
type Store interface {
	Get(ID int64) (*Job, error)
	Put(job Job) error
	Delete(ID int64) error
}

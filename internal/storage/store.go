package storage

const (

)

type Job struct {
	ID     string
	Status string
}

type Store interface {
	Get(ID int) (*Job, error)
	Put(job Job) error
	Delete(ID int) error
}

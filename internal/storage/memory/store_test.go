package memory

import (
	"testing"

	"github.com/fpapadopou/scheduler/internal/storage"
	"github.com/stretchr/testify/assert"
)

var s *Store

func TestStore_Get(t *testing.T) {
	purgeStore()

	s.jobs[1] = &storage.Job{
		ID:     1,
		Status: "a status",
	}

	_, err := s.Get(100)
	assert.Error(t, err, "no job with ID 100 exists")

	res, err := s.Get(1)
	assert.EqualValues(t, 1, res.ID)
	assert.Equal(t, "a status", res.Status)
}

func TestStore_Put(t *testing.T) {
	purgeStore()

	job := storage.Job{ID: 50, Status: "stat"}
	err := s.Put(job)

	assert.NoError(t, err)

	res, err := s.Get(50)
	assert.NoError(t, err)
	assert.IsType(t, &storage.Job{}, res)
}

func TestStore_Delete(t *testing.T) {
	purgeStore()

	job := storage.Job{ID: 5, Status: "stat"}
	err := s.Put(job)
	assert.NoError(t, err)

	err = s.Delete(5)
	assert.NoError(t, err)

	_, err = s.Get(5)
	assert.Error(t, err, "no job with ID 5 exists")

	err = s.Delete(5000)
	assert.Error(t, err, "no job with ID 5000 exists")
}

func TestNew(t *testing.T) {
	assert.Empty(t, New().jobs)
}

func purgeStore() {
	s = New()
}

package redis

import (
	"testing"

	"github.com/fpapadopou/scheduler/internal/storage"
	"github.com/go-redis/redis/v7"
	"github.com/stretchr/testify/assert"
)

var s Store

func TestStore_Get(t *testing.T) {
	_, err := s.Get(1)
	assert.Error(t, err, "not implemented yet")
}

func TestStore_Put(t *testing.T) {
	j := storage.Job{}

	assert.Error(t, s.Put(j), "not implemented yet")
}

func TestStore_Delete(t *testing.T) {
	assert.Error(t, s.Delete(1), "not implemented yet")
}

func TestNew(t *testing.T) {
	assert.Empty(t, New(redis.Options{}).c)
}

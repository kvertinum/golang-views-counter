package store

import (
	"strconv"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/assert"
)

const addSteps = 5

func TestAddByName(t *testing.T) {
	s := miniredis.RunT(t)
	rdbConfig := &RedisConfig{
		RedisURL: s.Addr(),
	}

	r := NewRedisStore(rdbConfig)
	r.ConfigureStore()

	for i := 0; i < addSteps; i++ {
		r.AddByName("testName")
	}

	newUserValue, _ := r.rdb.Get(r.rctx, "testName").Result()
	intNewUserValue, _ := strconv.Atoi(newUserValue)

	assert.Equal(t, addSteps, intNewUserValue)
}

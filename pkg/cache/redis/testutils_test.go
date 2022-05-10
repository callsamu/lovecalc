package redis

import (
	"os"
	"testing"

	"github.com/go-redis/redis/v8"
)

func newTestRedisClient(t *testing.T) (*redis.Client, func()) {
	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		t.Error(err)
	}

	rdb := redis.NewClient(opt)

	return rdb, func() {
		defer rdb.Close()
	}
}

package redis

import (
	"os"
	"testing"

	"github.com/go-redis/redis/v8"
)

func newTestRedisClient(t *testing.T) (*redis.Client, func()) {
	url := os.Getenv("REDIS_URL")
	if url == "" {
		t.Fatal("redis url not supplied")
	}

	opt, err := redis.ParseURL(url)
	if err != nil {
		t.Fatal(err)
	}

	rdb := redis.NewClient(opt)

	return rdb, func() {
		defer rdb.Close()
	}
}

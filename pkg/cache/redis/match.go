package redis

import (
	"github.com/callsamu/lovecalc/pkg/core"
	"github.com/go-redis/redis/v8"
)

type MatchCache struct {
	Client redis.Client
}

func NewMatchCache(rdc redis.Client) *MatchCache {
	return &MatchCache{
		Client: rdc,
	}
}

func (mc *MatchCache) Get(c core.Couple) (*core.Match, error) {
	return nil, nil
}

func (mc *MatchCache) Set(c core.Couple, m *core.Match) error {
	return nil
}

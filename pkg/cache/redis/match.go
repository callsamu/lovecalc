package redis

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/callsamu/lovecalc/pkg/cache"
	"github.com/callsamu/lovecalc/pkg/core"
	"github.com/go-redis/redis/v8"
)

type MatchCache struct {
	Client *redis.Client
}

func NewMatchCache(rdc *redis.Client) *MatchCache {
	return &MatchCache{
		Client: rdc,
	}
}

func (mc *MatchCache) Get(c core.Couple) (*core.Match, error) {
	ctx := context.Background()
	key := c.FirstName + c.SecondName

	val, err := mc.Client.Get(ctx, key).Bytes()
	if err != nil {
		switch {
		case errors.Is(err, redis.Nil):
			return nil, cache.ErrKeyNotFound
		default:
			return nil, err
		}
	}

	m := &core.Match{}
	if err = json.Unmarshal(val, m); err != nil {
		return nil, err
	}

	return m, nil
}

func (mc *MatchCache) Set(c core.Couple, m *core.Match) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}

	key := c.FirstName + c.SecondName
	ctx := context.Background()

	return mc.Client.Set(ctx, key, b, time.Hour*24).Err()
}

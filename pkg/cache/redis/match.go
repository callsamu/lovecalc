package redis

import (
	"bytes"
	"encoding/gob"
	"errors"

	"github.com/callsamu/lovecalc/pkg/cache"
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
	val, err := mc.Client.Get(nil, c.FirstName+c.SecondName).Bytes()
	if err != nil {
		switch {
		case errors.Is(err, redis.Nil):
			return nil, cache.ErrKeyNotFound
		default:
			return nil, err
		}
	}

	var match core.Match
	b := bytes.NewReader(val)

	if err = gob.NewDecoder(b).Decode(&match); err != nil {
		return nil, err
	}

	return &match, nil
}

func (mc *MatchCache) Set(c core.Couple, m *core.Match) error {
	return nil
}

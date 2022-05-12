package mock

import (
	"github.com/callsamu/lovecalc/pkg/cache"
	"github.com/callsamu/lovecalc/pkg/core"
)

type MatchCache map[core.Couple]*core.Match

func NewMatchCache() MatchCache {
	return MatchCache{}
}

func (mc MatchCache) Get(c core.Couple) (*core.Match, error) {
	match, ok := mc[c]

	if !ok {
		return nil, cache.ErrKeyNotFound
	}

	return match, nil
}

func (mc MatchCache) Set(c core.Couple, m *core.Match) error {
	mc[c] = m
	return nil
}

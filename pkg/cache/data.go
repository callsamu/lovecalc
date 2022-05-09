package cache

import (
	"errors"

	"github.com/callsamu/lovecalc/pkg/core"
)

var ErrKeyNotFound = errors.New("cache: no matching key was found")

type MatchCache interface {
	Get(core.Couple) (*core.Match, error)
	Set(core.Couple, *core.Match) error
}

package redis

import (
	"reflect"
	"testing"

	"github.com/callsamu/lovecalc/pkg/core"
)

func TestGetCache(t *testing.T) {
	if testing.Short() {
		t.Skip("cache/redis: skipping integration test")
	}

	rdb, close := newTestRedisClient(t)
	defer close()

	mc := NewMatchCache(rdb)

	key := core.Couple{FirstName: "foo", SecondName: "bar"}
	want := &core.Match{Couple: key, CoupleName: "foobar", Probability: 2.0}

	err := mc.Set(key, want)
	if err != nil {
		t.Error(err)
	}

	got, err := mc.Get(key)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

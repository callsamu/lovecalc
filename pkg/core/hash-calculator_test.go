package core

import (
	"testing"
	"testing/quick"
)

func TestCalculatorProperties(t *testing.T) {
	c, _ := NewCalculator(HashWithSHA1)
	f := func(s1, s2 string) bool {
		r1 := c.Compute(s1, s2)
		r2 := c.Compute(s2, s1)

		t.Logf("testing %s + %s = %f\n", s1, s2, r1)
		return r1 == r2 && 0 <= r1 && r1 <= 1
	}

	config := &quick.Config{MaxCount: 100000}
	if err := quick.Check(f, config); err != nil {
		t.Error(err)
	}
}

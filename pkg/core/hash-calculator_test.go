package core

import (
	"testing"
	"testing/quick"
)

func TestCalculatorProperties(t *testing.T) {
	c, _ := NewCalculator(HashWithFNV)
	f := func(s1, s2 string) bool {
		r1 := c.Compute(s1, s2)
		r2 := c.Compute(s2, s1)

		t.Logf("testing %s + %s = %f\n", s1, s2, r1)
		return r1 == r2 && 0 < r1 && r1 < 1
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

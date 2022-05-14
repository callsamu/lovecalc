package core

import (
	"testing"
	"testing/quick"
)

func TestCalculatorProperties(t *testing.T) {
	c, _ := NewCalculator(HashWithSHA1)
	f := func(s1, s2 string) bool {
		r1 := c.Compute(Couple{s1, s2}).Probability
		r2 := c.Compute(Couple{s2, s1}).Probability

		return r1 == r2 && 0 <= r1 && r1 <= 1
	}

	config := &quick.Config{MaxCount: 1000}
	if err := quick.Check(f, config); err != nil {
		t.Error(err)
	}
}

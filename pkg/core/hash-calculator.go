package core

import (
	"encoding/binary"
	"fmt"
	"hash"
	"math"
)

const maxValue = math.MaxUint64

type HashCalculator struct {
	Hash hash.Hash
}

// Distance between two points p and q in an arc
func distance(p, q float64) float64 {
	min := math.Min(p, q)
	max := math.Max(p, q)

	d1 := max - min
	d2 := maxValue - max + min

	return math.Min(d1, d2)
}

func (c *HashCalculator) hashString(s string) uint64 {
	digest := c.Hash.Sum([]byte(s))
	return binary.BigEndian.Uint64(digest[c.Hash.Size()-8:]) // Take last 8 bytes of hash
}

// Compute output value for two names
func (c *HashCalculator) Compute(first, second string) float64 {
	hash1 := c.hashString(first)
	hash2 := c.hashString(second)

	fmt.Printf("%d for %s and %d for %s\n", hash1, first, hash2, second)

	distance := distance(float64(hash1), float64(hash2))

	return 1 - distance*2/maxValue
}

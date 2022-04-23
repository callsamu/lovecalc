package core

import (
	"encoding/binary"
	"hash"
	"math"
)

const maxValue = math.MaxUint32

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

func (c *HashCalculator) hash(b []byte) uint32 {
	digest := c.Hash.Sum(b)
	return binary.LittleEndian.Uint32(digest)
}

// Compute output value for two names
func (c *HashCalculator) Compute(first, second string) float64 {
	hash1 := c.hash([]byte(first))
	hash2 := c.hash([]byte(second))

	bytes := []byte{0x0, 0x0, 0x0, 0x0}
	binary.LittleEndian.PutUint32(bytes, hash1+hash2)
	hash3 := c.hash(bytes)

	return float64(hash3) / maxValue
}

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

func (c *HashCalculator) hash(b []byte) uint32 {
	c.Hash.Write(b)
	digest := c.Hash.Sum(nil)
	c.Hash.Reset()
	return binary.LittleEndian.Uint32(digest)
}

// Compute output value for two names
func (c *HashCalculator) Compute(first, second string) float64 {
	hash1 := c.hash([]byte(first))
	hash2 := c.hash([]byte(second))

	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, hash1+hash2)
	hash3 := c.hash(bytes)

	return float64(hash3) / maxValue
}

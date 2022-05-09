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
	digest := c.Hash.Sum(b)
	return binary.LittleEndian.Uint32(digest)
}

// Compute output value for two names
func (c *HashCalculator) Compute(couple Couple) float64 {
	hash1 := c.hash([]byte(couple.FirstName))
	hash2 := c.hash([]byte(couple.SecondName))

	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, hash1+hash2)
	hash3 := c.hash(bytes)

	return float64(hash3) / maxValue
}

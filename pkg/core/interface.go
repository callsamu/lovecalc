package core

import (
	"crypto/sha1"
	"errors"
	"hash/fnv"
)

const (
	HashWithFNV  = "HashWithFNV"
	HashWithSHA1 = "HashWithSHA1"
)

type Calculator interface {
	Compute(string, string) float64
}

func NewCalculator(algorithm string) (Calculator, error) {
	switch algorithm {
	case HashWithFNV:
		return &HashCalculator{Hash: fnv.New64()}, nil
	case HashWithSHA1:
		return &HashCalculator{Hash: sha1.New()}, nil
	default:
		return nil, errors.New("Unrecognized algorithm selected")
	}
}

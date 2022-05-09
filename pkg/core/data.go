package core

import (
	"crypto/sha1"
	"errors"
	"hash/adler32"
	"hash/fnv"
)

const (
	HashWithFNV     = "HashWithFNV"
	HashWithSHA1    = "HashWithSHA1"
	HashWithAdler32 = "HashWithAdler32"
)

// Simple PODs for containing input and results
// used for serialization
type Couple struct {
	FirstName  string
	SecondName string
}

type Match struct {
	Couple
	CoupleName  string
	Probability float64
}

// The Calculator generic interface
type Calculator interface {
	Compute(Couple) *Match
}

func NewCalculator(algorithm string) (Calculator, error) {
	switch algorithm {
	case HashWithFNV:
		return &HashCalculator{Hash: fnv.New128()}, nil
	case HashWithSHA1:
		return &HashCalculator{Hash: sha1.New()}, nil
	case HashWithAdler32:
		return &HashCalculator{Hash: adler32.New()}, nil
	default:
		return nil, errors.New("Unrecognized algorithm selected")
	}
}

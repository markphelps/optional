// Code generated by go generate
// This file was generated by robots

package optional

import (
	"math/big"
	"errors"
)

// BigInt is an optional big.Int.
type BigInt struct {
	value *big.Int
}

// NewBigInt creates an optional.BigInt from a big.Int.
func NewBigInt(v big.Int) BigInt {
	return BigInt{&v}
}

// Set sets the big.Int value.
func (b *BigInt) Set(v big.Int) {
	b.value = &v
}

// Get returns the big.Int value or an error if not present.
func (b BigInt) Get() (big.Int, error) {
	if !b.Present() {
		var zero big.Int
		return zero, errors.New("value not present")
	}
	return *b.value, nil
}

// MustGet returns the big.Int value or panics if not present.
func (b BigInt) MustGet() big.Int {
	if !b.Present() {
		panic("value not present")
	}
	return *b.value
}

// Present returns whether or not the value is present.
func (b BigInt) Present() bool {
	return b.value != nil
}

// OrElse returns the big.Int value or a default value if the value is not present.
func (b BigInt) OrElse(v big.Int) big.Int {
	if b.Present() {
		return *b.value
	}
	return v
}

// If calls the function f with the value if the value is present.
func (b BigInt) If(fn func(big.Int)) {
	if b.Present() {
		fn(*b.value)
	}
}

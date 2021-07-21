// Code generated by go generate
// This file was generated by robots

package optional

import (
	"errors"
)

// Complex128 is an optional complex128.
type Complex128 struct {
	value *complex128
}

// NewComplex128 creates an optional.Complex128 from a complex128.
func NewComplex128(v complex128) Complex128 {
	return Complex128{&v}
}

// Set sets the complex128 value.
func (c *Complex128) Set(v complex128) {
	c.value = &v
}

// Get returns the complex128 value or an error if not present.
func (c Complex128) Get() (complex128, error) {
	if !c.Present() {
		var zero complex128
		return zero, errors.New("value not present")
	}
	return *c.value, nil
}

// MustGet returns the complex128 value or panics if not present.
func (c Complex128) MustGet() complex128 {
	if !c.Present() {
		panic("value not present")
	}
	return *c.value
}

// Present returns whether or not the value is present.
func (c Complex128) Present() bool {
	return c.value != nil
}

// OrElse returns the complex128 value or a default value if the value is not present.
func (c Complex128) OrElse(v complex128) complex128 {
	if c.Present() {
		return *c.value
	}
	return v
}

// If calls the function f with the value if the value is present.
func (c Complex128) If(fn func(complex128)) {
	if c.Present() {
		fn(*c.value)
	}
}

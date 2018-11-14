// Code generated by go generate
// This file was generated by robots at 2018-11-13 21:59:14.409667 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
)

// Uintptr is an optional uintptr
type Uintptr struct {
	value *uintptr
}

// NewUintptr creates a optional.Uintptr from a uintptr
func NewUintptr(v uintptr) Uintptr {
	return Uintptr{&v}
}

// Set sets the uintptr value
func (u *Uintptr) Set(v uintptr) {
	u.value = &v
}

// Get returns the uintptr value or an error if not present
func (u Uintptr) Get() (uintptr, error) {
	if !u.Present() {
		var zero uintptr
		return zero, errors.New("value not present")
	}
	return *u.value, nil
}

// Present returns whether or not the value is present
func (u Uintptr) Present() bool {
	return u.value != nil
}

// OrElse returns the uintptr value or a default value if the value is not present
func (u Uintptr) OrElse(v uintptr) uintptr {
	if u.Present() {
		return *u.value
	}
	return v
}

// If calls the function f with the value if the value is present
func (u Uintptr) If(fn func(uintptr)) {
	if u.Present() {
		fn(*u.value)
	}
}

func (u Uintptr) MarshalJSON() ([]byte, error) {
	if u.Present() {
		return json.Marshal(u.value)
	}
	var zero uintptr
	return json.Marshal(zero)
}

func (u *Uintptr) UnmarshalJSON(data []byte) error {
	var value uintptr

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	u.value = &value
	return nil
}

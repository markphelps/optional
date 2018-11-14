// Code generated by go generate
// This file was generated by robots at 2018-11-14 01:19:22.845542 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
)

// Int16 is an optional int16.
type Int16 struct {
	value *int16
}

// NewInt16 creates an optional.Int16 from a int16.
func NewInt16(v int16) Int16 {
	return Int16{&v}
}

// Set sets the int16 value.
func (i *Int16) Set(v int16) {
	i.value = &v
}

// Get returns the int16 value or an error if not present.
func (i Int16) Get() (int16, error) {
	if !i.Present() {
		var zero int16
		return zero, errors.New("value not present")
	}
	return *i.value, nil
}

// Must returns the int16 value or panics if not present.
func (i Int16) Must() int16 {
	if !i.Present() {
		panic("value not present")
	}
	return *i.value
}

// Present returns whether or not the value is present.
func (i Int16) Present() bool {
	return i.value != nil
}

// OrElse returns the int16 value or a default value if the value is not present.
func (i Int16) OrElse(v int16) int16 {
	if i.Present() {
		return *i.value
	}
	return v
}

// If calls the function f with the value if the value is present.
func (i Int16) If(fn func(int16)) {
	if i.Present() {
		fn(*i.value)
	}
}

func (i Int16) MarshalJSON() ([]byte, error) {
	if i.Present() {
		return json.Marshal(i.value)
	}
	var zero int16
	return json.Marshal(zero)
}

func (i *Int16) UnmarshalJSON(data []byte) error {
	var value int16

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	i.value = &value
	return nil
}

// Code generated by go generate
// This file was generated by robots at 2018-12-06 14:07:57.888442 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
)

// Int8 is an optional int8.
type Int8 struct {
	value *int8
}

// NewInt8 creates an optional.Int8 from a int8.
func NewInt8(v int8) Int8 {
	return Int8{&v}
}

// Set sets the int8 value.
func (i *Int8) Set(v int8) {
	i.value = &v
}

// Get returns the int8 value or an error if not present.
func (i Int8) Get() (int8, error) {
	if !i.Present() {
		var zero int8
		return zero, errors.New("value not present")
	}
	return *i.value, nil
}

// Present returns whether or not the value is present.
func (i Int8) Present() bool {
	return i.value != nil
}

// OrElse returns the int8 value or a default value if the value is not present.
func (i Int8) OrElse(v int8) int8 {
	if i.Present() {
		return *i.value
	}
	return v
}

// If calls the function f with the value if the value is present.
func (i Int8) If(fn func(int8)) {
	if i.Present() {
		fn(*i.value)
	}
}

func (i Int8) MarshalJSON() ([]byte, error) {
	if i.Present() {
		return json.Marshal(i.value)
	}
	return json.Marshal(nil)
}

func (i *Int8) UnmarshalJSON(data []byte) error {

	if string(data) == "null" {
		s.value = nil
		return nil
	}

	var value int8

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	i.value = &value
	return nil
}

// Code generated by go generate
// This file was generated by robots at 2021-05-03 20:21:08.402824 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
)

// Error is an optional error.
type Error struct {
	value *error
}

// NewError creates an optional.Error from a error.
func NewError(v error) Error {
	return Error{&v}
}

// Set sets the error value.
func (e *Error) Set(v error) {
	e.value = &v
}

// Get returns the error value or an error if not present.
func (e Error) Get() (error, error) {
	if !e.Present() {
		var zero error
		return zero, errors.New("value not present")
	}
	return *e.value, nil
}

// Present returns whether or not the value is present.
func (e Error) Present() bool {
	return e.value != nil
}

// OrElse returns the error value or a default value if the value is not present.
func (e Error) OrElse(v error) error {
	if e.Present() {
		return *e.value
	}
	return v
}

// If calls the function f with the value if the value is present.
func (e Error) If(fn func(error)) {
	if e.Present() {
		fn(*e.value)
	}
}

func (e Error) MarshalJSON() ([]byte, error) {
	if e.Present() {
		return json.Marshal(e.value)
	}
	return json.Marshal(nil)
}

func (e *Error) UnmarshalJSON(data []byte) error {

	if string(data) == "null" {
		e.value = nil
		return nil
	}

	var value error

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	e.value = &value
	return nil
}

//go:build 1.18
// +build 1.18

package optional

import (
	"encoding/json"
    "errors"
)

// Opt is an optional T.
type Opt[T any] struct {
	value *T
}

// New creates an optional.T from a T.
func New[T any](v T) Opt[T] {
	o := Opt[T]{value: &v}
	return o
}

// Set sets the value.
func (o *Opt[T]) Set(v T) {
	o.value = &v
}

// Get returns the value or an error if not present.
func (o Opt[T]) Get() (T, error) {
	if !o.Present() {
		var zero T
		return zero, errors.New("value not present")
	}
	return *o.value, nil
}

// MustGet returns the value or panics if not present.
func (o Opt[T]) MustGet() T {
	if !o.Present() {
		panic("value not present")
	}
	return *o.value
}

// Present returns whether or not the value is present.
func (o Opt[T]) Present() bool {
	return o.value != nil
}

// OrElse returns the value or a default value if the value is not present.
func (o Opt[T]) OrElse(v T) T {
	if o.Present() {
		return *o.value
	}
	return v
}

// If calls the function f with the value if the value is present.
func (o Opt[T]) If(fn func(T)) {
	if o.Present() {
		fn(*o.value)
	}
}

func (o Opt[T]) MarshalJSON() ([]byte, error) {
	if o.Present() {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

func (o *Opt[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = nil
		return nil
	}

	var value T

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	o.value = &value
	return nil
}

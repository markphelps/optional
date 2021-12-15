//go:build 1.18
// +build 1.18

package optional

import (
	"encoding/json"
    "errors"
)

// Optional is an optional T.
type Optional[T any] struct {
	value *T
}

// New creates an optional T from a T.
func New[T any](v T) Optional[T] {
	o := Optional[T]{value: &v}
	return o
}

// Set sets the value.
func (o *Optional[T]) Set(v T) {
	o.value = &v
}

// Get returns the value or an error if not present.
func (o Optional[T]) Get() (T, error) {
	if !o.Present() {
		var zero T
		return zero, errors.New("value not present")
	}
	return *o.value, nil
}

// MustGet returns the value or panics if not present.
func (o Optional[T]) MustGet() T {
	if !o.Present() {
		panic("value not present")
	}
	return *o.value
}

// Present returns whether or not the value is present.
func (o Optional[T]) Present() bool {
	return o.value != nil
}

// OrElse returns the value or a default value if the value is not present.
func (o Optional[T]) OrElse(v T) T {
	if o.Present() {
		return *o.value
	}
	return v
}

// If calls the function fn with the value if the value is present.
func (o Optional[T]) If(fn func(T)) {
	if o.Present() {
		fn(*o.value)
	}
}

// MarshalJSON marshals a type T to JSON if the value is present.
// Otherwise it returns the result of marshalling the nil value.
//
// Note: Types T of channel, complex, and function values cannot be encoded in JSON.
// Attempting to encode such a value causes Marshal to return
// a json.UnsupportedTypeError.
//
// See: https://github.com/golang/go/blob/129bb1917b4914f0743ec9b4ef0dfb74df39c07d/src/encoding/json/encode.go#L150
func (o Optional[T]) MarshalJSON() ([]byte, error) {
	if o.Present() {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result
// in the value pointed to by o.value. If v is nil or not a pointer,
// Unmarshal returns an InvalidUnmarshalError.
//
// Note: If a JSON value is not appropriate for a given target type,
// or if a JSON number overflows the target type, Unmarshal
// skips that field and completes the unmarshaling as best it can.
// If no more serious errors are encountered, Unmarshal returns
// an UnmarshalTypeError describing the earliest such error. In any
// case, it's not guaranteed that all the remaining fields following
// the problematic one will be unmarshaled into the target object.
//
// See: https://github.com/golang/go/blob/129bb1917b4914f0743ec9b4ef0dfb74df39c07d/src/encoding/json/decode.go#L78
func (o *Optional[T]) UnmarshalJSON(data []byte) error {
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

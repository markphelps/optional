// Code generated by 'go generate'

package bar

import (
	"encoding/json"
	"errors"
)

// optionalBar is an optional bar.
type optionalBar struct {
	value *bar
}

// NewoptionalBar creates an optional.optionalBar from a bar.
func NewoptionalBar(v bar) optionalBar {
	return optionalBar{&v}
}

// Set sets the bar value.
func (o *optionalBar) Set(v bar) {
	o.value = &v
}

// Get returns the bar value or an error if not present.
func (o optionalBar) Get() (bar, error) {
	if !o.Present() {
		var zero bar
		return zero, errors.New("value not present")
	}
	return *o.value, nil
}

// MustGet returns the bar value or panics if not present.
func (o optionalBar) MustGet() bar {
	if !o.Present() {
		panic("value not present")
	}
	return *o.value
}

// Present returns whether or not the value is present.
func (o optionalBar) Present() bool {
	return o.value != nil
}

// OrElse returns the bar value or a default value if the value is not present.
func (o optionalBar) OrElse(v bar) bar {
	if o.Present() {
		return *o.value
	}
	return v
}

// If calls the function f with the value if the value is present.
func (o optionalBar) If(fn func(bar)) {
	if o.Present() {
		fn(*o.value)
	}
}

func (o optionalBar) MarshalJSON() ([]byte, error) {
	if o.Present() {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

func (o *optionalBar) UnmarshalJSON(data []byte) error {

	if string(data) == "null" {
		o.value = nil
		return nil
	}

	var value bar

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	o.value = &value
	return nil
}

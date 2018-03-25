// Code generated by go generate
// This file was generated by robots at 2018-03-25 18:58:19.709353734 +0000 UTC

package foo

import (
    "encoding/json"
    "errors"
)

// OptionalFoo is an optional Foo
type OptionalFoo struct {
    value *Foo
}

// NewOptionalFoo creates a optional.OptionalFoo from a Foo
func NewOptionalFoo(v Foo) OptionalFoo {
    return OptionalFoo{&v}
}

// Set sets the Foo value
func (o OptionalFoo) Set(v Foo) {
    o.value = &v
}

// Get returns the Foo value or an error if not present
func (o OptionalFoo) Get() (Foo, error) {
    if !o.Present() {
        return *o.value, errors.New("value not present")
    }
    return *o.value, nil
}

// Present returns whether or not the value is present
func (o OptionalFoo) Present() bool {
    return o.value != nil
}

// OrElse returns the Foo value or a default value if the value is not present
func (o OptionalFoo) OrElse(v Foo) Foo {
    if o.Present() {
        return *o.value
    }
    return v
}

// If calls the function f with the value if the value is present
func (o OptionalFoo) If(fn func(Foo)) {
    if o.Present() {
        fn(*o.value)
    }
}

func (o OptionalFoo) MarshalJSON() ([]byte, error) {
    if o.Present() {
        return json.Marshal(o.value)
    }
    return nil, nil
}

func (o *OptionalFoo) UnmarshalJSON(data []byte) error {
    if len(data) < 1 {
        o.value = nil
        return nil
    }

    o.value = new(Foo)
    return json.Unmarshal(data, o.value)
}

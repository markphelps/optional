// Code generated by go generate
// This file was generated by robots at 2017-07-24 00:47:11.237626324 +0000 UTC

package optional

// Error is an optional error
type Error struct {
	error   error
	present bool
}

// EmptyError returns an empty optional.Error
func EmptyError() Error {
	return Error{}
}

// OfError creates a optional.Error from a error
func OfError(e error) Error {
	return Error{error: e, present: true}
}

// Set sets the error value
func (o *Error) Set(e error) {
	o.error = e
	o.present = true
}

// Error returns the error value
func (o *Error) Error() error {
	return o.error
}

// Present returns whether or not the value is present
func (o *Error) Present() bool {
	return o.present
}

// OrElse returns the error value or a default value if the value is not present
func (o *Error) OrElse(e error) error {
	if o.present {
		return o.error
	}
	return e
}

# Optional

[![Release](https://img.shields.io/github/release/markphelps/optional.svg?style=flat-square)](https://github.com/markphelps/optional/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/markphelps/optional)
[![Go Report Card](https://goreportcard.com/badge/github.com/markphelps/optional?style=flat-square)](https://goreportcard.com/report/github.com/markphelps/optional)
[![SayThanks.io](https://img.shields.io/badge/SayThanks.io-%E2%98%BC-1EAEDB.svg?style=flat-square)](https://saythanks.io/to/markphelps)

## Motivation

Ever had to write a test where you want to assert something only if a value is present?

```go
tests :=  []struct {
  schwifty string
  schwiftyPresent bool
} {
    { "wubba dub dub", true },
    { "", false },
}

...

if test.schwiftyPresent {
  assert.Equal(t, test.schwifty, schwiftyValue)
}
```

Now you can simplify all that with `optional` types:

```go
tests :=  []struct {
  schwifty optional.String
} {
    { optional.OfString("wubba dub dub") },
    { optional.EmptyString() },
  }
}

...

if test.schwifty.Present() {
  assert.Equal(t, test.schwifty.Get(), schwiftyValue)
}
```

## Inspiration

* Java [Optional](https://docs.oracle.com/javase/8/docs/api/java/util/Optional.html) type.
* [https://github.com/leighmcculloch/go-optional](https://github.com/leighmcculloch/go-optional)

## Install

`go get -u github.com/markphelps/optional/cmd/optional`

## Usage

Optional is a tool that generates 'optional' type wrappers around a given type T.

Typically this process would be run using go generate, like this:

	//go:generate optional -type=Foo

running this command:

	optional -type=Foo

in the same directory will create the file optional_foo.go
containing a definition of:

	type OptionalFoo struct {
		...
	}

The default type is OptionalT or optionalT (depending on if the type is exported)
and output file is optional_t.go. This can be overridden with the -output flag.

## Library

Optional is also a library that provides optional types for the basic Go types:

* bool
* byte
* complex128
* complex64
* float32
* float64
* int
* int16
* int32
* int64
* int8
* rune
* string
* uint
* uint16
* uint32
* uint64
* uint8
* uintptr
* error

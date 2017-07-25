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
  catchPhrase string
  catchPhrasePresent bool
} {
    { "wubba dub dub", true },
    { "", false },
}

...

if test.catchPhrasePresent {
  assert.Equal(t, test.catchPhrase, catchPhrase)
}
```

Now you can simplify all that with `optional` types:

```go
tests :=  []struct {
  catchPhrase optional.String
} {
    { optional.OfString("wubba dub dub") },
    { optional.EmptyString() },
  }
}

...

if test.catchPhrase.Present() {
  assert.Equal(t, test.catchPhrase.Get(), catchPhrase)
}
```

## Inspiration

* Java [Optional](https://docs.oracle.com/javase/8/docs/api/java/util/Optional.html)
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

* [bool](bool.go)
* [byte](byte.go)
* [complex128](complex128.go)
* [complex64](complex64.go)
* [float32](float32.go)
* [float64](float64.go)
* [int](int.go)
* [int16](int16.go)
* [int32](int32.go)
* [int64](int64.go)
* [int8](int8.go)
* [rune](rune.go)
* [string](string.go)
* [uint](uint.go)
* [uint16](uint16.go)
* [uint32](uint32.go)
* [uint64](uint64.go)
* [uint8](uint8.go)
* [uintptr](uintptr.go)
* [error](error.go)

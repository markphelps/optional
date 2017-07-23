# Optional

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

[![Release](https://img.shields.io/github/release/markphelps/optional.svg?style=flat-square)](https://github.com/markphelps/optional/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/markphelps/optional)
[![Go Report Card](https://goreportcard.com/badge/github.com/markphelps/optional?style=flat-square)](https://goreportcard.com/report/github.com/markphelps/optional)
[![SayThanks.io](https://img.shields.io/badge/SayThanks.io-%E2%98%BC-1EAEDB.svg?style=flat-square)](https://saythanks.io/to/markphelps)

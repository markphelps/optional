# Optional

[![Build Status](https://travis-ci.org/markphelps/optional.svg?branch=master)](https://travis-ci.org/markphelps/optional)
[![Release](https://img.shields.io/github/release/markphelps/optional.svg?style=flat-square)](https://github.com/markphelps/optional/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/markphelps/optional)
[![Go Report Card](https://goreportcard.com/badge/github.com/markphelps/optional?style=flat-square)](https://goreportcard.com/report/github.com/markphelps/optional)
[![SayThanks.io](https://img.shields.io/badge/SayThanks.io-%E2%98%BC-1EAEDB.svg?style=flat-square)](https://saythanks.io/to/markphelps)

Optional is a library that provides option types for Go using generics.

## Motivation

In Go, variables declared without an explicit initial value are given their zero value. Most of the time this is what you want, but sometimes you want to be able to tell if a variable was set or if it's just a zero value. That's where [option types](https://en.wikipedia.org/wiki/Option_type) come in handy.

## Inspiration

* Java [Optional](https://docs.oracle.com/javase/8/docs/api/java/util/Optional.html)
* [https://github.com/leighmcculloch/go-optional](https://github.com/leighmcculloch/go-optional)
* [https://github.com/golang/go/issues/7054](https://github.com/golang/go/issues/7054)

### Usage

```go
package main

import (
    "fmt"

    "github.com/markphelps/optional"
)

func main() {
    s := optional.New("foo")

    value, err := s.Get()
    if err != nil {
        // handle error!
    } else {
        fmt.Println(value)
    }

    t := optional.Optional[string]{}
    fmt.Println(t.OrElse("bar"))
}
```

See [example_test.go](example_test.go) and the [documentation](http://godoc.org/github.com/markphelps/optional) for more usage.

## Marshalling/Unmarshalling JSON

**Note:** v0.6.0 introduces a potential breaking change to anyone depending on marshalling non-present values to their zero values instead of null. See: [#9](https://github.com/markphelps/optional/pull/9) for more context.

Option types marshal to/from JSON as you would expect:

### Marshalling

```go
package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    var value = struct {
        Field optional.Optional[string] `json:"field,omitempty"`
    }{
        Field: optional.New("bar"),
    }

    out, _ := json.Marshal(value)

    fmt.Println(string(out))
    // outputs: {"field":"bar"}
}
```

### Unmarshalling

```go
package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    var value = &struct {
        Field optional.Optional[string] `json:"field,omitempty"`
    }{}

    _ = json.Unmarshal([]byte(`{"field":"bar"}`), value)

    value.Field.If(func(s string) {
        fmt.Println(s)
    })
    // outputs: bar
}
```

See [example_test.go](example_test.go) for more examples.

## Contributing

1. [Fork it](https://github.com/markphelps/optional/fork)
1. Create your feature branch (`git checkout -b my-new-feature`)
1. Commit your changes (`git commit -am 'Add some feature'`)
1. Push to the branch (`git push origin my-new-feature`)
1. Create a new Pull Request

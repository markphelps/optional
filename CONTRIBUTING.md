# Contributing

## Setup your machine

`optional` is written in [Go](https://golang.org/).

Prerequisites are:

- `make`
- [Go 1.18+](http://golang.org/doc/install)

Install the build and lint dependencies:

``` sh
make setup
```

A good way of making sure everything is all right is running the test suite:

``` sh
make test
```

## Test your change

When you are satisfied with the changes, we suggest you run:

``` sh
make ci
```

Which runs all the linters and tests.

## Submit a pull request

Push your branch to your `optional` fork and open a pull request against the
master branch.

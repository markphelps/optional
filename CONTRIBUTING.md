# Contributing

## Setup your machine

`optional` is written in [Go](https://golang.org/).

Prerequisites are:

* Build:
  * `make`
  * [Go 1.10+](http://golang.org/doc/install)

Clone `optional` from source into `$GOPATH`:

```sh
$ go get github.com/markphelps/optional
$ cd $GOPATH/src/github.com/markphelps/optional
```

Install the build and lint dependencies:

``` sh
$ make setup
```

A good way of making sure everything is all right is running the test suite:

``` sh
$ make test
```

## Test your change

You can create a branch for your changes and try to build from the source as you go:

``` sh
$ make build
```

### Goldenfiles

If you change the public API, you may have to update the goldenfiles in [cmd/optional](cmd/optional).

To do so, run:

```sh
$ TEST_OPTIONS=-update make test
```

When you are satisfied with the changes, we suggest you run:

``` sh
$ make ci
```

Which runs all the linters and tests.

## Submit a pull request

Push your branch to your `optional` fork and open a pull request against the
master branch.

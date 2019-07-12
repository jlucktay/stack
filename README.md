# `stack`

[![Build Status](https://cloud.drone.io/api/badges/jlucktay/stack/status.svg)](https://cloud.drone.io/jlucktay/stack)
[![Go Report Card](https://goreportcard.com/badge/github.com/jlucktay/stack)](https://goreportcard.com/report/github.com/jlucktay/stack)
[![GoDoc](https://godoc.org/github.com/jlucktay/stack?status.svg)](https://godoc.org/github.com/jlucktay/stack)

A support tool for use with Terraform stacks, Azure DevOps build pipelines, and GitHub projects/repos.

It currently has the following functions:

- queueing Terraform builds in Azure DevOps
- cancelling unneeded releases of aforementioned builds
- creating GitHub issues in corresponding projects

All of these functions are executed contextually against a specific Terraform stack directory.

## Installation

There are two installation options: building `stack` from the source code hosted here, or downloading a pre-built
binary for your desired platform.

### Building from source

#### Prerequisites

You should have a [working Go environment](https://golang.org/doc/install) and have `$GOPATH/bin` in your `$PATH`.

#### Compiling

To download the source, compile, and install the demo binary, run:

``` shell
go get github.com/jlucktay/stack/...
```

The source code will be located in `$GOPATH/src/github.com/jlucktay/stack/`.

A newly-compiled `stack` binary will be placed in `$GOPATH/bin/`.

### Binary download

<!--
TODO build darwin/amd64
via goreleaser
-->

Coming soon!

## Usage

`stack` has several subcommands:

- `build`
- `cancel`
- `issue`

### `stack build`

``` console
$ stack build

TODO
...build things happen...
TODO
```

## Further implementation ideas

- [Cobra - A Commander for modern Go CLI interactions](https://github.com/spf13/cobra)
- [Viper - Go configuration with fangs](https://github.com/spf13/viper)
- [go-github - Go library for accessing the GitHub API](https://github.com/google/go-github)

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)

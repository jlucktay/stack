# `stack`

[![Build Status][drone-badge]][drone]
[![Go Report Card][goreportcard-badge]][goreportcard]
[![GoDoc][godoc-badge]][godoc]

A support tool for use with Terraform stacks, Azure DevOps build pipelines, and GitHub projects/repos.

It currently has the following functions:

- initialising Terraform against remote state storage, for local execution
- queueing Terraform builds in an Azure DevOps CI/CD pipeline
- cancelling unneeded releases of aforementioned builds
- creating GitHub issues in corresponding projects

All of these functions are executed contextually against a specific Terraform stack directory.

## Installation

There are numerous installation options for `stack`:

- [Homebrew](https://brew.sh)
- building from the source code hosted here
- directly downloading a pre-built binary for your desired platform

### Homebrew

#### First time install

``` shell
brew tap jlucktay/tap
brew install jlucktay/tap/stack
```

#### Ongoing upgrades

``` shell
brew upgrade jlucktay/tap/stack
```

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

### Direct download of binary relases

Binary releases can be downloaded [here on GitHub](https://github.com/jlucktay/stack/releases/latest).

## Configuration

There is a sample JSON file `stack.config.example.json` that should be copied over to your
`$HOME/.config/stack/stack.config.json` directory and populated appropriately.

Filling out this config file will require the generation of two personal access tokens, one from Azure DevOps and one
from GitHub. Links to the appropriate pages on each site are in the example file.

## Usage

`stack` has several subcommands:

- `init`
- `build`
- `cancel`
- `issue`

### `stack init`

Coming soon!

### `stack build`

``` console
$ stack build
Build URL: https://dev.azure.com/MyAzureDevOpsOrg/12345678-90ab-cdef-1234-567890abcdef/_build/results?buildId=1234
```

### `stack cancel`

Coming soon!

### `stack issue`

Coming soon!

## Further implementation ideas

- [Cobra - A Commander for modern Go CLI interactions](https://github.com/spf13/cobra)
- ~~[Viper - Go configuration with fangs](https://github.com/spf13/viper)~~
- ~~[go-github - Go library for accessing the GitHub API](https://github.com/google/go-github)~~

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)

<!-- Badges and associated links -->
[drone-badge]: https://cloud.drone.io/api/badges/jlucktay/stack/status.svg
[drone]: https://cloud.drone.io/jlucktay/stack
[goreportcard-badge]: https://goreportcard.com/badge/github.com/jlucktay/stack
[goreportcard]: https://goreportcard.com/report/github.com/jlucktay/stack
[godoc-badge]: https://godoc.org/github.com/jlucktay/stack?status.svg
[godoc]: https://godoc.org/github.com/jlucktay/stack

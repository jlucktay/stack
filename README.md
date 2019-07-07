# `stack`

[![Build Status](https://cloud.drone.io/api/badges/jlucktay/stack/status.svg)](https://cloud.drone.io/jlucktay/stack)

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

## `curl` request example

``` shell
$ curl --include --header "Authorization: token $GHPAT" https://api.github.com/repos/Dentsu-Aegis-Network-Global-Technology/dan-migration-factory/issues --data '{"title":"Hello world","body":"POSTed via API","labels":["apparea_prod/appstacks/nordic-jenkins","migration"]}'
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Location: https://api.github.com/repos/Dentsu-Aegis-Network-Global-Technology/dan-migration-factory/issues/116
X-RateLimit-Limit: 5000
X-RateLimit-Remaining: 4994
X-RateLimit-Reset: 1562323709
...

{
  "html_url": "https://github.com/Dentsu-Aegis-Network-Global-Technology/dan-migration-factory/issues/116",
  "title": "Hello world",
  "body": "POSTed via API",
  "labels": [
    {
      "name": "apparea_prod/appstacks/nordic-jenkins",
      ...
    },
    {
      "name": "migration",
      ...
    }
  ],
  "user": {
    "login": "jlucktay",
    ...
  }
  ...
}
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)

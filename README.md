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

**Note:** `stack` will also look in the current working directory for this configuration file.

Filling out this config file will require the generation of two personal access tokens, one from Azure DevOps and one
from GitHub. Links to the appropriate pages on each site are in the example file.

### Mapping subscriptions to remote state storage containers and keys

Under the `.azure.subscriptions` section, all keys defined here will map verbatim to the parent directory's name when
`stack` is executed. The values need to be set as GUIDs for the corresponding subscriptions in Azure.

Assume - for this example's sake - that the working directory is as follows:

``` bash
/git/MyGitHubOrg/MyGitHubRepo/stack-prefix/subscription-alias/one/two/three/my-stack
```

Also assume that the following values are configured:

- `.azure.state.storageAccount` = `mytfstatestorage`
- `.azure.subscriptions.subscription-alias` = `01234567-89ab-cdef-0123-456789abcdef`
- `.github.org` = `MyGitHubOrg`
- `.github.repo` = `MyGitHubRepo`
- `.stackPrefix` = `stack-prefix`

The keys under `.azure.subscriptions` map to the first child directory underneath the directory set under
`.stackPrefix`.

For remote state storage within the storage account, the key value is made up of three components:

1. the `.stackPrefix` value (`stack-prefix` in this example)
1. the name of the stack's direct parent directory (`three`)
1. the name of the stack directory itself (`my-stack`)

The container within the remote state storage account maps to the GUID of the subscription.

``` bash
$ pwd
/git/MyGitHubOrg/MyGitHubRepo/stack-prefix/subscription-alias/one/two/three/my-stack
$ stack init
Switching subscriptions... done.
Retrieving storage account key... done.
Switching subscriptions again... done.
Initialising Terraform with following dynamic values:
        container_name:         01234567-89ab-cdef-0123-456789abcdef
        key:                    stack-prefix.three.my-stack
        storage_account:        mytfstatestorage
...
```

### Other tools in use

Some of the functionality in `stack` comes from executing other tools, which will need to be installed, configured,
authed, and available on your `$PATH`:

- [Azure CLI](https://docs.microsoft.com/cli/azure)
- [Git](https://git-scm.com)
- [Terraform](https://www.terraform.io)

## Usage

`stack` has several subcommands:

- `init`
- `build`
- `cancel`
- `issue`

### `stack init`

Initialises the current Terraform stack directory using the Azure storage account for the remote state backend.

``` bash
$ stack init
Switching subscriptions... done.
Retrieving storage account key... done.
Switching subscriptions again... done.
Initialising Terraform with following dynamic values:
...
```

#### `stack init` relevant config keys

- `.azure.state.keyPrefix`: the first of three segments that will make up the blob name for the Terraform backend state
in the Azure storage account
- `.azure.state.storageAccount`: the name of the storage account in Azure where Terraform stores state
- `.azure.state.subscription`: the GUID of the Azure subscription holding the Terraform state storage account
- `.azure.subscriptions.*`: populate this object with all relevant subscriptions, where the key is the name to map to
the directory structure of the Terraform stack, and the value is the GUID of said subscription
- `.stackPrefix`: the name of the parent directory holding all of the Terraform stacks

### `stack build`

Queues a build in Azure DevOps for the current Terraform stack directory.

``` bash
$ stack build
Build URL: https://dev.azure.com/MyAzureDevOpsOrg/12345678-90ab-cdef-1234-567890abcdef/_build/results?buildId=1234
```

- `.azureDevOps.buildDefID`: the build definition ID within Azure DevOps to queue
- `.azureDevOps.org`: the name of the organisation within Azure DevOps
- `.azureDevOps.pat`: the user's personal access token for Azure DevOps
- `.azureDevOps.project`: the name of the project under the organisation within Azure DevOps
- `.stackPrefix`: the name of the parent directory holding all of the Terraform stacks

#### `stack build` optional arguments

##### `--branch`

If given, build from this branch. Defaults to the current branch.

##### `--target`

If given, target these specific Terraform resources only. Delimit multiple target IDs with a semi-colon ';'.
For example:

``` bash
stack build --target="azurerm_resource_group.main;azurerm_virtual_machine.app;azurerm_virtual_machine.database"
```

### `stack cancel`

Cancels any pending releases in Azure DevOps.

**Coming soon!**

### `stack issue`

Creates an issue in GitHub with a label referring to the current Terraform stack directory.

``` bash
$ stack issue "There's a problem with this stack!"
New issue: https://github.com/MyGitHubOrg/MyGitHubRepo/issues/1234
```

- `.github.org`: the name of the organisation within GitHub
- `.github.pat`: the user's personal access token for GitHub
- `.github.repo`: the name of the repository under the organisation within GitHub

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

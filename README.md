# `stack`

[![Build Status][drone-badge]][drone]
[![Go Report Card][goreportcard-badge]][goreportcard]
[![GoDoc][godoc-badge]][godoc]

A support tool for use with Terraform stacks, Azure DevOps build pipelines, and GitHub projects/repos.

It currently has the following functions:

- initialising Terraform against remote state storage, for local execution
- queueing Terraform plans to build and destroy stacks in an Azure DevOps CI/CD pipeline
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
go get github.com/jlucktay/stack
```

The source code will be located in `$GOPATH/src/github.com/jlucktay/stack/`.

A newly-compiled `stack` binary will be placed in `$GOPATH/bin/`.

### Direct download of binary relases

Binary releases can be downloaded [here on GitHub](https://github.com/jlucktay/stack/releases/latest).

## Configuration

There is a sample JSON file named `stack.config.example.json` here in the root of this repo.

The search order that `stack` follows when looking for the fully populated `stack.config.json` config file is as
follows:

1. `<current working directory>/stack.config.json`
1. `$HOME/.config/stack/stack.config.json`
1. `/etc/stack/stack.config.json`

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

``` json
{
  "azure": {
    "state": {
      "storageAccount": "mytfstatestorage"
    },
    "subscriptions": {
      "subscription-alias": "01234567-89ab-cdef-0123-456789abcdef"
    }
  },
  "github": {
    "org": "MyGitHubOrg",
    "repo": "MyGitHubRepo"
  },
  "stackPrefix": "stack-prefix"
}
```

The keys under `.azure.subscriptions` map to the first child directory underneath the directory set under
`.stackPrefix` so the sub-directory `subscription-alias` (under `stack-prefix`) would map to the subscription with a
GUID of `01234567-89ab-cdef-0123-456789abcdef`.

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
Switching subscriptions... done.
Initialising Terraform with following dynamic values:
        container_name:         01234567-89ab-cdef-0123-456789abcdef
        key:                    stack-prefix.three.my-stack
        storage_account:        mytfstatestorage
...
```

### Other tools in use

Some of the functionality in `stack` comes from executing other tools, which will need to be installed, configured,
authed, and available on your `$PATH`:

- [Git](https://git-scm.com) - `git`
- [Terraform](https://www.terraform.io) - `terraform`

## Usage

`stack` itself has several subcommands:

- `init`
- `build`
- `destroy`
- `cancel`
- `issue`
- `version`

### `stack init`

Initialises the current Terraform stack directory using the Azure storage account for the remote state backend.

``` bash
$ stack init
Switching subscriptions... done.
Retrieving storage account key... done.
Switching subscriptions... done.
Initialising Terraform with following dynamic values:
...
```

#### `stack init` relevant config keys

``` json
{
  "azure": {
    "state": {
      "keyPrefix": "first of three segments for key names of state files within blob storage",
      "resourceGroup": "name of resource group on Azure which contains the storage account",
      "storageAccount": "name of Azure storage account where Terraform state is stored",
      "subscription": "GUID of Azure subscription holding the state storage account"
    },
    "subscriptions": {
      "a stack under '/<stackPrefix>/<this key>/<a stack name>/'": "will map to subscription associated with <this key>",
      "exampleSubName": "subscription GUIDs go here",
      "this key will be matched to a parent directory": "this value will map said directory to a specific subscription"
    }
  },
  "stackPrefix": "/some/segment/of/repo/directory/structure/"
}
```

### `stack build`

Queues a plan in Azure DevOps to build the Terraform stack in the current directory.

``` bash
$ stack build
Stack (plan) URL: https://dev.azure.com/MyAzureDevOpsOrg/12345678-90ab-cdef-1234-567890abcdef/_build/results?buildId=1234
```

#### `stack build` relevant config keys

``` json
{
  "azureDevOps": {
    "buildDefID": 5,
    "org": "the name of the organisation within Azure DevOps",
    "pat": "52 character alphanumeric, generated here: https://dev.azure.com/<org>/_usersSettings/tokens",
    "project": "the name of the project under the organisation within Azure DevOps"
  },
  "stackPrefix": "/some/segment/of/repo/directory/structure/"
}
```

#### `stack build` optional arguments

##### `--branch`

If given, build from this branch. Defaults to the current branch.

##### `--target`

If given, target these specific Terraform resources only. Delimit multiple target IDs with a comma `,`.
For example:

``` bash
stack build --target="azurerm_resource_group.main,azurerm_virtual_machine.app,azurerm_virtual_machine.database"
```

### `stack destroy`

Queues a plan in Azure DevOps to destroy the Terraform stack in the current directory.

Functionally identical to the `stack build` subcommand, including `--branch` and `--target` optional arguments, with
the singular difference being that this subcommand references `.azureDevOps.destroyDefID` in the config instead of
`.azureDevOps.buildDefID`.

### `stack cancel`

Cancels any pending releases in Azure DevOps.

**Not yet implemented - coming soon!**

### `stack issue`

Creates an issue in GitHub with a label referring to the current Terraform stack directory, and assigned to the
current user.

The issue's body text is gathered by way of an interactive editor, designated by the current environment's `EDITOR`
variable.

``` bash
$ stack issue -t "There's a problem with this stack!"
...
($EDITOR is launched to gather issue body)
...
New issue: https://github.com/MyGitHubOrg/MyGitHubRepo/issues/1234
```

#### `stack issue` relevant config keys

``` json
{
  "github": {
    "org": "the name of the organisation within GitHub",
    "pat": "<40 character hexadecimal, generated here: https://github.com/settings/tokens>",
    "repo": "the name of the repository under the organisation within GitHub"
  }
}
```

### `stack version`

Displays version and build information for the current `stack` binary.

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

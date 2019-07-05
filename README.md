# `stack`

A support tool for use with Terraform stacks.

It currently has the following functions:

- queueing builds
- cancelling builds
- creating issues

All of these functions are executed contextually against a specific stack.

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

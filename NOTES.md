# Notepad

## `curl` request example, to create a GitHub issue

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

# go-gitinfo
> Gitinfo for golang.

## installation
```sh
go get -u github.com/afeiship/go-gitinfo
```

## usage
```go
package main

func main() {
	info := gitinfo.Get()
}
```

```json
{
  "name": "afeiship",
  "email": "1290657123@qq.com",
  "current_branch": "main",
  "hash": "bd457285054d8f95f9d4b4840c76873b5d549569",
  "short_hash": "bd45728",
  "url_meta": {
    "original_url": "git@github.com:afeiship/go-gitinfo.git",
    "protocol": "https",
    "hostname": "github.com",
    "owner": "afeiship",
    "repo": "go-gitinfo",
    "ssh_url": "git@github.com:afeiship/go-gitinfo.git",
    "https_url": "https://github.com/afeiship/go-gitinfo.git",
    "url": "https://github.com/afeiship/go-gitinfo"
  }
}
```
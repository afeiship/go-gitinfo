# go-gitinfo
> Gitinfo for golang - Extract git repository information and parse git URLs for GitHub and GitLab.

## Installation
```sh
go get -u github.com/afeiship/go-gitinfo
```

## Usage

### Get Current Repository Info
```go
package main
import (
	"fmt"
	"github.com/afeiship/go-gitinfo"
)

func main() {
	hasGit, info := gitinfo.GetGitInfo()
	if !hasGit {
		fmt.Println("Not a git repository")
		return
	}
	fmt.Println(info)
}
```

**Output:**
```json
{
  "name": "afeiship",
  "email": "1290657123@qq.com",
  "current_branch": "main",
  "hash": "bd457285054d8f95f9d4b4840c76873b5d549569",
  "short_hash": "bd45728",
  "is_github": true,
  "meta": {
    "original_url": "git@github.com:afeiship/go-gitinfo.git",
    "protocol": "https",
    "hostname": "github",
    "owner": "afeiship",
    "repo": "go-gitinfo",
    "repo_name": "afeiship/go-gitinfo",
    "ssh_url": "git@github.com:afeiship/go-gitinfo.git",
    "https_url": "https://github.com/afeiship/go-gitinfo.git",
    "url": "https://github.com/afeiship/go-gitinfo",
    "base_url": "https://github.com",
    "actions_url": "https://github.com/afeiship/go-gitinfo/actions",
    "commits_url": "https://github.com/afeiship/go-gitinfo/commits",
    "tags_url": "https://github.com/afeiship/go-gitinfo/tags",
    "pages_url": "https://afeiship.github.io/go-gitinfo/",
    "issues_url": "https://github.com/afeiship/go-gitinfo/issues",
    "pull_requests_url": "https://github.com/afeiship/go-gitinfo/pulls"
  }
}
```

### Parse Git URLs
```go
package main
import (
	"fmt"
	"github.com/afeiship/go-gitinfo"
)

func main() {
	// Parse GitHub URL
	info, err := gitinfo.ParseGitUrl("git@github.com:afeiship/go-gitinfo.git")
	if err != nil {
		panic(err)
	}

	fmt.Println("Repository:", info.RepoName)
	fmt.Println("Actions:", info.ActionsUrl)
	fmt.Println("Pull Requests:", info.PullRequestsUrl)
}
```

## Supported URLs

### GitHub
- SSH: `git@github.com:owner/repo.git`
- HTTPS: `https://github.com/owner/repo.git`
- Enterprise: `git@github.company.com:owner/repo.git`

### GitLab
- SSH: `git@gitlab.com:owner/repo.git`
- HTTPS: `https://gitlab.com/owner/repo.git`
- Self-hosted: `git@git.example.com:owner/repo.git`

## Generated URLs

| Field | GitHub | GitLab |
|-------|--------|--------|
| `actions_url` | `/actions` | `/-/pipelines` |
| `commits_url` | `/commits` | `/-/commits` |
| `tags_url` | `/tags` | `/-/tags` |
| `issues_url` | `/issues` | `/-/issues` |
| `pull_requests_url` | `/pulls` | `/-/merge_requests` |
| `pages_url` | `{owner}.github.io/{repo}/` | `{owner}.pages.{hostname}/{repo}/` |

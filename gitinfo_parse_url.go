package gitinfo

import (
	"fmt"
	"regexp"
	"strings"
)

type GitUrl struct {
	OriginalUrl string `json:"original_url"`
	Protocol    string `json:"protocol"`
	Hostname    string `json:"hostname"`
	Owner       string `json:"owner"`
	Repo        string `json:"repo"`
	RepoName    string `json:"repo_name"`
	SshUrl      string `json:"ssh_url"`
	HttpsUrl    string `json:"https_url"`
	ActionsUrl  string `json:"actions_url"`
	CommitsUrl  string `json:"commits_url"`
	TagsUrl     string `json:"tags_url"`
	PagesUrl    string `json:"pages_url"`
	IssuesUrl   string `json:"issues_url"`
	Url         string `json:"url"`
	BaseUrl     string `json:"base_url"`
}

var reGithub = regexp.MustCompile(`^(git@github\.com[:/])?([\w.-]+)/([\w.-]+)(\.git)?$`)
var reGitHubHttps = regexp.MustCompile(`^(https://github\.com/)([\w.-]+)/([\w.-]+)(\.git)?$`)
var reGitlab = regexp.MustCompile(`^(git@(git\.saybot\.net|lab\.com)[:/])?([.\w-]+)/(.*?)(\.git)?$`)
var reGitlabHttps = regexp.MustCompile(`^(https://(git\.saybot\.net|lab\.com)/)([\w-]+)/(.*?)(\.git)?$`)

func ParseGitUrl(originalUrl string) (*GitUrl, error) {

	var gitUrl GitUrl
	gitUrl.OriginalUrl = originalUrl

	// 处理 GitHub
	if strings.Contains(originalUrl, "github.com") {
		var match []string

		if strings.HasPrefix(originalUrl, "git@") {
			match = reGithub.FindStringSubmatch(originalUrl)
			gitUrl.Protocol = "ssh"
		} else {
			match = reGitHubHttps.FindStringSubmatch(originalUrl)
			gitUrl.Protocol = "https"
		}

		if match != nil {
			// 如果是 GitHub HTTPS URL
			gitUrl.Protocol = "https"
			gitUrl.Hostname = "github.com"
			gitUrl.Owner = match[2]
			// 去掉 .git 后缀
			gitUrl.Repo = strings.TrimSuffix(match[3], ".git")
			gitUrl.RepoName = fmt.Sprintf("%s/%s", gitUrl.Owner, gitUrl.Repo)
			gitUrl.BaseUrl = "https://github.com"
			gitUrl.SshUrl = fmt.Sprintf("git@github.com:%s/%s.git", gitUrl.Owner, gitUrl.Repo)
			gitUrl.HttpsUrl = fmt.Sprintf("https://github.com/%s/%s.git", gitUrl.Owner, gitUrl.Repo)
			gitUrl.Url = fmt.Sprintf("https://github.com/%s/%s", gitUrl.Owner, gitUrl.Repo)
			gitUrl.ActionsUrl = fmt.Sprintf("%s/actions", gitUrl.Url)
			gitUrl.CommitsUrl = fmt.Sprintf("%s/commits", gitUrl.Url)
			gitUrl.TagsUrl = fmt.Sprintf("%s/tags", gitUrl.Url)
			gitUrl.PagesUrl = fmt.Sprintf("https://%s.github.io/%s/", gitUrl.Owner, gitUrl.Repo)
			gitUrl.IssuesUrl = fmt.Sprintf("%s/issues", gitUrl.Url)
			return &gitUrl, nil
		}
	}

	// 处理 GitLab
	if strings.Contains(originalUrl, "git.saybot.net") || strings.Contains(originalUrl, "gitlab.com") {
		var match []string

		if strings.HasPrefix(originalUrl, "git@") {
			match = reGitlab.FindStringSubmatch(originalUrl)
			gitUrl.Protocol = "ssh"
		} else {
			match = reGitlabHttps.FindStringSubmatch(originalUrl)
			gitUrl.Protocol = "https"
		}

		if match != nil {
			// 如果是 GitLab HTTPS URL
			gitUrl.Protocol = "https"
			gitUrl.Hostname = match[2]
			gitUrl.Owner = match[3]
			// 去掉 .git 后缀
			gitUrl.Repo = strings.TrimSuffix(match[4], ".git")
			gitUrl.RepoName = fmt.Sprintf("%s/%s", gitUrl.Owner, gitUrl.Repo)
			gitUrl.BaseUrl = fmt.Sprintf("https://%s", match[2])
			gitUrl.SshUrl = fmt.Sprintf("git@%s:%s/%s.git", match[2], gitUrl.Owner, gitUrl.Repo)
			gitUrl.HttpsUrl = fmt.Sprintf("https://%s/%s/%s.git", match[2], gitUrl.Owner, gitUrl.Repo)
			gitUrl.Url = fmt.Sprintf("https://%s/%s/%s", match[2], gitUrl.Owner, gitUrl.Repo)
			gitUrl.ActionsUrl = fmt.Sprintf("%s/-/pipelines", gitUrl.Url)
			gitUrl.CommitsUrl = fmt.Sprintf("%s/-/commits", gitUrl.Url)
			gitUrl.TagsUrl = fmt.Sprintf("%s/-/tags", gitUrl.Url)
			gitUrl.PagesUrl = fmt.Sprintf("https://%s.pages.%s/%s/", gitUrl.Owner, match[2], gitUrl.Repo)
			gitUrl.IssuesUrl = fmt.Sprintf("%s/-/issues", gitUrl.Url)
			return &gitUrl, nil
		}
	}

	return nil, fmt.Errorf("unable to parse the URL: %s", originalUrl)
}

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
	SshUrl      string `json:"ssh_url"`
	HttpsUrl    string `json:"https_url"`
	ActionsUrl  string `json:"actions_url"`
	CommitsUrl  string `json:"commits_url"`
	TagsUrl     string `json:"tags_url"`
	PagesUrl    string `json:"pages_url"`
	Url         string `json:"url"`
}

var reGithub = regexp.MustCompile(`^(git@github\.com[:/])?([\w-]+)/([\w-]+)(\.git)?$`)
var reGitHubHttps = regexp.MustCompile(`^(https://github\.com/)([\w-]+)/([\w-]+)(\.git)?$`)
var reGitlab = regexp.MustCompile(`^(git@(git\.saybot\.net|lab\.com)[:/])?([\w-]+)/(.*?)(\.git)?$`)
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
			gitUrl.Repo = match[3]
			gitUrl.SshUrl = "git@github.com:" + match[2] + "/" + match[3] + ".git"
			gitUrl.HttpsUrl = "https://github.com/" + match[2] + "/" + match[3] + ".git"
			gitUrl.Url = "https://github.com/" + match[2] + "/" + match[3]
			gitUrl.ActionsUrl = gitUrl.Url + "/actions"
			gitUrl.CommitsUrl = gitUrl.Url + "/commits"
			gitUrl.TagsUrl = gitUrl.Url + "/tags"
			gitUrl.PagesUrl = fmt.Sprintf("https://%s.github.io/%s/", match[2], match[3])
			return &gitUrl, nil
		}
	}

	// 处理 GitLab
	if strings.Contains(originalUrl, "git.saybot.net") || strings.Contains(originalUrl, "gitlab.com") {
		var match []string

		if strings.HasPrefix(originalUrl, "git@") {
			match = reGitlab.FindStringSubmatch(originalUrl)
			fmt.Println("just match: ", match)
			gitUrl.Protocol = "ssh"
		} else {
			match = reGitlabHttps.FindStringSubmatch(originalUrl)
			gitUrl.Protocol = "https"
		}

		if match != nil {
			// 如果是 GitLab HTTPS URL
			gitUrl.Url = "https://" + match[2] + "/" + match[3] + "/" + match[4]
			gitUrl.Protocol = "https"
			gitUrl.Hostname = match[2]
			gitUrl.Owner = match[3]
			gitUrl.Repo = match[4]
			gitUrl.SshUrl = "git@" + match[2] + ":" + match[3] + "/" + match[4] + ".git"
			gitUrl.HttpsUrl = "https://" + match[2] + "/" + match[3] + "/" + match[4] + ".git"
			gitUrl.Url = "https://" + match[2] + "/" + match[3] + "/" + match[4]
			gitUrl.ActionsUrl = gitUrl.Url + "/-/pipelines"
			gitUrl.CommitsUrl = gitUrl.Url + "/-/commits"
			gitUrl.TagsUrl = gitUrl.Url + "/-/tags"
			gitUrl.PagesUrl = fmt.Sprintf("https://%s.pages.%s/%s/", match[3], match[2], match[4])
			return &gitUrl, nil
		}
	}

	return nil, fmt.Errorf("unable to parse the URL: %s", originalUrl)
}

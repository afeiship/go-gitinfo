package parsers

import (
	"fmt"
	"strings"
)

// ParseGitlabUrl parses GitLab URLs
func ParseGitlabUrl(originalUrl string) (*GitUrl, error) {
	var gitUrl GitUrl
	gitUrl.OriginalUrl = originalUrl

	// Check if it's a GitLab URL
	if !strings.Contains(originalUrl, "git.saybot.net") && !strings.Contains(originalUrl, "gitlab.com") {
		return nil, fmt.Errorf("not a GitLab URL")
	}

	var match []string
	if strings.HasPrefix(originalUrl, "git@") {
		match = ReGitlab.FindStringSubmatch(originalUrl)
		gitUrl.Protocol = "https"
	} else {
		match = ReGitlabHttps.FindStringSubmatch(originalUrl)
		gitUrl.Protocol = "https"
	}

	if match == nil {
		return nil, fmt.Errorf("unable to parse GitLab URL: %s", originalUrl)
	}

	// Extract components from regex match
	gitUrl.Protocol = "https"
	gitUrl.Hostname = match[2]
	gitUrl.Owner = match[3]
	// Remove .git suffix
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

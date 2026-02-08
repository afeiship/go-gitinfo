package parsers

import (
	"fmt"
	"strings"
)

// ParseGithubUrl parses GitHub and GitHub Enterprise URLs
func ParseGithubUrl(originalUrl string) (*GitUrl, error) {
	var gitUrl GitUrl
	gitUrl.OriginalUrl = originalUrl

	// Check if it's a GitHub URL
	if !strings.Contains(originalUrl, "github") && !ReGithub.MatchString(originalUrl) && !ReGitHubHttps.MatchString(originalUrl) {
		return nil, fmt.Errorf("not a GitHub URL")
	}

	var match []string
	if strings.HasPrefix(originalUrl, "git@") {
		match = ReGithub.FindStringSubmatch(originalUrl)
		gitUrl.Protocol = "https"
	} else {
		match = ReGitHubHttps.FindStringSubmatch(originalUrl)
		gitUrl.Protocol = "https"
	}

	if match == nil {
		return nil, fmt.Errorf("unable to parse GitHub URL: %s", originalUrl)
	}

	// Extract components from regex match
	originalHostname := match[2]
	gitUrl.Owner = match[3]
	// Remove .git suffix
	gitUrl.Repo = strings.TrimSuffix(match[4], ".git")
	gitUrl.RepoName = fmt.Sprintf("%s/%s", gitUrl.Owner, gitUrl.Repo)

	// If it's a github-related domain, normalize to github.com
	if strings.Contains(originalHostname, "github") {
		gitUrl.Hostname = "github.com"
		gitUrl.BaseUrl = "https://github.com"
		gitUrl.SshUrl = fmt.Sprintf("git@github.com:%s/%s.git", gitUrl.Owner, gitUrl.Repo)
		gitUrl.HttpsUrl = fmt.Sprintf("https://github.com/%s/%s.git", gitUrl.Owner, gitUrl.Repo)
		gitUrl.Url = fmt.Sprintf("https://github.com/%s/%s", gitUrl.Owner, gitUrl.Repo)
		gitUrl.ActionsUrl = fmt.Sprintf("%s/actions", gitUrl.Url)
		gitUrl.CommitsUrl = fmt.Sprintf("%s/commits", gitUrl.Url)
		gitUrl.TagsUrl = fmt.Sprintf("%s/tags", gitUrl.Url)
		gitUrl.PagesUrl = fmt.Sprintf("https://%s.github.io/%s/", gitUrl.Owner, gitUrl.Repo)
		gitUrl.IssuesUrl = fmt.Sprintf("%s/issues", gitUrl.Url)
		gitUrl.PullRequestsUrl = fmt.Sprintf("%s/pulls", gitUrl.Url)
	} else {
		// For non-github domains, keep original hostname
		gitUrl.Hostname = originalHostname
		gitUrl.BaseUrl = fmt.Sprintf("https://%s", gitUrl.Hostname)
		gitUrl.SshUrl = fmt.Sprintf("git@%s:%s/%s.git", gitUrl.Hostname, gitUrl.Owner, gitUrl.Repo)
		gitUrl.HttpsUrl = fmt.Sprintf("https://%s/%s/%s.git", gitUrl.Hostname, gitUrl.Owner, gitUrl.Repo)
		gitUrl.Url = fmt.Sprintf("https://%s/%s/%s", gitUrl.Hostname, gitUrl.Owner, gitUrl.Repo)
		gitUrl.ActionsUrl = fmt.Sprintf("%s/-/pipelines", gitUrl.Url)
		gitUrl.CommitsUrl = fmt.Sprintf("%s/-/commits", gitUrl.Url)
		gitUrl.TagsUrl = fmt.Sprintf("%s/-/tags", gitUrl.Url)
		gitUrl.PagesUrl = fmt.Sprintf("https://%s.pages.%s/%s/", gitUrl.Owner, gitUrl.Hostname, gitUrl.Repo)
		gitUrl.IssuesUrl = fmt.Sprintf("%s/-/issues", gitUrl.Url)
		gitUrl.PullRequestsUrl = fmt.Sprintf("%s/-/merge_requests", gitUrl.Url)
	}

	return &gitUrl, nil
}


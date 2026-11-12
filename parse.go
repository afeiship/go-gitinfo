package gitinfo

import (
	"fmt"
	"regexp"

	"github.com/afeiship/go-gitinfo/parsers"
)

var reUserInfo = regexp.MustCompile(`^(https?://)[^/@]*@`)

// stripUserInfo removes credentials from an HTTP(S) URL authority.
func stripUserInfo(raw string) string {
	return reUserInfo.ReplaceAllString(raw, "$1")
}

// ParseGitUrl parses a Git repository URL and returns a GitUrl struct
// It supports GitHub (including GitHub Enterprise) and GitLab URLs
func ParseGitUrl(originalUrl string) (*GitUrl, error) {
	originalUrl = stripUserInfo(originalUrl)

	// Try GitHub first
	if gitUrl, err := parsers.ParseGithubUrl(originalUrl); err == nil {
		return gitUrl, nil
	}

	// Try GitLab
	if gitUrl, err := parsers.ParseGitlabUrl(originalUrl); err == nil {
		return gitUrl, nil
	}

	// If neither matches, return error
	return nil, fmt.Errorf("unable to parse the URL: %s", originalUrl)
}

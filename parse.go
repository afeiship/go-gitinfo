package gitinfo

import (
	"fmt"

	"github.com/afeiship/go-gitinfo/parsers"
)

// ParseGitUrl parses a Git repository URL and returns a GitUrl struct
// It supports GitHub (including GitHub Enterprise) and GitLab URLs
func ParseGitUrl(originalUrl string) (*GitUrl, error) {
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

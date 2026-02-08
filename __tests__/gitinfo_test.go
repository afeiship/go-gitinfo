package gitinfo_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/afeiship/go-gitinfo"
)

func TestParseUrl(t *testing.T) {
	// Test GitLab SSH URLs
	tests := []struct {
		name     string
		url      string
		hostname string
		owner    string
		repo     string
		protocol string
		baseUrl  string
	}{
		{
			name:     "GitLab SSH - Simple Path",
			url:      "git@git.saybot.net:web_app/rhino-h5.git",
			hostname: "git.saybot.net",
			owner:    "web_app",
			repo:     "rhino-h5",
			protocol: "https",
			baseUrl:  "https://git.saybot.net",
		},
		{
			name:     "GitLab SSH - Nested Path",
			url:      "git@git.saybot.net:ACE/courseware/pptWebPlayer.git",
			hostname: "git.saybot.net",
			owner:    "ACE",
			repo:     "courseware/pptWebPlayer",
			protocol: "https",
			baseUrl:  "https://git.saybot.net",
		},
		{
			name:     "GitLab SSH - Username with Dot",
			url:      "git@git.saybot.net:aric.zheng/frontend-ci.git",
			hostname: "git.saybot.net",
			owner:    "aric.zheng",
			repo:     "frontend-ci",
			protocol: "https",
			baseUrl:  "https://git.saybot.net",
		},
		{
			name:     "GitLab HTTPS",
			url:      "https://git.saybot.net/web_app/rhino-h5.git",
			hostname: "git.saybot.net",
			owner:    "web_app",
			repo:     "rhino-h5",
			protocol: "https",
			baseUrl:  "https://git.saybot.net",
		},
		{
			name:     "GitHub SSH",
			url:      "git@github.com:afeiship/nx.git",
			hostname: "github.com",
			owner:    "afeiship",
			repo:     "nx",
			protocol: "https",
			baseUrl:  "https://github.com",
		},
		{
			name:     "GitHub SSH",
			url:      "git@github.com:aric-jswork/pl.js.work.git",
			hostname: "github.com",
			owner:    "aric-jswork",
			repo:     "pl.js.work",
			protocol: "https",
			baseUrl:  "https://github.com",
		},
		{
			name:     "GitHub HTTPS",
			url:      "https://github.com/afeiship/nx.git",
			hostname: "github.com",
			owner:    "afeiship",
			repo:     "nx",
			protocol: "https",
			baseUrl:  "https://github.com",
		},
		{
			name:     "GitHub HTTPS without .git",
			url:      "https://github.com/afeiship/nx",
			hostname: "github.com",
			owner:    "afeiship",
			repo:     "nx",
			protocol: "https",
			baseUrl:  "https://github.com",
		},
		{
			name:     "GitHub Enterprise SSH",
			url:      "git@github.work:bosinc/katana-web.git",
			hostname: "github.com",
			owner:    "bosinc",
			repo:     "katana-web",
			protocol: "https",
			baseUrl:  "https://github.com",
		},
		{
			name:     "GitHub Enterprise HTTPS",
			url:      "https://github.company.com/enterprise/repo.git",
			hostname: "github.com",
			owner:    "enterprise",
			repo:     "repo",
			protocol: "https",
			baseUrl:  "https://github.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info, err := gitinfo.ParseGitUrl(tt.url)
			if err != nil {
				t.Errorf("ParseGitUrl(%s) error = %v", tt.url, err)
				return
			}

			if info.Hostname != tt.hostname {
				t.Errorf("Hostname = %v, want %v", info.Hostname, tt.hostname)
			}
			if info.Owner != tt.owner {
				t.Errorf("Owner = %v, want %v", info.Owner, tt.owner)
			}
			if info.RepoName != fmt.Sprintf("%s/%s", tt.owner, tt.repo) {
				t.Errorf("RepoName = %v, want %v", info.RepoName, fmt.Sprintf("%s/%s", tt.owner, tt.repo))
			}
			if info.Protocol != tt.protocol {
				t.Errorf("Protocol = %v, want %v", info.Protocol, tt.protocol)
			}
			if info.BaseUrl != tt.baseUrl {
				t.Errorf("BaseUrl = %v, want %v", info.BaseUrl, tt.baseUrl)
			}
			// Check ActionUrl based on repository type
			expectedActionUrl := ""
			if strings.Contains(info.Hostname, "github") {
				expectedActionUrl = fmt.Sprintf("%s/actions", info.Url)
			} else {
				expectedActionUrl = fmt.Sprintf("%s/-/pipelines", info.Url)
			}
			if info.ActionsUrl != expectedActionUrl {
				t.Errorf("ActionUrl = %v, want %v", info.ActionsUrl, expectedActionUrl)
			}

			// Check PullRequestsUrl based on repository type
			expectedPullRequestsUrl := ""
			if strings.Contains(info.Hostname, "github") {
				expectedPullRequestsUrl = fmt.Sprintf("%s/pulls", info.Url)
			} else {
				expectedPullRequestsUrl = fmt.Sprintf("%s/-/merge_requests", info.Url)
			}
			if info.PullRequestsUrl != expectedPullRequestsUrl {
				t.Errorf("PullRequestsUrl = %v, want %v", info.PullRequestsUrl, expectedPullRequestsUrl)
			}
		})
	}

	// Test invalid URL
	_, err := gitinfo.ParseGitUrl("invalid-url")
	if err == nil {
		t.Error("Expected error for invalid URL, got nil")
	}
}

func TestGithubDetection(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		isGithub bool
	}{
		{
			name:     "GitHub.com SSH",
			url:      "git@github.com:afeiship/nx.git",
			isGithub: true,
		},
		{
			name:     "GitHub.com HTTPS",
			url:      "https://github.com/afeiship/nx.git",
			isGithub: true,
		},
		{
			name:     "GitHub Enterprise SSH",
			url:      "git@github.work:bosinc/katana-web.git",
			isGithub: true,
		},
		{
			name:     "GitHub Enterprise HTTPS",
			url:      "https://github.company.com/enterprise/repo.git",
			isGithub: true,
		},
		{
			name:     "GitLab SSH",
			url:      "git@gitlab.com:user/repo.git",
			isGithub: false,
		},
		{
			name:     "GitLab HTTPS",
			url:      "https://gitlab.com/user/repo.git",
			isGithub: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info, err := gitinfo.ParseGitUrl(tt.url)
			if err != nil {
				t.Errorf("ParseGitUrl(%s) error = %v", tt.url, err)
				return
			}

			isGithub := strings.Contains(info.Hostname, "github")
			if isGithub != tt.isGithub {
				t.Errorf("IsGithub = %v, want %v for hostname %s", isGithub, tt.isGithub, info.Hostname)
			}

			// 额外验证：所有github相关域名都应该返回github.com
			if tt.isGithub && info.Hostname != "github.com" {
				t.Errorf("Expected github.com hostname, got %s for URL %s", info.Hostname, tt.url)
			}
		})
	}
}

func TestGetGitInfo(f *testing.T) {
	_, info := gitinfo.GetGitInfo()
	jsonInfo, _ := json.Marshal(info)
	fmt.Println("jsonInfo: ", string(jsonInfo))
}

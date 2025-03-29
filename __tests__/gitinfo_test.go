package gitinfo_test

import (
	"encoding/json"
	"fmt"
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
			if info.Repo != tt.repo {
				t.Errorf("Repo = %v, want %v", info.Repo, tt.repo)
			}
			if info.Protocol != tt.protocol {
				t.Errorf("Protocol = %v, want %v", info.Protocol, tt.protocol)
			}
			if info.BaseUrl != tt.baseUrl {
				t.Errorf("BaseUrl = %v, want %v", info.BaseUrl, tt.baseUrl)
			}
		})
	}

	// Test invalid URL
	_, err := gitinfo.ParseGitUrl("invalid-url")
	if err == nil {
		t.Error("Expected error for invalid URL, got nil")
	}
}

func TestGetGitInfo(f *testing.T) {
	_, info := gitinfo.GetGitInfo()
	jsonInfo, _ := json.Marshal(info)
	fmt.Println("jsonInfo: ", string(jsonInfo))
}

package gitinfo

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

// github: git@github.com:afeiship/go-gitinfo.git
// gitlab: git@git.saybot.net:web_app/rhino-h5.git

type ParseUrlResult struct {
	GitType  string `json:"git_type"`
	BaseUrl  string `json:"base_url"`
	SshUrl   string `json:"ssh_url"`
	HttpsUrl string `json:"https_url"`
	RepoUrl  string `json:"repo_url"`
	Owner    string `json:"owner"`
	Repo     string `json:"repo"`
	Hostname string `json:"hostname"`
}

type GitInfo struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	Owner         string `json:"owner"`
	Hash          string `json:"hash"`
	ShortHash     string `json:"short_hash"`
	Repo          string `json:"repo"`
	CurrentBranch string `json:"current_branch"`
	Hostname      string `json:"hostname"`
	GitType       string `json:"git_type"`
	OriginUrl     string `json:"origin_url"`
	BaseUrl       string `json:"base_url"`
	SshUrl        string `json:"ssh_url"`
	HttpsUrl      string `json:"https_url"`
	RepoUrl       string `json:"repo_url"`
}

func Get() GitInfo {
	var gitInfo GitInfo

	originUrl := runShell("git config --get remote.origin.url")
	currentBranch := runShell("git rev-parse --abbrev-ref HEAD")
	hash := runShell("git rev-parse --verify HEAD")
	shortHash := runShell("git rev-parse --short HEAD")
	email := runShell("git config user.email")
	name := runShell("git config user.name")

	// urlInfo := ParseGitUrl(originUrl)

	gitInfo.CurrentBranch = currentBranch
	gitInfo.OriginUrl = originUrl
	gitInfo.Hash = hash
	gitInfo.ShortHash = shortHash
	gitInfo.Email = email
	gitInfo.Name = name

	return gitInfo
}

func runShell(command string) string {
	execCmd := exec.Command("bash", "-c", command)
	out, err := execCmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

// GitUrl结构体用于存储解析出来的内容
type GitUrl struct {
	OriginalUrl string
	Protocol    string
	Hostname    string
	Owner       string
	Repo        string
	SshUrl      string
	HttpsUrl    string
}

// parseGitUrl 解析原始 Git URL，返回对应的 GitUrl 结构体
func ParseGitUrl(originalUrl string) (*GitUrl, error) {
	// 正则表达式用于提取 Git URL 中的主机名、协议、仓库路径等信息
	reGithub := regexp.MustCompile(`^(git@github\.com[:/])?([\w-]+)/([\w-]+)(\.git)?$`)
	reGitlab := regexp.MustCompile(`^(git@git\.(saybot\.net|lab\.com)[:/])?([\w-]+)/([\w-]+)(\.git)?$`)

	var gitUrl GitUrl
	gitUrl.OriginalUrl = originalUrl

	// 处理 GitHub
	if strings.Contains(originalUrl, "github.com") {
		match := reGithub.FindStringSubmatch(originalUrl)
		if match != nil {
			// 如果是 GitHub SSH URL
			gitUrl.Protocol = "ssh"
			gitUrl.Hostname = "github.com"
			gitUrl.Owner = match[2]
			gitUrl.Repo = match[3]
			gitUrl.SshUrl = "git@github.com:" + match[2] + "/" + match[3] + ".git"
			gitUrl.HttpsUrl = "https://github.com/" + match[2] + "/" + match[3] + ".git"
			return &gitUrl, nil
		}

		// 如果是 GitHub HTTPS URL
		match = reGithub.FindStringSubmatch(originalUrl)
		if match != nil {
			gitUrl.Protocol = "https"
			gitUrl.Hostname = "github.com"
			gitUrl.Owner = match[2]
			gitUrl.Repo = match[3]
			gitUrl.SshUrl = "git@github.com:" + match[2] + "/" + match[3] + ".git"
			gitUrl.HttpsUrl = "https://github.com/" + match[2] + "/" + match[3] + ".git"
			return &gitUrl, nil
		}
	}

	// 处理 GitLab
	if strings.Contains(originalUrl, "git.saybot.net") || strings.Contains(originalUrl, "gitlab.com") {
		match := reGitlab.FindStringSubmatch(originalUrl)
		if match != nil {
			// 如果是 GitLab SSH URL
			gitUrl.Protocol = "ssh"
			gitUrl.Hostname = match[2]
			gitUrl.Owner = match[3]
			gitUrl.Repo = match[4]
			gitUrl.SshUrl = "git@" + match[2] + ":" + match[3] + "/" + match[4] + ".git"
			gitUrl.HttpsUrl = "https://" + match[2] + "/" + match[3] + "/" + match[4] + ".git"
			return &gitUrl, nil
		}
	}

	return nil, fmt.Errorf("unable to parse the URL: %s", originalUrl)
}

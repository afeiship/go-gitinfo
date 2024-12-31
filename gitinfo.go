package gitinfo

import (
	"os/exec"
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

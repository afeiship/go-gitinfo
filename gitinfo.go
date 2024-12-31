package gitinfo

import (
	"os/exec"
	"strings"
)

type GitInfo struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	CurrentBranch string `json:"current_branch"`
	Hash          string `json:"hash"`
	ShortHash     string `json:"short_hash"`
	Meta          GitUrl `json:"meta"`
}

func GetGitInfo() GitInfo {
	var gitInfo GitInfo

	originUrl := runShell("git config --get remote.origin.url")
	name := runShell("git config user.name")
	email := runShell("git config user.email")
	currentBranch := runShell("git rev-parse --abbrev-ref HEAD")
	hash := runShell("git rev-parse --verify HEAD")
	shortHash := runShell("git rev-parse --short HEAD")

	gitInfo.Email = email
	gitInfo.Name = name
	gitInfo.CurrentBranch = currentBranch
	gitInfo.Hash = hash
	gitInfo.ShortHash = shortHash

	gitUrls, _ := ParseGitUrl(originUrl)
	gitInfo.Meta = *gitUrls

	return gitInfo
}

// ---- private functions ----
func runShell(command string) string {
	execCmd := exec.Command("bash", "-c", command)
	out, err := execCmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

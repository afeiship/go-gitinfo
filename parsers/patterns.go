package parsers

import "regexp"

// 域名后缀模式
const domainSuffix = `(?:com|net|org|io|work|dev|local)`

// 主机名模式：允许字母、数字、点、连字符，后跟域名后缀
const hostnamePattern = `[\w.-]+\.` + domainSuffix

// 用户名/组织名模式
const ownerPattern = `[\w.-]+`

// 仓库名模式
const repoPattern = `[\w.-]+`

// 可选的 .git 后缀
const gitSuffix = `(\.git)?`

// ReGithub matches GitHub SSH format: git@hostname:owner/repo.git
var ReGithub = regexp.MustCompile(
	`^(git@(` + hostnamePattern + `)[:/])?(` + ownerPattern + `)/(` + repoPattern + `)` + gitSuffix + `$`)

// ReGitHubHttps matches GitHub HTTPS format: https://hostname/owner/repo.git
var ReGitHubHttps = regexp.MustCompile(
	`^(https://(` + hostnamePattern + `)/)(` + ownerPattern + `)/(` + repoPattern + `)` + gitSuffix + `$`)

// ReGitlab matches GitLab SSH format: git@hostname:owner/repo.git
var ReGitlab = regexp.MustCompile(`^(git@(git\.saybot\.net|lab\.com)[:/])?([.\w-]+)/(.*?)(\.git)?$`)

// ReGitlabHttps matches GitLab HTTPS format: https://hostname/owner/repo.git
var ReGitlabHttps = regexp.MustCompile(`^(https://(git\.saybot\.net|lab\.com)/)([\w-]+)/(.*?)(\.git)?$`)


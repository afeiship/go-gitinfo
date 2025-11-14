package parsers

// GitUrl represents a parsed Git repository URL with all its components
type GitUrl struct {
	OriginalUrl string `json:"original_url"`
	Protocol    string `json:"protocol"`
	Hostname    string `json:"hostname"`
	Owner       string `json:"owner"`
	Repo        string `json:"repo"`
	RepoName    string `json:"repo_name"`
	SshUrl      string `json:"ssh_url"`
	HttpsUrl    string `json:"https_url"`
	ActionsUrl  string `json:"actions_url"`
	CommitsUrl  string `json:"commits_url"`
	TagsUrl     string `json:"tags_url"`
	PagesUrl    string `json:"pages_url"`
	IssuesUrl   string `json:"issues_url"`
	Url         string `json:"url"`
	BaseUrl     string `json:"base_url"`
}


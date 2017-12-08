package models

type GitlabProject struct {
	Id        int    `json:"id"`
	ReposName string `json:"name"`
	ReposPath string `json:"path"`
}

type GitlabFile struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Path string `json:"path"`
}

type AddGitlabCommitResp struct {
	ReposName     string `json:"repos_name"`
	FileName      string `json:"file_name"`
	FileContent   string `json:"file_content"`
	Branch        string `json:"branch"`
	CommitMessage string `json:"commit_message"`
}

type GitlabCommitAddAction struct {
	Action   string `json:"action"`
	FilePath string `json:"file_path"`
	Content  string `json:"content"`
}

type GitlabAddCommit struct {
	Actions       []GitlabCommitAddAction `json:"actions"`
	Branch        string                  `json:"branch"`
	CommitMessage string                  `json:"commit_message"`
}

type GitlabCommitDelAction struct {
	Action   string `json:"action"`
	FilePath string `json:"file_path"`
}

type Gitlab struct {
	Url   string `json:"url"`
	User  string `json:"user"`
	Token string `json:"token"`
	Repos string `json:"repos"`
}

type GitlabInterface interface {
	GetReposID(reposName string) (id int, err error)
	GetFileBlobID(reposID int, ref, filePath string) (blobID string, err error)
	GetFileContent(reposID int, blobId string) (content string, err error)
}

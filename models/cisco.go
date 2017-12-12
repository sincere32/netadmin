package models

type CiscoUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type CiscoCommandRec struct {
	Hosts   []string  `json:"hosts"`
	Port    int       `json:"port"`
	User    CiscoUser `json:"user"`
	Command string    `json:"command"`
}

type CiscoConfigRec struct {
	Hosts     []string  `json:"hosts"`
	User      CiscoUser `json:"user"`
	ReposName string    `json:"repos_name"`
	Ref       string    `json:"ref"`
	FilePath  string    `json:"file_path"`
}

type CiscoConfigPost struct {
	Hosts       []string  `json:"hosts"`
	User        CiscoUser `json:"user"`
	FileContent string    `json:"file_content"`
	BlobId      string    `json:"blob_id"`
}

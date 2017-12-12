package models

type JuniperUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type JuniperCommandRec struct {
	Hosts   []string    `json:"hosts"`
	Port    int         `json:"port"`
	User    JuniperUser `json:"user"`
	Command string      `json:"command"`
}

type JuniperConfigRec struct {
	Hosts     []string    `json:"hosts"`
	User      JuniperUser `json:"user"`
	ReposName string      `json:"repos_name"`
	Ref       string      `json:"ref"`
	FilePath  string      `json:"file_path"`
}

type JuniperConfigPost struct {
	Hosts       []string    `json:"hosts"`
	User        JuniperUser `json:"user"`
	FileContent string      `json:"file_content"`
}

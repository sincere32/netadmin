package models

type JunosUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type JunosCommand struct {
	Hosts   []string  `json:"hosts"`
	Port    int       `json:"port"`
	User    JunosUser `json:"user"`
	Command string    `json:"command"`
}

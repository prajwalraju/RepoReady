package dto

type Config struct {
	Github Github `yaml:"github"`
}
type Github struct {
	Token string `yaml:"token"`
}

// RepoInput is the Object that holds all the info required to create a remote repo
type RepoInput struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Private     bool     `json:"private"`
	Topics      []string `json:"topics"`
	Owner       string
	HtmlUrl     string
	GitUrl      string
	SshUrl      string
}

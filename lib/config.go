package lib

import (
	"github.com/BurntSushi/toml"
	"os"
)

type Config struct {
	Jira Jira
}

type ConfigArgs struct {
	JiraUsername      string `arg:"--jira-username,env:JIRA_USERNAME"`
	JiraPassword      string `arg:"--jira-password,env:JIRA_PASSWORD"`
	JiraUrl           string `arg:"--jira-url,env:JIRA_URL"`
	JiraProjectPrefix string `arg:"--jira-project-prefix,env:JIRA_PROJECT_PREFIX"`
}

type Jira struct {
	Username      string
	Password      string
	Url           string
	ProjectPrefix string
}

func ConfigMain(args *ConfigArgs) {
	config := Config{
		Jira: Jira{
			Username:      args.JiraUsername,
			Password:      args.JiraPassword,
			Url:           args.JiraUrl,
			ProjectPrefix: args.JiraProjectPrefix,
		},
	}
	err := config.Save(".rean")
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
	println("Config saved.")
	os.Exit(0)
}

func (config *Config) Save(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	encoder := toml.NewEncoder(file)
	return encoder.Encode(config)
}

func Load(filePath string) (*Config, error) {
	var config Config
	_, err := toml.DecodeFile(filePath, &config)
	if err != nil {
		return nil, err
	}

	_, err = toml.DecodeFile(filePath, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

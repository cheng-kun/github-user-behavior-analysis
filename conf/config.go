package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/github-user-behavior-analysis/logs"
)

type Config struct {
	GithubarchivePath string `json:"githubarchive_path"`
	Database          Database
	GithubCredential  []GithubCredentials
}

type Database struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
	Port     int    `json:"port"`
}

type GithubCredentials struct {
	Account string `json:"account"`
	Token   string `json:"token"`
}

func LoadConfigFile(fileName string) (*Config, error) {
	conf := &Config{}
	_, err := toml.DecodeFile(fileName, &conf)
	if err != nil {
		logs.PrintLogger().Error(err)
		return nil, err
	}

	return conf, nil
}

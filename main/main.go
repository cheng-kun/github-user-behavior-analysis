package main

import (
	"github.com/github-user-behavior-analysis/conf"
	"github.com/github-user-behavior-analysis/github-archive"
)

func SaveArchiveFiles(fileName string) {
	conf, err := conf.LoadConfigFile(fileName)
	if err == nil {
		github_archive.SaveFiles(conf.GithubarchivePath)
	}
}

func main()  {
	SaveArchiveFiles("./conf/config.toml")
}
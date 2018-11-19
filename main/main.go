package main

import (
	"github.com/github-user-behavior-analysis/conf"
	"github.com/github-user-behavior-analysis/controller"
	"github.com/github-user-behavior-analysis/db"
	"github.com/github-user-behavior-analysis/github-archive"
)

func SaveArchiveFiles(fileName string) {
	conf, err := conf.LoadConfigFile(fileName)
	if err == nil {
		github_archive.SaveFiles(conf.GithubarchivePath)
	}
}

func SaveTenTop(filename string, conn db.Database)  {


	file, err := controller.LoadJsonFile(filename)
	if err != nil{
		return
	}

	rankings, err := controller.ReadData(file)
	if err != nil {
		return
	}

	controller.SaveData(rankings, conn)
}

func main()  {
	//SaveArchiveFiles("./conf/config.toml")

	conf, err := conf.LoadConfigFile("./conf/config.toml")
	if err != nil {
		return
	}
	//logs.PrintLogger().Info(conf, err)


	db, err := db.Connect(*conf)
	//logs.PrintLogger().Info(db, err)




	SaveTenTop("./data/github3.txt", *db)
}

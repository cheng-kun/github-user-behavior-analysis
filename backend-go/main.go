package main

import (
	"github.com/github-user-behavior-analysis/backend-go/controller"
	"github.com/github-user-behavior-analysis/backend-go/db"
	"github.com/github-user-behavior-analysis/backend-go/web-service"
)


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

	//SaveTenTop("./data/github3.txt", *db)

	web_service.StartWebRequest()
}

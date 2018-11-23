package main

import (
	"github.com/github-user-behavior-analysis/backend-go/controller"
	"github.com/github-user-behavior-analysis/backend-go/db"
	"github.com/github-user-behavior-analysis/backend-go/logs"
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

func SaveTopUsers()  {
	file, err := controller.OpenCSVFile("./backend-go/data/top_100_users.csv")
	if err != nil {
		logs.PrintLogger().Error(err)
		return
	}

	err = controller.SaveCSVData(file)
	if err != nil {
		logs.PrintLogger().Error(err)
		return
	}
}


func main()  {

	//SaveTenTop("./data/github3.txt", *db)

	//SaveTopUsers()

	web_service.StartWebRequest()

}

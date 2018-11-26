package main

import (
	"github.com/github-user-behavior-analysis/backend-go/controller"
	"github.com/github-user-behavior-analysis/backend-go/db"
	"github.com/github-user-behavior-analysis/backend-go/logs"
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

	err = controller.SaveTopUser(file)
	if err != nil {
		logs.PrintLogger().Error(err)
		return
	}
}

func SaveCountryRepo()  {

	file, err := controller.OpenCSVFile("./backend-go/data/country_reposamount.csv")
	if err != nil {
		logs.PrintLogger().Error(err)
		return
	}

	err = controller.SaveCountryRepo(file)
	if err != nil {
		logs.PrintLogger().Error(err)
		return
	}
}

func SaveCountryUser()  {

	file, err := controller.OpenCSVFile("./backend-go/data/country_useramount.csv")
	if err != nil {
		logs.PrintLogger().Error(err)
		return
	}

	err = controller.SaveCountryUser(file)
	if err != nil {
		logs.PrintLogger().Error(err)
		return
	}
}

func SaveCountryPushs()  {

	file, err := controller.OpenCSVFile("./backend-go/data/country_pushes_oct.csv")
	if err != nil {
		logs.PrintLogger().Error(err)
		return
	}

	err = controller.SaveCountryPushs(file)
	if err != nil {
		logs.PrintLogger().Error(err)
		return
	}
}

func SaveProjectLanguage()  {
	file, err := controller.OpenCSVFile("./backend-go/data/language.csv")
	if err != nil {
		logs.PrintLogger().Error(err)
		return
	}

	err = controller.SaveProjectLanguage(file)
	if err != nil {
		logs.PrintLogger().Error(err)
		return
	}
}

func WriteMangoDB()  {
	lang := db.LanguageDAO{}
	lang.Server = "localhost:27017"
	lang.Database = "demoDB"

	lang.Connect()

	rankings, _ := db.ConnDB.GetAllDailyRanking()

	for _, rank :=range rankings {
		logs.PrintLogger().Info(rank)
		lang.Insert(rank)
	}

}

func main()  {

	//SaveTenTop("./data/github3.txt", *db)

	//SaveTopUsers()

	//SaveCountryRepo()

	//SaveCountryUser()

	//SaveCountryPushs()

	SaveProjectLanguage()

	//web_service.StartWebRequest()

	//WriteMangoDB()
}
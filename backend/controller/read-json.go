package controller

import (
	"encoding/json"
	"github.com/github-user-behavior-analysis/backend/db"
	"github.com/github-user-behavior-analysis/backend/logs"
	"github.com/github-user-behavior-analysis/backend/models"
	"io/ioutil"
	"os"
)

func LoadJsonFile(filename string) (*os.File, error)  {
	jsonFile, err := os.Open(filename)
	if err != nil {
		logs.PrintLogger().Error(err)
		return nil,err
	}
	//defer jsonFile.Close()

	return jsonFile,err
}

func ReadData(jsonFIle *os.File) ([]models.Ranking, error) {
	byteValue, err := ioutil.ReadAll(jsonFIle)
	if err != nil {
		logs.PrintLogger().Error(err)
		return nil, err
	}

	var rankings models.RankingsJSON

	//var rankings map[string]interface{}

	err = json.Unmarshal([]byte(byteValue),&rankings)
	if err != nil {
		logs.PrintLogger().Error(err)
	}

	//logs.PrintLogger().Info(len(rankings.Rankings))
	//
	//logs.PrintLogger().Info(rankings.Rankings[0].RepoNum)
	//
	//ranks := make([]*models.Ranking, 0)
	//
	//ranks = append(ranks, &models.Ranking{})
	//
	//return ranks, err

	return rankings.Rankings, err
}

func SaveData(rankings []models.Ranking, conn db.Database)  {

	if len(rankings) <= 0 {
		return
	}

	logs.PrintLogger().Info(rankings)


	for _, ranking := range rankings{
		err := conn.SaveTopTenRanking(&ranking)
		if err != nil {
			logs.PrintLogger().Error(err)
		}
		logs.PrintLogger().Info(rankings)
	}
}
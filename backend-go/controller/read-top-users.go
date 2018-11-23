package controller

import (
	"encoding/csv"
	"github.com/github-user-behavior-analysis/backend-go/db"
	"github.com/github-user-behavior-analysis/backend-go/logs"
	"github.com/github-user-behavior-analysis/backend-go/models"
	"os"
	"strconv"
)

func OpenCSVFile(filename string) (*os.File, error)  {
	file, err := os.Open(filename)
	if err != nil {
		logs.PrintLogger().Error(err)
		return nil, err
	}

	return file, err
}

func SaveCSVData(file *os.File) error {

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return err
	}

	var rank int64;

	for _, line := range lines {
		user := line[0]
		followers,_ := strconv.ParseInt(line[1],10,64)
		rank++;

		userFollower := &models.UserFollower{User:user, Follower:followers,Rank:rank}

		db.ConnDB.SaveTopUsers(userFollower)
	}

	return err

	//userFollowes := []*models.UserFollower{}
	//
	//err := gocsv.UnmarshalFile(file, &userFollowes)
	//if err != nil {
	//	logs.PrintLogger().Error(err)
	//	return err
	//}
	//
	//for _, user := range userFollowes {
	//	err := db.ConnDB.SaveTopUsers(user)
	//	if err != nil {
	//		logs.PrintLogger().Error(err)
	//		return err
	//	}
	//}
	//
	//return err


	//reader := csv.NewReader(file)
	//reader.Comma = ';'
	//lineCount := 0
	//
	//for  {
	//	record, err := reader.Read()
	//	if lineCount == 0 {
	//		continue
	//	}
	//
	//	if err == io.EOF {
	//		break
	//	} else if err != nil {
	//		logs.PrintLogger().Error(err)
	//		return
	//	}
	//
	//	db.ConnDB.SaveTopUsers()
	//
	//
	//
	//}




}
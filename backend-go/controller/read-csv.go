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

func SaveTopUser(file *os.File) error {

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
}

func SaveCountryRepo(file *os.File) error  {
	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return err
	}

	for _, line := range lines {
		country := line[0]
		//repos,_ := strconv.ParseInt(line[1],10,64)
		repos := line[1]

		db.ConnDB.SaveCountryRepos(country, repos)
	}

	return err
}

func SaveCountryUser(file *os.File) error  {
	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return err
	}

	for _, line := range lines {
		country := line[0]
		//repos,_ := strconv.ParseInt(line[1],10,64)
		user := line[1]

		db.ConnDB.SaveCountryUsers(country, user)
	}

	return err
}

func SaveCountryPushs(file *os.File) error  {
	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return err
	}

	for _, line := range lines {
		country := line[0]
		//repos,_ := strconv.ParseInt(line[1],10,64)
		push := line[1]

		db.ConnDB.SaveCountryPushs(country, push)
	}

	return err
}

func SaveProjectLanguage(file *os.File) error  {
	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return err
	}

	for _, line := range lines {
		timeStamp := line[0]
		languageName := line[1]
		amount := line[2]

		db.ConnDB.SaveProjectLanguage(timeStamp, languageName, amount)
	}

	return  err
}
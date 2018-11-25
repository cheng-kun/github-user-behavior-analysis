package db

import (
	"github.com/github-user-behavior-analysis/backend-go/logs"
	"github.com/github-user-behavior-analysis/backend-go/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type LanguageDAO struct {
	Server string
	Database string
}

var ConnDatabase *mgo.Database

const (
	COLLECTION = "demoDB"
)

func (l *LanguageDAO) Connect()  error {
	session, err := mgo.Dial(l.Server)
	if err != nil {
		logs.PrintLogger().Error(err)
		return err
	}
	ConnDatabase = session.DB(l.Database)

	return err
}

func (l *LanguageDAO) Insert (rank *models.Ranking) error  {
	err := ConnDatabase.C(COLLECTION).Insert(&rank)
	if err != nil {
		logs.PrintLogger().Error(err)
		return err
	}
	return err
}

func (l *LanguageDAO) FindAll() ([]models.Ranking, error)  {
	var rank []models.Ranking
	err := ConnDatabase.C(COLLECTION).Find(bson.M{}).All(&rank)
	if err != nil {
		logs.PrintLogger().Error(err)
		return nil, err
	}
	return rank, err
}

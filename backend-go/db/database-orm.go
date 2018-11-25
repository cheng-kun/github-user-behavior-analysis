package db

import (
	"fmt"
	"github.com/github-user-behavior-analysis/backend-go/conf"
	"github.com/github-user-behavior-analysis/backend-go/logs"
	"github.com/github-user-behavior-analysis/backend-go/models"
	"github.com/jinzhu/gorm"
)


//type DBPostgres struct {
//	*gorm.DB
//}

var ConnDBPostgres *gorm.DB

//func init()  {
//
//	logs.PrintLogger().Info("initilizing ConnDatabase connection ... ")
//
//	cfg, err := conf.LoadConfigFile("./config.toml")
//	if err != nil {
//		logs.PrintLogger().Error(err)
//		return
//	}
//
//	ConnDBPostgres, err = ConnectPostgre(*cfg)
//	if err != nil {
//		logs.PrintLogger().Error(err)
//		return
//	}
//
//	logs.PrintLogger().Info("Setup ConnDatabase connection sucessfully! ")
//}

// Connect to the ConnDatabase
func ConnectPostgre(cfg conf.Config) (*gorm.DB, error) {
	connStr := fmt.Sprintf("user='%s' password='%s' dbname='%s' host='%s' sslmode=disable",
		cfg.Database.Username, cfg.Database.Password, cfg.Database.DBName, cfg.Database.Host)
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}






func FindAll()(models.Ranking)  {

	rank := models.Ranking{}

	ConnDBPostgres.First(&rank)
	logs.PrintLogger().Info(rank)

	return rank
}
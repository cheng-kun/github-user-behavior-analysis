package db

import (
	"github.com/github-user-behavior-analysis/logs"
	"testing"
)

func TestDatabase_GetRankInfoByLanguage(t *testing.T) {
	//cfg, err := conf.LoadConfigFile("/home/nebula-ai-chengkun/gopath/src/github.com/github-user-behavior-analysis/conf/config.toml")
	//if err != nil {
	//	logs.PrintLogger().Error(err)
	//	return
	//}
	//
	//db, err := Connect(*cfg)
	//if err != nil {
	//	logs.PrintLogger().Error(err)
	//	return
	//}

	ranks, err := ConnDB.GetRankInfoByLanguage("php")
	if err != nil {
		logs.PrintLogger().Error(err)
		return
	}

	logs.PrintLogger().Info("len:",len(ranks))
}
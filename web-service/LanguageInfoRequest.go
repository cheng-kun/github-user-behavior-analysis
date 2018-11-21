package web_service

import (
	"github.com/gin-gonic/gin"
	"github.com/github-user-behavior-analysis/db"
	"github.com/github-user-behavior-analysis/logs"
	"net/http"
)

func GetLanguageRankByLanguage(r *gin.Engine){

	r.GET("/language/name/:name", func(c *gin.Context) {

		languageName := c.Param("name")

		ranks, err := db.ConnDB.GetRankInfoByLanguage(languageName)

		if err != nil {
			logs.PrintLogger().Error(err)
			c.String(http.StatusBadRequest, err.Error())
		}

		c.JSON(http.StatusOK, ranks)

	})

}

func GetDailyRankByDate(r *gin.Engine){

	r.GET("/language/date/:dateTime", func(c *gin.Context) {

		dateTime := c.Param("dateTime")

		ranks, err := db.ConnDB.GetDailyRankByDate(dateTime)

		if err != nil {
			logs.PrintLogger().Error(err)
			c.String(http.StatusBadRequest, err.Error())
		}

		c.JSON(http.StatusOK, ranks)

	})

}
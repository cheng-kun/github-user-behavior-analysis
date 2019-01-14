package web_service

import (
	"github.com/gin-gonic/gin"
	"github.com/github-user-behavior-analysis/backend/db"
	"github.com/github-user-behavior-analysis/backend/logs"
	"net/http"
)

func GetLanguageRankByLanguage(r *gin.Engine){

	r.GET("/language/nameandday/:name/:timeday", func(c *gin.Context) {

		languageName := c.Param("name")
		dateTime := c.Param("timeday")

		ranks, err := db.ConnDB.GetRankInfoByLanguage(languageName,dateTime)

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

func GetTopUserByAmount(r *gin.Engine)  {
	r.GET("/topuser/:amount", func(c *gin.Context) {
		amount := c.Param("amount")

		topUser, err := db.ConnDB.GetTopUsers(amount)
		if err != nil {
			logs.PrintLogger().Error(topUser)
			c.String(http.StatusBadRequest, err.Error())
		}

		c.JSON(http.StatusOK, topUser)

	})
}

func GetTopCountryByAmount(r *gin.Engine)  {
	r.GET("/topcountry/:amount", func(c *gin.Context) {

		amount := c.Param("amount")

		topUser, err := db.ConnDB.GetCountryUser(amount)
		if err != nil {
			logs.PrintLogger().Error(topUser)
			c.String(http.StatusBadRequest, err.Error())
		}

		c.JSON(http.StatusOK, topUser)

	})
}
package web_service

import (
	"github.com/gin-gonic/gin"
	"github.com/github-user-behavior-analysis/db"
	"github.com/github-user-behavior-analysis/logs"
	"net/http"
)

func GetLanguageRankByLanguage(r *gin.Engine){

	r.GET("/language/:id", func(c *gin.Context) {

		languageName := c.Param("id")

		ranks, err := db.ConnDB.GetRankInfoByLanguage(languageName)

		if err != nil {
			logs.PrintLogger().Error(err)
			c.String(http.StatusBadRequest, err.Error())
		}

		c.JSON(http.StatusOK, ranks)

	})

}
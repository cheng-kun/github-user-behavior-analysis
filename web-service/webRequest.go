package web_service

import (
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"time"
)

func StartWebRequest()  {

	r := gin.Default()

	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	GetLanguageRankByLanguage(r)

	r.Run(":8080")

}
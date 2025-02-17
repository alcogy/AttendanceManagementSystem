package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"app/models"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	// For HTML Template
	router.Static("/public", "./public")
	router.LoadHTMLGlob("public/*.tmpl")
	router.GET("/", models.HtmlHome)

	// For API set.
	api := router.Group("/api")
	{
		api.GET("/hello", models.ApiCheck)
		api.POST("/workin", models.ApiWorkIn)
		api.POST("/workout", models.ApiWorkOut)
		api.POST("/breakin", models.ApiBreakIn)
		api.POST("/breakout", models.ApiBreakOut)
	}

	router.Run(":9999")
}

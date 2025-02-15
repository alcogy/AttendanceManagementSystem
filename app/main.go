package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/public", "./public")
	router.LoadHTMLGlob("public/*.tmpl")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"movies": []string{"aaa", "bbb", "ccc"},
			"today":  "2/13 Mon.",
		})
	})
	router.Run(":8080")
}

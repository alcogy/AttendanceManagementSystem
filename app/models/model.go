package models

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Item struct {
	Date         string
	In           string
	Out          string
	BreakHours   string
	WorkingHours string
	Remarks      string
}

// Home View.
func HtmlHome(c *gin.Context) {
	current := time.Now()
	first := time.Date(current.Year(), current.Month(), 1, 0, 0, 0, 0, time.Local)
	end := first.AddDate(0, 1, -1)
	date := c.Query("d")
	fmt.Println(date)
	var items []Item
	for i := 1; i <= end.Day(); i++ {
		d := time.Date(current.Year(), current.Month(), i, 0, 0, 0, 0, time.Local)
		items = append(items, Item{
			Date:         d.Format("01/02 Mon"),
			In:           date,
			Out:          "10:12",
			BreakHours:   "10:12",
			WorkingHours: "10:12",
			Remarks:      "",
		})
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"items": items,
		"today": current.Format("2006/1/2 Mon"),
	})
}

// ApiCheck is Health check for api.
func ApiCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}

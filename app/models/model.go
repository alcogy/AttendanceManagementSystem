package models

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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
		})
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"userName": "Smith Get",
		"items":    items,
		"today":    current.Format("2006/1/2 Mon"),
	})
}

// ApiCheck is Health check for api.
func ApiCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}

func ApiWorkIn(c *gin.Context) {
	id := uuid.Must(uuid.NewRandom())
	now := time.Now()
	data := Attendance{
		ID:        id.String(),
		Date:      &now,
		InTime:    &now,
		OutTime:   nil,
		BreakHour: nil,
		UserID:    "abf874f5-cfef-4c17-9be5-532259e9be65",
	}

	RegistrateWorkIn(&data)

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func ApiWorkOut(c *gin.Context) {
	var json WorkOutBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	RegistrateWorkOut(json.ID)
	c.JSON(http.StatusOK, gin.H{
		"result": "ok",
	})
}

func ApiBreakIn(c *gin.Context) {
	// TODO Not Implement.
}

func ApiBreakOut(c *gin.Context) {
	// TODO Not Implement.
}

package main

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int    `json:"user_id" gorm:"primaryKey;column:user_id"`
	Username string `json:"username" gorm:"column:username"`
	//CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
}

type Memory struct {
	ID       int    `json:"memory_id" gorm:"primaryKey;column:memory_id"`
	Title    string `json:"title" gorm:"column:title"`
	Text     string `json:"text" gorm:"column:text"`
	PhotoURL string `json:"photo_url" gorm:"column:photo_url"`
	//CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
}

type Message struct {
	ID             int        `json:"message_id" gorm:"primaryKey;column:message_id"`
	NotificationID int        `json:"notification_id" gorm:"column:notification_id"`
	Title          string     `json:"title" gorm:"column:title"`
	Text           string     `json:"text" gorm:"column:description"`
	PhotoURL       string     `json:"photo_url" gorm:"column:photo_url"`
	ScheduledTime  *time.Time `json:"scheduled_time" gorm:"column:scheduled_time"`
	//CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
}

func GetUsers(c *gin.Context) {

	var users []User
	db, err := createDB()
	if err != nil {
		slog.Error(err.Error())
		c.JSON(500, nil)
	}
	err = db.Find(&users).Error
	if err != nil {
		slog.Error("User Find Error", err)
		c.JSON(500, nil)
	}

	c.JSON(200, users)
}

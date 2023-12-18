package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int    `json:"user_id" gorm:"primaryKey;column:user_id"`
	Username string `json:"username" gorm:"column:username"`
}

type Memory struct {
	ID        int       `json:"memory_id" gorm:"primaryKey;column:memory_id"`
	Title     string    `json:"title" gorm:"column:title"`
	Text      string    `json:"text" gorm:"column:text"`
	PhotoURL  string    `json:"photo_url" gorm:"column:photo_url"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}

type Message struct {
	ID            int       `json:"message_id" gorm:"primaryKey;column:message_id"`
	ToUser        int       `json:"to_user" gorm:"column:to_User"`
	FromUser      int       `json:"from_user" gorm:"column:from_User"`
	Title         string    `json:"title" gorm:"column:title"`
	Message       string    `json:"message" gorm:"column:message"`
	PhotoURL      string    `json:"photo_url" gorm:"column:photo_url"`
	ScheduledTime time.Time `json:"scheduled_time" gorm:"column:scheduled_time"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at"`
}

func PostMessage(c *gin.Context) {
	var message Message
	if err := c.ShouldBindJSON(&message); err != nil {
		slog.Error("Message JSON Bind Error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	slog.Info("Request ", message)

	db, err := createDB()
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	err = db.Create(&message).Error
	if err != nil {
		slog.Error("User Find Error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, nil)

}

func GetUsers(c *gin.Context) {

	var users []User
	db, err := createDB()
	if err != nil {
		slog.Error("DB Connection Error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	err = db.Find(&users).Error
	if err != nil {
		slog.Error("User Find Error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, users)
}

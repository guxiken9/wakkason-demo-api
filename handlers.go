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
	ID            int       `json:"memory_id" gorm:"primaryKey;column:memory_id"`
	Title         string    `json:"title" gorm:"column:title"`
	Memory        string    `json:"memory" gorm:"column:memory"`
	Image         string    `json:"image" gorm:"column:image"`
	PhotoOrignURL string    `json:"photo_origin_url" gorm:"column:photo_origin_url"`
	PhotoURL      string    `json:"photo_url" gorm:"column:photo_url"`
	CreatedBy     int       `json:"created_by" gorm:"column:created_by"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at"`
}

type Message struct {
	ID            int       `json:"message_id" gorm:"primaryKey;column:message_id"`
	ToUser        int       `json:"to_user" gorm:"column:to_User"`
	FromUser      int       `json:"from_user" gorm:"column:from_User"`
	Title         string    `json:"title" gorm:"column:title"`
	Message       string    `json:"message" gorm:"column:message"`
	Image         string    `json:"image" gorm:"column:image"`
	PhotoOrignURL string    `json:"photo_origin_url" gorm:"column:photo_origin_url"`
	PhotoURL      string    `json:"photo_url" gorm:"column:photo_url"`
	ScheduledTime time.Time `json:"scheduled_time" gorm:"column:scheduled_time"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at"`
}

func PostMemory(c *gin.Context) {
	var memory Memory
	if err := c.ShouldBindJSON(&memory); err != nil {
		slog.Error("Message JSON Bind Error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	db, err := CreateDB()
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if memory.Image != "" {
		b, err := DecodeBase64(memory.Image)
		if err != nil {
			slog.Error("Image Decode Error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		r, err := PutImage(b)
		if err != nil {
			slog.Error("S3 Put Image Error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		memory.PhotoOrignURL = r.Key
		memory.PhotoURL = r.PreSignedURL
	}

	err = db.Create(&memory).Error
	if err != nil {
		slog.Error("User Find Error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, nil)

}

func PostMessage(c *gin.Context) {
	var message Message
	if err := c.ShouldBindJSON(&message); err != nil {
		slog.Error("Message JSON Bind Error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	db, err := CreateDB()
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
	db, err := CreateDB()
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

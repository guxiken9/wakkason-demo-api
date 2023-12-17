package main

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserID    int       `json:"user_id" gorm:"primaryKey;column:user_id"`
	Username  string    `json:"username" gorm:"column:username"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}

func GetUsers(c *gin.Context) {

	var users []User
	db := createDB()
	if db.Error != nil {
		c.JSON(500, nil)
	}
	err := db.Find(users).Error
	if err != nil {
		slog.Error("User Find Error", err)
		c.JSON(500, nil)
	}

	c.JSON(200, users)
}

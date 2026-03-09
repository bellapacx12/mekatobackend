package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"bingo-backend/pkg/db"

	"github.com/gin-gonic/gin"
)

type TelegramUser struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type AuthRequest struct {
	InitData string `json:"initData"`
	User     string `json:"user"`
}

func TelegramLogin(c *gin.Context) {

	var req AuthRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	if !VerifyTelegram(req.InitData) {
		c.JSON(401, gin.H{"error": "telegram verification failed"})
		return
	}

	var tgUser TelegramUser

	json.Unmarshal([]byte(req.User), &tgUser)

	var userID int64

	err := db.Pool.QueryRow(context.Background(),
		`
		INSERT INTO users (telegram_id, username, first_name, last_name)
		VALUES ($1,$2,$3,$4)
		ON CONFLICT (telegram_id)
		DO UPDATE SET username=$2
		RETURNING id
		`,
		tgUser.ID,
		tgUser.Username,
		tgUser.FirstName,
		tgUser.LastName,
	).Scan(&userID)

	if err != nil {
		c.JSON(500, gin.H{"error": "database error"})
		return
	}

	token, _ := GenerateJWT(userID)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  tgUser,
	})
}
package auth

import (
	"context"
	"encoding/json"
	"fmt"
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
	fmt.Println("==== TelegramLogin called ====")

	var req AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("Failed to bind JSON:", err)
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}
	fmt.Println("Received initData:", req.InitData)
	fmt.Println("Received user JSON:", req.User)

	if !VerifyTelegram(req.InitData) {
		fmt.Println("Telegram verification failed")
		c.JSON(401, gin.H{"error": "telegram verification failed"})
		return
	}
	fmt.Println("Telegram verification passed")

	var tgUser TelegramUser
	if err := json.Unmarshal([]byte(req.User), &tgUser); err != nil {
		fmt.Println("Failed to unmarshal Telegram user:", err)
		c.JSON(400, gin.H{"error": "invalid user data"})
		return
	}
	fmt.Printf("Parsed Telegram user: %+v\n", tgUser)

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
		fmt.Println("Database error:", err)
		c.JSON(500, gin.H{"error": "database error"})
		return
	}
	fmt.Println("User inserted/updated with ID:", userID)

	token, err := GenerateJWT(userID)
	if err != nil {
		fmt.Println("JWT generation error:", err)
		c.JSON(500, gin.H{"error": "could not generate token"})
		return
	}
	fmt.Println("Generated JWT token:", token)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  tgUser,
	})
	fmt.Println("==== TelegramLogin finished ====")
}
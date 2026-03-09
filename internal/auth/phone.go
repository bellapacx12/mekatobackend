package auth

import (
	"context"

	"bingo-backend/pkg/db"

	"github.com/gin-gonic/gin"
)

type PhoneRequest struct {
	Phone string `json:"phone"`
}

func UpdatePhone(c *gin.Context) {

	userID := c.GetInt64("user_id")

	var req PhoneRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid phone"})
		return
	}

	_, err := db.Pool.Exec(context.Background(),
		`UPDATE users SET phone_number=$1 WHERE id=$2`,
		req.Phone,
		userID,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": "db error"})
		return
	}

	c.JSON(200, gin.H{"success": true})
}
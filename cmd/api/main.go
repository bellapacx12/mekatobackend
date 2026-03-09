package main

import (
	"log"

	"bingo-backend/internal/auth"
	"bingo-backend/internal/config"
	"bingo-backend/pkg/db"
	"bingo-backend/pkg/redis"

	"github.com/gin-gonic/gin"
)

func main() {

	config.Load()

	db.InitPostgres()

	redis.InitRedis()

	r := gin.Default()
    
	r.POST("/auth/telegram", auth.TelegramLogin)

authGroup := r.Group("/")
authGroup.Use(auth.AuthMiddleware())

authGroup.POST("/users/phone", auth.UpdatePhone)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	log.Println("Server running on port", config.App.Port)

	r.Run(":" + config.App.Port)
}
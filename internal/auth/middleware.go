package auth

import (
	"strings"

	"bingo-backend/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatus(401)
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]

		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.App.JWTSecret), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			userID := int64(claims["user_id"].(float64))

			c.Set("user_id", userID)

			c.Next()

		} else {
			c.AbortWithStatus(401)
		}
	}
}
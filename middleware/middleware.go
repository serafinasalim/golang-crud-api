package middleware

import (
	"net/http"
	"strings"

	"golang-crud-api/helper"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware() gin.HandlerFunc {
	const Bearer = "Bearer "
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is missing"})
			c.Abort()
			return
		}

		if strings.HasPrefix(tokenString, Bearer) {
			tokenString = tokenString[7:]
		}

		parsedToken, err := helper.ValidateJWT(tokenString)
		if err != nil || !parsedToken.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or Expired Token"})
			c.Abort()
			return
		}

		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or Expired Token"})
			c.Abort()
			return
		}

		userId := claims["userId"].(float64)
		username := claims["username"].(string)

		c.Set("userId", userId)
		c.Set("username", username)

		c.Next()
	}
}

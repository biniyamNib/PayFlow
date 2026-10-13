package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
	
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H {
				"error": "Invalid authorization header format",
			})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H {
				"error": "Bearer token is required",
			})
			c.Abort()
			return 
		}

		jwtSecret := os.Getenv("JWT_SECRET")

		if jwtSecret == "" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
			c.Abort()
			return
		}

        token, err := jwt.Parse(tokenString,func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(jwtSecret), nil
		    }, 
	    )

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H {
				"error": "Invalid or expired token",
			})
			c.Abort()
			return 
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token claims",
			})
			c.Abort()
			return
		}

		userID, userIDOK := claims["user_id"].(string)
		merchantID, merchantIDOK := claims["merchant_id"].(string)
		role, roleOK := claims["role"].(string)
        if !userIDOK || userID == "" ||
			!merchantIDOK || merchantID == "" ||
			!roleOK || role == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token claims",
			})
			c.Abort()
			return
		}

		c.Set("user_id", userID)
        c.Set("merchant_id", merchantID)
		c.Set("role", role) 
		
		c.Next()
	}
}
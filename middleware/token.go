package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/shuttlersIT/itsm-mvp/stafftokenization"
	"github.com/shuttlersIT/itsm-mvp/structs"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		// Validate the token
		claims := &structs.AccessToken{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("your-secret-key"), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Set user ID in the context for further use
		c.Set("userID", claims.UserID)

		c.Next()
	}
}

// authMiddleware is a middleware function for protecting routes that require authorization
func AuthMiddlewareToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the token from the request header or query parameter
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// Validate the token
		claims, err := stafftokenization.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Attach the user ID to the context for use in subsequent handlers
		c.Set("user_id", claims.UserID)

		// Continue to the next handler
		c.Next()
	}
}

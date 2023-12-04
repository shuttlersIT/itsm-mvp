package handlers

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/itsm-mvp/stafftokenization"
	"github.com/shuttlersIT/itsm-mvp/structs"
)

// Generate User Token
func JwtToken(userID int, username string) (bool, string, error) {
	status := false

	// Generate a token
	token, err := stafftokenization.GenerateToken(userID, username)
	if err != nil {
		fmt.Println("Error generating token:", err)
		status = false
		return status, token, fmt.Errorf("unable to generate token")
	}
	status = true

	//fmt.Println("Generated Token:", token)
	return status, token, nil
}

func RetrieveClaim(c *gin.Context, token string) (*stafftokenization.CustomClaims, bool, error) {
	status := false

	// Parse the token
	claims, err := stafftokenization.ParseToken(token)
	if err != nil {
		fmt.Println("Error parsing token:", err)
		status = false
		return nil, status, fmt.Errorf("unable to retrieve api session details")
	}

	status = true
	//fmt.Println("User ID:", claims.UserID)
	//fmt.Println("Username:", claims.Username)

	return claims, status, err
}

func RefreshTokenHandler(c *gin.Context) {
	refreshTokenString := c.PostForm("refresh_token")

	claims := &structs.RefreshToken{}
	token, err := jwt.ParseWithClaims(refreshTokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-secret-key"), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	// Generate a new access token
	newAccessToken, err := stafftokenization.GenerateAccessToken(claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": newAccessToken,
		// Include other necessary data in the response
	})
}

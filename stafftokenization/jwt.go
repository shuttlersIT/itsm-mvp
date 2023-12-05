package stafftokenization

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

//godotenv.Load(".env")

func GoDotEnv(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

// CustomClaims represents the JWT claims structure
type CustomClaims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken generates a new JWT token
func GenerateToken(userID int, username string) (string, error) {
	// godotenv package
	jToken := GoDotEnv("JTOKEN")

	claims := CustomClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jToken)
}

// ParseToken parses a JWT token and returns the claims
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return token, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

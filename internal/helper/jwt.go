package helper

import (
	"fmt"
	"os"
	"time"

	// "github.com/BoruTamena/models"
	"github.com/BoruTamena/go_chat/internal/constant/models/dto"
	"github.com/golang-jwt/jwt/v5"
)

const (
	accessTokenExpireDuration  = time.Minute * 15
	refreshTokenExpireDuration = time.Hour * 24 * 7 // for 7days or 1 week
)

func CreateToken(userClaim dto.User) (string, string, error) {
	// Creating access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": userClaim, // Change "UserId" to "userId" for consistency
		"exp":  time.Now().Add(accessTokenExpireDuration).Unix(),
	})

	// Creating refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": userClaim,
		"exp":  time.Now().Add(refreshTokenExpireDuration).Unix(),
	})

	// Signing tokens
	accessTokenStr, err := accessToken.SignedString([]byte(os.Getenv("SCERATEKEY")))
	if err != nil {
		return "", "", err
	}

	refreshTokenStr, err := refreshToken.SignedString([]byte(os.Getenv("SCERATEKEY")))
	if err != nil {
		return "", "", err
	}

	return accessTokenStr, refreshTokenStr, nil
}

func GenerateToken(user dto.User) (string, error) {
	claims := jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(accessTokenExpireDuration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SCERATEKEY")))
}

// func RefreshAccessToken(refreshToken string) (string, error) {
// 	// Parse refresh token
// 	token, err := parseToken(refreshToken)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Extract user ID from refresh token
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok {
// 		return "", fmt.Errorf("invalid refresh token claims")
// 	}
// 	userID, ok := claims["userId"].(float64)
// 	if !ok {
// 		return "", fmt.Errorf("invalid user ID in refresh token")
// 	}

// 	// Generate new access token
// 	newAccessToken, err := GenerateToken(int(userID))
// 	if err != nil {
// 		return "", err
// 	}

// 	return newAccessToken, nil
// }

func ParseAccessToken(accessToken string) (*jwt.Token, error) {
	return parseToken(accessToken)
}

func ParseRefreshToken(refreshToken string) (*jwt.Token, error) {
	return parseToken(refreshToken)
}

func parseToken(tokenString string) (*jwt.Token, error) {
	// Parse token
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SCERATEKEY")), nil
	})

	// Check for parsing errors
	if err != nil {
		return nil, err
	}

	// Check token validity
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

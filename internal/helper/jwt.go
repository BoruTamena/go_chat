package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/BoruTamena/go_chat/internal/constant/errors"
	"github.com/BoruTamena/go_chat/internal/constant/models/dto"
	"github.com/golang-jwt/jwt/v5"
)

const (
	accessTokenExpireDuration  = time.Minute * 15
	refreshTokenExpireDuration = time.Hour * 24 * 7 // for 7days or 1 week
)

func UserMarshal(userClaim dto.User) (string, error) {
	user, err := json.Marshal(userClaim)

	if err != nil {

		err = errors.MarshalErr.Wrap(err, "marshaling user claim err").
			WithProperty(errors.ErrorCode, http.StatusUnauthorized)

		fmt.Println(err.Error())
		return "", err
	}

	return string(user), nil

}

func CreateToken(userClaim dto.User) (string, string, error) {

	user_claim, err := UserMarshal(userClaim)

	if err != nil {
		return "", "", err
	}

	// Creating access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user_claim,
		"exp":  time.Now().Add(accessTokenExpireDuration).Unix(),
	})

	// Creating refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user_claim,
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

	_user, err := UserMarshal(user)

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"user": _user,
		"exp":  time.Now().Add(accessTokenExpireDuration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SCERATEKEY")))
}

func RefreshAccessToken(refreshToken string) (string, error) {

	var user dto.User
	// Parse refresh token
	token, err := parseToken(refreshToken)
	if err != nil {
		return "", err
	}

	// Extract user ID from refresh token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid refresh token claims")
	}
	user_claim, ok := claims["user"].(string)

	if !ok {
		return "", fmt.Errorf("invalid user in refresh token")
	}

	if err := json.Unmarshal([]byte(user_claim), &user); err != nil {

		err = errors.UnMarshalErr.Wrap(err, "unmarshalling user claim").
			WithProperty(errors.ErrorCode, 500)

		fmt.Println(err.Error())

		return "", err

	}

	// Generate new access token
	newAccessToken, err := GenerateToken(user)
	if err != nil {
		return "", err
	}

	return newAccessToken, nil
}

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

package jwt

import (
	"errors"
	"time"

	"example.com/events/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwtToken(userId int64, email string) (string, error) {
	jwtClaims := jwt.MapClaims{
		"userId": userId,
		"email":  email,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	return token.SignedString([]byte(config.AppConfig.SecretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return []byte(config.AppConfig.SecretKey), nil
	})

	if err != nil {
		return 0, errors.New("Could not parse jwt token")
	}

	if(!parsedToken.Valid) {
		return 0, errors.New("Invalid JWT token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if(!ok) {
		return 0, errors.New("Invalid Token claim")
	}

	userId := int64(claims["userId"].(float64))
	return userId, nil
}

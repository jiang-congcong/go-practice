package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	UserId   string
	Username string
	Claims   jwt.StandardClaims
}

var signKey = "jwt golang sign key"

func IssueToken(userId string, username string) (string, error) {
	if len(userId) == 0 || len(username) == 0 {
		return "", errors.New("issueToken failed")
	}

	claims := jwt.MapClaims{
		"userId":   userId,
		"username": username,
		"claims": jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 24*60*60, // 定义过期时间 1d
			IssuedAt:  time.Now().Unix(),
			Issuer:    username,
			Id:        userId,
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := t.SignedString(signKey)
	if err != nil {
		return "", errors.New("issueToken failed")
	}

	return token, nil
}

func PraseToken(tokenString string) (*jwt.Token, *jwt.MapClaims, error) {
	mapClaims := &jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, mapClaims, func(t *jwt.Token) (interface{}, error) {
		return signKey, nil
	})

	return token, mapClaims, err
}

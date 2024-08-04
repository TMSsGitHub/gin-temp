package utils

import (
	"gin-temp/conf"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	Subject uint64 `json:"sub"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(sub uint64, duration time.Duration) (string, error) {
	claims := &Claims{
		Subject: sub,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(conf.Cfg.App.AuthSecret))
}

func GenerateRefreshToken(sub uint64) (string, error) {
	return "", nil
}

func ValidateAccessToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.Cfg.App.AuthSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

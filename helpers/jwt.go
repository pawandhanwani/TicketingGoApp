package helpers

import (
	"fmt"
	"time"

	"../configs"
	"github.com/golang-jwt/jwt"
)

func GenerateJwt(userIdentifier uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Duration(configs.EnvJwtExpirySeconds()) * time.Second).Unix()
	claims["authorized"] = true
	claims["user"] = userIdentifier

	tokenString, err := token.SignedString([]byte(configs.EnvJwtSecret()))
	if err != nil {

		return "", err
	}

	return tokenString, nil
}

func ValidateJwt(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there's an error with the signing method")
		}
		return []byte(configs.EnvJwtSecret()), nil
	})

	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		username := claims["user"].(float64)
		return uint(username), nil
	}
	return 0, nil
}

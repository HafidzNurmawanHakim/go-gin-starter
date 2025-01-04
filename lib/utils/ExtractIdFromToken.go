package utils

import (
	"gin-template/lib/schema"

	"github.com/dgrijalva/jwt-go"
)

func ExtraxtIdFromToken(tokenString string, secret string) (int, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, schema.ErrUnexpectedSigningMethod
		}
		return []byte(secret), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if id, exists := claims["id"]; exists {
			if idFloat, ok := id.(float64); ok {
				return int(idFloat), nil
			}
		}
		return 0, schema.ErrUserNotFound
	}

	return 0, schema.ErrInvalidToken
}

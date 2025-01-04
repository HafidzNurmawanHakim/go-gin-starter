package utils

import (
	"gin-template/lib/schema"
	"log"

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
		log.Println(claims)
		if id, err := claims["sub"].(float64); !err {
			log.Println(id)
			return int(id), nil

		}
		return 0, schema.ErrUserNotFound
	}

	return 0, schema.ErrInvalidToken
}

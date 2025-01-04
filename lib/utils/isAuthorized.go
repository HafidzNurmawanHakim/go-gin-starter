package utils

import (
	"gin-template/lib/schema"

	"github.com/dgrijalva/jwt-go"
)

func IsAuthorized(token string, secret string) (bool, error) {
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, schema.ErrUnexpectedSigningMethod
		}
		return []byte(secret), nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

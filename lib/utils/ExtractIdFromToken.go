package utils

import (
	"gin-template/lib/schema"
	"log"
	"strconv"

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
		log.Println(err)
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        log.Println("claims:", claims)

        sub, ok := claims["sub"].(string)
        if !ok {
            log.Println("sub is missing or not a number")
            return 0, schema.ErrUserNotFound
        }

		id, err := strconv.Atoi(sub)

		if err != nil {
			log.Println("sub is missing or not a number")
            return 0, schema.ErrUserNotFound
		}

        log.Println("sub value:", sub)
        return int(id), nil
    }

	return 0, schema.ErrInvalidToken
}
